---
page_title: "gigavuecore_delete_all_apps_exporter_group Action - gigavuecore"
subcategory: ""
description: |-
  Delete all Apps Exporter Group
---

# gigavuecore_delete_all_apps_exporter_group Action

Delete all Apps Exporter Group

## Example Usage

```terraform
action "gigavuecore_delete_all_apps_exporter_group" "example" {
  config {
    cluster_id = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
