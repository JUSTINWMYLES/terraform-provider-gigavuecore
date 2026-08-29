---
page_title: "gigavuecore_get_connected_nodes Data Source - gigavuecore"
subcategory: ""
description: |-
  Retrieve connected nodes for troubleshooting flows
---

# gigavuecore_get_connected_nodes Data Source

Retrieve connected nodes for troubleshooting flows

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


