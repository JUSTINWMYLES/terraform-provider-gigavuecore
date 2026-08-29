---
page_title: "gigavuecore_vxlan_group Resource - gigavuecore"
subcategory: ""
description: |-
  Load all Circuit Tunnel Vxlan Groups
---

# gigavuecore_vxlan_group Resource

Load all Circuit Tunnel Vxlan Groups

## Example Usage

```terraform
resource "gigavuecore_vxlan_group" "example" {
  alias      = "example"
  box_id     = "example"
  cluster_id = "example"
  comment    = "example"
  page       = "example"
  sort       = "example"
  vxlan_ids  = [ 0 ]
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required)
* `box_id` (String, optional) - device box id. valid range 1 - 64.
* `cluster_id` (String, required) - Target Cluster ID
* `comment` (String, optional)
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)
* `vxlan_ids` (List of Number, optional)

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `box_id` (String, computed) - device box id. valid range 1 - 64.
* `comment` (String, computed)
* `vxlan_ids` (List of Number, computed)


## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_vxlan_group.example {alias}
```
