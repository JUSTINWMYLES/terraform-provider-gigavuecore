---
page_title: "gigavuecore_load_remote_auth_system_config Data Source - gigavuecore"
subcategory: ""
description: |-
  Load Remote Auth servers system configuration
---

# gigavuecore_load_remote_auth_system_config Data Source

Load Remote Auth servers system configuration

## Example Usage

```terraform
data "gigavuecore_load_remote_auth_system_config" "example" {
  cluster_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID

### Attributes

In addition to all arguments above, the following attributes are exported:

* `ldap_config` (Object({accept_user_roles, map_enable, referrals, remote_map_table, server_config_defaults}), computed) - System-level LDAP config
  * `accept_user_roles` (Bool, computed) - Enables the GigaVUE H Series node to accept user roles assigned in the LDAP server
  * `map_enable` (Bool, computed) - Enables mapping of remote base DN to local account
  * `referrals` (Bool, computed) - Toggles support for LDAP referrals
  * `remote_map_table` (List(Object({local_account_name, remote_base_dn})), computed)
  * `server_config_defaults` (Object({base_dn, bind_dn, bind_pwd, bind_timeout, group_attribute, group_dn, login_attribute, port, search_scope, search_timeout, ssl, version}), computed) - Remote LDAP Server default config
    * `base_dn` (String, computed) - Identifies the base distinguished name (location) of the user information in the LDAP server's schema
    * `bind_dn` (String, computed) - Specifies the distinguished name (dn) on the LDAP server with which to bind. By default, this is left empty for anonymous login
    * `bind_pwd` (String, computed) - Provides the credentials to be used for binding with the LDAP server. If bind-dn is left undefined for anonymous login (the default), bind-password should be left undefined, too
    * `bind_timeout` (Number, computed) - in seconds
    * `group_attribute` (String, computed) - Use this argument to specify the name of the attribute to check for group membership. If you specify a value for group-dn, the attribute you name here will be checked to see whether it contains the user's distinguished name as one of the values in the LDAP server
    * `group_dn` (String, computed) - You can use this option to require membership in the named group-dn for successful login to the H Series node
    * `login_attribute` (String, computed) - Specify the name of the LDAP attribute containing the login name
    * `port` (Number, computed)
    * `search_scope` (String, computed) - Specifies the search scope for the user under the base distinguished name
    * `search_timeout` (Number, computed) - in seconds
    * `ssl` (Object({ca_list, cert_verify, mode, server_port}), computed) - Remote LDAP Server SSL config
      * `ca_list` (String, computed) - LDAP to use a supplemental CA list.
      * `cert_verify` (Bool, computed) - Enable LDAP SSL/TLS certificate verification
      * `mode` (String, computed) - LDAP either in SSL or TLS
      * `server_port` (Number, computed) - LDAP SSL port number
    * `version` (String, computed)
* `radius_config` (Object({accept_user_roles, server_config_defaults}), computed) - System-level RADIUS config
  * `accept_user_roles` (Bool, computed) - Enables the GigaVUE H Series node to accept user roles assigned in the RADIUS server
  * `server_config_defaults` (Object({retries, secret_key, timeout}), computed) - Remote RADIUS Server default config
    * `retries` (Number, computed) - value of 0 disables retries
    * `secret_key` (String, computed)
    * `timeout` (Number, computed) - in seconds
* `tacacs_config` (Object({accept_user_roles, server_config_defaults}), computed) - System-level TACACS+ config
  * `accept_user_roles` (Bool, computed) - Enables the GigaVUE H Series node to accept user roles assigned in the TACACS+ server
  * `server_config_defaults` (Object({retries, secret_key, service, timeout}), computed) - Remote TACACS+ Server default config
    * `retries` (Number, computed) - value of 0 disables retries
    * `secret_key` (String, computed)
    * `service` (String, computed) - Specify which authorization service will be used for TACACS
    * `timeout` (Number, computed) - in seconds

