---
page_title: "gigavuecore_mobility List Resource - gigavuecore"
subcategory: ""
description: |-
  Load all mobility solutions configured
---

# gigavuecore_mobility List Resource

Load all mobility solutions configured

-> **Note:** This list resource requires Terraform 1.14 or later and is used through the `terraform query` command, not in configuration files.

## Example Usage

```terraform
list "gigavuecore_mobility" "example" {
  provider = gigavuecore
  limit    = 100
}
```
## Schema

### Identity Attributes

The following identity attributes are exported for each matching result:

* `solution_alias` (String, computed) - Alias of the solution


