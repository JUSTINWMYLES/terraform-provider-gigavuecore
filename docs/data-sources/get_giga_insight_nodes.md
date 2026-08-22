---
page_title: "gigavuecore_get_giga_insight_nodes Data Source - gigavuecore"
subcategory: ""
description: |-
  Get all GigaInsight Nodes
---

# gigavuecore_get_giga_insight_nodes Data Source

Get all GigaInsight Nodes

## Example Usage

```terraform
data "gigavuecore_get_giga_insight_nodes" "example" {
  alias = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, optional) - Filter by node alias

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({alias, created_date, ipv4_address, ipv6_address, last_config_updated_at, last_status_updated_at, node_id, node_version, prompt_bundle, status})), computed)

