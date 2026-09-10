---
page_title: "gigavuecore_load_inline_ssl_cert_valid_record Data Source - gigavuecore"
subcategory: ""
description: |-
  Get Inline SSL certificate validation record
---

# gigavuecore_load_inline_ssl_cert_valid_record Data Source

Get Inline SSL certificate validation record

## Example Usage

```terraform
data "gigavuecore_load_inline_ssl_cert_valid_record" "example" {
  fingerprint = "example"
  page        = "example"
  sort        = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `fingerprint` (String, required) - The fingerprint of the record to lookup
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `cert_cn` (String)
* `expiry` (String)
* `revocation_status` (String)
* `sha1` (String)

