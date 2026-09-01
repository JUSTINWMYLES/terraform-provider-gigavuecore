---
page_title: "gigavuecore_retry_device_upgrade Action - gigavuecore"
subcategory: ""
description: |-
  Retry Device Upgrade
---

# gigavuecore_retry_device_upgrade Action

Retry Device Upgrade

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_retry_device_upgrade" "example" {
  config {
    device_upgrade_retry_specs = [{
      cluster_name = "example"
      option       = "Current"
      task_id      = "example"
    }]
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `device_upgrade_retry_specs` (Attributes List, required) (see [below for nested schema](#nestedatt--device_upgrade_retry_specs))

<a id="nestedatt--device_upgrade_retry_specs"></a>
### Nested Schema for `device_upgrade_retry_specs`

Required:

* `cluster_name` (String) - Cluster Name
* `option` (String) - Option
* `task_id` (String) - Task ID

