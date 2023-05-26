output "arn" {
  description = "ARN of the IAM policy"
  value       = try(aws_iam_policy.main[0].arn, null)
}

output "name" {
  description = "Name of the IAM policy"
  value       = try(aws_iam_policy.main[0].name, null)
}

output "id" {
  description = "ID of the IAM policy"
  value       = try(aws_iam_policy.main[0].id, null)
}

output "json" {
  description = "Rendered JSON of the policy."
  value       = data.aws_iam_policy_document.combined.json
}
