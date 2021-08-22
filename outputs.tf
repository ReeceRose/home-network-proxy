# Output value definitions

output "lambda_bucket_name" {
  description = "Name of the S3 bucket used to store function code."

  value = aws_s3_bucket.lambda_bucket.id
}

output "get_all_ip_function_name" {
  description = "Name of the GetAllIP Lambda function."

  value = aws_lambda_function.get_all_ip.function_name
}

output "get_ip_function_name" {
  description = "Name of the GetIP Lambda function."

  value = aws_lambda_function.get_ip.function_name
}

output "health_function_name" {
  description = "Name of the Health Lambda function."

  value = aws_lambda_function.health.function_name
}

output "base_url" {
  description = "Base URL for API Gateway."

  value = aws_apigatewayv2_stage.production.invoke_url
}

