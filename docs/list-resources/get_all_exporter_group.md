---
page_title: "gigavuecore_get_all_exporter_group List Resource - gigavuecore"
subcategory: ""
description: |-
  Get all Apps Exporter Group
---

# gigavuecore_get_all_exporter_group List Resource

Get all Apps Exporter Group

## Example Usage

```terraform
list "gigavuecore_get_all_exporter_group" "example" {
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

* `cluster_id` (String, required) - Target Cluster ID


### Identity Attributes

The following identity attributes are exported for each matching result:

* `alias` (String, computed)


