---
page_title: "gigavuecore_generate_sysdump Action - gigavuecore"
subcategory: ""
description: |-
  Generate New Sysdump File
---

# gigavuecore_generate_sysdump Action

Generate New Sysdump File

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_generate_sysdump" "example" {
  config {
    box_id     = "example"
    cluster_id = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `box_id` (String, optional) - Box ID range from 1 to 64(inclusive). all is applicable
* `cluster_id` (String, required) - Target Cluster ID


