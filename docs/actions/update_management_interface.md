---
page_title: "gigavuecore_update_management_interface Action - gigavuecore"
subcategory: ""
description: |-
  since FM 5.8
---

# gigavuecore_update_management_interface Action

since FM 5.8

## Example Usage

```terraform
action "gigavuecore_update_management_interface" "example" {
  config {
    body_box_id         = 1
    body_cluster_id     = "example"
    body_interface_name = "example"
    box_id              = 1
    cluster_id          = "example"
    discovery_protocol  = "example"
    g_arp               = true
    interface_name      = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `body_box_id` (Number, required)
* `body_cluster_id` (String, optional) - id of the defining cluster
* `body_interface_name` (String, required) - Management Interface Name
* `box_id` (Number, required) - boxId
* `cluster_id` (String, required) - Target Cluster ID
* `discovery_protocol` (String, required)
* `g_arp` (Boolean, required) - Enable or disable Gratuitous ARP
* `interface_name` (String, required) - Interface Name


