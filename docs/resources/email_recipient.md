---
page_title: "gigavuecore_email_recipient Resource - gigavuecore"
subcategory: ""
description: |-
  Get recipients
---

# gigavuecore_email_recipient Resource

Get recipients

## Example Usage

```terraform
resource "gigavuecore_email_recipient" "example" {
  email_address = null
  enable_details = null
  enable_failures = null
  enable_infos = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `email_address` (String, required) - Email address
* `enable_details` (Bool, optional) - Get Details
* `enable_failures` (Bool, optional) - Get Failures
* `enable_infos` (Bool, optional) - Get Infos

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `context` (Object({page_no, page_size, sort, total_items}), computed) - Gigamon query result context
  * `page_no` (Number, computed) - page number of the returned result set
  * `page_size` (Number, computed) - page size of the returned result set
  * `sort` (List(String), computed) - sorting info of the returned result set. list of fields in the array indicate sorting order
  * `total_items` (Number, computed) - total number of items in the queried entity type
* `id` (String, computed)
* `recipients` (List(Object({email_address, enable_details, enable_failures, enable_infos})), computed)

