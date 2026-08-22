---
page_title: "gigavuecore_load_all_management_interfaces Data Source - gigavuecore"
subcategory: ""
description: |-
  since FM 5.8
---

# gigavuecore_load_all_management_interfaces Data Source

since FM 5.8

## Example Usage

```terraform
data "gigavuecore_load_all_management_interfaces" "example" {
  box_id = null
  cluster_id = null
  interface_name = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `box_id` (Number, optional) - boxId
* `cluster_id` (String, required) - Target Cluster ID
* `interface_name` (String, optional) - Interface Name

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({box_id, cluster_id, discovery_protocol, g_arp, interface_name})), computed)

