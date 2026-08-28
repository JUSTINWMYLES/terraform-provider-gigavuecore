---
page_title: "gigavuecore_install_node_locked_license_key_on_node Action - gigavuecore"
subcategory: ""
description: |-
  Install a node-locked license key on a node
---

# gigavuecore_install_node_locked_license_key_on_node Action

Install a node-locked license key on a node

## Example Usage

```terraform
action "gigavuecore_install_node_locked_license_key_on_node" "example" {
  config {
    box_id      = 1
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


