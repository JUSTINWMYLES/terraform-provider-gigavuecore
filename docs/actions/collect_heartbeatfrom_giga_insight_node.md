---
page_title: "gigavuecore_collect_heartbeatfrom_giga_insight_node Action - gigavuecore"
subcategory: ""
description: |-
  Accept Heartbeat from Gigamon Insights in FM
---

# gigavuecore_collect_heartbeatfrom_giga_insight_node Action

Accept Heartbeat from Gigamon Insights in FM

## Example Usage

```terraform
action "gigavuecore_collect_heartbeatfrom_giga_insight_node" "example" {
  config {
    body_node_id = "example"
    node_id = "example"
    node_version = "example"
    prompt_bundle = null
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `body_node_id` (String, optional) - The node identifier
* `node_id` (String, required) - The unique identifier of the GigaInsight Node
* `node_version` (String, optional) - The GigaInsight node version
* `prompt_bundle` (Dynamic, optional) - The prompt bundle information installed on the node
