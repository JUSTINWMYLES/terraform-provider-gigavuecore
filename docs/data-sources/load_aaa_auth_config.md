---
page_title: "gigavuecore_load_aaa_auth_config Data Source - gigavuecore"
subcategory: ""
description: |-
  Load Node Authentication config
---

# gigavuecore_load_aaa_auth_config Data Source

Load Node Authentication config

## Example Usage

```terraform
data "gigavuecore_load_aaa_auth_config" "example" {
  cluster_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID

### Attributes

In addition to all arguments above, the following attributes are exported:

* `auth_sequence` (Set(String), computed) - Authentication methods order for user login. Valid values are 'local', 'ldap', 'radius', 'tacacs'
* `external_login_mapping` (Object({default_local_user, user_map_order}), computed) - Specifies how externally authenticated logins (RADIUS, TACACS+, or LDAP) are mapped to local accounts. Private class
  * `default_local_user` (String, computed) - Specifies the account to which externally authenticated logins are mapped when map order is set to remote-first (if there is no matching local account) or local-only
  * `user_map_order` (String, computed) - Specifies how externally authenticated logins (RADIUS, TACACS+, or LDAP) are mapped to local accounts
* `non_local_users` (Object({hash_username, track_auth_failures}), computed) - Settings for treating usernames that are not recognized as real accounts (not a locally configured account). Private class
  * `hash_username` (Bool, computed) - Apply a hash function to the username and store the hashed result for usernames that are not recognized as real accounts (not a locally configured account)
  * `track_auth_failures` (Bool, computed) - Enables tracking authentication failures for usernames that are not recognized as real accounts (not a locally configured account)
* `password_expiration` (Object({duration, enabled}), computed) - Node Password Expiration config. Private class
  * `duration` (Number, computed) - the number of days password is valid
  * `enabled` (Bool, computed) - Enables password expiration
* `sshd_max_sessions` (Number, computed) - Maximum concurrent session that can be logged into devices
* `user_lockout` (Object({enable_admin_lockout, enable_lockout, lock_time, max_fail, track_auth_failures, unlock_time}), computed) - Node AAA user lockout settings. Private class
  * `enable_admin_lockout` (Bool, computed) - Overrides the global settings for tracking and lockouts for the 'admin' account. Disabling means that the admin user will never be locked out, though their authentication failure history will still be tracked if tracking is enabled overall. This applies only to the single account with the username 'admin'. It does not apply to any other users with administrative privileges
  * `enable_lockout` (Bool, computed) - Enables or disables locking out of user accounts based on authentication failures
  * `lock_time` (Number, computed) - In seconds. Specifies that no logins are permitted for this number of seconds following any login failure (not counting failures caused by the lockout mechanism, or the lockTime itself). This is not based on the number of consecutive failures. If both 'unlockTime' and 'lockTime' are set, the 'unlockTime' must be greater than the 'lockTime'
  * `max_fail` (Number, computed) - Sets the maximum number of consecutive authentication failures (attempts) permitted for a user account before the account is locked. After this number of failures, the account is locked and subsequent attempts are not permitted
  * `track_auth_failures` (Bool, computed) - Enables or disables tracking of authentication failures. Can be used for informational purposes of for user account lockout
  * `unlock_time` (Number, computed) - In seconds. Specifies that if a user account is locked due to authentication failures, another login attempt will be permitted if this number of seconds has elapsed since the last login failure. That does not count failures caused by the lockout mechanism itself. A user must have been permitted to attempt to login, and then failed. After this interval has elapsed, the account does not become unlocked, nor does its history reset. It simply permits one more login attempt even if the account is locked

