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

* `alias` (String) - Icap Server Alias
* `cluster_id` (String) - id of the defining cluster
* `comment` (String) - Icap Server Comment
* `l3_address` (String) - Icap Server IP Address
* `l4_port` (Number) - Icap Server l4 Port
* `options_service_url` (String) - Options Service URL
* `reqmod_service_url` (String) - Request Modification Service URL
* `respmod_service_url` (String) - Response Modification Service URL

