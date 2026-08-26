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
  page       = null
  sort       = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required)
* `cluster_id` (String, required) - id of the defining cluster
* `imsi_count` (Number, optional) - Number of IMSI entries in gtp-whitelist
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `id` (String, computed)
* `imsi_count` (Number, computed) - Number of IMSI entries in gtp-whitelist


