---
page_title: "gigavuecore_redefine_gs_group_sip_tcp_idle_timeout Action - gigavuecore"
subcategory: ""
description: |-
  Redefine GS Group's Sip Tcp Idle Timeout
---

# gigavuecore_redefine_gs_group_sip_tcp_idle_timeout Action

Redefine GS Group's Sip Tcp Idle Timeout

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_redefine_gs_group_sip_tcp_idle_timeout" "example" {
  config {
    alias = "example"
    time  = 20
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target GS Group
* `time` (Number, optional) - Sip tcp idle timeout value in seconds .Valid values: 20-600.


