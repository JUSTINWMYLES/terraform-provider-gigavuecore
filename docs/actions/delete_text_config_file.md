---
page_title: "gigavuecore_delete_text_config_file Action - gigavuecore"
subcategory: ""
description: |-
  Delete a system text configuration file
---

# gigavuecore_delete_text_config_file Action

Delete a system text configuration file

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_delete_text_config_file" "example" {
  config {
    cluster_id = "example"
    filename   = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `filename` (String, required) - filename of the text configuration file


