---
page_title: "gigavuecore_redefine_netflow_exporter_filter Action - gigavuecore"
subcategory: ""
description: |-
  Redefine Netflow Exporter Filter
---

# gigavuecore_redefine_netflow_exporter_filter Action

Redefine Netflow Exporter Filter

## Example Usage

```terraform
action "gigavuecore_redefine_netflow_exporter_filter" "example" {
  config {
    alias = "example"
    cluster_id = "example"
    rules = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target Netflow Exporter
* `cluster_id` (String, required) - Target Cluster ID
* `rules` (List(Dynamic), required)
