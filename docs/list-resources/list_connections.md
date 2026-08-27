---
page_title: "gigavuecore_list_connections List Resource - gigavuecore"
subcategory: ""
description: |-
  List all unified resource connections by environment id
---

# gigavuecore_list_connections List Resource

List all unified resource connections by environment id

## Example Usage

```terraform
list "gigavuecore_list_connections" "example" {
  provider = gigavuecore
  limit    = 100
  config {
    env_id = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `env_id` (String, required) - unified environment identifier


