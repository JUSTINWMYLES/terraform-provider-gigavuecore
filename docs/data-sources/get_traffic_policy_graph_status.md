---
page_title: "gigavuecore_get_traffic_policy_graph_status Data Source - gigavuecore"
subcategory: ""
description: |-
  Get the traffic policy graph status
---

# gigavuecore_get_traffic_policy_graph_status Data Source

Get the traffic policy graph status

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


