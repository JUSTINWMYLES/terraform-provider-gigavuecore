---
page_title: "gigavuecore_delete_deploy Action - gigavuecore"
subcategory: ""
description: |-
  Delete unified resource deployment by environment and unified resource id
---

# gigavuecore_delete_deploy Action

Delete unified resource deployment by environment and unified resource id

## Example Usage

```terraform
action "gigavuecore_delete_deploy" "example" {
  config {
    env_id = "example"
    unify_id = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `env_id` (String, required) - unified environment identifier
* `unify_id` (String, required) - unified resource identifier
