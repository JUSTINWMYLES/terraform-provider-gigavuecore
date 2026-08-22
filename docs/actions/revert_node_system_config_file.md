---
page_title: "gigavuecore_revert_node_system_config_file Action - gigavuecore"
subcategory: ""
description: |-
  Revert running configuration to last saved configuration
---

# gigavuecore_revert_node_system_config_file Action

Revert running configuration to last saved configuration

## Example Usage

```terraform
action "gigavuecore_revert_node_system_config_file" "example" {
  config {
    cluster_id = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
