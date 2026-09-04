---
page_title: "gigavuecore_add_devices_to_domain Action - gigavuecore"
subcategory: ""
description: |-
  add clusters/devices to the FM-managed domain
---

# gigavuecore_add_devices_to_domain Action

add clusters/devices to the FM-managed domain

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

~> **Warning:** This action accepts attributes whose names indicate secrets (password), but action schemas cannot mark attributes Sensitive, so their values are stored and displayed in plain text. Avoid passing real secrets through these attributes.

## Example Usage

```terraform
action "gigavuecore_add_devices_to_domain" "example" {
  config {
    async = true
    node_add_specs = [{
      https_port   = "example"
      node_address = "example"
      password     = "example"
      snmp_version = "v2"
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

