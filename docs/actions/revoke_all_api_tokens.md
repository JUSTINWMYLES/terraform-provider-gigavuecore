---
page_title: "gigavuecore_revoke_all_api_tokens Action - gigavuecore"
subcategory: ""
description: |-
  Revoke token for FM Users
---

# gigavuecore_revoke_all_api_tokens Action

Revoke token for FM Users

## Example Usage

```terraform
action "gigavuecore_revoke_all_api_tokens" "example" {
  config {
    context = null
    fm_api_token_user_entities = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `context` (Dynamic, optional) - Gigamon query result context
* `fm_api_token_user_entities` (List(Dynamic), required)
