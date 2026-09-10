---
page_title: "gigavuecore_load_ip_destination_statuses Data Source - gigavuecore"
subcategory: ""
description: |-
  Load IP destination status of all or those matching the IP Interface alias or IP type or GsGroup
---

# gigavuecore_load_ip_destination_statuses Data Source

Load IP destination status of all or those matching the IP Interface alias or IP type or GsGroup

## Example Usage

```terraform
data "gigavuecore_load_ip_destination_statuses" "example" {
  cluster_id         = "example"
  gs_group_alias     = "example"
  ip_interface_alias = "example"
  ip_type            = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `gs_group_alias` (String, optional) - GsGroup alias
* `ip_interface_alias` (String, optional) - IP Interface alias
* `ip_type` (String, optional) - IP type

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `destination` (String) - IP Destination
* `gs_group` (String) - Associated GsGroup
* `interface_alias` (String) - Interface Alias
* `status` (String) - Status
* `te_id` (String) - Te-ID

