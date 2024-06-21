locals {
  lambda_name = "notes-api-${var.environment}"
}

data "archive_file" "lambda" {
  type        = "zip"
  source_file = "../bin/bootstrap"
  output_path = "lambda.zip"
}

resource "aws_lambda_permission" "allow_api_gateway" {
  statement_id  = "AllowAPIGatewayInvoke"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.lambda_function.function_name
  principal     = "apigateway.amazonaws.com"
  source_arn    = "${aws_api_gateway_rest_api.notes.execution_arn}/*/*${aws_api_gateway_resource.notes.path}*"
}

resource "aws_lambda_function" "lambda_function" {
  filename         = data.archive_file.lambda.output_path
  source_code_hash = data.archive_file.lambda.output_base64sha256
  function_name    = local.lambda_name
  role             = aws_iam_role.lambda_role.arn
  handler          = "bootstrap"
  architectures    = ["arm64"]
  memory_size      = var.lambda_memory_size_megabytes
  runtime          = "provided.al2023"
  timeout          = var.lambda_timeout_seconds
  environment {
    variables = {
      LOG_LEVEL           = var.log_level
      DYNAMODB_REGION     = var.aws_region
      DYNAMODB_TABLE_NAME = var.dynamodb_table_name
    }
  }
}
