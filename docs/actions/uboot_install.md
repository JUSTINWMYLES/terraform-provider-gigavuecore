---
page_title: "gigavuecore_uboot_install Action - gigavuecore"
subcategory: ""
description: |-
  Install the binary bootloader code included with the active/booted image
---

# gigavuecore_uboot_install Action

Install the binary bootloader code included with the active/booted image

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_uboot_install" "example" {
  config {
    cluster_id = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID


