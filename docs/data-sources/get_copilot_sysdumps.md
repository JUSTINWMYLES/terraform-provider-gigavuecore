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

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({filename, size, timestamp})), computed) - List of copilot sysdump files

