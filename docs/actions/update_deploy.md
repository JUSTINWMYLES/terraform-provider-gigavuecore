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
    body     = "example"
    env_id   = "example"
    unify_id = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `body` (Dynamic, required) - Unified Resource deployment
* `env_id` (String, required) - unified environment identifier
* `unify_id` (String, required) - unified resource identifier


