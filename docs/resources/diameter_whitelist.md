---
page_title: "gigavuecore_diameter_whitelist Resource - gigavuecore"
subcategory: ""
description: |-
  Load available Diameter Whitelists
---

# gigavuecore_diameter_whitelist Resource

Load available Diameter Whitelists

## Example Usage

```terraform
resource "gigavuecore_diameter_whitelist" "example" {
  alias           = null
  user_name_count = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required)
* `user_name_count` (Number, optional) - Number of username entries in diameter-whitelist

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `context` (Attributes, computed) - Gigamon query result context (see [below for nested schema](#nestedatt--context))
* `diameter_whitelists` (Attributes List, computed) (see [below for nested schema](#nestedatt--diameter_whitelists))
* `id` (String, computed)

<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `page_no` (Number) - page number of the returned result set
* `page_size` (Number) - page size of the returned result set
* `sort` (List of String) - sorting info of the returned result set. list of fields in the array indicate sorting order
* `total_items` (Number) - total number of items in the queried entity type
<a id="nestedatt--diameter_whitelists"></a>
### Nested Schema for `diameter_whitelists`

Read-Only:

* `alias` (String)
* `user_name_count` (Number) - Number of username entries in diameter-whitelist

