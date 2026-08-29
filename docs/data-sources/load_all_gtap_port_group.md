---
page_title: "gigavuecore_load_all_gtap_port_group Data Source - gigavuecore"
subcategory: ""
description: |-
  load all port group details
---

# gigavuecore_load_all_gtap_port_group Data Source

load all port group details

## Example Usage

```terraform
data "gigavuecore_load_all_gtap_port_group" "example" {
  cluster_id = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `port_group_id` (String) - Port group Id
* `sfp` (Attributes) - Port group SFP compatible details (see [below for nested schema](#nestedatt--items--sfp))

<a id="nestedatt--items--sfp"></a>
### Nested Schema for `items.sfp`

Read-Only:

* `in_compatible` (Boolean) - SFP incompatible true/false
* `in_compatible_reason` (String) - Port group SFP incompatible reason

