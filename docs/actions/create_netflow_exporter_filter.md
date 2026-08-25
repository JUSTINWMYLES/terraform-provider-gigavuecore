---
page_title: "gigavuecore_create_netflow_exporter_filter Action - gigavuecore"
subcategory: ""
description: |-
  Create Netflow Exporter Filter
---

# gigavuecore_create_netflow_exporter_filter Action

Create Netflow Exporter Filter

## Example Usage

```terraform
action "gigavuecore_create_netflow_exporter_filter" "example" {
  config {
    alias      = "example"
    cluster_id = "example"
    rules      = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target Netflow Exporter
* `cluster_id` (String, required) - Target Cluster ID
* `rules` (List of Dynamic, required)


