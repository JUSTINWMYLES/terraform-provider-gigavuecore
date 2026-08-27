---
page_title: "gigavuecore_get_all_ptp_clock_states List Resource - gigavuecore"
subcategory: ""
description: |-
  Lists get all ptp clock states resources.
---

# gigavuecore_get_all_ptp_clock_states List Resource

Lists get all ptp clock states resources.

## Example Usage

```terraform
list "gigavuecore_get_all_ptp_clock_states" "example" {
  provider = gigavuecore
  limit    = 100
  config {
    box_id = 1
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `box_id` (Number, optional) - specify the cluster node by boxId. By default all nodes are selected.


### Identity Attributes

The following identity attributes are exported for each matching result:

* `alias` (String, computed)


