---
page_title: "gigavuecore_get_all_gta_profiles List Resource - gigavuecore"
subcategory: ""
description: |-
  get all gta profiles
---

# gigavuecore_get_all_gta_profiles List Resource

get all gta profiles

## Example Usage

```terraform
list "gigavuecore_get_all_gta_profiles" "example" {
  provider = gigavuecore
  limit    = 100
}

```
## Schema

### Identity Attributes

The following identity attributes are exported for each matching result:

* `alias` (String, computed) - alias of the GTA profile


