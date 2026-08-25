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
  env_id   = null
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

* `metadata_application_templates` (Attributes, computed) - app intel info metadata application templates for all apps (see [below for nested schema](#nestedatt--metadata_application_templates))
* `version` (String, computed)

<a id="nestedatt--metadata_application_templates"></a>
### Nested Schema for `metadata_application_templates`

Read-Only:

* `env` (Attributes Set) (see [below for nested schema](#nestedatt--metadata_application_templates--env))
<a id="nestedatt--metadata_application_templates--env"></a>
### Nested Schema for `metadata_application_templates.env`

Read-Only:

* `attributes` (Attributes) - app intel info metadata template attributes (see [below for nested schema](#nestedatt--metadata_application_templates--env--attributes))
* `description` (String)
* `family` (String)
* `meta_data_template_version` (String)
* `name` (String)
<a id="nestedatt--metadata_application_templates--env--attributes"></a>
### Nested Schema for `metadata_application_templates.env.attributes`

Read-Only:

* `env` (Attributes Set) (see [below for nested schema](#nestedatt--metadata_application_templates--env--attributes--env))
<a id="nestedatt--metadata_application_templates--env--attributes--env"></a>
### Nested Schema for `metadata_application_templates.env.attributes.env`

Read-Only:

* `name` (String)
* `value` (String)

