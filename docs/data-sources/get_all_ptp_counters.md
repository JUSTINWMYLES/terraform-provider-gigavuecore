---
page_title: "gigavuecore_get_all_ptp_counters Data Source - gigavuecore"
subcategory: ""
description: |-
  Reads the get all ptp counters data source.
---

# gigavuecore_get_all_ptp_counters Data Source

Reads the get all ptp counters data source.

## Example Usage

```terraform
data "gigavuecore_get_all_ptp_counters" "example" {
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

* `items` (List(Object({ipv4_ptp_rx_packets, ipv6_ptp_rx_packets, l2_ptp_rx_packets, queue_overflow_rx_packets, rcpu_encap_rx_packets, udp_ptp_rx_packets, alias, box_id, discarded_packets, rx_packets, tx_packets})), computed)

