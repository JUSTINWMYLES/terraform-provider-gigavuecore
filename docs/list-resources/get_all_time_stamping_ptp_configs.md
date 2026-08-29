---
page_title: "gigavuecore_get_all_time_stamping_ptp_configs List Resource - gigavuecore"
subcategory: ""
description: |-
  Lists get all time stamping ptp configs resources.
---

# gigavuecore_get_all_time_stamping_ptp_configs List Resource

Lists get all time stamping ptp configs resources.

## Example Usage

```terraform
list "gigavuecore_get_all_time_stamping_ptp_configs" "example" {
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

* `alias` (String, computed)


