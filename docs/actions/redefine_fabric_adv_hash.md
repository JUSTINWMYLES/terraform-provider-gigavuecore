---
page_title: "gigavuecore_redefine_fabric_adv_hash Action - gigavuecore"
subcategory: ""
description: |-
  Redefine fabric advanced hash
---

# gigavuecore_redefine_fabric_adv_hash Action

Redefine fabric advanced hash

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_redefine_fabric_adv_hash" "example" {
  config {
    box_id = 0
    fields = ["ethertype"]
    type   = "all"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `box_id` (Number, required) - box id
* `fields` (Set of String, optional) - Required when 'type' == 'fields'. 'gtpteid' requires fields (portsrc and portdst) or (port6src and port6dst)
* `type` (String, required)


