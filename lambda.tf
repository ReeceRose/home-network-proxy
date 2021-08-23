# get-all-ips
data "archive_file" "lambda_get_all_ip_addresses_zip" {
  type = "zip"

  source_dir  = "${path.module}/lambdas/get-all-ip-addresses"
  output_path = "${path.module}/get-all-ip-addresses.zip"
}

resource "aws_s3_bucket_object" "lambda_get_all_ip_addresses" {
  bucket = aws_s3_bucket.lambda_bucket.id

  key    = "get-all-ip-addresses.zip"
  source = data.archive_file.lambda_get_all_ip_addresses_zip.output_path

  etag = filemd5(data.archive_file.lambda_get_all_ip_addresses_zip.output_path)
}

resource "aws_lambda_function" "get_all_ip_addresses" {
  function_name = "Get_All_IP_Addresses"

  s3_bucket = aws_s3_bucket.lambda_bucket.id
  s3_key    = aws_s3_bucket_object.lambda_get_all_ip_addresses.key

  runtime = "nodejs14.x"
  handler = "index.handler"

  source_code_hash = data.archive_file.lambda_get_all_ip_addresses_zip.output_base64sha256

  role = aws_iam_role.lambda_exec.arn

  environment {
    variables = {
      TABLE_NAME                          = var.table_name
      AWS_NODEJS_CONNECTION_REUSE_ENABLED = "1"
    }
  }
}

resource "aws_cloudwatch_log_group" "get_all_ip_addresses" {
  name = "/aws/lambda/${aws_lambda_function.get_all_ip_addresses.function_name}"

  retention_in_days = 30
}

resource "aws_apigatewayv2_integration" "get_all_ip_addresses" {
  api_id = aws_apigatewayv2_api.home_network_proxy.id

  integration_uri    = aws_lambda_function.get_all_ip_addresses.invoke_arn
  integration_type   = "AWS_PROXY"
  integration_method = "POST"
}

resource "aws_apigatewayv2_route" "get_all_ip_addresses" {
  api_id = aws_apigatewayv2_api.home_network_proxy.id

  route_key = "GET /get_all_ip_addresses"
  target    = "integrations/${aws_apigatewayv2_integration.get_all_ip_addresses.id}"
}

resource "aws_lambda_permission" "get_all_ip_addresses_permission" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.get_all_ip_addresses.function_name
  principal     = "apigateway.amazonaws.com"

  source_arn = "${aws_apigatewayv2_api.home_network_proxy.execution_arn}/*/*"
}

# health
data "archive_file" "lambda_health_zip" {
  type = "zip"

  source_dir  = "${path.module}/lambdas/health"
  output_path = "${path.module}/health.zip"
}

resource "aws_s3_bucket_object" "lambda_health" {
  bucket = aws_s3_bucket.lambda_bucket.id

  key    = "health.zip"
  source = data.archive_file.lambda_health_zip.output_path

  etag = filemd5(data.archive_file.lambda_health_zip.output_path)
}

resource "aws_lambda_function" "health" {
  function_name = "Health"

  s3_bucket = aws_s3_bucket.lambda_bucket.id
  s3_key    = aws_s3_bucket_object.lambda_health.key

  runtime = "nodejs14.x"
  handler = "index.handler"

  source_code_hash = data.archive_file.lambda_health_zip.output_base64sha256

  role = aws_iam_role.lambda_exec.arn
}

resource "aws_cloudwatch_log_group" "health" {
  name = "/aws/lambda/${aws_lambda_function.health.function_name}"

  retention_in_days = 30
}

resource "aws_apigatewayv2_integration" "health" {
  api_id = aws_apigatewayv2_api.home_network_proxy.id

  integration_uri    = aws_lambda_function.health.invoke_arn
  integration_type   = "AWS_PROXY"
  integration_method = "POST"
}

resource "aws_apigatewayv2_route" "health" {
  api_id = aws_apigatewayv2_api.home_network_proxy.id

  route_key = "GET /health"
  target    = "integrations/${aws_apigatewayv2_integration.health.id}"
}

resource "aws_lambda_permission" "health_permission" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.health.function_name
  principal     = "apigateway.amazonaws.com"

  source_arn = "${aws_apigatewayv2_api.home_network_proxy.execution_arn}/*/*"
}
