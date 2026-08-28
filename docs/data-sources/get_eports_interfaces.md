---
page_title: "gigavuecore_get_eports_interfaces Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all GigaSMART ports interface details
---

# gigavuecore_get_eports_interfaces Data Source

Load all GigaSMART ports interface details

## Example Usage

```terraform
data "gigavuecore_get_eports_interfaces" "example" {
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

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `cluster_id` (String) - id of the defining cluster
* `dhcp` (Boolean)
* `dns` (String) - Only valid and required when 'dhcp' is false
* `eport` (String) - GigaSMART engine port
* `gateway` (String) - Only valid and required when 'dhcp' is false
* `hw_address` (String)
* `interface` (String)
* `ip_address` (String) - Only valid and required when 'dhcp' is false
* `ip_mask` (String) - Only valid and required when 'dhcp' is false
* `mtu` (Number) - Only valid when 'dhcp' is false
* `proxy_server_ping_status` (String)
* `proxy_server_profile` (String)
* `status` (String)
* `vlan` (Number)

