---
page_title: "gigavuecore_get_node_sys_stats Data Source - gigavuecore"
subcategory: ""
description: |-
  Obtain a node sysStats of unified deployment node via environment, connection and node id
---

# gigavuecore_get_node_sys_stats Data Source

Obtain a node sysStats of unified deployment node via environment, connection and node id

## Example Usage

```terraform
data "gigavuecore_get_node_sys_stats" "example" {
  env_id   = "example"
  node_id  = "example"
  unify_id = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `env_id` (String, required) - unified environment identifier
* `node_id` (String, required) - unified resource node identifier
* `unify_id` (String, required) - unified resource identifier

### Attributes

In addition to all arguments above, the following attributes are exported:

* `hosts` (Attributes, computed) - Node sysstat hosts struct of a node in unified deployment (see [below for nested schema](#nestedatt--hosts))

<a id="nestedatt--hosts"></a>
### Nested Schema for `hosts`

Read-Only:

* `date` (String)
* `machine` (String)
* `nodename` (String)
* `number_of_avail_cpus` (Number)
* `number_of_cpus` (Number)
* `release` (String)
* `sysname` (String)

