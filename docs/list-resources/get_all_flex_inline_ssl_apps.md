---
page_title: "gigavuecore_get_all_flex_inline_ssl_apps List Resource - gigavuecore"
subcategory: ""
description: |-
  Get All Inline ssl apps across FM
---

# gigavuecore_get_all_flex_inline_ssl_apps List Resource

Get All Inline ssl apps across FM

## Example Usage

```terraform
list "gigavuecore_get_all_flex_inline_ssl_apps" "example" {
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

* `cluster_id` (String, optional) - If provided, ssl apps only for that cluster are returned


### Identity Attributes

The following identity attributes are exported for each matching result:

* `alias` (String, computed) - Alias of the inline ssl app


