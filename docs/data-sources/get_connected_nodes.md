---
page_title: "gigavuecore_get_connected_nodes Data Source - gigavuecore"
subcategory: ""
description: |-
  Retrieve connected nodes for troubleshooting flows
---

# gigavuecore_get_connected_nodes Data Source

Retrieve connected nodes for troubleshooting flows

~> **Note:** This data source is not yet wired to a remote API endpoint. Invoking it fails with an explicit "not wired" diagnostic instead of calling the API. The OpenAPI operations it was inferred from could not be resolved into a complete mapping; consult the eidos generation warnings for the exact cause.

## Example Usage

```terraform
data "gigavuecore_get_connected_nodes" "example" {
  cluster_name = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_name` (String, optional) - Target Cluster Name


