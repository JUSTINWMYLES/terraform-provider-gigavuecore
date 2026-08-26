---
page_title: "gigavuecore_load_ip_destination_status Data Source - gigavuecore"
subcategory: ""
description: |-
  Load IP destination status of the IP Address specified
---

# gigavuecore_load_ip_destination_status Data Source

Load IP destination status of the IP Address specified

## Example Usage

```terraform
data "gigavuecore_load_ip_destination_status" "example" {
  cluster_id = null
  ip_address = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `ip_address` (String, required) - IP Address

### Attributes

In addition to all arguments above, the following attributes are exported:

* `destination` (String, computed) - IP Destination
* `gs_group` (String, computed) - Associated GsGroup
* `interface_alias` (String, computed) - Interface Alias
* `status` (String, computed) - Status
* `te_id` (String, computed) - Te-ID


