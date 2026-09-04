---
page_title: "gigavuecore_get_user_tokens_for_privilege_users Data Source - gigavuecore"
subcategory: ""
description: |-
  Returns API tokens of the other users. Users with FM Security Management role with write access can access API.
---

# gigavuecore_get_user_tokens_for_privilege_users Data Source

Returns API tokens of the other users. Users with FM Security Management role with write access can access API.

## Example Usage

```terraform
data "gigavuecore_get_user_tokens_for_privilege_users" "example" {
}
```

## Schema

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `authentication_type` (String) - Authentication Type local/radius/tacacs+/external
* `created_by` (String) - FM user who has created the token
* `created_ts` (Number) - Token creation timestamp
* `expiry_time` (String) - Expiry Time in number of days. Default would be 30 and Maximum of 105 days.
* `expiry_ts` (String) - Expiry Timestamp
* `groups` (List of String) - FM User Groups
* `token` (String) - FM generated JWT token for FM REST API access
* `token_id` (String) - Random alpha-numeric 64 digit string
* `token_name` (String) - User defined token name
* `usage_count` (Number) - Token Tracking count. Usagecount keep track of number times token was used to access API's
* `username` (String) - FM Username

