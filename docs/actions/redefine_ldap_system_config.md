---
page_title: "gigavuecore_redefine_ldap_system_config Action - gigavuecore"
subcategory: ""
description: |-
  Redefine LDAP system configuration
---

# gigavuecore_redefine_ldap_system_config Action

Redefine LDAP system configuration

## Example Usage

```terraform
action "gigavuecore_redefine_ldap_system_config" "example" {
  config {
    accept_user_roles = true
    cluster_id = "example"
    map_enable = true
    referrals = true
    remote_map_table = null
    server_config_defaults = null
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `accept_user_roles` (Bool, optional) - Enables the GigaVUE H Series node to accept user roles assigned in the LDAP server
* `cluster_id` (String, required) - Target Cluster ID
* `map_enable` (Bool, optional) - Enables mapping of remote base DN to local account
* `referrals` (Bool, optional) - Toggles support for LDAP referrals
* `remote_map_table` (List(Dynamic), optional)
* `server_config_defaults` (Dynamic, optional) - Remote LDAP Server default config
