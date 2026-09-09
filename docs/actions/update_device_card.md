---
page_title: "gigavuecore_update_device_card Action - gigavuecore"
subcategory: ""
description: |-
  Update Device Card configuration
---

# gigavuecore_update_device_card Action

Update Device Card configuration

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_update_device_card" "example" {
  config {
    admin_status           = "up"
    alarm_buffer_threshold = 0
    body_slot_id           = "example"
    cluster_id             = "example"
    fabric_hash_adv        = true
    filter_template        = "example"
    mode                   = "32x"
    node_id                = "example"
    pld_upgrade            = true
    slot_id                = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `admin_status` (String, optional) - up: bring card up; down: shutdown the card
* `alarm_buffer_threshold` (Number, optional) - card micro burst threshold
* `body_slot_id` (String, required) - device card slot id. used to identify target card. not updatable
* `cluster_id` (String, required) - Target Cluster ID. Either 'clusterId' or 'nodeId' is required
* `fabric_hash_adv` (Boolean, optional) - Advanced Fabric Hash. Supported only for Q02X32/Q08 cards
* `filter_template` (String, optional) - alias of filter template or 'defaults'.'defaults' is special alias for predefined filter templates
* `mode` (String, optional)
* `node_id` (String, optional) - ID of the target device
* `pld_upgrade` (Boolean, optional) - Upgrade PLD image
* `slot_id` (String, required) - Device card slot ID


