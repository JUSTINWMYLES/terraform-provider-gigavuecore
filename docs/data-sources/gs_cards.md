---
page_title: "gigavuecore_gs_cards Data Source - gigavuecore"
subcategory: ""
description: |-
  Get Gigasmart card information
---

# gigavuecore_gs_cards Data Source

Get Gigasmart card information

## Example Usage

```terraform
data "gigavuecore_gs_cards" "example" {
  cluster_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, optional) - if provided, GS card information of the requested cluster will be returned

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({cards, cluster_id})), computed)

