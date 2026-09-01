---
page_title: "gigavuecore_get_connection_status Data Source - gigavuecore"
subcategory: ""
description: |-
  Obtain the connection status for underlying unified environment and unified resource id
---

# gigavuecore_get_connection_status Data Source

Obtain the connection status for underlying unified environment and unified resource id

~> **Note:** This data source is not yet wired to a remote API endpoint. Invoking it fails with an explicit "not wired" diagnostic instead of calling the API. The OpenAPI operations it was inferred from could not be resolved into a complete mapping; consult the eidos generation warnings for the exact cause.

## Example Usage

```terraform
data "gigavuecore_get_connection_status" "example" {
  env_id   = "example"
  unify_id = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `env_id` (String, required) - unified environment identifier
* `unify_id` (String, required) - unified resource identifier


