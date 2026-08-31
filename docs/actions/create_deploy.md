---
page_title: "gigavuecore_create_deploy Action - gigavuecore"
subcategory: ""
description: |-
  Create unified resource deployment by environment id
---

# gigavuecore_create_deploy Action

Create unified resource deployment by environment id

## Example Usage

```terraform
action "gigavuecore_create_deploy" "example" {
  config {
    body = {
      name          = "example"
      platform_type = "aws"
    }
    env_id = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `body` (Attributes, required) - Unified Resource deployment (see [below for nested schema](#nestedatt--body))
* `env_id` (String, required) - unified environment identifier

<a id="nestedatt--body"></a>
### Nested Schema for `body`

Required:

* `name` (String)
* `platform_type` (String)

