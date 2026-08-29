---
page_title: "gigavuecore_delete_all_sx_ips Action - gigavuecore"
subcategory: ""
description: |-
  new in H 5.8
---

# gigavuecore_delete_all_sx_ips Action

new in H 5.8

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


