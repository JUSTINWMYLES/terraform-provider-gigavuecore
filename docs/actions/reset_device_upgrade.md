---
page_title: "gigavuecore_reset_device_upgrade Action - gigavuecore"
subcategory: ""
description: |-
  Reset Device Upgrade
---

# gigavuecore_reset_device_upgrade Action

Reset Device Upgrade

## Example Usage

```terraform
action "gigavuecore_reset_device_upgrade" "example" {
  config {
    device_upgrade_reset_specs = [{
      cluster_name = "example"
      task_id      = "example"
    }]
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `device_upgrade_reset_specs` (Attributes List, required) (see [below for nested schema](#nestedatt--device_upgrade_reset_specs))

<a id="nestedatt--device_upgrade_reset_specs"></a>
### Nested Schema for `device_upgrade_reset_specs`

Required:

* `cluster_name` (String) - Cluster Name
* `task_id` (String) - Task ID

