resource "aws_cloudwatch_log_group" "api_gw" {
  name = "/aws/api_gw/${aws_apigatewayv2_api.home_network_proxy.name}"

  retention_in_days = 30
}
