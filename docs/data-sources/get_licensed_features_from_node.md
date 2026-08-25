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

* `context` (Attributes, computed) (see [below for nested schema](#nestedatt--context))
* `licenses` (Attributes List, computed) (see [below for nested schema](#nestedatt--licenses))

<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `total_items` (Number) - number of items in 'licenses' property below
<a id="nestedatt--licenses"></a>
### Nested Schema for `licenses`

Read-Only:

* `box_id` (Number)
* `features` (Attributes) (see [below for nested schema](#nestedatt--licenses--features))
* `slot_id` (Number)
<a id="nestedatt--licenses--features"></a>
### Nested Schema for `licenses.features`

Read-Only:

* `apf` (Attributes) (see [below for nested schema](#nestedatt--licenses--features--apf))
* `erspan` (Attributes) (see [below for nested schema](#nestedatt--licenses--features--erspan))
* `masking` (Attributes) (see [below for nested schema](#nestedatt--licenses--features--masking))
<a id="nestedatt--licenses--features--apf"></a>
### Nested Schema for `licenses.features.apf`

Read-Only:

* `expiry` (String)
<a id="nestedatt--licenses--features--erspan"></a>
### Nested Schema for `licenses.features.erspan`

Read-Only:

* `expiry` (String)
<a id="nestedatt--licenses--features--masking"></a>
### Nested Schema for `licenses.features.masking`

Read-Only:

* `expiry` (String)

