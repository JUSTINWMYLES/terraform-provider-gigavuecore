---
page_title: "gigavuecore_delete_filter_templates Action - gigavuecore"
subcategory: ""
description: |-
  Delete all Filter Templates
---

# gigavuecore_delete_filter_templates Action

Delete all Filter Templates

## Example Usage

```terraform
action "gigavuecore_delete_filter_templates" "example" {
  config {
    cluster_id = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
