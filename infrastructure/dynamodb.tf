resource "aws_dynamodb_table" "notes" {
  name         = var.dynamodb_table_name
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "ID"
  attribute {
    name = "ID"
    type = "S"
  }
}
