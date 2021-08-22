data "archive_file" "lambda_get_ip_zip" {
  type = "zip"

  source_dir  = "${path.module}/lambdas/getIP"
  output_path = "${path.module}/getIP.zip"
}

resource "aws_s3_bucket_object" "lambda_get_ip" {
  bucket = aws_s3_bucket.lambda_bucket.id

  key    = "getIP.zip"
  source = data.archive_file.lambda_get_ip_zip.output_path

  etag = filemd5(data.archive_file.lambda_get_ip_zip.output_path)
}

resource "aws_lambda_function" "get_ip" {
  function_name = "Get_IP"

  s3_bucket = aws_s3_bucket.lambda_bucket.id
  s3_key    = aws_s3_bucket_object.lambda_get_ip.key

  runtime = "nodejs14.x"
  handler = "index.handler"

  source_code_hash = data.archive_file.lambda_get_ip_zip.output_base64sha256

  role = aws_iam_role.lambda_exec.arn

  environment {
    variables = {
      TABLE_NAME                          = var.table_name
      AWS_NODEJS_CONNECTION_REUSE_ENABLED = "1"
    }
  }
}

resource "aws_cloudwatch_log_group" "get_ip" {
  name = "/aws/lambda/${aws_lambda_function.get_ip.function_name}"

  retention_in_days = 30
}

resource "aws_apigatewayv2_integration" "get_ip" {
  api_id = aws_apigatewayv2_api.home_network_proxy.id

  integration_uri    = aws_lambda_function.get_ip.invoke_arn
  integration_type   = "AWS_PROXY"
  integration_method = "POST"
}

resource "aws_apigatewayv2_route" "get_ip" {
  api_id = aws_apigatewayv2_api.home_network_proxy.id

  route_key = "GET /get_ip"
  target    = "integrations/${aws_apigatewayv2_integration.get_ip.id}"
}

resource "aws_lambda_permission" "get_ip_permission" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.get_ip.function_name
  principal     = "apigateway.amazonaws.com"

  source_arn = "${aws_apigatewayv2_api.home_network_proxy.execution_arn}/*/*"
}
