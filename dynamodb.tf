resource "aws_dynamodb_table" "home-network-proxy" {
  name         = var.table_name
  billing_mode = var.table_billing_mode
  hash_key     = "Id"

  attribute {
    name = "Id"
    type = "S"
  }

  tags = {
    Environment = "${var.environment}"
  }
}
