---
page_title: "gigavuecore_revoke_api_tokens_for_privilege_users Action - gigavuecore"
subcategory: ""
description: |-
  Revoke other user token of FM Users. Users with FM Security Management role with write access can access API.
---

# gigavuecore_revoke_api_tokens_for_privilege_users Action

Revoke other user token of FM Users. Users with FM Security Management role with write access can access API.

## Example Usage

```terraform
action "gigavuecore_revoke_api_tokens_for_privilege_users" "example" {
  config {
    context                    = null
    fm_api_token_user_entities = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `context` (Dynamic, optional) - Gigamon query result context
* `fm_api_token_user_entities` (List of Dynamic, required)


