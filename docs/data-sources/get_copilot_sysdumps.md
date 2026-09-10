---
page_title: "gigavuecore_get_copilot_sysdumps Data Source - gigavuecore"
subcategory: ""
description: |-
  Load Available Sysdump Files
---

# gigavuecore_get_copilot_sysdumps Data Source

Load Available Sysdump Files

## Example Usage

```terraform
data "gigavuecore_get_copilot_sysdumps" "example" {
}
```

## Schema

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) - List of copilot sysdump files (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `filename` (String) - The sysdump filename
* `size` (Number) - The sysdump file size in bytes
* `timestamp` (String) - File creation date and time in ISO 8601 format

