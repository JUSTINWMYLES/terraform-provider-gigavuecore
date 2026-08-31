---
page_title: "gigavuecore_update_l2_circuit_resource_config Action - gigavuecore"
subcategory: ""
description: |-
  Update L2Circuit resource pool configuration
---

# gigavuecore_update_l2_circuit_resource_config Action

Update L2Circuit resource pool configuration

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_update_l2_circuit_resource_config" "example" {
  config {
    vlan_ids = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `vlan_ids` (String, required) - VLAN id ranges used for L2CIRCUIT resource type.


