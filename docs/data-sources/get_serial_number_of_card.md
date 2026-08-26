---
page_title: "gigavuecore_get_serial_number_of_card Data Source - gigavuecore"
subcategory: ""
description: |-
  Gets the serial numbers of card specified by input parameters
---

# gigavuecore_get_serial_number_of_card Data Source

Gets the serial numbers of card specified by input parameters

## Example Usage

```terraform
data "gigavuecore_get_serial_number_of_card" "example" {
  box_slash_slot_id = null
  cluster_name      = null
  node_id           = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `box_slash_slot_id` (String, required) - <box-id>/<slot-id>
* `cluster_name` (String, required) - name of chassis cluster
* `node_id` (String, required) - node ID, the IP address of the chassis


