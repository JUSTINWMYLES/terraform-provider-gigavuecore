---
page_title: "gigavuecore_redefine_gs_group_generic_session_timeout Action - gigavuecore"
subcategory: ""
description: |-
  Redefine GS Group's Session Timeout
---

# gigavuecore_redefine_gs_group_generic_session_timeout Action

Redefine GS Group's Session Timeout

## Example Usage

```terraform
action "gigavuecore_redefine_gs_group_generic_session_timeout" "example" {
  config {
    alias = "example"
    time  = 1
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target GS Group
* `time` (Number, optional) - Maximum timeout for session entry


