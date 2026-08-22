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
  page = null
  sort = null
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

* `items` (List(Object({cluster_id, dhcp, dns, eport, gateway, hw_address, interface, ip_address, ip_mask, mtu, proxy_server_ping_status, proxy_server_profile, status, vlan})), computed)

