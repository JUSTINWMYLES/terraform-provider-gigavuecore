---
page_title: "gigavuecore_get_traffic_policy_graph_status Data Source - gigavuecore"
subcategory: ""
description: |-
  Get the traffic policy graph status
---

# gigavuecore_get_traffic_policy_graph_status Data Source

Get the traffic policy graph status

~> **Note:** This data source is not yet wired to a remote API endpoint. Invoking it fails with an explicit "not wired" diagnostic instead of calling the API. The OpenAPI operations it was inferred from could not be resolved into a complete mapping; consult the eidos generation warnings for the exact cause.

## Example Usage

```terraform
data "gigavuecore_get_traffic_policy_graph_status" "example" {
  alias = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Traffic Policy Graph alias


