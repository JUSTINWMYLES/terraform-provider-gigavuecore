---
page_title: "gigavuecore_delete_device_config_file Action - gigavuecore"
subcategory: ""
description: |-
  Delete specified device-side config file
---

# gigavuecore_delete_device_config_file Action

Delete specified device-side config file

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_delete_device_config_file" "example" {
  config {
    config_file_name = "example"
    node_id          = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `config_file_name` (String, required) - Name of the target configuration file
* `node_id` (String, required) - ID of the target device


