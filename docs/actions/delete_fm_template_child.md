---
page_title: "gigavuecore_delete_fm_template_child Action - gigavuecore"
subcategory: ""
description: |-
  Delete child entry from template configuration
---

# gigavuecore_delete_fm_template_child Action

Delete child entry from template configuration

## Example Usage

```terraform
action "gigavuecore_delete_fm_template_child" "example" {
  config {
    body_child_type    = "example"
    body_config_type   = "example"
    child_type         = "example"
    config             = null
    config_level       = "example"
    config_level_value = [ "example" ]
    config_type        = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `body_child_type` (String, optional) - Child Type of the FM template Configuration
* `body_config_type` (String, optional) - Configuration Type of the FM template
* `child_type` (String, required) - childType of the fm template
* `config` (Dynamic, optional)
* `config_level` (String, optional) - Scope of the applied FM template
* `config_level_value` (List of String, optional)
* `config_type` (String, required) - configType of the fm template


