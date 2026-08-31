---
page_title: "gigavuecore_get_all_apps_listener List Resource - gigavuecore"
subcategory: ""
description: |-
  Get all Apps Listener
---

# gigavuecore_get_all_apps_listener List Resource

Get all Apps Listener

## Example Usage

```terraform
list "gigavuecore_get_all_apps_listener" "example" {
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

* `alias` (String, computed) - Alias of the listener


