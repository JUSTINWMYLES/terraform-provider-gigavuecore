---
page_title: "gigavuecore_redefine_gs_group_sip_ports Action - gigavuecore"
subcategory: ""
description: |-
  Redefine GS Group's Sip Ports
---

# gigavuecore_redefine_gs_group_sip_ports Action

Redefine GS Group's Sip Ports

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_redefine_gs_group_sip_ports" "example" {
  config {
    alias = "example"
    ports = [1]
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target GS Group
* `ports` (List of Number, optional) - list of TCP/UDP ports. Valid ports 1 - 65535. maximum 10 ports


