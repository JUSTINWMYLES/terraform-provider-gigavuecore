---
page_title: "gigavuecore_get_all_ptp_ports_counters List Resource - gigavuecore"
subcategory: ""
description: |-
  Lists get all ptp ports counters resources.
---

# gigavuecore_get_all_ptp_ports_counters List Resource

Lists get all ptp ports counters resources.

## Example Usage

```terraform
list "gigavuecore_get_all_ptp_ports_counters" "example" {
  provider = gigavuecore
  limit    = 100
  config {
    box_id     = 0
    cluster_id = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `box_id` (Number, optional) - specify the cluster node by boxId. By default all nodes are selected.
* `cluster_id` (String, required) - Target Cluster ID


### Identity Attributes

The following identity attributes are exported for each matching result:

* `port_id` (String, computed)


