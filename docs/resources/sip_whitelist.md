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
  page  = null
  sort  = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Sip Whitelist alias
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `caller_id_count` (Number, computed) - Number of caller-id entries in sip-whitelist
* `id` (String, computed)


