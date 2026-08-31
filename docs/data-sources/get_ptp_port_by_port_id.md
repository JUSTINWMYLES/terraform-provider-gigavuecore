---
page_title: "gigavuecore_get_ptp_port_by_port_id Data Source - gigavuecore"
subcategory: ""
description: |-
  Reads the get ptp port by port id data source.
---

# gigavuecore_get_ptp_port_by_port_id Data Source

Reads the get ptp port by port id data source.

## Example Usage

```terraform
data "gigavuecore_get_ptp_port_by_port_id" "example" {
  port_id = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `port_id` (String, required) - Port ID in \[box/slot/portid\] format

### Attributes

In addition to all arguments above, the following attributes are exported:

* `announce_interval` (Number, computed) - Logarithm to the base 2 of the mean announce Interval
* `announce_receipt_timeout` (Number, computed) - Specify the number of announceInterval that has to pass without receipt of an Announce message before the occurrence of the expire event
* `clock_identity` (String, computed) - Clock Identifier as per the IEEE 1588 standard which is represented by an integer 8-octet array
* `clock_port_id` (Number, computed) - The value of the portNumber for a port on a PTP node supporting a single PTP port shall be 1. The values of the port numbers for the N ports on a PTP node supporting N PTP ports shall be 1, 2, ...N, respectively. The all-zeros and all-ones portNumber values are reserved.
* `clock_port_state` (String, computed) - Value of the current state of the protocol engine associated with the given PTP port. (deprecated: use clockPortStateAlias)
* `clock_port_state_alias` (String, computed) - Value of the current state of the protocol engine associated with the given PTP port
* `delay_mechanism` (String, computed) - The propagation delay measuring option used by the port in computing mean path delay
* `delay_request_interval` (Number, computed) - Logarithm to the base 2 of the permitted mean time interval between successive Delay\_req messages
* `enabled` (Boolean, computed)
* `local_priority` (Number, computed)
* `peer_delay_request_interval` (Number, computed) - Logarithm to the base 2 of the delayReqInterval
* `peer_mean_path_delay` (Number, computed) - If the value of the delay Mechanism member is peer-to-peer (P2P), the value of peerMeanPathDelay shall be an estimate of the current one-way propagation delay on the link, i.e.,meanPathDelay, attached to this port computed using the peer delay mechanism
* `ptp_version` (String, computed) - Indicates the version of the PTP standard implemented on the port. IEEE Std 1588-2008 corresponds to PTP version 2 whereas 1588-2002 is PTP version 1. For this implementation, PTPv2 will be supported
* `sync_interval` (Number, computed) - Logarithm to the base 2 of the mean SyncInterval for multicast messages
* `vlan_id` (Number, computed)


