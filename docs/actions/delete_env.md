---
page_title: "gigavuecore_delete_env Action - gigavuecore"
subcategory: ""
description: |-
  Delete unified resource environment
---

# gigavuecore_delete_env Action

Delete unified resource environment

## Example Usage

```terraform
action "gigavuecore_delete_env" "example" {
  config {
    env_id = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `env_id` (String, required) - envId is the unified resource environment id
