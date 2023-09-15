package service

import (
	"context"
	"log"
	"mime/multipart"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type AwsService interface {
	Upload(file *multipart.FileHeader) (*manager.UploadOutput, error)
}

type awsService struct {
	S3Client *s3.Client
}

func createS3() *s3.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Printf("error: %v", err)
		return nil
	}
	return s3.NewFromConfig(cfg)
}

func NewAwsService() AwsService {
	return &awsService{
		S3Client: createS3(),
	}
}

func (service awsService) Upload(file *multipart.FileHeader) (*manager.UploadOutput, error) {
	uploader := manager.NewUploader(service.S3Client)
	f, err := file.Open()
	if err != nil {
		return nil, err
	}

	result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String("my-first-bucket-kit"),
		Key:    aws.String(file.Filename),
		Body:   f,
	})

	if err != nil {
		return nil, err
	}
	return result, nil
}
