---
page_title: "gigavuecore_retry_device_upgrade Action - gigavuecore"
subcategory: ""
description: |-
  Retry Device Upgrade
---

# gigavuecore_retry_device_upgrade Action

Retry Device Upgrade

## Example Usage

```terraform
action "gigavuecore_retry_device_upgrade" "example" {
  config {
    device_upgrade_retry_specs = [{
      cluster_name = "example"
      option       = "example"
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

