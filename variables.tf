## NAMING
variable "name" {
  description = "Name of the policy."
  type        = string
}

variable "description" {
  description = "Description of the Security Group."
  default     = null
  type        = string
}

variable "tags" {
  description = "Tags to add to the Security Group."
  default     = {}
  type        = map(any)
}

## POLICY FROM TEMPLATES
variable "create_policy" {
  description = "Whether to create the actual policy resource or to only render it."
  default     = true
  type        = bool
}

variable "templates" {
  description = "A list of templates. Multiple templates will be combined into a single policy."
  default     = []
  type = list(object({
    name = string
    vars = optional(map(any))
  }))
}

## POLICY
variable "use_name_prefix" {
  description = "Use the `name` attribute as prefix for the role name."
  default     = true
  type        = bool
}

variable "path" {
  description = "Path in which to create the policy."
  default     = "/"
  type        = string
}

variable "policy_id" {
  description = "ID for the policy document."
  default     = null
  type        = string
}

variable "statements" {
  description = "A map of principals which can assume the role."
  default     = []
  type = list(object({
    sid    = optional(string)
    effect = optional(string, "Allow")

    actions     = optional(list(string))
    not_actions = optional(list(string))

    resources     = optional(list(string))
    not_resources = optional(list(string))

    conditions = optional(list(object({
      test     = string
      variable = string
      values   = list(string)
    })))

    principals = optional(list(object({
      type        = string
      identifiers = list(string)
    })))

    not_principals = optional(list(object({
      type        = string
      identifiers = list(string)
    })))
  }))
}
