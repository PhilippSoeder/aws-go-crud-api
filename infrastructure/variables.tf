variable "environment" {
  type        = string
  default     = "undefined"
  description = "Name of the environment. This variable will be used in the AWS tags."
}

variable "aws_region" {
  type        = string
  default     = "eu-central-1"
  description = "AWS region in which the infrastructure will be deployed."
}

variable "lambda_memory_size_megabytes" {
  type        = number
  default     = 128
  description = "Lambda Memory size in MB. This value will also impact CPU performance."
}

variable "lambda_timeout_seconds" {
  type        = number
  default     = 3
  description = "Time until the lambda will stop due to timeout in seconds."
}

variable "log_level" {
  type        = string
  default     = "INFO"
  description = "Log level for the lambda function. Possible values: DEBUG, INFO, WARN, ERROR"
  validation {
    condition     = contains(["DEBUG", "INFO", "WARN", "ERROR"], var.log_level)
    error_message = "Valid values for variable log_level are (DEBUG, INFO, WARN, ERROR)."
  }
}

variable "dynamodb_table_name" {
  type        = string
  default     = "notes"
  description = "Name of the DynamoDB table."
}
