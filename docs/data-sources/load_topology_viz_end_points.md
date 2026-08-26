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
  cluster_id  = null
  hostname    = null
  ignore_used = null
  type        = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Cluster Name
* `hostname` (String, required) - Node hostname
* `ignore_used` (Boolean, optional) - Set this to 'true' to display endpoints that are not part of anyother link
* `type` (String, optional) - EndPoint type

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `alias` (String) - PortId if endpointType is port, alias if it is a gigastream
* `endpoint_type` (String)
* `port_type` (String)
* `ports` (List of String) - List of portIds in the gigastream

