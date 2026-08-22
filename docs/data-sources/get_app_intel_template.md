---
page_title: "gigavuecore_get_app_intel_template Data Source - gigavuecore"
subcategory: ""
description: |-
  List all appIntel Info for unified deployment via environment and connection id
---

# gigavuecore_get_app_intel_template Data Source

List all appIntel Info for unified deployment via environment and connection id

## Example Usage

```terraform
data "gigavuecore_get_app_intel_template" "example" {
  env_id = null
  unify_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `env_id` (String, required) - unified environment identifier
* `unify_id` (String, required) - unified resource identifier

### Attributes

In addition to all arguments above, the following attributes are exported:

* `metadata_application_templates` (Object({env}), computed) - app intel info metadata application templates for all apps
  * `env` (Set(Object({attributes, description, family, meta_data_template_version, name})), computed)
* `version` (String, computed)

