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
    cluster_id        = "example"
    server_config_defaults = {
      retries    = 0
      secret_key = "example"
      timeout    = 0
    }
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `accept_user_roles` (Boolean, optional) - Enables the GigaVUE H Series node to accept user roles assigned in the RADIUS server
* `cluster_id` (String, required) - Target Cluster ID
* `server_config_defaults` (Attributes, optional) - Remote RADIUS Server default config (see [below for nested schema](#nestedatt--server_config_defaults))

<a id="nestedatt--server_config_defaults"></a>
### Nested Schema for `server_config_defaults`

Optional:

* `retries` (Number) - value of 0 disables retries
* `secret_key` (String)
* `timeout` (Number) - in seconds

