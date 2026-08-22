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
  alias = null
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

* `context` (Object({page_no, page_size, sort, total_items}), computed) - Gigamon query result context
  * `page_no` (Number, computed) - page number of the returned result set
  * `page_size` (Number, computed) - page size of the returned result set
  * `sort` (List(String), computed) - sorting info of the returned result set. list of fields in the array indicate sorting order
  * `total_items` (Number, computed) - total number of items in the queried entity type
* `diameter_whitelists` (List(Object({alias, user_name_count})), computed)
* `id` (String, computed)

