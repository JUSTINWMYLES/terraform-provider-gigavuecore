---
page_title: "gigavuecore_activet_device_config_file Action - gigavuecore"
subcategory: ""
description: |-
  Activate the specified device-stored config file
---

# gigavuecore_activet_device_config_file Action

Activate the specified device-stored config file

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_activet_device_config_file" "example" {
  config {
    file_name = "example"
    node_id   = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `file_name` (String, required) - Name of the config file to activate
* `node_id` (String, required) - NodeId of the target device. Note that while this is a device-level operation, activation only makes sense on cluster leader nodes


