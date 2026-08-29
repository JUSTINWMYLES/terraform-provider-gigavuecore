---
page_title: "gigavuecore_apply_precryption_policy_config Action - gigavuecore"
subcategory: ""
description: |-
  Apply Precryption Policy Config
---

# gigavuecore_apply_precryption_policy_config Action

Apply Precryption Policy Config

## Example Usage

```terraform
action "gigavuecore_apply_precryption_policy_config" "example" {
  config {
    id   = "example"
    name = "example"
    rules = [{
      action    = "example"
      direction = "example"
      filters   = [ "example" ]
      priority  = "example"
      rule_name = "example"
    }]
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `id` (String, required) - policy graph id
* `name` (String, required) - Unique name of this traffic policy. It can be between 1 and 32 characters long and may contain only alpha numeric characters, underscores and dashes.
* `rules` (Attributes List, required) - Rules defined for this traffic policy. At least one rule has to be specified and a maximum of 16 rules could be specified. (see [below for nested schema](#nestedatt--rules))

<a id="nestedatt--rules"></a>
### Nested Schema for `rules`

Required:

* `action` (String) - Action associated with the rule.
* `direction` (String) - Direction associated with the rule.
* `filters` (List of String) - Currently it is a pass-all Rule, no filter is applicable
* `priority` (String) - Priority of the rule is 1, Cannot be null.
* `rule_name` (String) - Name of the Precryption traffic policy rule. It can be between 1 and 20 characters long and may contain only alpha numeric characters, underscores and dashes.

