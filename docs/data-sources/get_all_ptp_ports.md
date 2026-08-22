---
page_title: "gigavuecore_get_all_ptp_ports Data Source - gigavuecore"
subcategory: ""
description: |-
  Reads the get all ptp ports data source.
---

# gigavuecore_get_all_ptp_ports Data Source

Reads the get all ptp ports data source.

## Example Usage

```terraform
data "gigavuecore_get_all_ptp_ports" "example" {
  box_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `box_id` (Number, optional) - specify the cluster node by boxId. By default all nodes are selected.

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({announce_interval, announce_receipt_timeout, clock_identity, clock_port_id, clock_port_state, clock_port_state_alias, delay_mechanism, delay_request_interval, enabled, local_priority, peer_delay_request_interval, peer_mean_path_delay, port_id, ptp_version, sync_interval, vlan_id})), computed)

