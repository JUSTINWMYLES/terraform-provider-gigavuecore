---
page_title: "gigavuecore_flex_inline List Resource - gigavuecore"
subcategory: ""
description: |-
  Get All flexInline solutions present on FM
---

# gigavuecore_flex_inline List Resource

Get All flexInline solutions present on FM

-> **Note:** This list resource requires Terraform 1.14 or later and is used through the `terraform query` command, not in configuration files.

## Example Usage

```terraform
list "gigavuecore_flex_inline" "example" {
  provider = gigavuecore
  limit    = 100
}
```
## Schema

### Identity Attributes

The following identity attributes are exported for each matching result:

* `alias` (String, computed) - Alias of the solution


