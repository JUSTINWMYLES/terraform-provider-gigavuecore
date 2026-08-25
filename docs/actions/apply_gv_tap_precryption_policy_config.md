---
page_title: "gigavuecore_apply_gv_tap_precryption_policy_config Action - gigavuecore"
subcategory: ""
description: |-
  Apply Precryption Policy Config \[marked for deprecation\]
---

# gigavuecore_apply_gv_tap_precryption_policy_config Action

Apply Precryption Policy Config [marked for deprecation]

## Example Usage

```terraform
action "gigavuecore_apply_gv_tap_precryption_policy_config" "example" {
  config {
    id    = "example"
    name  = "example"
    rules = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `id` (String, required) - policy graph id
* `name` (String, required) - Unique name of this traffic policy. It can be between 1 and 32 characters long and may contain only alpha numeric characters, underscores and dashes.
* `rules` (List of Dynamic, required) - Rules defined for this traffic policy. At least one rule has to be specified and a maximum of 16 rules could be specified.


