---
page_title: "gigavuecore_redefine_radius_system_config Action - gigavuecore"
subcategory: ""
description: |-
  Redefine RADIUS system configuration
---

# gigavuecore_redefine_radius_system_config Action

Redefine RADIUS system configuration

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

~> **Warning:** This action accepts attributes whose names indicate secrets (secretKey), but action schemas cannot mark attributes Sensitive, so their values are stored and displayed in plain text. Avoid passing real secrets through these attributes.

## Example Usage

```terraform
action "gigavuecore_redefine_radius_system_config" "example" {
  config {
    accept_user_roles = true
    cluster_id        = "example"
    server_config_defaults = {
      retries    = 0
      secret_key = "example"
      timeout    = 1
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

