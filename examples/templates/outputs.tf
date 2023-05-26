output "arn" {
  description = "IAM role arn"
  value       = module.codedeploy_policy.arn
}

output "name" {
  description = "IAM role name"
  value       = module.codedeploy_policy.name
}

output "id" {
  description = "ID of the IAM policy"
  value       = module.codedeploy_policy.id
}

output "json" {
  description = "Rendered JSON of the policy."
  value       = module.codedeploy_policy.json
}
