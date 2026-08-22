---
page_title: "gigavuecore_get_connected_links_traffic_flows_troubleshoot_flows_topology Data Source - gigavuecore"
subcategory: ""
description: |-
  Retrieve traversed nodes with levels for a cluster and port
---

# gigavuecore_get_connected_links_traffic_flows_troubleshoot_flows_topology Data Source

Retrieve traversed nodes with levels for a cluster and port

## Example Usage

```terraform
data "gigavuecore_get_connected_links_traffic_flows_troubleshoot_flows_topology" "example" {
  cluster_name = null
  start_point = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_name` (String, optional) - Target Cluster Name
* `start_point` (String, optional) - Start port

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Dynamic), computed)

