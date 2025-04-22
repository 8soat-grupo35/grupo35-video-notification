package awsresources

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Client struct {
	Client *s3.Client
}

func NewS3Client(cfg aws.Config) S3Client {
	return S3Client{
		Client: s3.NewFromConfig(cfg),
	}
}
