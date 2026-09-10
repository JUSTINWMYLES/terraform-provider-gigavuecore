---
page_title: "gigavuecore_inline_ssl_app List Resource - gigavuecore"
subcategory: ""
description: |-
  Get All Inline ssl apps across FM
---

# gigavuecore_inline_ssl_app List Resource

Get All Inline ssl apps across FM

-> **Note:** This list resource requires Terraform 1.14 or later and is used through the `terraform query` command, not in configuration files.

~> **Warning:** This list resource accepts attributes whose names indicate secrets (password, password, password, password), but list resource schemas cannot mark attributes Sensitive, so their values are stored and displayed in plain text. Avoid passing real secrets through these attributes.

## Example Usage

```terraform
list "gigavuecore_inline_ssl_app" "example" {
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


