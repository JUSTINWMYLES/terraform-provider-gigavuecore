---
page_title: "gigavuecore_activation Resource - gigavuecore"
subcategory: ""
description: |-
  Get license activations(for a specific entitlement,or for all entitlements)
---

# gigavuecore_activation Resource

Get license activations(for a specific entitlement,or for all entitlements)

## Example Usage

```terraform
resource "gigavuecore_activation" "example" {
  eli_id   = null
  quantity = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `eli_id` (String, optional)
* `quantity` (Number, optional)

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `activations` (Attributes List, computed) (see [below for nested schema](#nestedatt--activations))
* `context` (Attributes, computed) (see [below for nested schema](#nestedatt--context))
* `id` (String, computed)

<a id="nestedatt--activations"></a>
### Nested Schema for `activations`

Read-Only:

* `eid` (String)
* `entl_item_id` (String)
* `gid` (String)
* `imported` (Boolean)
* `num_licenses` (Number)
* `owner_fm` (Boolean)
* `vmac` (String)
<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `page_no` (Number)
* `page_size` (Number)
* `sort` (List of String)
* `total_items` (Number)

