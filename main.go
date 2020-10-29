package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
)

func listObjectsV2(input *s3.ListObjectsV2Input) (*s3.ListObjectsV2Output, error) {
	svc := s3.New(session.New())
	result, err := svc.ListObjectsV2(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchBucket:
				fmt.Println(s3.ErrCodeNoSuchBucket, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
	}

	return result, nil
}

func main() {
	input := &s3.ListObjectsV2Input{
		Bucket:  aws.String(os.Getenv("S3_BUCKET")),
		MaxKeys: aws.Int64(5),
	}

	result, err := listObjectsV2(input)
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range result.Contents {
		fmt.Println(*v.Key)
	}
}
