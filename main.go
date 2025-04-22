package main

import (
	"context"
	"log"

	"github.com/8soat-grupo35/grupo35-video-notification/internal/adapters/email"
	"github.com/8soat-grupo35/grupo35-video-notification/internal/adapters/handler"
	"github.com/8soat-grupo35/grupo35-video-notification/internal/adapters/storage"
	"github.com/8soat-grupo35/grupo35-video-notification/internal/usecase"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
)

func handleRequest(ctx context.Context, sqsEvent events.SQSEvent) error {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	sesClient := sesv2.NewFromConfig(cfg)
	s3Client := s3.NewFromConfig(cfg)

	emailService := email.NewSESService(sesClient)
	presigner := storage.NewPresignS3(s3Client)

	sendEmailUC := usecase.NewSendEmailUseCase(emailService)
	sqsHandler := handler.NewSQSHandler(sendEmailUC, presigner.GeneratePresignedURL)

	sqsHandler.Handle(ctx, sqsEvent)

	log.Printf("Lambda processado com sucesso")

	return nil
}

func main() {
	lambda.Start(handleRequest)
}
