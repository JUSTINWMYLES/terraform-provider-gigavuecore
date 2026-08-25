---
page_title: "gigavuecore_load_inline_ssl_url_record Data Source - gigavuecore"
subcategory: ""
description: |-
  Get Inline SSL URL cache record
---

# gigavuecore_load_inline_ssl_url_record Data Source

Get Inline SSL URL cache record

## Example Usage

```terraform
data "gigavuecore_load_inline_ssl_url_record" "example" {
  domain_pattern = null
  page           = null
  sort           = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `domain_pattern` (String, required) - The domain name of the record to lookup
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `context` (Attributes, computed) - Gigamon query result context (see [below for nested schema](#nestedatt--context))
* `records` (Attributes List, computed) (see [below for nested schema](#nestedatt--records))

<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `page_no` (Number) - page number of the returned result set
* `page_size` (Number) - page size of the returned result set
* `sort` (List of String) - sorting info of the returned result set. list of fields in the array indicate sorting order
* `total_items` (Number) - total number of items in the queried entity type
<a id="nestedatt--records"></a>
### Nested Schema for `records`

Read-Only:

* `added_on` (String)
* `categories` (List of String)
* `domain` (String)
* `expiry` (String)

