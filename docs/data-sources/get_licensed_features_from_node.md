---
page_title: "gigavuecore_get_licensed_features_from_node Data Source - gigavuecore"
subcategory: ""
description: |-
  Get licensed features from the node (consider using newer preferred API /licensing/module/all)
---

# gigavuecore_get_licensed_features_from_node Data Source

Get licensed features from the node (consider using newer preferred API /licensing/module/all)

## Example Usage

```terraform
data "gigavuecore_get_licensed_features_from_node" "example" {
  cluster_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - name of the cluster

### Attributes

In addition to all arguments above, the following attributes are exported:

* `context` (Object({total_items}), computed)
  * `total_items` (Number, computed) - number of items in 'licenses' property below
* `licenses` (List(Object({box_id, features, slot_id})), computed)

