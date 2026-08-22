---
page_title: "gigavuecore_restore_node_system_config_text_file Action - gigavuecore"
subcategory: ""
description: |-
  Apply the text configuration to the running system from a remote host or local storage
---

# gigavuecore_restore_node_system_config_text_file Action

Apply the text configuration to the running system from a remote host or local storage

## Example Usage

```terraform
action "gigavuecore_restore_node_system_config_text_file" "example" {
  config {
    clear_config = true
    cluster_id = "example"
    destination = "example"
    fail_continue = true
    filename = "example"
    protocol = "example"
    remote = true
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `clear_config` (Bool, optional) - If true, clear the existing traffic configuration before applying the text configuration file.
* `cluster_id` (String, required) - Target Cluster ID
* `destination` (Dynamic, required)
* `fail_continue` (Bool, required) - If true, while applying commands, continue execution even if one of them fails
* `filename` (String, required) - File name of the text configuration file to be applied on system
* `protocol` (String, required) - File Source/File Destination protocol. http and https only applicable for fileSource
* `remote` (Bool, required) - If true, fetch the text configuration file from remote host.If false, fetch file from local storage
