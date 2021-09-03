# get-all-ip-addresses
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

  route_key = "GET /ip/all"
  target    = "integrations/${aws_apigatewayv2_integration.get_all_ip_addresses.id}"

  authorization_type = "JWT"
  authorizer_id      = aws_apigatewayv2_authorizer.home_network_proxy.id
}

resource "aws_lambda_permission" "get_all_ip_addresses_permission" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.get_all_ip_addresses.function_name
  principal     = "apigateway.amazonaws.com"

  source_arn = "${aws_apigatewayv2_api.home_network_proxy.execution_arn}/*/*"
}

# upsert-ip-address
data "archive_file" "lambda_upsert_ip_address_zip" {
  type = "zip"

  source_dir  = "${path.module}/lambdas/upsert-ip-address"
  output_path = "${path.module}/upsert-ip-address.zip"
}

resource "aws_s3_bucket_object" "lambda_upsert_ip_address" {
  bucket = aws_s3_bucket.lambda_bucket.id

  key    = "upsert-ip-address.zip"
  source = data.archive_file.lambda_upsert_ip_address_zip.output_path

  etag = filemd5(data.archive_file.lambda_upsert_ip_address_zip.output_path)
}

resource "aws_lambda_function" "upsert_ip_address" {
  function_name = "Upsert_IP_Address"

  s3_bucket = aws_s3_bucket.lambda_bucket.id
  s3_key    = aws_s3_bucket_object.lambda_upsert_ip_address.key

  runtime = "nodejs14.x"
  handler = "index.handler"

  source_code_hash = data.archive_file.lambda_upsert_ip_address_zip.output_base64sha256

  role = aws_iam_role.lambda_exec.arn

  environment {
    variables = {
      TABLE_NAME                          = var.table_name
      AWS_NODEJS_CONNECTION_REUSE_ENABLED = "1"
    }
  }
}

resource "aws_cloudwatch_log_group" "upsert_ip_address" {
  name = "/aws/lambda/${aws_lambda_function.upsert_ip_address.function_name}"

  retention_in_days = 30
}

resource "aws_apigatewayv2_integration" "upsert_ip_address" {
  api_id = aws_apigatewayv2_api.home_network_proxy.id

  integration_uri    = aws_lambda_function.upsert_ip_address.invoke_arn
  integration_type   = "AWS_PROXY"
  integration_method = "POST"
}

resource "aws_apigatewayv2_route" "upsert_ip_address" {
  api_id = aws_apigatewayv2_api.home_network_proxy.id

  route_key = "POST /ip"
  target    = "integrations/${aws_apigatewayv2_integration.upsert_ip_address.id}"

  authorization_type = "JWT"
  authorizer_id      = aws_apigatewayv2_authorizer.home_network_proxy.id
}

resource "aws_lambda_permission" "upsert_ip_address_permission" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.upsert_ip_address.function_name
  principal     = "apigateway.amazonaws.com"

  source_arn = "${aws_apigatewayv2_api.home_network_proxy.execution_arn}/*/*"
}

# delete-ip-address
data "archive_file" "lambda_delete_ip_address_zip" {
  type = "zip"

  source_dir  = "${path.module}/lambdas/delete-ip-address"
  output_path = "${path.module}/delete-ip-address.zip"
}

resource "aws_s3_bucket_object" "lambda_delete_ip_address" {
  bucket = aws_s3_bucket.lambda_bucket.id

  key    = "delete-ip-address.zip"
  source = data.archive_file.lambda_delete_ip_address_zip.output_path

  etag = filemd5(data.archive_file.lambda_delete_ip_address_zip.output_path)
}

resource "aws_lambda_function" "delete_ip_address" {
  function_name = "Delete_IP_Address"

  s3_bucket = aws_s3_bucket.lambda_bucket.id
  s3_key    = aws_s3_bucket_object.lambda_delete_ip_address.key

  runtime = "nodejs14.x"
  handler = "index.handler"

  source_code_hash = data.archive_file.lambda_delete_ip_address_zip.output_base64sha256

  role = aws_iam_role.lambda_exec.arn

  environment {
    variables = {
      TABLE_NAME                          = var.table_name
      AWS_NODEJS_CONNECTION_REUSE_ENABLED = "1"
    }
  }
}

resource "aws_cloudwatch_log_group" "delete_ip_address" {
  name = "/aws/lambda/${aws_lambda_function.delete_ip_address.function_name}"

  retention_in_days = 30
}

resource "aws_apigatewayv2_integration" "delete_ip_address" {
  api_id = aws_apigatewayv2_api.home_network_proxy.id

  integration_uri    = aws_lambda_function.delete_ip_address.invoke_arn
  integration_type   = "AWS_PROXY"
  integration_method = "POST"
}

resource "aws_apigatewayv2_route" "delete_ip_address" {
  api_id = aws_apigatewayv2_api.home_network_proxy.id

  route_key = "DELETE /ip"
  target    = "integrations/${aws_apigatewayv2_integration.delete_ip_address.id}"

  authorization_type = "JWT"
  authorizer_id      = aws_apigatewayv2_authorizer.home_network_proxy.id
}

resource "aws_lambda_permission" "delete_ip_address_permission" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.delete_ip_address.function_name
  principal     = "apigateway.amazonaws.com"

  source_arn = "${aws_apigatewayv2_api.home_network_proxy.execution_arn}/*/*"
}

# api-key-auth
data "archive_file" "lambda_api_key_auth_zip" {
  type = "zip"

  source_dir  = "${path.module}/lambdas/api-key-auth"
  output_path = "${path.module}/api-key-auth.zip"
}

resource "aws_s3_bucket_object" "lambda_api_key_auth" {
  bucket = aws_s3_bucket.lambda_bucket.id

  key    = "api-key-auth.zip"
  source = data.archive_file.lambda_api_key_auth_zip.output_path

  etag = filemd5(data.archive_file.lambda_api_key_auth_zip.output_path)
}

resource "aws_lambda_function" "api_key_auth" {
  function_name = "API_Key_Auth"

  s3_bucket = aws_s3_bucket.lambda_bucket.id
  s3_key    = aws_s3_bucket_object.lambda_api_key_auth.key

  runtime = "nodejs14.x"
  handler = "index.handler"

  source_code_hash = data.archive_file.lambda_api_key_auth_zip.output_base64sha256

  role = aws_iam_role.lambda_exec.arn

  environment {
    variables = {
      TABLE_NAME                          = var.table_name_auth
      AWS_NODEJS_CONNECTION_REUSE_ENABLED = "1"
    }
  }
}

resource "aws_cloudwatch_log_group" "api_key_auth" {
  name = "/aws/lambda/${aws_lambda_function.api_key_auth.function_name}"

  retention_in_days = 30
}

resource "aws_apigatewayv2_integration" "api_key_auth" {
  api_id = aws_apigatewayv2_api.home_network_proxy.id

  integration_uri    = aws_lambda_function.api_key_auth.invoke_arn
  integration_type   = "AWS_PROXY"
  integration_method = "POST"
}

resource "aws_apigatewayv2_route" "api_key_auth" {
  api_id = aws_apigatewayv2_api.home_network_proxy.id

  route_key = "POST /auth"
  target    = "integrations/${aws_apigatewayv2_integration.api_key_auth.id}"

  authorization_type = "CUSTOM"
  authorizer_id      = aws_apigatewayv2_authorizer.home_network_proxy-lambda.id
}

resource "aws_lambda_permission" "api_key_auth_permission" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.api_key_auth.function_name
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
