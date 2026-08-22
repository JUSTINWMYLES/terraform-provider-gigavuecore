---
page_title: "gigavuecore_redefine_gs_group_sip_session Action - gigavuecore"
subcategory: ""
description: |-
  Redefine GS Group's Sip Session
---

# gigavuecore_redefine_gs_group_sip_session Action

Redefine GS Group's Sip Session

## Example Usage

```terraform
action "gigavuecore_redefine_gs_group_sip_session" "example" {
  config {
    alias = "example"
    timeout = 1
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target GS Group
* `timeout` (Number, optional) - Sip session inactivity timer, value in seconds .Valid values: 30-300.
