---
page_title: "gigavuecore_get_all_tunnel_application Data Source - gigavuecore"
subcategory: ""
description: |-
  Get all Tunnel Apps
---

# gigavuecore_get_all_tunnel_application Data Source

Get all Tunnel Apps

## Example Usage

```terraform
data "gigavuecore_get_all_tunnel_application" "example" {
  cluster_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID

### Attributes

In addition to all arguments above, the following attributes are exported:

* `context` (Object({page_no, page_size, sort, total_items}), computed) - Gigamon query result context
  * `page_no` (Number, computed) - page number of the returned result set
  * `page_size` (Number, computed) - page size of the returned result set
  * `sort` (List(String), computed) - sorting info of the returned result set. list of fields in the array indicate sorting order
  * `total_items` (Number, computed) - total number of items in the queried entity type
* `tunnel_apps` (List(Object({alias, cluster_id, config_status, config_status_reasons, decap, description, dscp, encap, gsgroup, ip_interface_in, ip_interface_out, local_key_alias, remote_key_alias, ssl_profile, ssl_profile_alias, tcp_profile, tcp_profile_alias, traffic_dir, ttl, tunnel_type})), computed)

