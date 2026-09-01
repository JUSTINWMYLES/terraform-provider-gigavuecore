---
page_title: "gigavuecore_update_giga_insight_node_prompt_upgrade Action - gigavuecore"
subcategory: ""
description: |-
  Upload prompt bundle upgrade for a GigaInsight Node
---

# gigavuecore_update_giga_insight_node_prompt_upgrade Action

Upload prompt bundle upgrade for a GigaInsight Node

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

~> **Note:** This action is not yet wired to a remote API endpoint. Invoking it fails with an explicit "not wired" diagnostic instead of calling the API. The OpenAPI operations it was inferred from could not be resolved into a complete mapping; consult the eidos generation warnings for the exact cause.

## Example Usage

```terraform
action "gigavuecore_update_giga_insight_node_prompt_upgrade" "example" {
  config {
    bundle  = "example"
    node_id = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `bundle` (String, optional)
* `node_id` (String, required)


