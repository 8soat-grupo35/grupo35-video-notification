package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/8soat-grupo35/grupo35-video-notification/internal/domain"
	"github.com/8soat-grupo35/grupo35-video-notification/internal/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockEmailService struct {
	mock.Mock
}

func (m *MockEmailService) SendEmail(ctx context.Context, from string, email *domain.Email) error {
	args := m.Called(ctx, from, email)
	return args.Error(0)
}

func TestSendEmailUseCase_Execute_Success(t *testing.T) {
	mockEmailService := new(MockEmailService)
	useCase := usecase.NewSendEmailUseCase(mockEmailService)

	ctx := context.Background()
	from := "test@example.com"
	email := domain.NewEmail("recipient@example.com", "www.google.com")

	mockEmailService.On("SendEmail", ctx, from, email).Return(nil)

	err := useCase.Execute(ctx, from, email)

	assert.NoError(t, err)
	mockEmailService.AssertExpectations(t)
}

func TestSendEmailUseCase_Execute_Error(t *testing.T) {
	mockEmailService := new(MockEmailService)
	useCase := usecase.NewSendEmailUseCase(mockEmailService)

	ctx := context.Background()
	from := "test@example.com"
	email := domain.NewEmail("recipient@example.com", "www.google.com")

	mockError := errors.New("failed to send email")
	mockEmailService.On("SendEmail", ctx, from, email).Return(mockError)

	err := useCase.Execute(ctx, from, email)

	assert.Error(t, err)
	assert.Equal(t, mockError, err)
	mockEmailService.AssertExpectations(t)
}
