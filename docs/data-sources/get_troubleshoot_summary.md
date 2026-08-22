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
  alias = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Traffic flow alias

### Attributes

In addition to all arguments above, the following attributes are exported:

* `cluster_config_summaries` (List(Object({cluster_id, green_map_health_count, green_tunnel_health_count, hostname, map_count, red_map_health_count, red_tunnel_health_count, tunnel_count, yellow_map_health_count, yellow_tunnel_health_count})), computed)
* `cluster_links` (List(Object({destination_cluster, source_cluster})), computed)

