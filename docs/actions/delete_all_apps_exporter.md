---
page_title: "gigavuecore_delete_all_apps_exporter Action - gigavuecore"
subcategory: ""
description: |-
  new in H 5.8
---

# gigavuecore_delete_all_apps_exporter Action

new in H 5.8

## Example Usage

```terraform
action "gigavuecore_delete_all_apps_exporter" "example" {
  config {
    cluster_id = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
