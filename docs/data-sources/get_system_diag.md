---
page_title: "gigavuecore_get_system_diag Data Source - gigavuecore"
subcategory: ""
description: |-
  get system diagnostics information
---

# gigavuecore_get_system_diag Data Source

get system diagnostics information

## Example Usage

```terraform
data "gigavuecore_get_system_diag" "example" {
  cluster_id = null
  detail     = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `detail` (Boolean, optional) - Include extra detail in the diagnostics

### Attributes

In addition to all arguments above, the following attributes are exported:

* `diag` (String, computed)


