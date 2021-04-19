package models

import (
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type Bucket struct {
	BucketName string `json:"bucket_name"`
}

type GeneralResponse struct {
	Error  string      `json:"error"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func CreateNewBucket(bucketname string) GeneralResponse {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2")},
	)
	if err != nil {
		fmt.Println("Error Bucket:", err)
	}

	// Create S3 service client
	svc := s3.New(sess)
	_, err = svc.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(bucketname),
	})
	if err != nil {
		errors.ExitErrorf("Unable to create bucket %q, %v", bucketname, err)
	}

	// Wait until bucket is created before finishing
	fmt.Printf("Waiting for bucket %q to be created...\n", bucketname)

	err = svc.WaitUntilBucketExists(&s3.HeadBucketInput{
		Bucket: aws.String(bucketname),
	})
	responseResult := GeneralResponse{
		Error:  err,
		Status: "200",
		Data:   bucketname,
	}
	return responseResult
}
