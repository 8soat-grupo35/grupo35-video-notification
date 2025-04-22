package usecase

import (
	"context"

	domain "github.com/8soat-grupo35/grupo35-video-notification/internal/domain/entity"
	"github.com/8soat-grupo35/grupo35-video-notification/internal/domain/service"
)

type SendEmailUseCase struct {
	EmailService service.EmailService
}

func NewSendEmailUseCase(emailService service.EmailService) *SendEmailUseCase {
	return &SendEmailUseCase{EmailService: emailService}
}

func (uc *SendEmailUseCase) Execute(ctx context.Context, from string, email *domain.Email) error {
	return uc.EmailService.SendEmail(ctx, from, email)
}
