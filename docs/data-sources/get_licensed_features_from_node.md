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
  cluster_id = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - name of the cluster

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `box_id` (Number)
* `features` (Attributes) (see [below for nested schema](#nestedatt--items--features))
* `slot_id` (Number)

<a id="nestedatt--items--features"></a>
### Nested Schema for `items.features`

Read-Only:

* `apf` (Attributes) (see [below for nested schema](#nestedatt--items--features--apf))
* `erspan` (Attributes) (see [below for nested schema](#nestedatt--items--features--erspan))
* `masking` (Attributes) (see [below for nested schema](#nestedatt--items--features--masking))

<a id="nestedatt--items--features--apf"></a>
### Nested Schema for `items.features.apf`

Read-Only:

* `expiry` (String)

<a id="nestedatt--items--features--erspan"></a>
### Nested Schema for `items.features.erspan`

Read-Only:

* `expiry` (String)

<a id="nestedatt--items--features--masking"></a>
### Nested Schema for `items.features.masking`

Read-Only:

* `expiry` (String)

