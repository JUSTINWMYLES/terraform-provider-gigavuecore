---
page_title: "gigavuecore_load_cards_details List Resource - gigavuecore"
subcategory: ""
description: |-
  Load all Device Cards details
---

# gigavuecore_load_cards_details List Resource

Load all Device Cards details

## Example Usage

```terraform
list "gigavuecore_load_cards_details" "example" {
  provider = gigavuecore
  limit    = 100
  config {
    cluster_id = "example"
    node_id    = "example"
    page       = "example"
    sort       = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, optional) - Target Cluster ID. Either 'clusterId' or 'nodeId' is required
* `node_id` (String, optional) - ID of the target device. Either 'clusterId' or 'nodeId' is required
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)


### Identity Attributes

The following identity attributes are exported for each matching result:

* `slot_id` (String, computed)


