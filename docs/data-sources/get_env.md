---
page_title: "gigavuecore_get_env Data Source - gigavuecore"
subcategory: ""
description: |-
  Obtain a single unified resource environment by id
---

# gigavuecore_get_env Data Source

Obtain a single unified resource environment by id

## Example Usage

```terraform
data "gigavuecore_get_env" "example" {
  env_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `env_id` (String, required) - unified environment identifier

### Attributes

In addition to all arguments above, the following attributes are exported:

* `name` (String, computed)
* `type` (String, computed) - Unified Environment platform types

