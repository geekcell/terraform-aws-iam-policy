output "arn" {
  description = "IAM role arn"
  value       = module.s3_policy.arn
}

output "name" {
  description = "IAM role name"
  value       = module.s3_policy.name
}

output "id" {
  description = "ID of the IAM policy"
  value       = module.s3_policy.id
}

output "json" {
  description = "Rendered JSON of the policy."
  value       = module.s3_policy.json
}
