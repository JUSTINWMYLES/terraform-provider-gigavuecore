---
page_title: "gigavuecore_sysdump List Resource - gigavuecore"
subcategory: ""
description: |-
  Load Available Sysdump Files
---

# gigavuecore_sysdump List Resource

Load Available Sysdump Files

-> **Note:** This list resource requires Terraform 1.14 or later and is used through the `terraform query` command, not in configuration files.

## Example Usage

```terraform
list "gigavuecore_sysdump" "example" {
  provider = gigavuecore
  limit    = 100
}
```
## Schema

### Identity Attributes

The following identity attributes are exported for each matching result:

* `filename` (String, computed) - The sysdump filename


