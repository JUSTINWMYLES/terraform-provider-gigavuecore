---
page_title: "gigavuecore_load_diameter_whitelist_entry Data Source - gigavuecore"
subcategory: ""
description: |-
  Check if Whitelist Entry exists
---

# gigavuecore_load_diameter_whitelist_entry Data Source

Check if Whitelist Entry exists

## Example Usage

```terraform
data "gigavuecore_load_diameter_whitelist_entry" "example" {
  alias     = "example"
  user_name = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target Diameter Whitelist
* `user_name` (String, required) - imsi-based whitelist entry being queried

### Attributes

In addition to all arguments above, the following attributes are exported:

* `active_sessions` (Number, computed) - Number of active sessions


