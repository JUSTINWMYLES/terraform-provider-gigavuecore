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
    protocol   = "scp"
    source = {
      hostname = "example"
      password = "example"
      path     = "example"
      username = "example"
    }
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `protocol` (String, required)
* `source` (Attributes, required) (see [below for nested schema](#nestedatt--source))

<a id="nestedatt--source"></a>
### Nested Schema for `source`

Required:

* `hostname` (String) - server address
* `path` (String) - configuration file path on server

Optional:

* `password` (String) - password to use for server login
* `username` (String) - user name to use for server login

