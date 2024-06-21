output "lambda_arn" {
  value       = aws_lambda_function.lambda_function.arn
  description = "Amazon Ressource Name for Lambda."
}

output "example_api_endpoint_curl" {
  value       = "curl -X GET ${aws_api_gateway_deployment.notes-deployment.invoke_url}${aws_api_gateway_stage.dev.stage_name}${aws_api_gateway_resource.notes.path}"
  description = "Use this command to test the API."
}
