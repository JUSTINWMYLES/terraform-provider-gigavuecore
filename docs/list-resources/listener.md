---
page_title: "gigavuecore_listener List Resource - gigavuecore"
subcategory: ""
description: |-
  Get all Apps Listener
---

# gigavuecore_listener List Resource

Get all Apps Listener

-> **Note:** This list resource requires Terraform 1.14 or later and is used through the `terraform query` command, not in configuration files.

## Example Usage

```terraform
list "gigavuecore_listener" "example" {
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


