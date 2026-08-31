---
page_title: "gigavuecore_redefine_gs_group_load_balance_params Action - gigavuecore"
subcategory: ""
description: |-
  Redefine GS Group's Load Balancing Params
---

# gigavuecore_redefine_gs_group_load_balance_params Action

Redefine GS Group's Load Balancing Params

## Example Usage

```terraform
action "gigavuecore_redefine_gs_group_load_balance_params" "example" {
  config {
    alias      = "example"
    cluster_id = "example"
    failover = {
      enabled               = true
      threshold_lt_bw       = 50
      threshold_lt_pkt_rate = 500
    }
    link_weight_type = "speed"
    replicate_gtpc   = true
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target GS Group
* `cluster_id` (String, required) - Target Cluster ID
* `failover` (Attributes, optional) - Private class. Failover part of the GsGroup Load Balancing Parameters (see [below for nested schema](#nestedatt--failover))
* `link_weight_type` (String, optional) - Controls LB port weight resolution for 'weighted' load balancing (wtRoundRobin, wtLeastBw, wtLeastPktRate, wtLeastConn, wtLeastTotalTraffic,wtImsi, wtSupi). For 'assigned', explicit user-configured port weights on the port group are used. For 'speed', link speed is used and user-assigned port weights are ignored. Ex: if a port group consists of four tool ports, and one of them is 100G and the others are 10G, the 100G link will be selected about 10 times more than the 10G links
* `replicate_gtpc` (Boolean, optional) - Enables replication of GTP control packets (GTP-c)

<a id="nestedatt--failover"></a>
### Nested Schema for `failover`

Optional:

* `enabled` (Boolean) - Enables or disables failover when tool ports are down or thresholds to other tool ports in the load balancing port group are exceeded. A GigaSMART application failover will occur no more than once in 30 seconds
* `threshold_lt_bw` (Number) - Mutually exclusive with 'thresholdLtPktRate'. Failover threshold for Least Bandwidth. In percent's of the maximum bandwidth of a tool port. Ex: for a 1G port, a failover threshold of 90% means that failover to another tool port occurs when the bandwidth reaches 900Mbps. The default is 80%
* `threshold_lt_pkt_rate` (Number) - Mutually exclusive with 'thresholdLtBw'. Failover threshold for Least Packet Rate. In In kilo-packets-per-second: \[500k..5M\] pps. A tool port will failover to another tool port when the packet rate is over the specified threshold, in packets per second. The default is 1M

