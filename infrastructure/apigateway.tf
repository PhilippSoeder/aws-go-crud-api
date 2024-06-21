resource "aws_api_gateway_rest_api" "notes" {
  name = "Notes"
  endpoint_configuration {
    types = ["REGIONAL"]
  }
}

resource "aws_api_gateway_resource" "notes" {
  rest_api_id = aws_api_gateway_rest_api.notes.id
  parent_id   = aws_api_gateway_rest_api.notes.root_resource_id
  path_part   = "notes"
}

resource "aws_api_gateway_resource" "note" {
  rest_api_id = aws_api_gateway_rest_api.notes.id
  parent_id   = aws_api_gateway_resource.notes.id
  path_part   = "{note-id}"
}

resource "aws_api_gateway_method" "getNotes" {
  rest_api_id   = aws_api_gateway_rest_api.notes.id
  resource_id   = aws_api_gateway_resource.notes.id
  http_method   = "GET"
  authorization = "NONE"
}

resource "aws_api_gateway_method" "getNote" {
  rest_api_id   = aws_api_gateway_rest_api.notes.id
  resource_id   = aws_api_gateway_resource.note.id
  http_method   = "GET"
  authorization = "NONE"
  request_parameters = {
    "method.request.path.note-id" = true
  }
}

resource "aws_api_gateway_method" "postNote" {
  rest_api_id   = aws_api_gateway_rest_api.notes.id
  resource_id   = aws_api_gateway_resource.note.id
  http_method   = "POST"
  authorization = "NONE"
  request_parameters = {
    "method.request.path.note-id" = true
  }
}

resource "aws_api_gateway_method" "putNote" {
  rest_api_id   = aws_api_gateway_rest_api.notes.id
  resource_id   = aws_api_gateway_resource.note.id
  http_method   = "PUT"
  authorization = "NONE"
  request_parameters = {
    "method.request.path.note-id" = true
  }
}

resource "aws_api_gateway_method" "deleteNote" {
  rest_api_id   = aws_api_gateway_rest_api.notes.id
  resource_id   = aws_api_gateway_resource.note.id
  http_method   = "DELETE"
  authorization = "NONE"
  request_parameters = {
    "method.request.path.note-id" = true
  }
}

resource "aws_api_gateway_integration" "getNotes" {
  rest_api_id             = aws_api_gateway_rest_api.notes.id
  resource_id             = aws_api_gateway_resource.notes.id
  http_method             = aws_api_gateway_method.getNotes.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.lambda_function.invoke_arn
  depends_on              = [aws_lambda_permission.allow_api_gateway]
}

resource "aws_api_gateway_integration" "getNote" {
  rest_api_id = aws_api_gateway_rest_api.notes.id
  resource_id = aws_api_gateway_resource.note.id
  http_method = aws_api_gateway_method.getNote.http_method

  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.lambda_function.invoke_arn

  depends_on = [aws_lambda_permission.allow_api_gateway]
}

resource "aws_api_gateway_integration" "postNote" {
  rest_api_id             = aws_api_gateway_rest_api.notes.id
  resource_id             = aws_api_gateway_resource.note.id
  http_method             = aws_api_gateway_method.postNote.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.lambda_function.invoke_arn
  depends_on              = [aws_lambda_permission.allow_api_gateway]
}

resource "aws_api_gateway_integration" "putNote" {
  rest_api_id             = aws_api_gateway_rest_api.notes.id
  resource_id             = aws_api_gateway_resource.note.id
  http_method             = aws_api_gateway_method.putNote.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.lambda_function.invoke_arn
  depends_on              = [aws_lambda_permission.allow_api_gateway]
}

resource "aws_api_gateway_integration" "deleteNote" {
  rest_api_id             = aws_api_gateway_rest_api.notes.id
  resource_id             = aws_api_gateway_resource.note.id
  http_method             = aws_api_gateway_method.deleteNote.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.lambda_function.invoke_arn
  depends_on              = [aws_lambda_permission.allow_api_gateway]
}

resource "aws_api_gateway_deployment" "notes-deployment" {
  rest_api_id = aws_api_gateway_rest_api.notes.id

  triggers = {
    redeployment = sha256(jsonencode(aws_api_gateway_rest_api.notes.body))
  }

  lifecycle {
    create_before_destroy = true
  }

  depends_on = [
    aws_api_gateway_method.getNote,
    aws_api_gateway_method.postNote,
    aws_api_gateway_method.putNote,
    aws_api_gateway_method.deleteNote,
    aws_api_gateway_integration.getNote,
    aws_api_gateway_integration.postNote,
    aws_api_gateway_integration.putNote,
    aws_api_gateway_integration.deleteNote
  ]
}

resource "aws_api_gateway_stage" "dev" {
  deployment_id = aws_api_gateway_deployment.notes-deployment.id
  rest_api_id   = aws_api_gateway_rest_api.notes.id
  stage_name    = var.environment
}
