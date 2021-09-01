
resource "aws_amplify_app" "web" {
  name       = "Home Network Proxy Web UI"
  repository = "https://github.com/reecerose/home-network-proxy"

  # The default build_spec added by the Amplify Console for React.
  build_spec = <<-EOT
    version: 0.1
    frontend:
      phases:
        preBuild:
          commands:
            - cd web
            - yarn install
        build:
          commands:
            - yarn build
            - yarn export
      artifacts:
        baseDirectory: ./web/out
        files:
          - '**/*'
      cache:
        paths:
          - node_modules/**/*
  EOT

  # The default rewrites and redirects added by the Amplify Console.
  custom_rule {
    source = "/<*>"
    status = "404"
    target = "/index.html"
  }

  access_token = var.access_token

  environment_variables = {
    ENV = "production"
  }
}

resource "aws_amplify_branch" "production" {
  app_id      = aws_amplify_app.web.id
  branch_name = "main"

  display_name = "production"

  framework = "React"
  stage     = "PRODUCTION"
}
