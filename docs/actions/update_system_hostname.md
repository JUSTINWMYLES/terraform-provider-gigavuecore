---
page_title: "gigavuecore_update_system_hostname Action - gigavuecore"
subcategory: ""
description: |-
  Update system hostname configurations
---

# gigavuecore_update_system_hostname Action

Update system hostname configurations

## Example Usage

```terraform
action "gigavuecore_update_system_hostname" "example" {
  config {
    alternate_hostname_with_dhcp = "example"
    body_cluster_id              = "example"
    cluster_id                   = "example"
    hostname                     = "example"
    hostname_with_dhcp           = true
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alternate_hostname_with_dhcp` (String, optional) - indicates alternate host name with DHCP client request
* `body_cluster_id` (String, optional) - Gigamon cluster Id
* `cluster_id` (String, required) - Target Cluster ID
* `hostname` (String, optional) - hostname of system
* `hostname_with_dhcp` (Boolean, optional) - enable to use host name with DHCP client request


