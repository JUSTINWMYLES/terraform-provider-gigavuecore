---
page_title: "gigavuecore_get_ptp_ports_counters_by_port_id Data Source - gigavuecore"
subcategory: ""
description: |-
  Reads the get ptp ports counters by port id data source.
---

# gigavuecore_get_ptp_ports_counters_by_port_id Data Source

Reads the get ptp ports counters by port id data source.

## Example Usage

```terraform
data "gigavuecore_get_ptp_ports_counters_by_port_id" "example" {
  cluster_id = "example"
  port_id    = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `port_id` (String, required) - Port ID in \[box/slot/portid\] format

### Attributes

In addition to all arguments above, the following attributes are exported:

* `discarded_packets` (Number, computed)
* `rx_packets` (Number, computed)
* `tx_packets` (Number, computed)


