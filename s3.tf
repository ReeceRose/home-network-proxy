resource "aws_s3_bucket" "lambda_bucket" {
  bucket = "home-network-proxy-bucket"

  acl           = "private"
  force_destroy = true
}
