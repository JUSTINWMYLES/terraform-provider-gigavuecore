---
page_title: "gigavuecore_generate_report Action - gigavuecore"
subcategory: ""
description: |-
  Customer Deployed Assets Report Generation
---

# gigavuecore_generate_report Action

Customer Deployed Assets Report Generation

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_generate_report" "example" {
  config {
  }
}
```
