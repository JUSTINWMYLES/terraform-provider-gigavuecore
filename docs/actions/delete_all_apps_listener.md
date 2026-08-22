---
page_title: "gigavuecore_delete_all_apps_listener Action - gigavuecore"
subcategory: ""
description: |-
  Delete all Apps Listener
---

# gigavuecore_delete_all_apps_listener Action

Delete all Apps Listener

## Example Usage

```terraform
action "gigavuecore_delete_all_apps_listener" "example" {
  config {
    cluster_id = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
