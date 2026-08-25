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
  email_address   = null
  enable_details  = null
  enable_failures = null
  enable_infos    = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `email_address` (String, required) - Email address
* `enable_details` (Boolean, optional) - Get Details
* `enable_failures` (Boolean, optional) - Get Failures
* `enable_infos` (Boolean, optional) - Get Infos

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `context` (Attributes, computed) - Gigamon query result context (see [below for nested schema](#nestedatt--context))
* `id` (String, computed)
* `recipients` (Attributes List, computed) (see [below for nested schema](#nestedatt--recipients))

<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `page_no` (Number) - page number of the returned result set
* `page_size` (Number) - page size of the returned result set
* `sort` (List of String) - sorting info of the returned result set. list of fields in the array indicate sorting order
* `total_items` (Number) - total number of items in the queried entity type
<a id="nestedatt--recipients"></a>
### Nested Schema for `recipients`

Read-Only:

* `email_address` (String) - Email address
* `enable_details` (Boolean) - Get Details
* `enable_failures` (Boolean) - Get Failures
* `enable_infos` (Boolean) - Get Infos

