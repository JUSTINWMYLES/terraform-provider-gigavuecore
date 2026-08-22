---
page_title: "gigavuecore_update_syslog_config Action - gigavuecore"
subcategory: ""
description: |-
  Update Syslog Config
---

# gigavuecore_update_syslog_config Action

Update Syslog Config

## Example Usage

```terraform
action "gigavuecore_update_syslog_config" "example" {
  config {
    cluster_id = "example"
    cluster_name = "example"
    syslog_config_list = null
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `cluster_name` (String, optional) - Name of the cluster
* `syslog_config_list` (List(Dynamic), optional) - List of the syslog configuration specification for every node in the cluster
