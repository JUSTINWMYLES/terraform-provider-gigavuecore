---
page_title: "gigavuecore_get_all_giga_flex_inline_network_group Data Source - gigavuecore"
subcategory: ""
description: |-
  Get All FM Inline Network Groups
---

# gigavuecore_get_all_giga_flex_inline_network_group Data Source

Get All FM Inline Network Groups

## Example Usage

```terraform
data "gigavuecore_get_all_giga_flex_inline_network_group" "example" {
  cluster_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, optional) - if provided, Network groups only for that cluster is returned

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({alias, cluster_id, health_state, health_state_reasons, members})), computed)

