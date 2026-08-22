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
  cluster_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({port_group_id, sfp})), computed)

