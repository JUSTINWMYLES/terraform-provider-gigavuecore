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
  alias = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, optional) - Filter by node alias

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `alias` (String) - Alias name for the GI node
* `created_date` (String) - Date when the node was created
* `ipv4_address` (String) - IPv4 address of the GI node
* `ipv6_address` (String) - IPv6 address of the GI node
* `last_config_updated_at` (String) - Date when the node was last online
* `last_status_updated_at` (String) - Date when the node was last updated
* `node_id` (String) - Unique identifier for the GI node
* `node_version` (String) - Version of the GigaInsight node software
* `prompt_bundle` (Attributes) - Prompt bundle information for the GI node (see [below for nested schema](#nestedatt--items--prompt_bundle))
* `status` (String) - Current status of the GI node
<a id="nestedatt--items--prompt_bundle"></a>
### Nested Schema for `items.prompt_bundle`

Read-Only:

* `last_updated_at` (String) - Date when the prompt bundle was last updated
* `prompt_bundle_version` (String) - Version of the prompt bundle

