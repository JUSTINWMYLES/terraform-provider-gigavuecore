---
page_title: "gigavuecore_get_tools Data Source - gigavuecore"
subcategory: ""
description: |-
  Get Tools
---

# gigavuecore_get_tools Data Source

Get Tools

## Example Usage

```terraform
data "gigavuecore_get_tools" "example" {
  node_alias = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `node_alias` (String, optional) - manual node alias

### Attributes

In addition to all arguments above, the following attributes are exported:

* `alias` (String, computed)
* `tools_info` (List(Object({comment, compression_ratio, giga_streams, is_tool, is_used_in_deployed_policy, max_throughput, model, node_alias, ports, topo_node_id, total_storage, type, vendor})), computed)
* `type` (String, computed)

