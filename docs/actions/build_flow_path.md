---
page_title: "gigavuecore_build_flow_path Action - gigavuecore"
subcategory: ""
description: |-
  Build traffic flow path for a cluster
---

# gigavuecore_build_flow_path Action

Build traffic flow path for a cluster

## Example Usage

```terraform
action "gigavuecore_build_flow_path" "example" {
  config {
    cluster_name = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `cluster_name` (String, required) - Target Cluster Name


