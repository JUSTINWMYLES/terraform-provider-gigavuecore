---
page_title: "gigavuecore_list_device_config_files Data Source - gigavuecore"
subcategory: ""
description: |-
  List existing configuration files on device
---

# gigavuecore_list_device_config_files Data Source

List existing configuration files on device

## Example Usage

```terraform
data "gigavuecore_list_device_config_files" "example" {
  node_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `node_id` (String, required) - Node ID

### Attributes

In addition to all arguments above, the following attributes are exported:

* `box_id` (String, computed) - Cluster BoxId of the node
* `config_files` (List(Object({description, file_name, file_size, last_modified, next_boot_file, running})), computed)
* `device_model` (String, computed) - Gigamon physical device models
* `sw_version` (String, computed) - Device's current Software version

