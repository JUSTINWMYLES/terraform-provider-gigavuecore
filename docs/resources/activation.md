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
  eli_id   = "example"
  page     = "example"
  quantity = 0
  sort     = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `eli_id` (String, optional)
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `quantity` (Number, optional)
* `sort` (String, optional) - parentheses-enclosed pair of values in a (fieldName:sortOrder) format. The default sort order is DESC. Example: sort=(numLicenses:DESC)

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `eid` (String, computed)
* `entl_item_id` (String, computed)
* `gid` (String, computed)
* `imported` (Boolean, computed)
* `num_licenses` (Number, computed)
* `owner_fm` (Boolean, computed)
* `vmac` (String, computed)

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

* `create` (String) - A create timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `read` (String) - A read timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).
* `update` (String) - A update timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `delete` (String) - A delete timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).

