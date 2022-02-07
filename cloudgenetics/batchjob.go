package cloudgenetics

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/batch"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Job struct {
	gorm.Model
	ID               uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"uuid,omitempty"`
	JobDefinition    string    `json:"jobDefinition"`
	JobName          string    `json:"jobName"`
	JobArn           string    `json:"jobArn,omitempty"`
	JobQueue         string    `json:"jobQueue"`
	NextflowPipeline string    `json:"nextflowPipeline"`
	ContainerCommand []string  `gorm:"-" json:"containerCommand,omitempty"`
	Command          string    `json:"Command,omitempty"`
	ResultsDir       string    `json:"resultsDir,omitempty"`
	OwnerID          uuid.UUID `gorm:"type:uint" json:"omitempty"`
	User             User      `gorm:"foreignKey:OwnerID" json:"omitempty"`
}

type JobDetails struct {
	JobArn  string `json:"jobArn"`
	JobId   string `json:"jobId"`
	JobName string `json:"jobName"`
}

func jobParameters(job *Job) *batch.SubmitJobInput {
	commands := []*string{}
	commands = append(commands, aws.String(job.NextflowPipeline))
	job.ResultsDir = uuid.NewString()
	outdir := "s3://" + os.Getenv("AWS_S3_BUCKET") + "/" + job.ResultsDir
	commands = append(commands, aws.String("--outdir"))
	commands = append(commands, aws.String(outdir))
	for _, cmd := range job.ContainerCommand {
		commands = append(commands, aws.String(cmd))
		job.Command = job.Command + "," + cmd
	}

	input := &batch.SubmitJobInput{
		JobDefinition: aws.String(job.JobDefinition),
		JobName:       aws.String(job.JobName),
		JobQueue:      aws.String(job.JobQueue),

		ContainerOverrides: &batch.ContainerOverrides{
			Command: commands,
		},
	}
	return input
}

func submitJob(input *batch.SubmitJobInput) string {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := batch.New(sess, &aws.Config{Region: aws.String("us-east-2")})
	result, err := svc.SubmitJob(input)
	if err != nil {
		log.Print("Submit job: ", err.Error())
	}
	job_arn := *result.JobArn
	return job_arn
}

func submitBatchJob(c *gin.Context, db *gorm.DB) {
	var job Job
	c.BindJSON(&job)
	job.OwnerID = userid(c, db)
	params := jobParameters(&job)
	job.JobArn = submitJob(params)
	dbresp := db.Save(&job)
	if dbresp.Error != nil {
		log.Print("Create job: ", dbresp.Error)
	}
}
