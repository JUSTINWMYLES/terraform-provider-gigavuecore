---
page_title: "gigavuecore_gta_profile List Resource - gigavuecore"
subcategory: ""
description: |-
  get all gta profiles
---

# gigavuecore_gta_profile List Resource

get all gta profiles

-> **Note:** This list resource requires Terraform 1.14 or later and is used through the `terraform query` command, not in configuration files.

## Example Usage

```terraform
list "gigavuecore_gta_profile" "example" {
  provider = gigavuecore
  limit    = 100
}
```
## Schema

### Identity Attributes

The following identity attributes are exported for each matching result:

* `alias` (String, computed) - alias of the GTA profile


