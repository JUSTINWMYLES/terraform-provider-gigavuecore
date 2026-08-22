---
page_title: "gigavuecore_get_active_config Data Source - gigavuecore"
subcategory: ""
description: |-
  List only the active config
---

# gigavuecore_get_active_config Data Source

List only the active config

## Example Usage

```terraform
data "gigavuecore_get_active_config" "example" {
  cluster_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID

### Attributes

In addition to all arguments above, the following attributes are exported:

* `active` (Bool, computed) - Is config active
* `filename` (String, computed) - File name
* `modified` (Bool, computed) - Has unsaved config changes. True: need to save.

