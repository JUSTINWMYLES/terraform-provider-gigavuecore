---
page_title: "gigavuecore_get_all_giga_flex_inline_network_group List Resource - gigavuecore"
subcategory: ""
description: |-
  Get All FM Inline Network Groups
---

# gigavuecore_get_all_giga_flex_inline_network_group List Resource

Get All FM Inline Network Groups

## Example Usage

```terraform
list "gigavuecore_get_all_giga_flex_inline_network_group" "example" {
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

* `cluster_id` (String, optional) - if provided, Network groups only for that cluster is returned


### Identity Attributes

The following identity attributes are exported for each matching result:

* `alias` (String, computed) - Alias of the network group


