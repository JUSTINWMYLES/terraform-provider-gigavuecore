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
  box_id     = 0
  cluster_id = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `box_id` (Number, optional) - specify the cluster node by boxId. By default all nodes are selected.
* `cluster_id` (String, required) - Target Cluster ID

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `ipv4_ptp_rx_packets` (Number)
* `ipv6_ptp_rx_packets` (Number)
* `l2_ptp_rx_packets` (Number)
* `queue_overflow_rx_packets` (Number)
* `rcpu_encap_rx_packets` (Number)
* `udp_ptp_rx_packets` (Number)
* `alias` (String)
* `box_id` (String)
* `discarded_packets` (Number)
* `rx_packets` (Number)
* `tx_packets` (Number)

