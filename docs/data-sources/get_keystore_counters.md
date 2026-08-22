---
page_title: "gigavuecore_get_keystore_counters Data Source - gigavuecore"
subcategory: ""
description: |-
  Load keystore hit counters
---

# gigavuecore_get_keystore_counters Data Source

Load keystore hit counters

## Example Usage

```terraform
data "gigavuecore_get_keystore_counters" "example" {
  cluster_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID

### Attributes

In addition to all arguments above, the following attributes are exported:

* `hits` (Number, computed)
* `key_alias` (String, computed)

