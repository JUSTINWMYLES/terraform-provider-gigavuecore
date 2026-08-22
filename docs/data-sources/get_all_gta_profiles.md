---
page_title: "gigavuecore_get_all_gta_profiles Data Source - gigavuecore"
subcategory: ""
description: |-
  get all gta profiles
---

# gigavuecore_get_all_gta_profiles Data Source

get all gta profiles

## Example Usage

```terraform
data "gigavuecore_get_all_gta_profiles" "example" {
}
```

## Schema

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({alias, control_node, core_network_nodes, dst_port, src_port, user_node})), computed)

