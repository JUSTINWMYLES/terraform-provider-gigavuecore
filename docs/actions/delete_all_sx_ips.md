---
page_title: "gigavuecore_delete_all_sx_ips Action - gigavuecore"
subcategory: ""
description: |-
  new in H 5.8
---

# gigavuecore_delete_all_sx_ips Action

new in H 5.8

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_delete_all_sx_ips" "example" {
  config {
    alias        = "example"
    ip_interface = "example"
    port         = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the sffp profile
* `ip_interface` (String, required) - IP Interface address
* `port` (String, required) - Port or Port List separated by commas


