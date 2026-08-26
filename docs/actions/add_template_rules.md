---
page_title: "gigavuecore_add_template_rules Action - gigavuecore"
subcategory: ""
description: |-
  Add template rules to a traffic flow
---

# gigavuecore_add_template_rules Action

Add template rules to a traffic flow

## Example Usage

```terraform
action "gigavuecore_add_template_rules" "example" {
  config {
    alias          = "example"
    cluster_id     = "example"
    node_id        = "example"
    template_alias = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Traffic flow alias
* `cluster_id` (String, optional) - Target Cluster ID
* `node_id` (String, optional) - Node ID for cluster resolution
* `template_alias` (String, required) - Template alias to apply


