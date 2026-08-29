---
page_title: "gigavuecore_get_all_ptp_ports List Resource - gigavuecore"
subcategory: ""
description: |-
  Lists get all ptp ports resources.
---

# gigavuecore_get_all_ptp_ports List Resource

Lists get all ptp ports resources.

## Example Usage

```terraform
list "gigavuecore_get_all_ptp_ports" "example" {
  provider = gigavuecore
  limit    = 100
  config {
    box_id = 0
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `box_id` (Number, optional) - specify the cluster node by boxId. By default all nodes are selected.


### Identity Attributes

The following identity attributes are exported for each matching result:

* `port_id` (String, computed)


