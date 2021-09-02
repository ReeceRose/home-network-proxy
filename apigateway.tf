resource "aws_apigatewayv2_api" "home_network_proxy" {
  name          = "home_network_proxy_api"
  protocol_type = "HTTP"

  cors_configuration {
    allow_headers = ["*"]
    allow_methods = ["*"]
    allow_origins = ["*"]
  }
}

resource "aws_apigatewayv2_authorizer" "home_network_proxy" {
  api_id           = aws_apigatewayv2_api.home_network_proxy.id
  authorizer_type  = "JWT"
  identity_sources = ["$request.header.Authorization"]

  name = "jwt-auth"

  jwt_configuration {
    audience = ["5q42kpv9mh7spcjg63l88sqpah"]
    issuer   = "https://${aws_cognito_user_pool.pool.endpoint}"
  }
}

resource "aws_apigatewayv2_stage" "production" {
  api_id = aws_apigatewayv2_api.home_network_proxy.id

  name        = "production"
  auto_deploy = true

  access_log_settings {
    destination_arn = aws_cloudwatch_log_group.api_gw.arn

    format = jsonencode({
      requestId               = "$context.requestId"
      sourceIp                = "$context.identity.sourceIp"
      requestTime             = "$context.requestTime"
      protocol                = "$context.protocol"
      httpMethod              = "$context.httpMethod"
      resourcePath            = "$context.resourcePath"
      routeKey                = "$context.routeKey"
      status                  = "$context.status"
      responseLength          = "$context.responseLength"
      integrationErrorMessage = "$context.integrationErrorMessage"
      }
    )
  }
}
