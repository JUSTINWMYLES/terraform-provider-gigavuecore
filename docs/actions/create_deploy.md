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
    body   = "example"
    env_id = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `body` (Dynamic, required) - Unified Resource deployment
* `env_id` (String, required) - unified environment identifier


