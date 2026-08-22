---
page_title: "gigavuecore_get_all_ptp_ports_counters Data Source - gigavuecore"
subcategory: ""
description: |-
  Reads the get all ptp ports counters data source.
---

# gigavuecore_get_all_ptp_ports_counters Data Source

Reads the get all ptp ports counters data source.

## Example Usage

```terraform
data "gigavuecore_get_all_ptp_ports_counters" "example" {
  box_id = null
  cluster_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `box_id` (Number, optional) - specify the cluster node by boxId. By default all nodes are selected.
* `cluster_id` (String, required) - Target Cluster ID

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({discarded_packets, port_id, rx_packets, tx_packets})), computed)

