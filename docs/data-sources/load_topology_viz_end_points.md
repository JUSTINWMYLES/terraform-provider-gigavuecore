---
page_title: "gigavuecore_load_topology_viz_end_points Data Source - gigavuecore"
subcategory: ""
description: |-
  Load ports and gigaStream information required for manual link creation
---

# gigavuecore_load_topology_viz_end_points Data Source

Load ports and gigaStream information required for manual link creation

## Example Usage

```terraform
data "gigavuecore_load_topology_viz_end_points" "example" {
  cluster_id = null
  hostname = null
  ignore_used = null
  type = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Cluster Name
* `hostname` (String, required) - Node hostname
* `ignore_used` (Bool, optional) - Set this to 'true' to display endpoints that are not part of anyother link
* `type` (String, optional) - EndPoint type

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({alias, endpoint_type, port_type, ports})), computed)

