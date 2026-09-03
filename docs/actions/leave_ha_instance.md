---
page_title: "gigavuecore_leave_ha_instance Action - gigavuecore"
subcategory: ""
description: |-
  Join FM Instance to HA Group
---

# gigavuecore_leave_ha_instance Action

Join FM Instance to HA Group

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_leave_ha_instance" "example" {
  config {
    ha_group_name = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `ha_group_name` (String, required) - HA group name


