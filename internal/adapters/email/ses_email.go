package email

import (
	"context"
	"fmt"
	"log"
	"net/smtp"
	"os"

	domain "github.com/8soat-grupo35/grupo35-video-notification/internal/domain/entity"
)

const (
	SMTP     = "smtp.gmail.com"
	SMTPPort = "587"
)

type EmailService struct{}

func NewEmailService() *EmailService {
	return &EmailService{}
}

func (s *EmailService) SendEmail(ctx context.Context, from string, email *domain.Email) error {
	auth := smtp.PlainAuth("", from, os.Getenv("FROM_EMAIL_PASSWORD"), SMTP)

	msg := "Subject: " + email.GetSubject() + "\n" + email.Template()

	err := smtp.SendMail(SMTP+":"+SMTPPort, auth, from, []string{email.GetToEmail()}, []byte(msg))
	if err != nil {
		log.Printf("Erro ao enviar e-mail: %v", err)
		return fmt.Errorf("falha ao enviar e-mail: %w", err)
	}

	return nil
}
