---
page_title: "gigavuecore_redefine_syslog_config Action - gigavuecore"
subcategory: ""
description: |-
  Redefine Syslog config
---

# gigavuecore_redefine_syslog_config Action

Redefine Syslog config

## Example Usage

```terraform
action "gigavuecore_redefine_syslog_config" "example" {
  config {
    cluster_id         = "example"
    cluster_name       = "example"
    syslog_config_list = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `cluster_name` (String, required) - Name of the cluster
* `syslog_config_list` (List of Dynamic, required) - List of the syslog configuration specification for every node in the cluster


