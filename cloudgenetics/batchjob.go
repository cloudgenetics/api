package cloudgenetics

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/batch"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BatchJob struct {
	JobDefinition    string   `json:"jobDefinition"`
	JobName          string   `json:"jobName"`
	JobQueue         string   `json:"jobQueue"`
	NextflowPipeline string   `json:"nextflowPipeline"`
	ContainerCommand []string `json:"containerCommand,omitempty"`
}

type JobDetails struct {
	JobArn  string `json:"jobArn"`
	JobId   string `json:"jobId"`
	JobName string `json:"jobName"`
}

func jobParameters(job BatchJob) *batch.SubmitJobInput {
	commands := []*string{}
	commands = append(commands, aws.String(job.NextflowPipeline))
	for _, cmd := range job.ContainerCommand {
		commands = append(commands, aws.String(cmd))
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

func submitJob(input *batch.SubmitJobInput) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := batch.New(sess, &aws.Config{Region: aws.String("us-east-2")})
	result, err := svc.SubmitJob(input)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(result)
}

func submitBatchJob(c *gin.Context, db *gorm.DB) {
	var job BatchJob
	c.BindJSON(&job)
	params := jobParameters(job)
	submitJob(params)
}
