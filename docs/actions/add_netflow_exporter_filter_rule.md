---
page_title: "gigavuecore_add_netflow_exporter_filter_rule Action - gigavuecore"
subcategory: ""
description: |-
  Add a new rule to a Netflow Exporter Filter
---

# gigavuecore_add_netflow_exporter_filter_rule Action

Add a new rule to a Netflow Exporter Filter

## Example Usage

```terraform
action "gigavuecore_add_netflow_exporter_filter_rule" "example" {
  config {
    alias      = "example"
    cluster_id = "example"
    matches    = null
    rule_id    = 1
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target Netflow Exporter
* `cluster_id` (String, required) - Target Cluster ID
* `matches` (List of Dynamic, optional) - Set of rule's matching elements. Within a rule, matching elements of the the same type MAY be used more than once, However, their matching positions MUST be unique
* `rule_id` (Number, required)


