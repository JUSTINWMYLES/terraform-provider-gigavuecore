---
page_title: "gigavuecore_ptp_config List Resource - gigavuecore"
subcategory: ""
description: |-
  Lists get all time stamping ptp configs resources.
---

# gigavuecore_ptp_config List Resource

Lists get all time stamping ptp configs resources.

-> **Note:** This list resource requires Terraform 1.14 or later and is used through the `terraform query` command, not in configuration files.

## Example Usage

```terraform
list "gigavuecore_ptp_config" "example" {
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

* `alias` (String, computed) - alias of the time stamping ptp configuration


