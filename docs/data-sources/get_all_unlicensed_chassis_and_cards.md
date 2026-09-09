---
page_title: "gigavuecore_get_all_unlicensed_chassis_and_cards Data Source - gigavuecore"
subcategory: ""
description: |-
  Get all the chassis and cards that need GVOS and GVOS-module license respectively
---

# gigavuecore_get_all_unlicensed_chassis_and_cards Data Source

Get all the chassis and cards that need GVOS and GVOS-module license respectively

## Example Usage

```terraform
data "gigavuecore_get_all_unlicensed_chassis_and_cards" "example" {
  use_db = true
}
```

## Schema

### Arguments

The following arguments are supported:

* `use_db` (Boolean, optional) - if true, use FM database as information provider, else retrieve from node

### Attributes

In addition to all arguments above, the following attributes are exported:

* `unlicensed_cards` (List of String, computed) - list of cards that are missing GVOS-module license
* `unlicensed_chassis` (List of String, computed) - list of chassis that are missing GVOS license


