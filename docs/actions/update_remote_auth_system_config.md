---
page_title: "gigavuecore_update_remote_auth_system_config Action - gigavuecore"
subcategory: ""
description: |-
  Update Remote Auth servers system configuration
---

# gigavuecore_update_remote_auth_system_config Action

Update Remote Auth servers system configuration

## Example Usage

```terraform
action "gigavuecore_update_remote_auth_system_config" "example" {
  config {
    cluster_id = "example"
    ldap_config = {
      accept_user_roles = true
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
        bind_timeout    = 0
        group_attribute = "example"
        group_dn        = "example"
        login_attribute = "example"
        port            = 0
        search_scope    = "example"
        search_timeout  = 0
        ssl = {
          ca_list     = "example"
          cert_verify = true
          mode        = "example"
          server_port = 0
        }
        version = "example"
      }
    }
    radius_config = {
      accept_user_roles = true
      server_config_defaults = {
        retries    = 0
        secret_key = "example"
        timeout    = 0
      }
    }
    tacacs_config = {
      accept_user_roles = true
      server_config_defaults = {
        retries    = 0
        secret_key = "example"
        service    = "example"
        timeout    = 0
      }
    }
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `ldap_config` (Attributes, optional) - System-level LDAP config (see [below for nested schema](#nestedatt--ldap_config))
* `radius_config` (Attributes, optional) - System-level RADIUS config (see [below for nested schema](#nestedatt--radius_config))
* `tacacs_config` (Attributes, optional) - System-level TACACS+ config (see [below for nested schema](#nestedatt--tacacs_config))

<a id="nestedatt--ldap_config"></a>
### Nested Schema for `ldap_config`

Optional:

* `accept_user_roles` (Boolean) - Enables the GigaVUE H Series node to accept user roles assigned in the LDAP server
* `map_enable` (Boolean) - Enables mapping of remote base DN to local account
* `referrals` (Boolean) - Toggles support for LDAP referrals
* `remote_map_table` (Attributes List) (see [below for nested schema](#nestedatt--ldap_config--remote_map_table))
* `server_config_defaults` (Attributes) - Remote LDAP Server default config (see [below for nested schema](#nestedatt--ldap_config--server_config_defaults))

<a id="nestedatt--ldap_config--remote_map_table"></a>
### Nested Schema for `ldap_config.remote_map_table`

Optional:

* `local_account_name` (String) - Specifies local account to which remote base-dn needs to be mapped.
* `remote_base_dn` (String) - Specifies the base-dn of the remote user group to be mapped to local account

<a id="nestedatt--ldap_config--server_config_defaults"></a>
### Nested Schema for `ldap_config.server_config_defaults`

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
* `ssl` (Attributes) - Remote LDAP Server SSL config (see [below for nested schema](#nestedatt--ldap_config--server_config_defaults--ssl))
* `version` (String)

<a id="nestedatt--ldap_config--server_config_defaults--ssl"></a>
### Nested Schema for `ldap_config.server_config_defaults.ssl`

Optional:

* `ca_list` (String) - LDAP to use a supplemental CA list.
* `cert_verify` (Boolean) - Enable LDAP SSL/TLS certificate verification
* `mode` (String) - LDAP either in SSL or TLS
* `server_port` (Number) - LDAP SSL port number

<a id="nestedatt--radius_config"></a>
### Nested Schema for `radius_config`

Optional:

* `accept_user_roles` (Boolean) - Enables the GigaVUE H Series node to accept user roles assigned in the RADIUS server
* `server_config_defaults` (Attributes) - Remote RADIUS Server default config (see [below for nested schema](#nestedatt--radius_config--server_config_defaults))

<a id="nestedatt--radius_config--server_config_defaults"></a>
### Nested Schema for `radius_config.server_config_defaults`

Optional:

* `retries` (Number) - value of 0 disables retries
* `secret_key` (String)
* `timeout` (Number) - in seconds

<a id="nestedatt--tacacs_config"></a>
### Nested Schema for `tacacs_config`

Optional:

* `accept_user_roles` (Boolean) - Enables the GigaVUE H Series node to accept user roles assigned in the TACACS+ server
* `server_config_defaults` (Attributes) - Remote TACACS+ Server default config (see [below for nested schema](#nestedatt--tacacs_config--server_config_defaults))

<a id="nestedatt--tacacs_config--server_config_defaults"></a>
### Nested Schema for `tacacs_config.server_config_defaults`

Optional:

* `retries` (Number) - value of 0 disables retries
* `secret_key` (String)
* `service` (String) - Specify which authorization service will be used for TACACS
* `timeout` (Number) - in seconds

