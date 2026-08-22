---
page_title: "gigavuecore_get_all_allocation_maps Data Source - gigavuecore"
subcategory: ""
description: |-
  Get inventory of licenses assigned on the physical devices monitored by the FM
---

# gigavuecore_get_all_allocation_maps Data Source

Get inventory of licenses assigned on the physical devices monitored by the FM

## Example Usage

```terraform
data "gigavuecore_get_all_allocation_maps" "example" {
  encode = null
  exclude_host_name = null
  exclude_ip = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `encode` (Bool, optional) - encode the JSON output
* `exclude_host_name` (String, optional) - exclude the host name in the inventory output
* `exclude_ip` (String, optional) - exclude the host IP in the inventory output

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({chassis, deployment_mac, slot_to_licenses})), computed)

