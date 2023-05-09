<!-- BEGIN_TF_DOCS -->
[![Geek Cell GmbH](https://raw.githubusercontent.com/geekcell/template-terraform-module/main/docs/assets/logo.svg)](https://www.geekcell.io/)

### Code Quality
[![License](https://img.shields.io/github/license/geekcell/terraform-aws-iam-policy)](https://github.com/geekcell/terraform-aws-iam-policy/blob/master/LICENSE)
[![GitHub release (latest tag)](https://img.shields.io/github/v/release/geekcell/terraform-aws-iam-policy?logo=github&sort=semver)](https://github.com/geekcell/terraform-aws-iam-policy/releases)
[![Release](https://github.com/geekcell/terraform-aws-iam-policy/actions/workflows/release.yaml/badge.svg)](https://github.com/geekcell/terraform-aws-iam-policy/actions/workflows/release.yaml)
[![Validate](https://github.com/geekcell/terraform-aws-iam-policy/actions/workflows/validate.yaml/badge.svg)](https://github.com/geekcell/terraform-aws-iam-policy/actions/workflows/validate.yaml)
[![Lint](https://github.com/geekcell/terraform-aws-iam-policy/actions/workflows/linter.yaml/badge.svg)](https://github.com/geekcell/terraform-aws-iam-policy/actions/workflows/linter.yaml)

### Security
[![Infrastructure Tests](https://www.bridgecrew.cloud/badges/github/geekcell/terraform-aws-iam-policy/general)](https://www.bridgecrew.cloud/link/badge?vcs=github&fullRepo=geekcell%2Fterraform-aws-iam-policy&benchmark=INFRASTRUCTURE+SECURITY)

#### Cloud
[![Infrastructure Tests](https://www.bridgecrew.cloud/badges/github/geekcell/terraform-aws-iam-policy/cis_aws)](https://www.bridgecrew.cloud/link/badge?vcs=github&fullRepo=geekcell%2Fterraform-aws-iam-policy&benchmark=CIS+AWS+V1.2)
[![Infrastructure Tests](https://www.bridgecrew.cloud/badges/github/geekcell/terraform-aws-iam-policy/cis_aws_13)](https://www.bridgecrew.cloud/link/badge?vcs=github&fullRepo=geekcell%2Fterraform-aws-iam-policy&benchmark=CIS+AWS+V1.3)
[![Infrastructure Tests](https://www.bridgecrew.cloud/badges/github/geekcell/terraform-aws-iam-policy/cis_azure)](https://www.bridgecrew.cloud/link/badge?vcs=github&fullRepo=geekcell%2Fterraform-aws-iam-policy&benchmark=CIS+AZURE+V1.1)
[![Infrastructure Tests](https://www.bridgecrew.cloud/badges/github/geekcell/terraform-aws-iam-policy/cis_azure_13)](https://www.bridgecrew.cloud/link/badge?vcs=github&fullRepo=geekcell%2Fterraform-aws-iam-policy&benchmark=CIS+AZURE+V1.3)
[![Infrastructure Tests](https://www.bridgecrew.cloud/badges/github/geekcell/terraform-aws-iam-policy/cis_gcp)](https://www.bridgecrew.cloud/link/badge?vcs=github&fullRepo=geekcell%2Fterraform-aws-iam-policy&benchmark=CIS+GCP+V1.1)

##### Container
[![Infrastructure Tests](https://www.bridgecrew.cloud/badges/github/geekcell/terraform-aws-iam-policy/cis_kubernetes_16)](https://www.bridgecrew.cloud/link/badge?vcs=github&fullRepo=geekcell%2Fterraform-aws-iam-policy&benchmark=CIS+KUBERNETES+V1.6)
[![Infrastructure Tests](https://www.bridgecrew.cloud/badges/github/geekcell/terraform-aws-iam-policy/cis_eks_11)](https://www.bridgecrew.cloud/link/badge?vcs=github&fullRepo=geekcell%2Fterraform-aws-iam-policy&benchmark=CIS+EKS+V1.1)
[![Infrastructure Tests](https://www.bridgecrew.cloud/badges/github/geekcell/terraform-aws-iam-policy/cis_gke_11)](https://www.bridgecrew.cloud/link/badge?vcs=github&fullRepo=geekcell%2Fterraform-aws-iam-policy&benchmark=CIS+GKE+V1.1)
[![Infrastructure Tests](https://www.bridgecrew.cloud/badges/github/geekcell/terraform-aws-iam-policy/cis_kubernetes)](https://www.bridgecrew.cloud/link/badge?vcs=github&fullRepo=geekcell%2Fterraform-aws-iam-policy&benchmark=CIS+KUBERNETES+V1.5)

#### Data protection
[![Infrastructure Tests](https://www.bridgecrew.cloud/badges/github/geekcell/terraform-aws-iam-policy/soc2)](https://www.bridgecrew.cloud/link/badge?vcs=github&fullRepo=geekcell%2Fterraform-aws-iam-policy&benchmark=SOC2)
[![Infrastructure Tests](https://www.bridgecrew.cloud/badges/github/geekcell/terraform-aws-iam-policy/pci)](https://www.bridgecrew.cloud/link/badge?vcs=github&fullRepo=geekcell%2Fterraform-aws-iam-policy&benchmark=PCI-DSS+V3.2)
[![Infrastructure Tests](https://www.bridgecrew.cloud/badges/github/geekcell/terraform-aws-iam-policy/pci_dss_v321)](https://www.bridgecrew.cloud/link/badge?vcs=github&fullRepo=geekcell%2Fterraform-aws-iam-policy&benchmark=PCI-DSS+V3.2.1)
[![Infrastructure Tests](https://www.bridgecrew.cloud/badges/github/geekcell/terraform-aws-iam-policy/iso)](https://www.bridgecrew.cloud/link/badge?vcs=github&fullRepo=geekcell%2Fterraform-aws-iam-policy&benchmark=ISO27001)
[![Infrastructure Tests](https://www.bridgecrew.cloud/badges/github/geekcell/terraform-aws-iam-policy/nist)](https://www.bridgecrew.cloud/link/badge?vcs=github&fullRepo=geekcell%2Fterraform-aws-iam-policy&benchmark=NIST-800-53)
[![Infrastructure Tests](https://www.bridgecrew.cloud/badges/github/geekcell/terraform-aws-iam-policy/hipaa)](https://www.bridgecrew.cloud/link/badge?vcs=github&fullRepo=geekcell%2Fterraform-aws-iam-policy&benchmark=HIPAA)
[![Infrastructure Tests](https://www.bridgecrew.cloud/badges/github/geekcell/terraform-aws-iam-policy/fedramp_moderate)](https://www.bridgecrew.cloud/link/badge?vcs=github&fullRepo=geekcell%2Fterraform-aws-iam-policy&benchmark=FEDRAMP+%28MODERATE%29)

# Terraform AWS IAM Policy

Introducing the AWS IAM Policy Collection Terraform Module, a comprehensive solution for managing your AWS Identity
and Access Management (IAM) policies. This module has been carefully crafted to include the most commonly used
policies in our setups, making it easier for you to manage and secure your AWS resources.

Our team of experts has worked with AWS IAM policies for years and has a deep understanding of the best practices
and configurations. By using this Terraform module, you can be sure that your policies are created and managed in
a secure, efficient, and scalable manner.

This module offers a one-stop-shop for all your IAM policy needs, saving you time and effort in the process. Whether
you're looking to grant access to specific AWS services or to limit the actions that can be performed on your
resources, this module has you covered.

So, if you're looking for a convenient and reliable solution for managing your IAM policies, look no further than
the AWS IAM Policy Collection Terraform Module. Give it a try and see the difference it can make in your AWS setup!

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_create_policy"></a> [create\_policy](#input\_create\_policy) | Whether to create the actual policy resource or to only render it. | `bool` | `true` | no |
| <a name="input_description"></a> [description](#input\_description) | Description of the Security Group. | `string` | `null` | no |
| <a name="input_name"></a> [name](#input\_name) | Name of the policy. | `string` | n/a | yes |
| <a name="input_policy_id"></a> [policy\_id](#input\_policy\_id) | ID for the policy document. | `string` | `null` | no |
| <a name="input_statements"></a> [statements](#input\_statements) | A map of principals which can assume the role. | <pre>list(object({<br>    sid    = optional(string)<br>    effect = optional(string, "Allow")<br><br>    actions     = optional(list(string))<br>    not_actions = optional(list(string))<br><br>    resources     = optional(list(string))<br>    not_resources = optional(list(string))<br><br>    conditions = optional(list(object({<br>      test     = string<br>      variable = string<br>      values   = list(string)<br>    })))<br><br>    principals = optional(list(object({<br>      type        = string<br>      identifiers = list(string)<br>    })))<br><br>    not_principals = optional(list(object({<br>      type        = string<br>      identifiers = list(string)<br>    })))<br>  }))</pre> | `[]` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | Tags to add to the Security Group. | `map(any)` | `{}` | no |
| <a name="input_templates"></a> [templates](#input\_templates) | A list of templates. Multiple templates will be combined into a single policy. | <pre>list(object({<br>    name = string<br>    vars = optional(map(any))<br>  }))</pre> | `[]` | no |
| <a name="input_use_name_prefix"></a> [use\_name\_prefix](#input\_use\_name\_prefix) | Use the `name` attribute as prefix for the role name. | `bool` | `true` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_arn"></a> [arn](#output\_arn) | ARN of the IAM policy |
| <a name="output_id"></a> [id](#output\_id) | ID of the IAM policy |
| <a name="output_json"></a> [json](#output\_json) | Rendered JSON of the policy. |
| <a name="output_name"></a> [name](#output\_name) | Name of the IAM policy |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_aws"></a> [aws](#provider\_aws) | >= 4.36 |

## Resources

- resource.aws_iam_policy.main (main.tf#101)
- data source.aws_caller_identity.current (main.tf#40)
- data source.aws_iam_policy_document.combined (main.tf#93)
- data source.aws_iam_policy_document.statement (main.tf#48)
- data source.aws_iam_policy_document.template (main.tf#44)
- data source.aws_region.current (main.tf#36)

# Examples
### Statements
```hcl
module "s3_policy" {
  source = "../../"

  name = var.name
  statements = [
    {
      effect = "Allow"
      actions = [
        "s3:GetObject",
        "s3:PutObject",
        "s3:PutObjectAcl",
        "s3:ListBucket",
        "s3:DeleteObject"
      ]
      resources = [
        "arn:aws:s3:::my-bucket",
        "arn:aws:s3:::my-bucket/*"
      ]
    }
  ]
}
```
### Templates
```hcl
module "codedeploy_policy" {
  source = "../../"

  name      = "web-app"
  templates = [
    {
      name = "codedeploy/ecs-blue-green-deployment"
      vars = {
        codedeploy_app_name              = "my-project"
        codedeploy_deployment_group_name = "web-app"

        ecs_cluster_name = "my-project"
        ecs_service_name = "web-app"

        task_definition_task_role_name      = "web-app"
        task_definition_execution_role_name = "web-app-exec"
      }
    },
    {
      name = "ecr/push-and-pull"
      vars = {
        ecr_repository_name = "web-app"
      }
    }
  ]
}
```
<!-- END_TF_DOCS -->
