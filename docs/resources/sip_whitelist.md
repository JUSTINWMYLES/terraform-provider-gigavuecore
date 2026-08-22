---
page_title: "gigavuecore_sip_whitelist Resource - gigavuecore"
subcategory: ""
description: |-
  Load available SIP Whitelists
---

# gigavuecore_sip_whitelist Resource

Load available SIP Whitelists

## Example Usage

```terraform
resource "gigavuecore_sip_whitelist" "example" {
  alias = null
  caller_id_count = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Sip Whitelist alias
* `caller_id_count` (Number, optional) - Number of caller-id entries in sip-whitelist

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `context` (Object({page_no, page_size, sort, total_items}), computed) - Gigamon query result context
  * `page_no` (Number, computed) - page number of the returned result set
  * `page_size` (Number, computed) - page size of the returned result set
  * `sort` (List(String), computed) - sorting info of the returned result set. list of fields in the array indicate sorting order
  * `total_items` (Number, computed) - total number of items in the queried entity type
* `id` (String, computed)
* `sip_whitelists` (List(Object({alias, caller_id_count})), computed)

