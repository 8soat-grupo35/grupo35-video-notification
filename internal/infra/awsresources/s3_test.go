package awsresources_test

import (
	"testing"

	"github.com/8soat-grupo35/grupo35-video-notification/internal/infra/awsresources"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/stretchr/testify/assert"
)

func TestNewS3Client(t *testing.T) {
	// Mock AWS Config
	mockConfig := aws.Config{}

	// Call the function
	s3Client := awsresources.NewS3Client(mockConfig)

	// Assertions
	assert.NotNil(t, s3Client, "S3Client should not be nil")
	assert.NotNil(t, s3Client.Client, "S3Client.Client should not be nil")
	assert.IsType(t, &s3.Client{}, s3Client.Client, "S3Client.Client should be of type *s3.Client")
}
