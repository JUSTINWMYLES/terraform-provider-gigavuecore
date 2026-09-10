---
page_title: "gigavuecore_redefine_gs_group_rtp_ports Action - gigavuecore"
subcategory: ""
description: |-
  Redefine GS Group's Rtp Ports
---

# gigavuecore_redefine_gs_group_rtp_ports Action

Redefine GS Group's Rtp Ports

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_redefine_gs_group_rtp_ports" "example" {
  config {
    alias = "example"
    range = {
      port     = 0
      port_max = 0
    }
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target GS Group
* `range` (Attributes, optional) - GsGroup Rtp Port Range Parameters (see [below for nested schema](#nestedatt--range))

<a id="nestedatt--range"></a>
### Nested Schema for `range`

Required:

* `port` (Number)

Optional:

* `port_max` (Number) - If specified should be greater than 'port'

