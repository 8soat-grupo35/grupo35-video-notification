package domain

import "fmt"

type Email struct {
	to           string
	subject      string
	downloadLink string
}

func NewEmail(to, downloadLink string) *Email {
	return &Email{
		to:           to,
		downloadLink: downloadLink,
		subject:      "Obaa seu arquivo esta disponivel!",
	}
}

func (e *Email) GetToEmail() string {
	return e.to
}

func (e *Email) GetSubject() string {
	return e.subject
}

func (e *Email) Template() string {
	return fmt.Sprintf(`<html>
			<body>
				<p>Olá,</p>
				<p>Seu arquivo ficou pronto, você pode baixar o arquivo clicando no link abaixo:</p>
				<p><a href="%s" download>Baixar Arquivo</a></p>
				<p>Atenciosamente,<br>Equipe</p>
			</body>
		</html>`,
		e.downloadLink)
}
