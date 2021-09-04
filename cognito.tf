resource "aws_cognito_identity_pool" "pool" {
  identity_pool_name               = "home-network-proxy-pool"
  allow_unauthenticated_identities = false

}

resource "aws_cognito_user_pool" "pool" {
  name = "home-network-proxy-pool"
}

resource "aws_cognito_identity_pool_roles_attachment" "main" {
  identity_pool_id = aws_cognito_identity_pool.pool.id
  roles = {
    "authenticated" = aws_iam_role.lambda_exec.arn
  }
}

resource "aws_cognito_user_pool_domain" "main" {
  domain       = "home-network-proxy"
  user_pool_id = aws_cognito_user_pool.pool.id
}
