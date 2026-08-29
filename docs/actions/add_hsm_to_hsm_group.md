---
page_title: "gigavuecore_add_hsm_to_hsm_group Action - gigavuecore"
subcategory: ""
description: |-
  Add HSM alias to HSM Group
---

# gigavuecore_add_hsm_to_hsm_group Action

Add HSM alias to HSM Group

## Example Usage

```terraform
action "gigavuecore_add_hsm_to_hsm_group" "example" {
  config {
    alias     = "example"
    hsm_alias = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of target HSM Group
* `hsm_alias` (String, required) - HSM alias


