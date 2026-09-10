---
page_title: "gigavuecore_delete_report Action - gigavuecore"
subcategory: ""
description: |-
  Delete CDA Audit Report
---

# gigavuecore_delete_report Action

Delete CDA Audit Report

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_delete_report" "example" {
  config {
    file_name = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `file_name` (String, required) - Name of the CDA Audit Report


