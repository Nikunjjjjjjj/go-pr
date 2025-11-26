// package utils

// import (
// 	"context"

// 	"github.com/aws/aws-sdk-go-v2/config"
// 	"github.com/aws/aws-sdk-go-v2/service/s3"
// )

// func AwsUpload(path, fileName string) (string,error){
// 	bucket:=""
// 	region:=""

// 	cfg,err:=config.LoadDefaultConfig(context.TODO(),config.WithRegion(region))
// 	if err !=nil{
// 		return "",err
// 	}

//		client:=s3
//	}
package utils

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func AwsUpload(path, fileName string) (string, error) {
	// Load AWS config
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return "", fmt.Errorf("unable to load AWS config: %w", err)
	}

	// Open the file
	//fullPath := filepath.Join(path, fileName)
	file, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Create S3 client
	client := s3.NewFromConfig(cfg)

	// Create uploader
	uploader := manager.NewUploader(client)

	// Upload
	//var lConfig =co
	bucket := os.Getenv("BUCKET")
	//fmt.Println(bucket)
	key := fileName // or "uploads/" + fileName

	_, err = uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   file,
		ACL:    "public-read", // optional
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload to S3: %w", err)
	}

	// Construct region-specific public URL
	url := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", bucket, cfg.Region, key)
	return url, nil
}
