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
  cluster_id         = null
  gs_group_alias     = null
  ip_interface_alias = null
  ip_type            = null
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

* `context` (Attributes, computed) - Gigamon query result context (see [below for nested schema](#nestedatt--context))
* `ip_interfaces` (Attributes List, computed) (see [below for nested schema](#nestedatt--ip_interfaces))

<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `page_no` (Number) - page number of the returned result set
* `page_size` (Number) - page size of the returned result set
* `sort` (List of String) - sorting info of the returned result set. list of fields in the array indicate sorting order
* `total_items` (Number) - total number of items in the queried entity type
<a id="nestedatt--ip_interfaces"></a>
### Nested Schema for `ip_interfaces`

Read-Only:

* `destination` (String) - IP Destination
* `gs_group` (String) - Associated GsGroup
* `interface_alias` (String) - Interface Alias
* `status` (String) - Status
* `te_id` (String) - Te-ID

