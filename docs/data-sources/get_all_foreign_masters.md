---
page_title: "gigavuecore_get_all_foreign_masters Data Source - gigavuecore"
subcategory: ""
description: |-
  deprecated: use GET /ptp/portState/foreignSource
---

# gigavuecore_get_all_foreign_masters Data Source

deprecated: use GET /ptp/portState/foreignSource

## Example Usage

```terraform
data "gigavuecore_get_all_foreign_masters" "example" {
  box_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `box_id` (Number, optional) - specify the cluster node by boxId. By default all nodes are selected.

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({best_master, foreign_masters, port_id})), computed)

