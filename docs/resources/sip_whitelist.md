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
  alias           = null
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

* `context` (Attributes, computed) - Gigamon query result context (see [below for nested schema](#nestedatt--context))
* `id` (String, computed)
* `sip_whitelists` (Attributes List, computed) (see [below for nested schema](#nestedatt--sip_whitelists))

<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `page_no` (Number) - page number of the returned result set
* `page_size` (Number) - page size of the returned result set
* `sort` (List of String) - sorting info of the returned result set. list of fields in the array indicate sorting order
* `total_items` (Number) - total number of items in the queried entity type
<a id="nestedatt--sip_whitelists"></a>
### Nested Schema for `sip_whitelists`

Read-Only:

* `alias` (String) - Sip Whitelist alias
* `caller_id_count` (Number) - Number of caller-id entries in sip-whitelist

