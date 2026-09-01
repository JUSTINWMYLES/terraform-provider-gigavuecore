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
  cluster_id = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID

### Attributes

In addition to all arguments above, the following attributes are exported:

* `auth_sequence` (Set of String, computed) - Authentication methods order for user login. Valid values are 'local', 'ldap', 'radius', 'tacacs'
* `external_login_mapping` (Attributes, computed) - Specifies how externally authenticated logins (RADIUS, TACACS+, or LDAP) are mapped to local accounts. Private class (see [below for nested schema](#nestedatt--external_login_mapping))
* `non_local_users` (Attributes, computed) - Settings for treating usernames that are not recognized as real accounts (not a locally configured account). Private class (see [below for nested schema](#nestedatt--non_local_users))
* `password_expiration` (Attributes, computed) - Node Password Expiration config. Private class (see [below for nested schema](#nestedatt--password_expiration))
* `sshd_max_sessions` (Number, computed) - Maximum concurrent session that can be logged into devices
* `user_lockout` (Attributes, computed) - Node AAA user lockout settings. Private class (see [below for nested schema](#nestedatt--user_lockout))

<a id="nestedatt--external_login_mapping"></a>
### Nested Schema for `external_login_mapping`

Read-Only:

* `default_local_user` (String) - Specifies the account to which externally authenticated logins are mapped when map order is set to remote-first (if there is no matching local account) or local-only
* `user_map_order` (String) - Specifies how externally authenticated logins (RADIUS, TACACS+, or LDAP) are mapped to local accounts

<a id="nestedatt--non_local_users"></a>
### Nested Schema for `non_local_users`

Read-Only:

* `hash_username` (Boolean) - Apply a hash function to the username and store the hashed result for usernames that are not recognized as real accounts (not a locally configured account)
* `track_auth_failures` (Boolean) - Enables tracking authentication failures for usernames that are not recognized as real accounts (not a locally configured account)

<a id="nestedatt--password_expiration"></a>
### Nested Schema for `password_expiration`

Read-Only:

* `duration` (Number) - the number of days password is valid
* `enabled` (Boolean) - Enables password expiration

<a id="nestedatt--user_lockout"></a>
### Nested Schema for `user_lockout`

Read-Only:

* `enable_admin_lockout` (Boolean) - Overrides the global settings for tracking and lockouts for the 'admin' account. Disabling means that the admin user will never be locked out, though their authentication failure history will still be tracked if tracking is enabled overall. This applies only to the single account with the username 'admin'. It does not apply to any other users with administrative privileges
* `enable_lockout` (Boolean) - Enables or disables locking out of user accounts based on authentication failures
* `lock_time` (Number) - In seconds. Specifies that no logins are permitted for this number of seconds following any login failure (not counting failures caused by the lockout mechanism, or the lockTime itself). This is not based on the number of consecutive failures. If both 'unlockTime' and 'lockTime' are set, the 'unlockTime' must be greater than the 'lockTime'
* `max_fail` (Number) - Sets the maximum number of consecutive authentication failures (attempts) permitted for a user account before the account is locked. After this number of failures, the account is locked and subsequent attempts are not permitted
* `track_auth_failures` (Boolean) - Enables or disables tracking of authentication failures. Can be used for informational purposes of for user account lockout
* `unlock_time` (Number) - In seconds. Specifies that if a user account is locked due to authentication failures, another login attempt will be permitted if this number of seconds has elapsed since the last login failure. That does not count failures caused by the lockout mechanism itself. A user must have been permitted to attempt to login, and then failed. After this interval has elapsed, the account does not become unlocked, nor does its history reset. It simply permits one more login attempt even if the account is locked

