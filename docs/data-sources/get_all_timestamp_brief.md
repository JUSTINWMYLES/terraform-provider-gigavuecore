---
page_title: "gigavuecore_get_all_timestamp_brief Data Source - gigavuecore"
subcategory: ""
description: |-
  Reads the get all timestamp brief data source.
---

# gigavuecore_get_all_timestamp_brief Data Source

Reads the get all timestamp brief data source.

## Example Usage

```terraform
data "gigavuecore_get_all_timestamp_brief" "example" {
  box_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `box_id` (Number, optional) - specify the cluster node by boxId. By default all nodes are selected.

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({admin, egress_src_id, egress_timestamp, ingress_src_id, ingress_timestamp, port_id})), computed)

