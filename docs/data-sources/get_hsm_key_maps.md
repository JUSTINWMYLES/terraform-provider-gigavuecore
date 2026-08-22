---
page_title: "gigavuecore_get_hsm_key_maps Data Source - gigavuecore"
subcategory: ""
description: |-
  Get HSM key maps
---

# gigavuecore_get_hsm_key_maps Data Source

Get HSM key maps

## Example Usage

```terraform
data "gigavuecore_get_hsm_key_maps" "example" {
  alias = null
  cluster_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of target HSM Group
* `cluster_id` (String, required) - Target cluster ID.

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({address, cluster_id, key_name, key_token, port, rfs_match, rule_id})), computed)

