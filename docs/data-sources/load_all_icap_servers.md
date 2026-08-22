---
page_title: "gigavuecore_load_all_icap_servers Data Source - gigavuecore"
subcategory: ""
description: |-
  Load All ICAP Servers
---

# gigavuecore_load_all_icap_servers Data Source

Load All ICAP Servers

## Example Usage

```terraform
data "gigavuecore_load_all_icap_servers" "example" {
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

* `items` (List(Object({alias, cluster_id, comment, l3_address, l4_port, options_service_url, reqmod_service_url, respmod_service_url})), computed)

