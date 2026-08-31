---
page_title: "gigavuecore_audit_traffic_flow_generated_f_map Action - gigavuecore"
subcategory: ""
description: |-
  Check the integrity of traffic flow configurations
---

# gigavuecore_audit_traffic_flow_generated_f_map Action

Check the integrity of traffic flow configurations

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_audit_traffic_flow_generated_f_map" "example" {
  config {
  }
}
```
