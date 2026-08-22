---
page_title: "gigavuecore_get_node_system_config_text_file Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all text configuration files available on device
---

# gigavuecore_get_node_system_config_text_file Data Source

Load all text configuration files available on device

## Example Usage

```terraform
data "gigavuecore_get_node_system_config_text_file" "example" {
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

* `items` (List(Object({filename, size, timestamp})), computed) - List of available text configuration files

