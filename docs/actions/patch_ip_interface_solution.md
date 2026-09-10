---
page_title: "gigavuecore_patch_ip_interface_solution Action - gigavuecore"
subcategory: ""
description: |-
  Modify existing Ip Interface solution configuration
---

# gigavuecore_patch_ip_interface_solution Action

Modify existing Ip Interface solution configuration

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_patch_ip_interface_solution" "example" {
  config {
    ip_interface_configs = [{
      alias                 = "example"
      applications          = ["CPN"]
      cluster_name          = "example"
      config_status         = "PARTIAL_SUCCESS"
      config_status_reasons = "example"
      gateway               = "example"
      interfaces            = ["example"]
      ip_address            = "example"
      ip_mask               = "example"
      managed_status        = "ACTIVE"
      mtu                   = 1500
      ref_count             = 0
    }]
    tags = [{
      tag_key    = "example"
      tag_values = ["example"]
    }]
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `ip_interface_configs` (Attributes List, required) (see [below for nested schema](#nestedatt--ip_interface_configs))
* `tags` (Attributes List, optional) - RBAC Tags (see [below for nested schema](#nestedatt--tags))

<a id="nestedatt--ip_interface_configs"></a>
### Nested Schema for `ip_interface_configs`

Required:

* `alias` (String) - Alias of the ip interface solution
* `applications` (List of String) - GigaSMART applications for which the ipInterface is used
* `cluster_name` (String) - clusterId
* `gateway` (String) - gateway ipv4 or ipv6 address
* `interfaces` (List of String) - network ports ,tool ports or circuit ports
* `ip_address` (String) - ipv4/ipv6 address
* `ip_mask` (String) - ipAddress netmask required with ipAddress

Optional:

* `config_status` (String)
* `config_status_reasons` (String)
* `managed_status` (String) - managed status of the ip interface configuration. ACTIVE if the cluster is managed in this FM , else it will be marked as INACTIVE
* `mtu` (Number)
* `ref_count` (Number) - number of control or user nodes using this ip interface configuration

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Required:

* `tag_key` (String) - Name of the tag
* `tag_values` (List of String) - All possible values of the tag

