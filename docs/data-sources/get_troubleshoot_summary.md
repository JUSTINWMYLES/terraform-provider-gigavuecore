---
page_title: "gigavuecore_get_troubleshoot_summary Data Source - gigavuecore"
subcategory: ""
description: |-
  Retrieve troubleshoot summary for a traffic flow
---

# gigavuecore_get_troubleshoot_summary Data Source

Retrieve troubleshoot summary for a traffic flow

## Example Usage

```terraform
data "gigavuecore_get_troubleshoot_summary" "example" {
  alias = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Traffic flow alias

### Attributes

In addition to all arguments above, the following attributes are exported:

* `cluster_config_summaries` (Attributes List, computed) (see [below for nested schema](#nestedatt--cluster_config_summaries))
* `cluster_links` (Attributes List, computed) (see [below for nested schema](#nestedatt--cluster_links))

<a id="nestedatt--cluster_config_summaries"></a>
### Nested Schema for `cluster_config_summaries`

Read-Only:

* `cluster_id` (String)
* `green_map_health_count` (Number)
* `green_tunnel_health_count` (Number)
* `hostname` (String)
* `map_count` (Number)
* `red_map_health_count` (Number)
* `red_tunnel_health_count` (Number)
* `tunnel_count` (Number)
* `yellow_map_health_count` (Number)
* `yellow_tunnel_health_count` (Number)
<a id="nestedatt--cluster_links"></a>
### Nested Schema for `cluster_links`

Read-Only:

* `destination_cluster` (String)
* `source_cluster` (String)

