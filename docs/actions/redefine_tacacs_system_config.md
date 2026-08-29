---
page_title: "gigavuecore_redefine_tacacs_system_config Action - gigavuecore"
subcategory: ""
description: |-
  Redefine TACACS+ system configuration
---

# gigavuecore_redefine_tacacs_system_config Action

Redefine TACACS+ system configuration

## Example Usage

```terraform
action "gigavuecore_redefine_tacacs_system_config" "example" {
  config {
    accept_user_roles = true
    cluster_id        = "example"
    server_config_defaults = {
      retries    = 0
      secret_key = "example"
      service    = "example"
      timeout    = 0
    }
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `accept_user_roles` (Boolean, optional) - Enables the GigaVUE H Series node to accept user roles assigned in the TACACS+ server
* `cluster_id` (String, required) - Target Cluster ID
* `server_config_defaults` (Attributes, optional) - Remote TACACS+ Server default config (see [below for nested schema](#nestedatt--server_config_defaults))

<a id="nestedatt--server_config_defaults"></a>
### Nested Schema for `server_config_defaults`

Optional:

* `retries` (Number) - value of 0 disables retries
* `secret_key` (String)
* `service` (String) - Specify which authorization service will be used for TACACS
* `timeout` (Number) - in seconds

