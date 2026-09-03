---
page_title: "gigavuecore_redefine_aaa_auth_config Action - gigavuecore"
subcategory: ""
description: |-
  Redefine Node Authentication config
---

# gigavuecore_redefine_aaa_auth_config Action

Redefine Node Authentication config

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_redefine_aaa_auth_config" "example" {
  config {
    auth_sequence = ["local"]
    cluster_id    = "example"
    external_login_mapping = {
      default_local_user = "example"
      user_map_order     = "localOnly"
    }
    non_local_users = {
      hash_username       = true
      track_auth_failures = true
    }
    password_expiration = {
      duration = 1
      enabled  = true
    }
    sshd_max_sessions = 1
    user_lockout = {
      enable_admin_lockout = true
      enable_lockout       = true
      lock_time            = 0
      max_fail             = 0
      track_auth_failures  = true
      unlock_time          = 0
    }
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `auth_sequence` (Set of String, required) - Authentication methods order for user login. Valid values are 'local', 'ldap', 'radius', 'tacacs'
* `cluster_id` (String, required) - Target Cluster ID
* `external_login_mapping` (Attributes, optional) - Specifies how externally authenticated logins (RADIUS, TACACS+, or LDAP) are mapped to local accounts. Private class (see [below for nested schema](#nestedatt--external_login_mapping))
* `non_local_users` (Attributes, optional) - Settings for treating usernames that are not recognized as real accounts (not a locally configured account). Private class (see [below for nested schema](#nestedatt--non_local_users))
* `password_expiration` (Attributes, optional) - Node Password Expiration config. Private class (see [below for nested schema](#nestedatt--password_expiration))
* `sshd_max_sessions` (Number, optional) - Maximum concurrent session that can be logged into devices
* `user_lockout` (Attributes, optional) - Node AAA user lockout settings. Private class (see [below for nested schema](#nestedatt--user_lockout))

<a id="nestedatt--external_login_mapping"></a>
### Nested Schema for `external_login_mapping`

Optional:

* `default_local_user` (String) - Specifies the account to which externally authenticated logins are mapped when map order is set to remote-first (if there is no matching local account) or local-only
* `user_map_order` (String) - Specifies how externally authenticated logins (RADIUS, TACACS+, or LDAP) are mapped to local accounts

<a id="nestedatt--non_local_users"></a>
### Nested Schema for `non_local_users`

Optional:

* `hash_username` (Boolean) - Apply a hash function to the username and store the hashed result for usernames that are not recognized as real accounts (not a locally configured account)
* `track_auth_failures` (Boolean) - Enables tracking authentication failures for usernames that are not recognized as real accounts (not a locally configured account)

<a id="nestedatt--password_expiration"></a>
### Nested Schema for `password_expiration`

Optional:

* `duration` (Number) - the number of days password is valid
* `enabled` (Boolean) - Enables password expiration

<a id="nestedatt--user_lockout"></a>
### Nested Schema for `user_lockout`

Optional:

* `enable_admin_lockout` (Boolean) - Overrides the global settings for tracking and lockouts for the 'admin' account. Disabling means that the admin user will never be locked out, though their authentication failure history will still be tracked if tracking is enabled overall. This applies only to the single account with the username 'admin'. It does not apply to any other users with administrative privileges
* `enable_lockout` (Boolean) - Enables or disables locking out of user accounts based on authentication failures
* `lock_time` (Number) - In seconds. Specifies that no logins are permitted for this number of seconds following any login failure (not counting failures caused by the lockout mechanism, or the lockTime itself). This is not based on the number of consecutive failures. If both 'unlockTime' and 'lockTime' are set, the 'unlockTime' must be greater than the 'lockTime'
* `max_fail` (Number) - Sets the maximum number of consecutive authentication failures (attempts) permitted for a user account before the account is locked. After this number of failures, the account is locked and subsequent attempts are not permitted
* `track_auth_failures` (Boolean) - Enables or disables tracking of authentication failures. Can be used for informational purposes of for user account lockout
* `unlock_time` (Number) - In seconds. Specifies that if a user account is locked due to authentication failures, another login attempt will be permitted if this number of seconds has elapsed since the last login failure. That does not count failures caused by the lockout mechanism itself. A user must have been permitted to attempt to login, and then failed. After this interval has elapsed, the account does not become unlocked, nor does its history reset. It simply permits one more login attempt even if the account is locked

