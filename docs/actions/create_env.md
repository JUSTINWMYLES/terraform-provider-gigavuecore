---
page_title: "gigavuecore_create_env Action - gigavuecore"
subcategory: ""
description: |-
  Create a unified resource environment
---

# gigavuecore_create_env Action

Create a unified resource environment

## Example Usage

```terraform
action "gigavuecore_create_env" "example" {
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


