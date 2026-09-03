---
page_title: "gigavuecore_clear_keystore_counters Action - gigavuecore"
subcategory: ""
description: |-
  Clear keystore hit counters
---

# gigavuecore_clear_keystore_counters Action

Clear keystore hit counters

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_clear_keystore_counters" "example" {
  config {
    cluster_id = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID


