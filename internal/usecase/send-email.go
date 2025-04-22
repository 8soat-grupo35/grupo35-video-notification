package usecase

import (
	"context"

	"github.com/8soat-grupo35/grupo35-video-notification/internal/domain"
)

type EmailService interface {
	SendEmail(ctx context.Context, from string, email *domain.Email) error
}

type SendEmailUseCase struct {
	emailService EmailService
}

func NewSendEmailUseCase(emailService EmailService) *SendEmailUseCase {
	return &SendEmailUseCase{emailService: emailService}
}

func (uc *SendEmailUseCase) Execute(ctx context.Context, from string, email *domain.Email) error {
	return uc.emailService.SendEmail(ctx, from, email)
}
