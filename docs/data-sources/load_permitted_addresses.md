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
  address      = "example"
  address_type = "example"
  alias        = "example"
  page         = "example"
  sort         = "example"
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

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `address` (String) - Permitted address
* `address_type` (String) - Address type
* `alias` (String) - Alias for the address
* `subscribed_events` (List of String) - List of events subscribed by the address

