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
  alias = null
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

* `context` (Object({page_no, page_size, sort, total_items}), computed) - Gigamon query result context
  * `page_no` (Number, computed) - page number of the returned result set
  * `page_size` (Number, computed) - page size of the returned result set
  * `sort` (List(String), computed) - sorting info of the returned result set. list of fields in the array indicate sorting order
  * `total_items` (Number, computed) - total number of items in the queried entity type
* `gtp_whitelists` (List(Object({alias, cluster_id, imsi_count})), computed)
* `id` (String, computed)

