---
page_title: "gigavuecore_load_gdp_node_reports Data Source - gigavuecore"
subcategory: ""
description: |-
  Get GDP Neighbor Report from every Node in a cluster
---

# gigavuecore_load_gdp_node_reports Data Source

Get GDP Neighbor Report from every Node in a cluster

## Example Usage

```terraform
data "gigavuecore_load_gdp_node_reports" "example" {
  cluster_id = "example"
  page       = "example"
  sort       = "example"
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

* `gdp_neighbors` (Attributes List) - list of discovered GDP neighbors discovered on local ports (see [below for nested schema](#nestedatt--items--gdp_neighbors))
* `local_chassis_id` (String) - ChassisId of the reporting local node
<a id="nestedatt--items--gdp_neighbors"></a>
### Nested Schema for `items.gdp_neighbors`

Read-Only:

* `hostname` (String) - Hostname of the GDP neighbor node
* `local_port_id` (String) - PortId of the local port reporting the GDP Neighbor
* `local_port_type` (String) - Port type of the local port reporting the GDP Neighbor
* `mgmt_address` (String) - Management Address of the GDP neighbor node
* `product_type` (String) - Product type of the GDP neighbor node
* `remote_chassis_id` (String) - ChassisId of the attached GDP neighbor node
* `remote_port_id` (String) - PortId of the remote port on the attached GDP neighbor node
* `remote_port_type` (String) - Port type of the remote port on the attached GDP neighbor node
* `serial_number` (String) - Serial Number of the GDP neighbor node

