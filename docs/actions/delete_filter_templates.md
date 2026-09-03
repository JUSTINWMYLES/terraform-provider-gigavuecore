---
page_title: "gigavuecore_delete_filter_templates Action - gigavuecore"
subcategory: ""
description: |-
  Delete all Filter Templates
---

# gigavuecore_delete_filter_templates Action

Delete all Filter Templates

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_delete_filter_templates" "example" {
  config {
    cluster_id = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID


