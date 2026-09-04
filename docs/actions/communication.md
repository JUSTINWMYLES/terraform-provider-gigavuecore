---
page_title: "gigavuecore_communication Action - gigavuecore"
subcategory: ""
description: |-
  stop/resume communcation by cluster IDs.
---

# gigavuecore_communication Action

stop/resume communcation by cluster IDs.

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_communication" "example" {
  config {
    cluster_id = "example"
    disconnect = true
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Cluster ID
* `disconnect` (Boolean, required) - The boolean to stop/resume communication to device.


