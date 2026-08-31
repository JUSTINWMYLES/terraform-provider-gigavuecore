---
page_title: "gigavuecore_update_env Action - gigavuecore"
subcategory: ""
description: |-
  Update a unified resource environment
---

# gigavuecore_update_env Action

Update a unified resource environment

## Example Usage

```terraform
action "gigavuecore_update_env" "example" {
  config {
    env_id = "example"
    name   = "example"
    type   = "aws"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `env_id` (String, required) - unified environment id
* `name` (String, required)
* `type` (String, required) - Unified Environment platform types


