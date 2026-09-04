---
page_title: "gigavuecore_delete_sip_whitelist_entry Action - gigavuecore"
subcategory: ""
description: |-
  Delete SIP Whitelist Entry
---

# gigavuecore_delete_sip_whitelist_entry Action

Delete SIP Whitelist Entry

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_delete_sip_whitelist_entry" "example" {
  config {
    alias      = "example"
    caller_id  = "example"
    id_range   = "example"
    ip_address = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target SIP Whitelist
* `caller_id` (String, optional) - callerId-based whitelist entry being deleted. required till H 5.6
* `id_range` (String, optional) - idrange based whitelist entry being deleted. Example:idRange=110..120
* `ip_address` (String, optional) - ipAddress based whitelist entry being deleted. Example:ipAddress=1.1.1.1


