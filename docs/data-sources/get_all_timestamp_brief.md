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
  box_id = 0
}
```

## Schema

### Arguments

The following arguments are supported:

* `box_id` (Number, optional) - specify the cluster node by boxId. By default all nodes are selected.

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `admin` (Boolean)
* `egress_src_id` (Number)
* `egress_timestamp` (Boolean)
* `ingress_src_id` (Number)
* `ingress_timestamp` (Boolean)
* `port_id` (String)

