/**
 * # Terraform AWS IAM Policy Module
 *
 * IAM policy module which allows creating policies from statements or use remote
 * templates to render policies.
 */

locals {
  # Gather global variables passed to all templates
  templates_globals = length(var.templates) > 0 ? {
    aws_region     = data.aws_region.current[0].name
    aws_account_id = data.aws_caller_identity.current[0].account_id
  } : {}

  # Render any given templates
  templates_rendered = [
    for policy in var.templates :
    templatefile("${path.module}/templates/${policy.name}.tftpl", merge(local.templates_globals, policy.vars))
  ]

  # Render templates into HCL by trying JSON and YAML decoding
  templates_hcl = [for policy in local.templates_rendered : try(yamldecode(policy), jsondecode(policy))]
}

data "aws_region" "current" {
  count = length(var.templates) > 0 ? 1 : 0
}

data "aws_caller_identity" "current" {
  count = length(var.templates) > 0 ? 1 : 0
}

data "aws_iam_policy_document" "template" {
  override_policy_documents = [for policy in local.templates_hcl : jsonencode(policy)]
}

data "aws_iam_policy_document" "statement" {
  dynamic "statement" {
    for_each = var.statements

    content {
      sid    = statement.value.sid
      effect = statement.value.effect

      actions     = statement.value.actions
      not_actions = statement.value.not_actions

      resources     = statement.value.resources
      not_resources = statement.value.not_resources

      dynamic "principals" {
        for_each = coalesce(statement.value.principals, [])

        content {
          type        = principals.value.type
          identifiers = principals.value.identifiers
        }
      }

      dynamic "not_principals" {
        for_each = coalesce(statement.value.not_principals, [])

        content {
          type        = not_principals.value.type
          identifiers = not_principals.value.identifiers
        }
      }

      dynamic "condition" {
        for_each = coalesce(statement.value.conditions, [])

        content {
          test     = condition.value.test
          variable = condition.value.variable
          values   = condition.value.values
        }
      }
    }
  }
}

data "aws_iam_policy_document" "combined" {
  policy_id = var.policy_id
  override_policy_documents = [
    data.aws_iam_policy_document.template.json,
    data.aws_iam_policy_document.statement.json
  ]
}

resource "aws_iam_policy" "main" {
  count = var.create_policy ? 1 : 0

  name        = var.use_name_prefix ? null : var.name
  name_prefix = var.use_name_prefix ? var.name : null
  description = var.description
  policy      = data.aws_iam_policy_document.combined.json

  tags = merge({ Name = var.name }, var.tags)
}
