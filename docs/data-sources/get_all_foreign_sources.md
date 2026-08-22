---
page_title: "gigavuecore_get_all_foreign_sources Data Source - gigavuecore"
subcategory: ""
description: |-
  Reads the get all foreign sources data source.
---

# gigavuecore_get_all_foreign_sources Data Source

Reads the get all foreign sources data source.

## Example Usage

```terraform
data "gigavuecore_get_all_foreign_sources" "example" {
  box_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `box_id` (Number, optional) - specify the cluster node by boxId. By default all nodes are selected.

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({best_source, foreign_sources, port_id})), computed)

