---
page_title: "gigavuecore_get_monitoring_domain Data Source - gigavuecore"
subcategory: ""
description: |-
  Obtain the monitoring domain for underlying unified environment and unified resource id
---

# gigavuecore_get_monitoring_domain Data Source

Obtain the monitoring domain for underlying unified environment and unified resource id

## Example Usage

```terraform
data "gigavuecore_get_monitoring_domain" "example" {
  env_id   = null
  unify_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `env_id` (String, required) - unified environment identifier
* `unify_id` (String, required) - unified resource identifier


