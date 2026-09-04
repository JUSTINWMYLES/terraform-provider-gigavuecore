---
page_title: "gigavuecore_clear_ipv6_neighbor_entries Action - gigavuecore"
subcategory: ""
description: |-
  Clear Ipv6 Neighbor Entries
---

# gigavuecore_clear_ipv6_neighbor_entries Action

Clear Ipv6 Neighbor Entries

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_clear_ipv6_neighbor_entries" "example" {
  config {
  }
}
```
