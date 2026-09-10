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
  node_id = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `node_id` (String, required) - Node ID

### Attributes

In addition to all arguments above, the following attributes are exported:

* `box_id` (String, computed) - Cluster BoxId of the node
* `config_files` (Attributes List, computed) (see [below for nested schema](#nestedatt--config_files))
* `device_model` (String, computed) - Gigamon physical device models
* `sw_version` (String, computed) - Device's current Software version

<a id="nestedatt--config_files"></a>
### Nested Schema for `config_files`

Read-Only:

* `description` (String) - Description of the device configuration file
* `file_name` (String) - Config file name
* `file_size` (Number) - Config file size
* `last_modified` (String) - Last modification timestamp
* `next_boot_file` (Boolean) - indicates whether this config file should become active after reboot
* `running` (Boolean) - indicates whether this is the currently active config file

