package awsresources_test

import (
	"testing"

	"github.com/8soat-grupo35/grupo35-video-notification/internal/infra/awsresources"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/stretchr/testify/assert"
)

func TestNewSESClient(t *testing.T) {
	// Mock AWS Config
	mockConfig := aws.Config{}

	// Call the function
	sesClient := awsresources.NewSESClient(mockConfig)

	// Assertions
	assert.NotNil(t, sesClient, "SESClient should not be nil")
	assert.NotNil(t, sesClient.Client, "SESClient.Client should not be nil")
	assert.IsType(t, &sesv2.Client{}, sesClient.Client, "SESClient.Client should be of type *sesv2.Client")
}
