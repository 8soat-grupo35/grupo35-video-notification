package storage

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Presigner interface {
	GeneratePresignedURL(ctx context.Context, bucket, key string, expiration time.Duration) (string, error)
}

type PresignS3 struct {
	Client *s3.PresignClient
}

func NewPresignS3(client *s3.Client) *PresignS3 {
	return &PresignS3{
		Client: s3.NewPresignClient(client),
	}
}

func (p *PresignS3) GeneratePresignedURL(ctx context.Context, bucket, key string, expiration time.Duration) (string, error) {
	req, err := p.Client.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}, func(opts *s3.PresignOptions) {
		opts.Expires = expiration
	})

	if err != nil {
		log.Printf("Erro ao gerar URL assinada: %v", err)
		return "", err
	}

	return req.URL, nil
}
