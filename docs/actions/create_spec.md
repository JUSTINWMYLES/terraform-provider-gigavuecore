---
page_title: "gigavuecore_create_spec Action - gigavuecore"
subcategory: ""
description: |-
  Create Device Upgrade Specification
---

# gigavuecore_create_spec Action

Create Device Upgrade Specification

## Example Usage

```terraform
action "gigavuecore_create_spec" "example" {
  config {
    activate_stage       = true
    cluster_ids          = [ "example" ]
    config_backup        = true
    exec_cluster_counter = 0
    fetch_stage          = true
    image_file_specs = [{
      device_model           = "HC2"
      device_model_qualifier = "HDCCV1"
      file_path              = "example"
      file_type              = "example"
      target_slot            = "example"
      update_uboot           = "example"
    }]
    image_server               = "example"
    install_stage              = true
    node_ids                   = [ "example" ]
    post_check_stage           = true
    reboot                     = true
    skip_not_reachable_devices = true
    tags = [{
      tag_key    = "example"
      tag_values = [ "example" ]
    }]
    task_id       = "example"
    task_name     = "example"
    upgrade_state = "ACTIVE"
    version       = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `activate_stage` (Boolean, optional) - Activate Stage
* `cluster_ids` (List of Dynamic, required) - Cluster IDs
* `config_backup` (Boolean, optional) - Config Backup
* `exec_cluster_counter` (Number, optional) - Exec Cluster Counter
* `fetch_stage` (Boolean, optional) - Fetch Stage
* `image_file_specs` (Attributes List, required) - Image File Specs (see [below for nested schema](#nestedatt--image_file_specs))
* `image_server` (String, required) - Image Server
* `install_stage` (Boolean, optional) - Install Stage
* `node_ids` (List of Dynamic, optional) - Node IDs
* `post_check_stage` (Boolean, optional) - Post Check Stage
* `reboot` (Boolean, optional) - Reboot
* `skip_not_reachable_devices` (Boolean, optional) - Skip Not Reachable Devices
* `tags` (Attributes List, optional) (see [below for nested schema](#nestedatt--tags))
* `task_id` (String, required) - Task ID
* `task_name` (String, optional) - Task Name
* `upgrade_state` (String, optional) - Upgrade State
* `version` (String, optional) - Version

<a id="nestedatt--image_file_specs"></a>
### Nested Schema for `image_file_specs`

Required:

* `device_model` (String) - Device Model
* `device_model_qualifier` (String) - Device Model Qualifier
* `file_path` (String) - File Path
* `file_type` (String) - File Type
* `target_slot` (String) - Target Slot

Optional:

* `update_uboot` (String) - Update UBoot

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Required:

* `tag_key` (String) - Tag Key
* `tag_values` (List of Dynamic) - Tag Values

