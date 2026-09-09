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
  box_id         = 0
  cluster_id     = "example"
  interface_name = "example"
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

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `box_id` (Number)
* `cluster_id` (String) - id of the defining cluster
* `discovery_protocol` (String)
* `g_arp` (Boolean) - Enable or disable Gratuitous ARP
* `interface_name` (String) - Management Interface Name

