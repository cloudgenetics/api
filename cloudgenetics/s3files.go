package cloudgenetics

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func listS3Objects(c *gin.Context, db *gorm.DB) []string {
	awsregion := os.Getenv("AWS_REGION")
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String(awsregion)},
	)

	uuid, _ := uuid.Parse(c.Param("uuid"))
	// Create S3 service client
	svc := s3.New(sess)
	bucket := os.Getenv("AWS_S3_BUCKET")
	params := &s3.ListObjectsInput{
		Bucket: aws.String(bucket),
		Prefix: aws.String(uuid.String()),
	}

	resp, _ := svc.ListObjects(params)
	files := []string{}
	for _, key := range resp.Contents {
		files = append(files, *key.Key)
	}
	return files
}
