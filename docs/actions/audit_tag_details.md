---
page_title: "gigavuecore_audit_tag_details Action - gigavuecore"
subcategory: ""
description: |-
  Reapply cluster and node level tags to ports
---

# gigavuecore_audit_tag_details Action

Reapply cluster and node level tags to ports

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_audit_tag_details" "example" {
  config {
    cluster_id = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID


