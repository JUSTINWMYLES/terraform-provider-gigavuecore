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
  email_address   = "example"
  enable_details  = true
  enable_failures = true
  enable_infos    = true
  page            = "example"
  sort            = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `email_address` (String, required) - Email address
* `enable_details` (Boolean, optional) - Get Details
* `enable_failures` (Boolean, optional) - Get Failures
* `enable_infos` (Boolean, optional) - Get Infos
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `enable_details` (Boolean, computed) - Get Details
* `enable_failures` (Boolean, computed) - Get Failures
* `enable_infos` (Boolean, computed) - Get Infos


## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_email_recipient.example {email_address}
```
