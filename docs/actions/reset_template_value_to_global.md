---
page_title: "gigavuecore_reset_template_value_to_global Action - gigavuecore"
subcategory: ""
description: |-
  new in FM 5.14
---

# gigavuecore_reset_template_value_to_global Action

new in FM 5.14

## Example Usage

```terraform
action "gigavuecore_reset_template_value_to_global" "example" {
  config {
    body_config_type = "example"
    config = null
    config_level = "example"
    config_level_value = [ "example" ]
    config_resource = null
    config_type = "example"
    modifiable = true
    ref_count = 1
    ref_object = null
    template_name = "example"
    update_time = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `body_config_type` (String, optional) - Configuration Type of the FM template
* `config` (Dynamic, optional)
* `config_level` (String, optional) - Scope of the applied FM template
* `config_level_value` (List(String), optional)
* `config_resource` (Dynamic, optional) - For a particular config type, there may be multiple templates with different levels. For example, cluster level the user can have more than one template and in same way in tag level there may be more than one template for each tag combination. In such cases, the configResource property can be used to contain the list of clusters associated to the template or the tag key values combination of the template. The data structure changes based on the level hence it is kept as type Object.
* `config_type` (String, required) - configType of the fm template
* `modifiable` (Bool, optional)
* `ref_count` (Number, optional)
* `ref_object` (Dynamic, optional)
* `template_name` (String, optional)
* `update_time` (String, optional)
