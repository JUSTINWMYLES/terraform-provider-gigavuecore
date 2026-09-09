---
page_title: "gigavuecore_load_gtap_port_group Data Source - gigavuecore"
subcategory: ""
description: |-
  load port group details
---

# gigavuecore_load_gtap_port_group Data Source

load port group details

## Example Usage

```terraform
data "gigavuecore_load_gtap_port_group" "example" {
  cluster_id    = "example"
  port_group_id = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `port_group_id` (String, required) - Port group Id

### Attributes

In addition to all arguments above, the following attributes are exported:

* `sfp` (Attributes, computed) - Port group SFP compatible details (see [below for nested schema](#nestedatt--sfp))

<a id="nestedatt--sfp"></a>
### Nested Schema for `sfp`

Read-Only:

* `in_compatible` (Boolean) - SFP incompatible true/false
* `in_compatible_reason` (String) - Port group SFP incompatible reason

