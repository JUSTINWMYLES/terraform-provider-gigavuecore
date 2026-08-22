---
page_title: "gigavuecore_redefine_remote_auth_system_config Action - gigavuecore"
subcategory: ""
description: |-
  Redefine Remote Auth servers system configuration
---

# gigavuecore_redefine_remote_auth_system_config Action

Redefine Remote Auth servers system configuration

## Example Usage

```terraform
action "gigavuecore_redefine_remote_auth_system_config" "example" {
  config {
    cluster_id = "example"
    ldap_config = null
    radius_config = null
    tacacs_config = null
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `ldap_config` (Dynamic, optional) - System-level LDAP config
* `radius_config` (Dynamic, optional) - System-level RADIUS config
* `tacacs_config` (Dynamic, optional) - System-level TACACS+ config
