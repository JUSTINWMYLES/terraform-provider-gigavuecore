---
page_title: "gigavuecore_regenerate_secret Ephemeral Resource - gigavuecore"
subcategory: ""
description: |-
  Revoke other user tokens of FM Users. Users with FM Security Management role with write access can access API.
---

# gigavuecore_regenerate_secret Ephemeral Resource

Revoke other user tokens of FM Users. Users with FM Security Management role with write access can access API.

~> **Note:** Ephemeral resources are only available within the context of a single Terraform operation and are never persisted to state or plan files.

## Example Usage

```terraform
ephemeral "gigavuecore_regenerate_secret" "example" {
  context = {
    page_no     = 0
    page_size   = 0
    sort        = [ "example" ]
    total_items = 0
  }
  fm_api_token_user_entities = [{
    authentication_type = "example"
    created_by          = "example"
    created_ts          = "example"
    expiry_time         = "example"
    expiry_ts           = "example"
    groups              = [ "example" ]
    token               = "example"
    token_id            = "example"
    token_name          = "example"
    usage_count         = "example"
    username            = "example"
  }]
}

```
## Schema

### Arguments

The following arguments are supported:

* `context` (Attributes, optional) - Gigamon query result context (see [below for nested schema](#nestedatt--context))
* `fm_api_token_user_entities` (Attributes List, required) (see [below for nested schema](#nestedatt--fm_api_token_user_entities))

<a id="nestedatt--context"></a>
### Nested Schema for `context`

Required:

* `total_items` (Number) - total number of items in the queried entity type
Optional:

* `page_no` (Number) - page number of the returned result set
* `page_size` (Number) - page size of the returned result set
* `sort` (List of String) - sorting info of the returned result set. list of fields in the array indicate sorting order
<a id="nestedatt--fm_api_token_user_entities"></a>
### Nested Schema for `fm_api_token_user_entities`

Required:

* `expiry_time` (String) - Expiry Time in number of days. Default would be 30 and Maximum of 105 days.
* `groups` (List of String) - FM User Groups
* `token_name` (String) - User defined token name
* `username` (String) - FM Username
Optional:

* `authentication_type` (String) - Authentication Type local/radius/tacacs+/external
* `created_by` (String) - FM user who has created the token
* `created_ts` (Dynamic) - Token creation timestamp
* `expiry_ts` (String) - Expiry Timestamp
* `token` (String) - FM generated JWT token for FM REST API access
* `token_id` (String) - Random alpha-numeric 64 digit string
* `usage_count` (Dynamic) - Token Tracking count. Usagecount keep track of number times token was used to access API's

