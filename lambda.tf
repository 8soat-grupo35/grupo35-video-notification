provider "aws" {
  region = "us-east-1"  # Substitua pela sua região
}

resource "aws_lambda_function" "my_lambda_function" {
  function_name = "my_lambda_function"
  role          = aws_iam_role.lambda_exec_role.arn
  handler       = "main.handler"
  runtime       = "provided.al2"  # Substitua com o tempo de execução desejado (Go neste caso)

  # Defina as variáveis de ambiente
  environment {
    variables = {
      MY_CUSTOM_VAR         = "some-value"
    }
  }
}

# Política IAM para conceder permissões de execução à função Lambda
resource "aws_iam_role" "lambda_exec_role" {
  name = "lambda_exec_role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action    = "sts:AssumeRole"
        Effect    = "Allow"
        Principal = {
          Service = "lambda.amazonaws.com"
        }
      }
    ]
  })
}

# Permissões para a função Lambda (adapte conforme necessário)
resource "aws_iam_policy" "lambda_policy" {
  name        = "lambda_policy"
  description = "Permissões de execução para Lambda"
  policy      = jsonencode({
    Version   = "2012-10-17"
    Statement = [
      {
        Action   = "logs:*"
        Resource = "*"
        Effect   = "Allow"
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "lambda_policy_attachment" {
  role       = aws_iam_role.lambda_exec_role.name
  policy_arn = aws_iam_policy.lambda_policy.arn
}
