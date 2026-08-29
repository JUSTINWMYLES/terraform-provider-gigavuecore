---
page_title: "gigavuecore_get_all_app_info List Resource - gigavuecore"
subcategory: ""
description: |-
  List all appInfo for unified deployment via environment and connection id
---

# gigavuecore_get_all_app_info List Resource

List all appInfo for unified deployment via environment and connection id

## Example Usage

```terraform
list "gigavuecore_get_all_app_info" "example" {
  provider = gigavuecore
  limit    = 100
  config {
    env_id   = "example"
    unify_id = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `env_id` (String, required) - unified environment identifier
* `unify_id` (String, required) - unified resource identifier


### Identity Attributes

The following identity attributes are exported for each matching result:

* `env_id` (String, computed)
* `unify_id` (String, computed)
* `appinfo` (String, computed)


