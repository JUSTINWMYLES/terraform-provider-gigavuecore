---
page_title: "gigavuecore_app_visibility List Resource - gigavuecore"
subcategory: ""
description: |-
  Load all Application Intelligence Solutions deployed on V Series
---

# gigavuecore_app_visibility List Resource

Load all Application Intelligence Solutions deployed on V Series

-> **Note:** This list resource requires Terraform 1.14 or later and is used through the `terraform query` command, not in configuration files.

## Example Usage

```terraform
list "gigavuecore_app_visibility" "example" {
  provider = gigavuecore
  limit    = 100
}
```
## Schema

### Identity Attributes

The following identity attributes are exported for each matching result:

* `solution_alias` (String, computed) - Alias of the solution


