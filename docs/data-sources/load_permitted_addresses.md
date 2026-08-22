---
page_title: "gigavuecore_load_permitted_addresses Data Source - gigavuecore"
subcategory: ""
description: |-
  Load Permitted Addresses
---

# gigavuecore_load_permitted_addresses Data Source

Load Permitted Addresses

## Example Usage

```terraform
data "gigavuecore_load_permitted_addresses" "example" {
  address = null
  address_type = null
  alias = null
  page = null
  sort = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `address` (String, optional) - address
* `address_type` (String, optional) - addressType
* `alias` (String, optional) - alias
* `page` (String, optional) - page
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({address, address_type, alias, subscribed_events})), computed)

