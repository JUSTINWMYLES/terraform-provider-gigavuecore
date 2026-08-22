---
page_title: "gigavuecore_update_traffic_flows_draft Action - gigavuecore"
subcategory: ""
description: |-
  Update an existing traffic flows draft
---

# gigavuecore_update_traffic_flows_draft Action

Update an existing traffic flows draft

## Example Usage

```terraform
action "gigavuecore_update_traffic_flows_draft" "example" {
  config {
    alias = "example"
    body_alias = "example"
    comment = "example"
    deployment_type = "example"
    enable = true
    flows = "example"
    has_draft = true
    priority_type = "example"
    sources_and_rules = "example"
    tags = null
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Traffic flow alias
* `body_alias` (String, required) - unique map alias
* `comment` (String, optional)
* `deployment_type` (String, required) - 'regular' maps are from network/hybrid ports to tool/hybrid/gigastream; 'inline' maps are from inline ports to inline/tool/hybrid/gigastream; 'firstLevel' are from network/hybrid ports to vPorts/tool/hybrid/gigastream; 'secondLevel' maps are from vPorts to tool/hybrid/gigastream; 'inlineFirstLevel is from inline network to vport; 'inlineSecondLevel' is from vport to inline tool; 'transitLevel' is from vport to vport
* `enable` (Bool, optional) - enable/disable map, applicable only to first level maps
* `flows` (List(Dynamic), required)
* `has_draft` (Bool, optional)
* `priority_type` (String, optional) - Define the map priority to be LOWEST or HIGHEST. Default priority is LOWEST.
* `sources_and_rules` (List(Dynamic), required)
* `tags` (List(Dynamic), optional)
