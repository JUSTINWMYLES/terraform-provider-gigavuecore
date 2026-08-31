---
page_title: "gigavuecore_get_policy_fabric_maps Data Source - gigavuecore"
subcategory: ""
description: |-
  Get Tools
---

# gigavuecore_get_policy_fabric_maps Data Source

Get Tools

~> **Note:** This data source is not yet wired to a remote API endpoint. Invoking it fails with an explicit "not wired" diagnostic instead of calling the API. The OpenAPI operations it was inferred from could not be resolved into a complete mapping; consult the eidos generation warnings for the exact cause.

## Example Usage

```terraform
data "gigavuecore_get_policy_fabric_maps" "example" {
  name = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `name` (Dynamic, required) - policy name


