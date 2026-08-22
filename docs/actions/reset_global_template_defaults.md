---
page_title: "gigavuecore_reset_global_template_defaults Action - gigavuecore"
subcategory: ""
description: |-
  new in FM 5.14
---

# gigavuecore_reset_global_template_defaults Action

new in FM 5.14

## Example Usage

```terraform
action "gigavuecore_reset_global_template_defaults" "example" {
  config {
    config_type = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `config_type` (String, required) - configType of the fm template
