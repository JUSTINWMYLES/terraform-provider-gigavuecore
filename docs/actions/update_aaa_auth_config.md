---
page_title: "gigavuecore_update_aaa_auth_config Action - gigavuecore"
subcategory: ""
description: |-
  Update Node Authentication config
---

# gigavuecore_update_aaa_auth_config Action

Update Node Authentication config

## Example Usage

```terraform
action "gigavuecore_update_aaa_auth_config" "example" {
  config {
    auth_sequence          = [ "example" ]
    cluster_id             = "example"
    external_login_mapping = null
    non_local_users        = null
    password_expiration    = null
    sshd_max_sessions      = 1
    user_lockout           = null
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `auth_sequence` (Set of String, required) - Authentication methods order for user login. Valid values are 'local', 'ldap', 'radius', 'tacacs'
* `cluster_id` (String, required) - Target Cluster ID
* `external_login_mapping` (Dynamic, optional) - Specifies how externally authenticated logins (RADIUS, TACACS+, or LDAP) are mapped to local accounts. Private class
* `non_local_users` (Dynamic, optional) - Settings for treating usernames that are not recognized as real accounts (not a locally configured account). Private class
* `password_expiration` (Dynamic, optional) - Node Password Expiration config. Private class
* `sshd_max_sessions` (Number, optional) - Maximum concurrent session that can be logged into devices
* `user_lockout` (Dynamic, optional) - Node AAA user lockout settings. Private class


