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
    activate_stage = true
    cluster_ids = "example"
    config_backup = true
    exec_cluster_counter = 1
    fetch_stage = true
    image_file_specs = "example"
    image_server = "example"
    install_stage = true
    node_ids = null
    post_check_stage = true
    reboot = true
    skip_not_reachable_devices = true
    tags = null
    task_id = "example"
    task_name = "example"
    upgrade_state = "example"
    version = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `activate_stage` (Bool, optional) - Activate Stage
* `cluster_ids` (List(Dynamic), required) - Cluster IDs
* `config_backup` (Bool, optional) - Config Backup
* `exec_cluster_counter` (Number, optional) - Exec Cluster Counter
* `fetch_stage` (Bool, optional) - Fetch Stage
* `image_file_specs` (List(Dynamic), required) - Image File Specs
* `image_server` (String, required) - Image Server
* `install_stage` (Bool, optional) - Install Stage
* `node_ids` (List(Dynamic), optional) - Node IDs
* `post_check_stage` (Bool, optional) - Post Check Stage
* `reboot` (Bool, optional) - Reboot
* `skip_not_reachable_devices` (Bool, optional) - Skip Not Reachable Devices
* `tags` (List(Dynamic), optional)
* `task_id` (String, required) - Task ID
* `task_name` (String, optional) - Task Name
* `upgrade_state` (String, optional) - Upgrade State
* `version` (String, optional) - Version
