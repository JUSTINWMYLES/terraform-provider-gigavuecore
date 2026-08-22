---
page_title: "gigavuecore_delete_text_config_file Action - gigavuecore"
subcategory: ""
description: |-
  Delete a system text configuration file
---

# gigavuecore_delete_text_config_file Action

Delete a system text configuration file

## Example Usage

```terraform
action "gigavuecore_delete_text_config_file" "example" {
  config {
    cluster_id = "example"
    filename = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `filename` (String, required) - filename of the text configuration file
