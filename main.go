package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/8soat-grupo35/grupo35-video-notification/internal/domain"
	"github.com/8soat-grupo35/grupo35-video-notification/internal/infra/awsresources"
	"github.com/8soat-grupo35/grupo35-video-notification/internal/usecase"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const (
	EXPIRATION = 60 * time.Minute
)

type User struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

type Message struct {
	User    User   `json:"user"`
	ZipPath string `json:"zip_path"`
}

func handleRequest(ctx context.Context, sqsEvent events.SQSEvent) error {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// Configurações do serviço de envio de e-mail
	s3Client := awsresources.NewS3Client(cfg)
	sesClient := awsresources.NewSESClient(cfg)
	emailService := awsresources.NewSESService(&sesClient)

	sendEmailUseCase := usecase.NewSendEmailUseCase(emailService)

	bucketName := os.Getenv("BUCKET_NAME")
	fromEmail := os.Getenv("FROM_EMAIL")

	// Processando mensagens da fila SQS
	for _, message := range sqsEvent.Records {
		// Aqui você pode realizar o parse da mensagem da fila conforme necessário
		var messageBody Message

		// Fazendo o Unmarshal do JSON para a struct User
		errr := json.Unmarshal([]byte(message.Body), &messageBody)
		if errr != nil {
			log.Fatalf("Erro ao fazer unmarshal: %v", errr)
		}

		// Gerar a URL assinada
		presignClient := s3.NewPresignClient(s3Client.Client)

		req, errorpre := presignClient.PresignGetObject(ctx, &s3.GetObjectInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(messageBody.ZipPath),
		}, func(o *s3.PresignOptions) {
			o.Expires = EXPIRATION
		})
		if errorpre != nil {
			log.Printf("Erro ao gerar URL assinada: %v", errorpre)
		}

		log.Printf("URL assinada: %s", req.URL)

		email := domain.NewEmail(messageBody.User.Email, req.URL)

		err := sendEmailUseCase.Execute(ctx, fromEmail, email)
		if err != nil {
			log.Printf("Erro ao enviar e-mail: %v", err)
			return err
		}

		log.Printf("Mensagem processada: %s", messageBody.ZipPath)
	}

	return nil
}

func main() {
	lambda.Start(handleRequest)
}
