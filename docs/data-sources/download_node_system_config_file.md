---
page_title: "gigavuecore_download_node_system_config_file Data Source - gigavuecore"
subcategory: ""
description: |-
  Download a system configuration file
---

# gigavuecore_download_node_system_config_file Data Source

Download a system configuration file

## Example Usage

```terraform
data "gigavuecore_download_node_system_config_file" "example" {
  cluster_id = "example"
  filename   = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `filename` (String, required) - filename of the target config file


