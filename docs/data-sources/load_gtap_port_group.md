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
  cluster_id = null
  port_group_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `port_group_id` (String, required) - Port group Id

### Attributes

In addition to all arguments above, the following attributes are exported:

* `sfp` (Object({in_compatible, in_compatible_reason}), computed) - Port group SFP compatible details
  * `in_compatible` (Bool, computed) - SFP incompatible true/false
  * `in_compatible_reason` (String, computed) - Port group SFP incompatible reason

