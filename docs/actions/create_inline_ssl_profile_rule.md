---
page_title: "gigavuecore_create_inline_ssl_profile_rule Action - gigavuecore"
subcategory: ""
description: |-
  This must match one of the rule objects: InlineSslProfileRuleDomain, InlineSslProfileRuleCategory, InlineSslProfileRuleIssuer, InlineSslProfileRuleIpv4, InlineSslProfileRuleL4Port, InlineSslProfileRuleVlan
---

# gigavuecore_create_inline_ssl_profile_rule Action

This must match one of the rule objects: InlineSslProfileRuleDomain, InlineSslProfileRuleCategory, InlineSslProfileRuleIssuer, InlineSslProfileRuleIpv4, InlineSslProfileRuleL4Port, InlineSslProfileRuleVlan

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_create_inline_ssl_profile_rule" "example" {
  config {
    alias      = "example"
    body       = "example"
    cluster_id = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Alias of the inline SSL profile
* `body` (Dynamic, required) - inline SSL profile rule
* `cluster_id` (String, required) - Target Cluster ID


