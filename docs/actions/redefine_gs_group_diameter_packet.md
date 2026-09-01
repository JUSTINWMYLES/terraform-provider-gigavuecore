---
page_title: "gigavuecore_redefine_gs_group_diameter_packet Action - gigavuecore"
subcategory: ""
description: |-
  Redefine GS Group's Diameter packet
---

# gigavuecore_redefine_gs_group_diameter_packet Action

Redefine GS Group's Diameter packet

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_redefine_gs_group_diameter_packet" "example" {
  config {
    alias   = "example"
    timeout = 1
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target GS Group
* `timeout` (Number, optional) - timeout in seconds to remove OOO or fragmented packet


