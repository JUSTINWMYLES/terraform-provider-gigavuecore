---
page_title: "gigavuecore_delete_all_header_strip Action - gigavuecore"
subcategory: ""
description: |-
  Delete header strip for all boxes
---

# gigavuecore_delete_all_header_strip Action

Delete header strip for all boxes

## Example Usage

```terraform
action "gigavuecore_delete_all_header_strip" "example" {
  config {
    cluster_id = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target cluster ID.
