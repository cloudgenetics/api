package cloudgenetics

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/batch"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func listJobs(c *gin.Context, db *gorm.DB) []Job {
	var jobs []Job
	userid := userid(c, db)
	db.Find(&jobs, Dataset{OwnerID: userid})
	return jobs
}

func getJobDescription(c *gin.Context, db *gorm.DB) batch.DescribeJobsOutput {
	var job Job
	jobid, err := uuid.Parse(c.Param("uuid"))
	owner_id := userid(c, db)
	if err != nil {
		log.Print("Get job command: ", err)
	}
	db.Find(&job, Job{ID: jobid, OwnerID: owner_id})

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := batch.New(sess, &aws.Config{Region: aws.String("us-east-2")})

	input := &batch.DescribeJobsInput{
		Jobs: []*string{
			aws.String(job.JobId),
		},
	}

	result, _ := svc.DescribeJobs(input)
	return *result
}

func getJobInfo(c *gin.Context, db *gorm.DB) Job {
	var job Job
	jobid, err := uuid.Parse(c.Param("uuid"))
	owner_id := userid(c, db)
	if err != nil {
		log.Print("Get files dataset: ", err)
	}
	db.Find(&job, Job{ID: jobid, OwnerID: owner_id})
	return job
}
