---
page_title: "gigavuecore_delete_all_circuit_tunnels Action - gigavuecore"
subcategory: ""
description: |-
  Delete all Circuit Tunnels
---

# gigavuecore_delete_all_circuit_tunnels Action

Delete all Circuit Tunnels

## Example Usage

```terraform
action "gigavuecore_delete_all_circuit_tunnels" "example" {
  config {
    cluster_id = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID


