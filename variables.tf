# Input variable definitions

variable "aws_region" {
  description = "AWS region for all resources."

  type    = string
  default = "us-east-1"
}

variable "table_name" {
  description = "Dynamodb table name (space is not allowed)"
  default     = "home-network-proxy"
}

variable "table_billing_mode" {
  description = "Controls how you are charged for read and write throughput and how you manage capacity."
  default     = "PAY_PER_REQUEST"
}

variable "environment" {
  description = "Name of environment"
  default     = "production"
}
