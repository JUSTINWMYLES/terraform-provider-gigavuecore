---
page_title: "gigavuecore_load_all_circuit_tunnels List Resource - gigavuecore"
subcategory: ""
description: |-
  Load all Circuit Tunnels
---

# gigavuecore_load_all_circuit_tunnels List Resource

Load all Circuit Tunnels

## Example Usage

```terraform
list "gigavuecore_load_all_circuit_tunnels" "example" {
  provider = gigavuecore
  limit    = 100
  config {
    cluster_id = "example"
    mode       = "example"
    page       = "example"
    sort       = "example"
    type       = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `mode` (String, optional) - Filter circuit tunnels by mode
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)
* `type` (String, optional) - Filter circuit tunnels by type


### Identity Attributes

The following identity attributes are exported for each matching result:

* `alias` (String, computed)


