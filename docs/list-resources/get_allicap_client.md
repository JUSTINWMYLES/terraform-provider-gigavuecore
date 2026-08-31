---
page_title: "gigavuecore_get_allicap_client List Resource - gigavuecore"
subcategory: ""
description: |-
  Get All icap client apps across FM
---

# gigavuecore_get_allicap_client List Resource

Get All icap client apps across FM

## Example Usage

```terraform
list "gigavuecore_get_allicap_client" "example" {
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


