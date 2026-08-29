---
page_title: "gigavuecore_add_devices_to_domain Action - gigavuecore"
subcategory: ""
description: |-
  add clusters/devices to the FM-managed domain
---

# gigavuecore_add_devices_to_domain Action

add clusters/devices to the FM-managed domain

## Example Usage

```terraform
action "gigavuecore_add_devices_to_domain" "example" {
  config {
    async = true
    node_add_specs = [{
      https_port   = "example"
      node_address = "example"
      password     = "example"
      snmp_version = "example"
      username     = "example"
    }]
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `async` (Boolean, optional) - if provided, the call returns immediately with the \[202 Accepted\] HTTP status code, and the discovery process will complete in the background
* `node_add_specs` (Attributes List, required) (see [below for nested schema](#nestedatt--node_add_specs))

<a id="nestedatt--node_add_specs"></a>
### Nested Schema for `node_add_specs`

Required:

* `node_address` (String)

Optional:

* `https_port` (String) - The default value used is 443. If changed in the device same should be given here
* `password` (String)
* `snmp_version` (String) - SNMP version to use.
* `username` (String)

