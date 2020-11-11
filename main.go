package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
	"net/http"
	"os"
)

func listObjectsV2() ([]string ,error) {
	input := &s3.ListObjectsV2Input{
		Bucket:  aws.String(os.Getenv("S3_BUCKET")),
		MaxKeys: aws.Int64(5),
	}

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

	var objects []string

	for _, v := range result.Contents {
		objects = append(objects, *v.Key)
	}

	return objects, nil
}

func httpListen(w http.ResponseWriter, r *http.Request) {
	objects, err := listObjectsV2()
	if err != nil {
		fmt.Println(err)
	}

	log.Println(objects)
	log.Println(r)
}

func main() {
	http.HandleFunc("/s3", httpListen)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
