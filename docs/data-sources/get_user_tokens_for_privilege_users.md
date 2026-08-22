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

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `context` (Object({page_no, page_size, sort, total_items}), computed) - Gigamon query result context
  * `page_no` (Number, computed) - page number of the returned result set
  * `page_size` (Number, computed) - page size of the returned result set
  * `sort` (List(String), computed) - sorting info of the returned result set. list of fields in the array indicate sorting order
  * `total_items` (Number, computed) - total number of items in the queried entity type
* `fm_user_token_entities` (List(Object({authentication_type, created_by, created_ts, expiry_time, expiry_ts, groups, token, token_id, token_name, usage_count, username})), computed)

