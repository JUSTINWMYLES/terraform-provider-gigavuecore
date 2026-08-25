---
page_title: "gigavuecore_delete_inline_ssl_profile_rule Action - gigavuecore"
subcategory: ""
description: |-
  Delete inline SSL profile rule
---

# gigavuecore_delete_inline_ssl_profile_rule Action

Delete inline SSL profile rule

## Example Usage

```terraform
action "gigavuecore_delete_inline_ssl_profile_rule" "example" {
  config {
    alias      = "example"
    cluster_id = "example"
    rule_id    = 1.0
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Alias of the inline SSL profile
* `cluster_id` (String, required) - Target Cluster ID
* `rule_id` (Number, required) - rule ID of the inline SSL profile rule


