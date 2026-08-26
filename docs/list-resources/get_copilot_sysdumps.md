---
page_title: "gigavuecore_get_copilot_sysdumps List Resource - gigavuecore"
subcategory: ""
description: |-
  Load Available Sysdump Files
---

# gigavuecore_get_copilot_sysdumps List Resource

Load Available Sysdump Files

## Example Usage

```terraform
list "gigavuecore_get_copilot_sysdumps" "example" {
  provider = gigavuecore
  limit    = 100
}

```
## Schema

### Identity Attributes

The following identity attributes are exported for each matching result:

* `filename` (String, computed)


