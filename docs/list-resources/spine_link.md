---
page_title: "gigavuecore_spine_link List Resource - gigavuecore"
subcategory: ""
description: |-
  Load spine-link configuration
---

# gigavuecore_spine_link List Resource

Load spine-link configuration

-> **Note:** This list resource requires Terraform 1.14 or later and is used through the `terraform query` command, not in configuration files.

## Example Usage

```terraform
list "gigavuecore_spine_link" "example" {
  provider = gigavuecore
  limit    = 100
}
```
## Schema

### Identity Attributes

The following identity attributes are exported for each matching result:

* `alias` (String, computed)


