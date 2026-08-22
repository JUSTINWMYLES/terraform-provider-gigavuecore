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
    alias = "example"
    cluster_id = "example"
    failover = null
    link_weight_type = "example"
    replicate_gtpc = true
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target GS Group
* `cluster_id` (String, required) - Target Cluster ID
* `failover` (Dynamic, optional) - Private class. Failover part of the GsGroup Load Balancing Parameters
* `link_weight_type` (String, optional) - Controls LB port weight resolution for 'weighted' load balancing (wtRoundRobin, wtLeastBw, wtLeastPktRate, wtLeastConn, wtLeastTotalTraffic,wtImsi, wtSupi). For 'assigned', explicit user-configured port weights on the port group are used. For 'speed', link speed is used and user-assigned port weights are ignored. Ex: if a port group consists of four tool ports, and one of them is 100G and the others are 10G, the 100G link will be selected about 10 times more than the 10G links
* `replicate_gtpc` (Bool, optional) - Enables replication of GTP control packets (GTP-c)
