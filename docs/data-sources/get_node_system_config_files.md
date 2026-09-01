---
page_title: "gigavuecore_get_node_system_config_files Data Source - gigavuecore"
subcategory: ""
description: |-
  Load available system configuration files
---

# gigavuecore_get_node_system_config_files Data Source

Load available system configuration files

## Example Usage

```terraform
data "gigavuecore_get_node_system_config_files" "example" {
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

* `items` (Attributes List, computed) - list of the config files (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `active` (Boolean) - Is config active
* `filename` (String) - File name
* `modified` (Boolean) - Has unsaved config changes. True: need to save.

