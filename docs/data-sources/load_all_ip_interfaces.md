---
page_title: "gigavuecore_load_all_ip_interfaces Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all IP Interfaces
---

# gigavuecore_load_all_ip_interfaces Data Source

Load all IP Interfaces

## Example Usage

```terraform
data "gigavuecore_load_all_ip_interfaces" "example" {
  cluster_id = null
  page       = null
  sort       = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)

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

* `alias` (String) - ip interface name
* `attach` (List of String) - network ports ,tool ports or circuit ports
* `comment` (String)
* `gateway` (String) - gateway ipv4 or ipv6 address
* `gs_groups` (List of String) - Gs Groups associated with the IP Interface
* `hw_address` (String)
* `ip_address` (String) - ipv4/ipv6 address
* `ip_mask` (String) - ipAddress netmask required with ipAddress
* `ip_type` (String)
* `mtu` (Number)
* `netflow_exporters` (List of String) - Netflow Exporters associated with the IP Interface
* `tags` (Attributes List) (see [below for nested schema](#nestedatt--ip_interfaces--tags))
<a id="nestedatt--ip_interfaces--tags"></a>
### Nested Schema for `ip_interfaces.tags`

Read-Only:

* `tag_key` (String) - Name of the tag
* `tag_values` (List of String) - All possible values of the tag

