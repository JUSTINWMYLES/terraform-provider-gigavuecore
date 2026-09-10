---
page_title: "gigavuecore_get_ptp_counters_by_alias Data Source - gigavuecore"
subcategory: ""
description: |-
  Reads the get ptp counters by alias data source.
---

# gigavuecore_get_ptp_counters_by_alias Data Source

Reads the get ptp counters by alias data source.

## Example Usage

```terraform
data "gigavuecore_get_ptp_counters_by_alias" "example" {
  alias      = "example"
  box_id     = "example"
  cluster_id = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the time stamping PTP configuration
* `box_id` (String, optional) - specify the cluster node by boxId. By default all nodes are selected.
* `cluster_id` (String, required) - Target Cluster ID

### Attributes

In addition to all arguments above, the following attributes are exported:

* `discarded_packets` (Number, computed)
* `ipv4_ptp_rx_packets` (Number, computed)
* `ipv6_ptp_rx_packets` (Number, computed)
* `l2_ptp_rx_packets` (Number, computed)
* `queue_overflow_rx_packets` (Number, computed)
* `rcpu_encap_rx_packets` (Number, computed)
* `rx_packets` (Number, computed)
* `tx_packets` (Number, computed)
* `udp_ptp_rx_packets` (Number, computed)


