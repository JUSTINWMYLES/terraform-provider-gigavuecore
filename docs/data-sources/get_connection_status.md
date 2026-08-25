---
page_title: "gigavuecore_get_connection_status Data Source - gigavuecore"
subcategory: ""
description: |-
  Obtain the connection status for underlying unified environment and unified resource id
---

# gigavuecore_get_connection_status Data Source

Obtain the connection status for underlying unified environment and unified resource id

## Example Usage

```terraform
data "gigavuecore_get_connection_status" "example" {
  env_id   = null
  unify_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `env_id` (String, required) - unified environment identifier
* `unify_id` (String, required) - unified resource identifier


