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
  context                    = null
  fm_api_token_user_entities = "example"
}

```
## Schema

### Arguments

The following arguments are supported:

* `context` (Dynamic, optional) - Gigamon query result context
* `fm_api_token_user_entities` (List of Dynamic, required)


