package awsresources

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
)

type SESClient struct {
	Client *sesv2.Client
}

func NewSESClient(cfg aws.Config) SESClient {
	return SESClient{
		Client: sesv2.NewFromConfig(cfg),
	}
}
