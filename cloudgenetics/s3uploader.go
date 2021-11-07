package cloudgenetics

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"log"
	"time"

	"github.com/gin-gonic/gin"
)

type FileUpload struct {
	// Unique ID
	Name string `json:"name, omitempty"`
	// Title of project
	Type string `json:"mime, omitempty"`
}

func presignedUrl(c *gin.Context) string {

	var file FileUpload
	c.BindJSON(&file)
	log.Println(file.Name, file.Type)
	// Initialize a session in us-west-2 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)

	// Create S3 service client
	svc := s3.New(sess)

	req, _ := svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String("presigned-uploader"),
		Key:    aws.String(file.Name),
	})
	urlStr, err := req.Presign(15 * time.Minute)

	if err != nil {
		log.Println("Failed to sign request", err)
	}

	log.Println("The URL is", urlStr)
	return urlStr
}
