---
page_title: "gigavuecore_get_keystore_preference Data Source - gigavuecore"
subcategory: ""
description: |-
  Load keystore preference
---

# gigavuecore_get_keystore_preference Data Source

Load keystore preference

## Example Usage

```terraform
data "gigavuecore_get_keystore_preference" "example" {
  cluster_id = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID

### Attributes

In addition to all arguments above, the following attributes are exported:

* `auto_delete` (Boolean, computed)
* `auto_enable` (Boolean, computed)
* `auto_purge` (Boolean, computed)
* `max_keys` (Number, computed)
* `retention_time` (Number, computed)


