---
page_title: "gigavuecore_switch_node_system_config_file Action - gigavuecore"
subcategory: ""
description: |-
  Load a configuration file and make it the active configuration
---

# gigavuecore_switch_node_system_config_file Action

Load a configuration file and make it the active configuration

## Example Usage

```terraform
action "gigavuecore_switch_node_system_config_file" "example" {
  config {
    cluster_id = "example"
    filename = "example"
    keep_stack = true
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `filename` (String, optional) - Configuration file
* `keep_stack` (Bool, optional) - Keep the stack configuration for inband cluster configs. Default: false
