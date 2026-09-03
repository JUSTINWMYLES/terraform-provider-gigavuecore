---
page_title: "gigavuecore_redefine_gs_group_sip_media Action - gigavuecore"
subcategory: ""
description: |-
  Redefine GS Group's Sip Media
---

# gigavuecore_redefine_gs_group_sip_media Action

Redefine GS Group's Sip Media

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_redefine_gs_group_sip_media" "example" {
  config {
    alias   = "example"
    timeout = 30
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target GS Group
* `timeout` (Number, optional) - Sip media timeout value in seconds .Valid values: 30-300.


