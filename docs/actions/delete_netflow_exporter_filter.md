---
page_title: "gigavuecore_delete_netflow_exporter_filter Action - gigavuecore"
subcategory: ""
description: |-
  Delete Netflow Exporter Filter
---

# gigavuecore_delete_netflow_exporter_filter Action

Delete Netflow Exporter Filter

## Example Usage

```terraform
action "gigavuecore_delete_netflow_exporter_filter" "example" {
  config {
    alias = "example"
    cluster_id = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target Netflow Exporter
* `cluster_id` (String, required) - Target Cluster ID
