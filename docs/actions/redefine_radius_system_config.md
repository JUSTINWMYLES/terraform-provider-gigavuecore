---
page_title: "gigavuecore_redefine_radius_system_config Action - gigavuecore"
subcategory: ""
description: |-
  Redefine RADIUS system configuration
---

# gigavuecore_redefine_radius_system_config Action

Redefine RADIUS system configuration

## Example Usage

```terraform
action "gigavuecore_redefine_radius_system_config" "example" {
  config {
    accept_user_roles = true
    cluster_id = "example"
    server_config_defaults = null
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `accept_user_roles` (Bool, optional) - Enables the GigaVUE H Series node to accept user roles assigned in the RADIUS server
* `cluster_id` (String, required) - Target Cluster ID
* `server_config_defaults` (Dynamic, optional) - Remote RADIUS Server default config
