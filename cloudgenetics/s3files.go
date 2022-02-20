package cloudgenetics

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func getPresignedPlots(c *gin.Context, db *gorm.DB) []string {
	uuid, _ := uuid.Parse(c.Param("uuid"))

	// Create S3 service client
	awsregion := os.Getenv("AWS_REGION")
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String(awsregion)},
	)
	svc := s3.New(sess)
	bucket := os.Getenv("AWS_S3_BUCKET")

	params := &s3.ListObjectsInput{
		Bucket: aws.String(bucket),
		Prefix: aws.String(uuid.String() + "/123_reads/plots/"),
	}

	resp, _ := svc.ListObjects(params)
	s3urls := []string{}
	for _, key := range resp.Contents {
		file := *key.Key
		if filepath.Ext(file) == ".png" {
			req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
				Bucket: aws.String(bucket),
				Key:    aws.String(file),
			})
			url, err := req.Presign(15 * time.Minute)
			if err != nil {
				log.Println("Failed to sign request", err)
			}
			s3urls = append(s3urls, url)
		}

	}
	return s3urls
}

func getPresignedReports(c *gin.Context, db *gorm.DB) []string {
	uuid, _ := uuid.Parse(c.Param("uuid"))

	// Create S3 service client
	awsregion := os.Getenv("AWS_REGION")
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String(awsregion)},
	)
	svc := s3.New(sess)
	bucket := os.Getenv("AWS_S3_BUCKET")

	params := &s3.ListObjectsInput{
		Bucket: aws.String(bucket),
		Prefix: aws.String(uuid.String() + "/pipeline_info/"),
	}

	resp, _ := svc.ListObjects(params)
	s3urls := []string{}
	for _, key := range resp.Contents {
		file := *key.Key
		if filepath.Ext(file) == ".html" {
			req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
				Bucket: aws.String(bucket),
				Key:    aws.String(file),
			})
			url, err := req.Presign(15 * time.Minute)
			if err != nil {
				log.Println("Failed to sign request", err)
			}
			s3urls = append(s3urls, url)
		}
	}
	return s3urls
}

func listS3Objects(c *gin.Context, db *gorm.DB) []string {
	uuid, _ := uuid.Parse(c.Param("uuid"))

	// Create S3 service client
	awsregion := os.Getenv("AWS_REGION")
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String(awsregion)},
	)
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
