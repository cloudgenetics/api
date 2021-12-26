package cloudgenetics

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type FileUpload struct {
	// Unique ID
	Name string `json:"name, omitempty"`
	// Title of project
	Type string `json:"mime, omitempty"`
}

func presignedUrl(c *gin.Context) (string, string) {

	var file FileUpload
	c.BindJSON(&file)
	// Initialize a session the SDK will use credentials in
	// ~/.aws/credentials.
	awsregion := os.Getenv("AWS_REGION")
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(awsregion)},
	)

	// Create S3 service client
	svc := s3.New(sess)
	bucket := os.Getenv("AWS_S3_BUCKET")
	datsetid := uuid.New().String()
	filename := datsetid + "/" + file.Name
	req, _ := svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
	})
	url, err := req.Presign(15 * time.Minute)

	if err != nil {
		log.Println("Failed to sign request", err)
	}
	return datsetid, url
}
