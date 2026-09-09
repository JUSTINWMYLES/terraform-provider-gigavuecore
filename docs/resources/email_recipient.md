---
page_title: "gigavuecore_email_recipient Resource - gigavuecore"
subcategory: ""
description: |-
  Create email recipients
---

# gigavuecore_email_recipient Resource

Create email recipients

~> **Note:** The update operation for this resource is not wired to a remote API endpoint: the API spec exposes no usable update mapping. Changing any configuration attribute triggers a resource replacement (all config-settable attributes carry a RequiresReplace plan modifier); create, read, and delete remain functional.

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

* `create` (Number) - A create timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `read` (Number) - A read timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `update` (Number) - An update timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `delete` (Number) - A delete timeout in seconds for this operation. Overrides the generator default (1200 seconds).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_email_recipient.example {email_address}
```
