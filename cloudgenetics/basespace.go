package cloudgenetics

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/google/uuid"

	"encoding/json"
)

type Basespace struct {
	AccessToken string    `json:"accessToken"`
	Projectid   string    `json:"projectId"`
	Uuid        uuid.UUID `json:"uuid,omitempty"`
}

type Response struct {
	Status int    `json:"statusCode"`
	Files  string `json:"files"`
}

func basespace_s3upload(bsaccount Basespace) []File {
	// Create Lambda service client
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	client := lambda.New(sess, &aws.Config{Region: aws.String("us-east-2")})

	// Send request
	p := Basespace{
		AccessToken: bsaccount.AccessToken,
		Projectid:   bsaccount.Projectid,
		Uuid:        bsaccount.Uuid,
	}

	payload, err := json.Marshal(p)
	// This is the required format for the lambda request body.
	if err != nil {
		log.Println("Json Marshalling error")
	}

	result, err := client.Invoke(&lambda.InvokeInput{
		FunctionName:   aws.String("basespace-s3"),
		InvocationType: aws.String("RequestResponse"),
		Payload:        payload})

	if err != nil {
		log.Println("Error calling basespace-s3: ", err)
	}
	// Get result from Lambda
	var resp Response
	err = json.Unmarshal(result.Payload, &resp)
	if err != nil {
		log.Println("Error unmarshalling File response")
	}

	var files []File
	if resp.Status == 200 {
		fileString := resp.Files
		jerr := json.Unmarshal([]byte(fileString), &files)
		if jerr != nil {
			log.Println(jerr)
		}
	}
	return files
}
