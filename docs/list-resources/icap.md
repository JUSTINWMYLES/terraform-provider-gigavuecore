---
page_title: "gigavuecore_icap List Resource - gigavuecore"
subcategory: ""
description: |-
  Get All icap client apps across FM
---

# gigavuecore_icap List Resource

Get All icap client apps across FM

-> **Note:** This list resource requires Terraform 1.14 or later and is used through the `terraform query` command, not in configuration files.

## Example Usage

```terraform
list "gigavuecore_icap" "example" {
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

* `cluster_id` (String, optional) - If provided, icap client only for that cluster are returned


### Identity Attributes

The following identity attributes are exported for each matching result:

* `alias` (String, computed) - Icap Alias


