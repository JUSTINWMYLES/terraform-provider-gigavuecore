---
page_title: "gigavuecore_deactivate_by_aid Action - gigavuecore"
subcategory: ""
description: |-
  Deactivate a license (floating or VBL) by its Activation ID; floating licenses still assigned to at least one card or chassis will be skipped
---

# gigavuecore_deactivate_by_aid Action

Deactivate a license (floating or VBL) by its Activation ID; floating licenses still assigned to at least one card or chassis will be skipped

## Example Usage

```terraform
action "gigavuecore_deactivate_by_aid" "example" {
  config {
    aid = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `aid` (String, required) - Activation ID (created when license is generated) of the feature activation
