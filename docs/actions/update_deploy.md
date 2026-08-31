---
page_title: "gigavuecore_update_deploy Action - gigavuecore"
subcategory: ""
description: |-
  Update unified resource deployment by environment and unified resource id
---

# gigavuecore_update_deploy Action

Update unified resource deployment by environment and unified resource id

## Example Usage

```terraform
action "gigavuecore_update_deploy" "example" {
  config {
    body = {
      name          = "example"
      platform_type = "aws"
    }
    env_id   = "example"
    unify_id = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `body` (Attributes, required) - Unified Resource deployment (see [below for nested schema](#nestedatt--body))
* `env_id` (String, required) - unified environment identifier
* `unify_id` (String, required) - unified resource identifier

<a id="nestedatt--body"></a>
### Nested Schema for `body`

Required:

* `name` (String)
* `platform_type` (String)

