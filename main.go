package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
)

const (
	FROM = "grupo35soat8@gmail.com"
)

type SESService struct {
	client *sesv2.Client
}

type User struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

type Message struct {
	User    User   `json:"user"`
	ZipPath string `json:"zip_path"`
}

func NewSESService(client *sesv2.Client) *SESService {
	return &SESService{client: client}
}

func (s *SESService) SendEmail(ctx context.Context, from, to, content string) error {
	// Envio do e-mail usando o SES
	input := &sesv2.SendEmailInput{
		FromEmailAddress: aws.String(from),
		Destination: &types.Destination{
			ToAddresses: []string{to},
		},
		Content: &types.EmailContent{
			Simple: &types.Message{
				Subject: &types.Content{
					Data: aws.String("Obaa seu arquivo esta disponivel!"),
				},
				Body: &types.Body{
					Html: &types.Content{
						Data: aws.String(content),
					},
				},
			},
		},
	}

	_, err := s.client.SendEmail(ctx, input)
	if err != nil {
		log.Printf("Erro ao enviar e-mail: %v", err)
		return fmt.Errorf("falha ao enviar e-mail: %w", err)
	}

	return nil
}

func NewSESClient() *sesv2.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	return sesv2.NewFromConfig(cfg)
}

type EmailService interface {
	SendEmail(ctx context.Context, from, to, content string) error
}

type SendEmailUseCase struct {
	emailService EmailService
}

func NewSendEmailUseCase(emailService EmailService) *SendEmailUseCase {
	return &SendEmailUseCase{emailService: emailService}
}

func (uc *SendEmailUseCase) Execute(ctx context.Context, to, content string) error {
	// Lógica do caso de uso de envio de e-mail
	// Aqui você pode processar o conteúdo do e-mail e enviar

	// Enviar e-mail utilizando o serviço injetado
	return uc.emailService.SendEmail(ctx, FROM, to, content)
}

func init() {
	// Initialize the S3 client outside of the handler, during the init phase
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	fmt.Println("cfg: ", cfg)
}

func handleRequest(ctx context.Context, sqsEvent events.SQSEvent) error {
	// Configurações do serviço de envio de e-mail
	sesClient := NewSESClient()
	emailService := NewSESService(sesClient)
	sendEmailUseCase := NewSendEmailUseCase(emailService)

	// Processando mensagens da fila SQS
	for _, message := range sqsEvent.Records {
		// Aqui você pode realizar o parse da mensagem da fila conforme necessário
		var messageBody Message

		// Fazendo o Unmarshal do JSON para a struct User
		errr := json.Unmarshal([]byte(message.Body), &messageBody)
		if errr != nil {
			log.Fatalf("Erro ao fazer unmarshal: %v", errr)
		}

		err := sendEmailUseCase.Execute(ctx, messageBody.User.Email, templateEmail(messageBody.ZipPath))
		if err != nil {
			log.Printf("Erro ao enviar e-mail: %v", err)
			return err
		}

		log.Printf("Mensagem processada: %s", messageBody)
	}

	return nil
}

func templateEmail(downloadLink string) string {
	return fmt.Sprintf(`<html>
			<body>
				<p>Olá,</p>
				<p>Seu arquivo ficou pronto, você pode baixar o arquivo clicando no link abaixo:</p>
				<p><a href="%s" download>Baixar Arquivo</a></p>
				<p>Atenciosamente,<br>Equipe</p>
			</body>
		</html>`,
		downloadLink)
}

func main() {
	lambda.Start(handleRequest)
}
