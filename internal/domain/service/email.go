package service

import (
	"context"

	domain "github.com/8soat-grupo35/grupo35-video-notification/internal/domain/entity"
)

type EmailService interface {
	SendEmail(ctx context.Context, from string, email *domain.Email) error
}
