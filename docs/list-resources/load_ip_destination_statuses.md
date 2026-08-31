---
page_title: "gigavuecore_load_ip_destination_statuses List Resource - gigavuecore"
subcategory: ""
description: |-
  Load IP destination status of all or those matching the IP Interface alias or IP type or GsGroup
---

# gigavuecore_load_ip_destination_statuses List Resource

Load IP destination status of all or those matching the IP Interface alias or IP type or GsGroup

## Example Usage

```terraform
list "gigavuecore_load_ip_destination_statuses" "example" {
  provider = gigavuecore
  limit    = 100
  config {
    cluster_id         = "example"
    gs_group_alias     = "example"
    ip_interface_alias = "example"
    ip_type            = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `gs_group_alias` (String, optional) - GsGroup alias
* `ip_interface_alias` (String, optional) - IP Interface alias
* `ip_type` (String, optional) - IP type


### Identity Attributes

The following identity attributes are exported for each matching result:

* `ip_address` (String, computed)


