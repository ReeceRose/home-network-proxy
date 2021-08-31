# Output value definitions

output "lambda_bucket_name" {
  description = "Name of the S3 bucket used to store function code."

  value = aws_s3_bucket.lambda_bucket.id
}

output "get_all_ip_addresses_function_name" {
  description = "Name of the Get all IP addresses Lambda function."

  value = aws_lambda_function.get_all_ip_addresses.function_name
}

output "upsert_ip_address_function_name" {
  description = "Name of the Upsert IP Address Lambda function."

  value = aws_lambda_function.upsert_ip_address.function_name
}

output "health_function_name" {
  description = "Name of the health Lambda function."

  value = aws_lambda_function.health.function_name
}

output "base_url" {
  description = "Base URL for API Gateway."

  value = aws_apigatewayv2_stage.production.invoke_url
}

