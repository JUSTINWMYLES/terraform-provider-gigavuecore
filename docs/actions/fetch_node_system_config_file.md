---
page_title: "gigavuecore_fetch_node_system_config_file Action - gigavuecore"
subcategory: ""
description: |-
  Instruct the device to download a configuration file from a remote host
---

# gigavuecore_fetch_node_system_config_file Action

Instruct the device to download a configuration file from a remote host

## Example Usage

```terraform
action "gigavuecore_fetch_node_system_config_file" "example" {
  config {
    cluster_id = "example"
    protocol = "example"
    source = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `protocol` (String, required)
* `source` (Dynamic, required)
