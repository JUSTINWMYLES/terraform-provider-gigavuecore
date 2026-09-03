---
page_title: "gigavuecore_install_node_locked_license_key_on_node Action - gigavuecore"
subcategory: ""
description: |-
  Install a node-locked license key on a node
---

# gigavuecore_install_node_locked_license_key_on_node Action

Install a node-locked license key on a node

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_install_node_locked_license_key_on_node" "example" {
  config {
    box_id      = 0
    cluster_id  = "example"
    license_key = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `box_id` (Number, optional) - id of the box within the chasis cluster
* `cluster_id` (String, required) - name of the cluster
* `license_key` (String, optional) - the license key to install


