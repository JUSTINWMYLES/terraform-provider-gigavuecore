---
page_title: "gigavuecore_update_device_card Action - gigavuecore"
subcategory: ""
description: |-
  Update Device Card configuration
---

# gigavuecore_update_device_card Action

Update Device Card configuration

## Example Usage

```terraform
action "gigavuecore_update_device_card" "example" {
  config {
    admin_status = "example"
    alarm_buffer_threshold = 1
    body_slot_id = "example"
    cluster_id = "example"
    fabric_hash_adv = true
    filter_template = "example"
    mode = "example"
    node_id = "example"
    pld_upgrade = true
    slot_id = "example"
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
* `fabric_hash_adv` (Bool, optional) - Advanced Fabric Hash. Supported only for Q02X32/Q08 cards
* `filter_template` (String, optional) - alias of filter template or 'defaults'.'defaults' is special alias for predefined filter templates
* `mode` (String, optional)
* `node_id` (String, optional) - ID of the target device
* `pld_upgrade` (Bool, optional) - Upgrade PLD image
* `slot_id` (String, required) - Device card slot ID
