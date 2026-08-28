---
page_title: "gigavuecore_get_deploy Data Source - gigavuecore"
subcategory: ""
description: |-
  Get unified resource deployment by environment and unified resource id
---

# gigavuecore_get_deploy Data Source

Get unified resource deployment by environment and unified resource id

## Example Usage

```terraform
data "gigavuecore_get_deploy" "example" {
  env_id   = null
  unify_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `env_id` (String, required) - unified environment identifier
* `unify_id` (String, required) - unified resource connection identifier

### Attributes

In addition to all arguments above, the following attributes are exported:

* `name` (String, computed)
* `platform_type` (String, computed)
* `unified_resources_deploy_get_gv_tap_query_response` (Dynamic, computed) - Unified Resource deployment


