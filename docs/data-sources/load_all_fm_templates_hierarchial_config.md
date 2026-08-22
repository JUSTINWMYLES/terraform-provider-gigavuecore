---
page_title: "gigavuecore_load_all_fm_templates_hierarchial_config Data Source - gigavuecore"
subcategory: ""
description: |-
  new in FM 5.7
---

# gigavuecore_load_all_fm_templates_hierarchial_config Data Source

new in FM 5.7

## Example Usage

```terraform
data "gigavuecore_load_all_fm_templates_hierarchial_config" "example" {
  config_type = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `config_type` (String, optional) - configType of the fm template

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({config, config_level, config_level_value, config_resource, config_type, modifiable, ref_count, ref_object, template_name, update_time})), computed) - All FM Template  Configurations

