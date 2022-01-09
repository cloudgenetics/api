package cloudgenetics

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"gorm.io/gorm"

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
	// UUID
	UUID string `json:"uuid, omitempty"`
}

func presignedUrl(c *gin.Context, db *gorm.DB) (string, string) {

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
	// Set UUID if not found in the request
	datasetid := file.UUID
	filename := datasetid + "/" + file.Name
	req, _ := svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
	})
	url, err := req.Presign(15 * time.Minute)

	if err != nil {
		log.Println("Failed to sign request", err)
	}
	return datasetid, url
}

func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}
