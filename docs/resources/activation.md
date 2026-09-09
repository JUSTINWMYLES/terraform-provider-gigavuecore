---
page_title: "gigavuecore_activation Resource - gigavuecore"
subcategory: ""
description: |-
  Create a new activation from an entitlement
---

# gigavuecore_activation Resource

Create a new activation from an entitlement

~> **Note:** The update operation for this resource is not wired to a remote API endpoint: the API spec exposes no usable update mapping. Changing any configuration attribute triggers a resource replacement (all config-settable attributes carry a RequiresReplace plan modifier); create, read, and delete remain functional.

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
* `aid` (String, computed)

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

* `create` (Number) - A create timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `read` (Number) - A read timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `update` (Number) - An update timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `delete` (Number) - A delete timeout in seconds for this operation. Overrides the generator default (1200 seconds).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_activation.example {eli_id}
```
