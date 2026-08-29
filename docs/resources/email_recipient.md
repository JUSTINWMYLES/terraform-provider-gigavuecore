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

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

* `create` (String) - A create timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `read` (String) - A read timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).
* `update` (String) - A update timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `delete` (String) - A delete timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_email_recipient.example {email_address}
```
