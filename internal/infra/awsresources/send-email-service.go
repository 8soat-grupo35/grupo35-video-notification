package awsresources

import (
	"context"
	"fmt"
	"log"

	"github.com/8soat-grupo35/grupo35-video-notification/internal/domain"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
)

type SESService struct {
	client *SESClient
}

func NewSESService(client *SESClient) *SESService {
	return &SESService{client: client}
}

func (s *SESService) SendEmail(ctx context.Context, from string, email *domain.Email) error {
	// Envio do e-mail usando o SES
	input := &sesv2.SendEmailInput{
		FromEmailAddress: aws.String(from),
		Destination: &types.Destination{
			ToAddresses: []string{email.GetToEmail()},
		},
		Content: &types.EmailContent{
			Simple: &types.Message{
				Subject: &types.Content{
					Data: aws.String(email.GetSubject()),
				},
				Body: &types.Body{
					Html: &types.Content{
						Data: aws.String(email.Template()),
					},
				},
			},
		},
	}

	_, err := s.client.Client.SendEmail(ctx, input)
	if err != nil {
		log.Printf("Erro ao enviar e-mail: %v", err)
		return fmt.Errorf("falha ao enviar e-mail: %w", err)
	}

	return nil
}
