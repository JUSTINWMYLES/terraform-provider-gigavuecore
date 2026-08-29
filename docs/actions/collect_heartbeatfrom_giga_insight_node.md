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
    node_id      = "example"
    node_version = "example"
    prompt_bundle = {
      prompt_bundle_last_updated_at = "example"
      prompt_bundle_version         = "example"
    }
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `body_node_id` (String, optional) - The node identifier
* `node_id` (String, required) - The unique identifier of the GigaInsight Node
* `node_version` (String, optional) - The GigaInsight node version
* `prompt_bundle` (Attributes, optional) - The prompt bundle information installed on the node (see [below for nested schema](#nestedatt--prompt_bundle))

<a id="nestedatt--prompt_bundle"></a>
### Nested Schema for `prompt_bundle`

Optional:

* `prompt_bundle_last_updated_at` (String) - The date and time when the prompt bundle was last updated on the node
* `prompt_bundle_version` (String) - The prompt bundle version installed on the node

