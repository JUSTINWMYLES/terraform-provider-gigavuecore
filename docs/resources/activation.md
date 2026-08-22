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
  eli_id = null
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

* `activations` (List(Object({eid, entl_item_id, gid, imported, num_licenses, owner_fm, vmac})), computed)
* `context` (Object({page_no, page_size, sort, total_items}), computed)
  * `page_no` (Number, computed)
  * `page_size` (Number, computed)
  * `sort` (List(String), computed)
  * `total_items` (Number, computed)
* `id` (String, computed)

