---
page_title: "gigavuecore_create_node_system_config_file Action - gigavuecore"
subcategory: ""
description: |-
  Save the running configuration
---

# gigavuecore_create_node_system_config_file Action

Save the running configuration

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_create_node_system_config_file" "example" {
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
* `filename` (String, optional) - Save the configuration to a new file and make it active, if filename specified


