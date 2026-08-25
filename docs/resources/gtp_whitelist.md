---
page_title: "gigavuecore_gtp_whitelist Resource - gigavuecore"
subcategory: ""
description: |-
  Load available GTP Whitelists
---

# gigavuecore_gtp_whitelist Resource

Load available GTP Whitelists

## Example Usage

```terraform
resource "gigavuecore_gtp_whitelist" "example" {
  alias      = null
  cluster_id = null
  imsi_count = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required)
* `cluster_id` (String, optional) - id of the defining cluster
* `imsi_count` (Number, optional) - Number of IMSI entries in gtp-whitelist

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `context` (Attributes, computed) - Gigamon query result context (see [below for nested schema](#nestedatt--context))
* `gtp_whitelists` (Attributes List, computed) (see [below for nested schema](#nestedatt--gtp_whitelists))
* `id` (String, computed)

<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `page_no` (Number) - page number of the returned result set
* `page_size` (Number) - page size of the returned result set
* `sort` (List of String) - sorting info of the returned result set. list of fields in the array indicate sorting order
* `total_items` (Number) - total number of items in the queried entity type
<a id="nestedatt--gtp_whitelists"></a>
### Nested Schema for `gtp_whitelists`

Read-Only:

* `alias` (String)
* `cluster_id` (String) - id of the defining cluster
* `imsi_count` (Number) - Number of IMSI entries in gtp-whitelist

