---
page_title: "gigavuecore_get_tunnel_logical_group_stats Data Source - gigavuecore"
subcategory: ""
description: |-
  Get Stats for Tunnel Logical Groups
---

# gigavuecore_get_tunnel_logical_group_stats Data Source

Get Stats for Tunnel Logical Groups

## Example Usage

```terraform
data "gigavuecore_get_tunnel_logical_group_stats" "example" {
  alias = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Tunnel Logical Group Alias

### Attributes

In addition to all arguments above, the following attributes are exported:

* `deviation` (Dynamic, computed) - Traffic Deviation in Percentage
* `last_successful_attempt_time` (String, computed) - Last Stats Polled Timestamp UTC Format
* `octets_rx` (Dynamic, computed) - Decapsulation end Octets Tx
* `octets_tx` (Dynamic, computed) - Encapsulation end Octets Tx
* `packets_rx` (Dynamic, computed) - Decapsulation end Packets Tx
* `packets_tx` (Dynamic, computed) - Encapsulation end Packets Tx
* `processed_traffic_percentage` (Dynamic, computed) - Amount of Traffic Processed based on Encap Packets Tx and Decap Packets Rx


