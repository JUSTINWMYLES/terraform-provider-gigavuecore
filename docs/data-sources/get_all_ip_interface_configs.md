---
page_title: "gigavuecore_get_all_ip_interface_configs Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all ipInterface solution configurations
---

# gigavuecore_get_all_ip_interface_configs Data Source

Load all ipInterface solution configurations

## Example Usage

```terraform
data "gigavuecore_get_all_ip_interface_configs" "example" {
  application = "example"
  cluster_id  = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `application` (String, optional) - Gets ipInterface solution that supports the specified GigaSmart application
* `cluster_id` (String, optional) - Gets ipInterface solution configs for the provided cluster

### Attributes

In addition to all arguments above, the following attributes are exported:

* `ip_interface_configs` (Attributes List, computed) (see [below for nested schema](#nestedatt--ip_interface_configs))
* `tags` (Attributes List, computed) - RBAC Tags (see [below for nested schema](#nestedatt--tags))

<a id="nestedatt--ip_interface_configs"></a>
### Nested Schema for `ip_interface_configs`

Read-Only:

* `alias` (String) - Alias of the ip interface solution
* `applications` (List of String) - GigaSMART applications for which the ipInterface is used
* `cluster_name` (String) - clusterId
* `config_status` (String)
* `config_status_reasons` (String)
* `gateway` (String) - gateway ipv4 or ipv6 address
* `interfaces` (List of String) - network ports ,tool ports or circuit ports
* `ip_address` (String) - ipv4/ipv6 address
* `ip_mask` (String) - ipAddress netmask required with ipAddress
* `managed_status` (String) - managed status of the ip interface configuration. ACTIVE if the cluster is managed in this FM , else it will be marked as INACTIVE
* `mtu` (Number)
* `ref_count` (Number) - number of control or user nodes using this ip interface configuration

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

* `tag_key` (String) - Name of the tag
* `tag_values` (List of String) - All possible values of the tag

