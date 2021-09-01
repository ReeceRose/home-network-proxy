
resource "aws_amplify_app" "web" {
  name       = "Home Network Proxy Web UI"
  repository = "https://github.com/reecerose/home-network-proxy"

  auto_branch_creation_config {
    # Enable auto build for the created branch.
    enable_auto_build = true
  }

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
      artifacts:
        baseDirectory: build
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

resource "aws_amplify_branch" "main" {
  app_id      = aws_amplify_app.web.id
  branch_name = "main"

  framework = "React"
  stage     = "PRODUCTION"
}
