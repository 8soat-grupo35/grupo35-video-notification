# grupo35-video-notification

## Compile o código Go
Compile o código para a arquitetura Linux e AMD64 (ou ARM64, dependendo da configuração do Lambda) porque o runtime do AWS Lambda usa um ambiente Linux.

```ssh
GOOS=linux GOARCH=arm64 go build -tags lambda.norpc -o bootstrap main.go
```

## Empacote o binário
Coloque o arquivo executável em um arquivo zip
```ssh
zip myFunction.zip bootstrap
```