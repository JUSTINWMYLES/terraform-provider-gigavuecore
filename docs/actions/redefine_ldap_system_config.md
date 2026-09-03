---
page_title: "gigavuecore_redefine_ldap_system_config Action - gigavuecore"
subcategory: ""
description: |-
  Redefine LDAP system configuration
---

# gigavuecore_redefine_ldap_system_config Action

Redefine LDAP system configuration

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_redefine_ldap_system_config" "example" {
  config {
    accept_user_roles = true
    cluster_id        = "example"
    map_enable        = true
    referrals         = true
    remote_map_table = [{
      local_account_name = "example"
      remote_base_dn     = "example"
    }]
    server_config_defaults = {
      base_dn         = "example"
      bind_dn         = "example"
      bind_pwd        = "example"
      bind_timeout    = 1
      group_attribute = "example"
      group_dn        = "example"
      login_attribute = "example"
      port            = 1
      search_scope    = "oneLevel"
      search_timeout  = 1
      ssl = {
        ca_list     = "none"
        cert_verify = true
        mode        = "none"
        server_port = 0
      }
      version = "v2"
    }
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `accept_user_roles` (Boolean, optional) - Enables the GigaVUE H Series node to accept user roles assigned in the LDAP server
* `cluster_id` (String, required) - Target Cluster ID
* `map_enable` (Boolean, optional) - Enables mapping of remote base DN to local account
* `referrals` (Boolean, optional) - Toggles support for LDAP referrals
* `remote_map_table` (Attributes List, optional) (see [below for nested schema](#nestedatt--remote_map_table))
* `server_config_defaults` (Attributes, optional) - Remote LDAP Server default config (see [below for nested schema](#nestedatt--server_config_defaults))

<a id="nestedatt--remote_map_table"></a>
### Nested Schema for `remote_map_table`

Optional:

* `local_account_name` (String) - Specifies local account to which remote base-dn needs to be mapped.
* `remote_base_dn` (String) - Specifies the base-dn of the remote user group to be mapped to local account

<a id="nestedatt--server_config_defaults"></a>
### Nested Schema for `server_config_defaults`

Optional:

* `base_dn` (String) - Identifies the base distinguished name (location) of the user information in the LDAP server's schema
* `bind_dn` (String) - Specifies the distinguished name (dn) on the LDAP server with which to bind. By default, this is left empty for anonymous login
* `bind_pwd` (String) - Provides the credentials to be used for binding with the LDAP server. If bind-dn is left undefined for anonymous login (the default), bind-password should be left undefined, too
* `bind_timeout` (Number) - in seconds
* `group_attribute` (String) - Use this argument to specify the name of the attribute to check for group membership. If you specify a value for group-dn, the attribute you name here will be checked to see whether it contains the user's distinguished name as one of the values in the LDAP server
* `group_dn` (String) - You can use this option to require membership in the named group-dn for successful login to the H Series node
* `login_attribute` (String) - Specify the name of the LDAP attribute containing the login name
* `port` (Number)
* `search_scope` (String) - Specifies the search scope for the user under the base distinguished name
* `search_timeout` (Number) - in seconds
* `ssl` (Attributes) - Remote LDAP Server SSL config (see [below for nested schema](#nestedatt--server_config_defaults--ssl))
* `version` (String)

<a id="nestedatt--server_config_defaults--ssl"></a>
### Nested Schema for `server_config_defaults.ssl`

Optional:

* `ca_list` (String) - LDAP to use a supplemental CA list.
* `cert_verify` (Boolean) - Enable LDAP SSL/TLS certificate verification
* `mode` (String) - LDAP either in SSL or TLS
* `server_port` (Number) - LDAP SSL port number

