variable "name" {
  default     = "test-iam-policy"
  description = "Role name."
  type        = string
}

variable "account" {
  default     = "arn:aws:iam::1234567890:root"
  description = "Account user ARN."
  type        = string
}
