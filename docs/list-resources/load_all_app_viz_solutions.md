---
page_title: "gigavuecore_load_all_app_viz_solutions List Resource - gigavuecore"
subcategory: ""
description: |-
  Load all Apps Visibility Solutions
---

# gigavuecore_load_all_app_viz_solutions List Resource

Load all Apps Visibility Solutions

## Example Usage

```terraform
list "gigavuecore_load_all_app_viz_solutions" "example" {
  provider = gigavuecore
  limit    = 100
  config {
    cluster_id = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - clusterId


### Identity Attributes

The following identity attributes are exported for each matching result:

* `solution_alias` (String, computed) - user defined application visibility solution alias


