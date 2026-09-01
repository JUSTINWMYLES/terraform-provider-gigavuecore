---
page_title: "gigavuecore_update_keystore_preference Action - gigavuecore"
subcategory: ""
description: |-
  Set keystore preference
---

# gigavuecore_update_keystore_preference Action

Set keystore preference

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_update_keystore_preference" "example" {
  config {
    auto_delete     = true
    auto_enable     = true
    auto_purge      = true
    body_cluster_id = "example"
    cluster_id      = "example"
    max_keys        = 1
    retention_time  = 1
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `auto_delete` (Boolean, optional)
* `auto_enable` (Boolean, optional)
* `auto_purge` (Boolean, optional)
* `body_cluster_id` (String, required)
* `cluster_id` (String, required) - Target Cluster ID
* `max_keys` (Number, optional)
* `retention_time` (Number, optional)


