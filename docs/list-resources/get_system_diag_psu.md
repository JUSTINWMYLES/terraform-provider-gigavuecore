---
page_title: "gigavuecore_get_system_diag_psu List Resource - gigavuecore"
subcategory: ""
description: |-
  get system diagnostics PSU information
---

# gigavuecore_get_system_diag_psu List Resource

get system diagnostics PSU information

## Example Usage

```terraform
list "gigavuecore_get_system_diag_psu" "example" {
  provider = gigavuecore
  limit = 100
}

```

## Schema

### Arguments

The following arguments are supported:


### Identity Attributes

The following identity attributes are exported for each matching result:

* `slot_id` (String, computed)
