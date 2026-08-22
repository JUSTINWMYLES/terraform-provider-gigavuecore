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
  env_id = null
  node_id = null
  unify_id = null
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

* `hosts` (Object({date, machine, nodename, number_of_avail_cpus, number_of_cpus, release, sysname}), computed) - Node sysstat hosts struct of a node in unified deployment
  * `date` (String, computed)
  * `machine` (String, computed)
  * `nodename` (String, computed)
  * `number_of_avail_cpus` (Number, computed)
  * `number_of_cpus` (Number, computed)
  * `release` (String, computed)
  * `sysname` (String, computed)

