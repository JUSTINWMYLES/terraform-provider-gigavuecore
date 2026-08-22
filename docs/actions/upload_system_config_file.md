---
page_title: "gigavuecore_upload_system_config_file Action - gigavuecore"
subcategory: ""
description: |-
  Upload a configuration file from local file.
---

# gigavuecore_upload_system_config_file Action

Upload a configuration file from local file.

## Example Usage

```terraform
action "gigavuecore_upload_system_config_file" "example" {
  config {
    cluster_id = "example"
    config = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `config` (String, required) - User uploaded file
