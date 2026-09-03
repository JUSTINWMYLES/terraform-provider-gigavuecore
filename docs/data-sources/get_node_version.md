---
page_title: "gigavuecore_get_node_version Data Source - gigavuecore"
subcategory: ""
description: |-
  Obtain a node version of unified deployment node via environment, connection and node id
---

# gigavuecore_get_node_version Data Source

Obtain a node version of unified deployment node via environment, connection and node id

## Example Usage

```terraform
data "gigavuecore_get_node_version" "example" {
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

* `build_info` (String, computed)
* `hostname` (String, computed)
* `mgmt_if` (String, computed)
* `platform` (String, computed)
* `state` (String, computed)
* `uuid` (String, computed)
* `version` (String, computed)


