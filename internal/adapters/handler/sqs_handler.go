package handler

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	domain "github.com/8soat-grupo35/grupo35-video-notification/internal/domain/entity"
	"github.com/8soat-grupo35/grupo35-video-notification/internal/usecase"
	"github.com/aws/aws-lambda-go/events"
)

// 24 hours to expire the presigned URL
const EXPIRATION = 24 * time.Hour

type User struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

type Message struct {
	User    User   `json:"user"`
	ZipPath string `json:"zip_path"`
}

type SQSHandler struct {
	SendEmailUseCase *usecase.SendEmailUseCase
	PresignURLFunc   func(ctx context.Context, bucket, key string, expiration time.Duration) (string, error)
}

func NewSQSHandler(
	sendEmailUC *usecase.SendEmailUseCase,
	presignFunc func(ctx context.Context, bucket, key string, expiration time.Duration) (string, error),
) *SQSHandler {
	return &SQSHandler{
		SendEmailUseCase: sendEmailUC,
		PresignURLFunc:   presignFunc,
	}
}

func (h *SQSHandler) Handle(ctx context.Context, sqsEvent events.SQSEvent) error {
	bucketName := os.Getenv("BUCKET_NAME")
	fromEmail := os.Getenv("FROM_EMAIL")

	for _, record := range sqsEvent.Records {
		var msg Message
		if err := json.Unmarshal([]byte(record.Body), &msg); err != nil {
			log.Printf("Erro ao fazer unmarshal: %v", err)
			continue
		}

		url, err := h.PresignURLFunc(ctx, bucketName, msg.ZipPath, EXPIRATION)
		if err != nil {
			log.Printf("Erro ao gerar URL: %v", err)
			continue
		}

		email := domain.NewEmail(msg.User.Email, url)

		err = h.SendEmailUseCase.Execute(ctx, fromEmail, email)
		if err != nil {
			log.Printf("Erro ao enviar e-mail: %v", err)
		}

		log.Printf("Email enviado a %s com o link: %s", msg.User.Email, url)
	}

	return nil
}
