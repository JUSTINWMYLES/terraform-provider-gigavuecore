---
page_title: "gigavuecore_load_mobility_sam_node_configs Data Source - gigavuecore"
subcategory: ""
description: |-
  Load mobility sam node configurations by alias
---

# gigavuecore_load_mobility_sam_node_configs Data Source

Load mobility sam node configurations by alias

## Example Usage

```terraform
data "gigavuecore_load_mobility_sam_node_configs" "example" {
  samnode_alias = null
  solution_alias = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `samnode_alias` (String, required) - Alias of the sam node
* `solution_alias` (String, required) - Alias of the configured solution

### Attributes

In addition to all arguments above, the following attributes are exported:

* `alias` (String, computed) - Alias of the Sam node
* `engine_level_app_vz_configs` (List(Object({app_vz_config, engine_port})), computed)
* `ip_interface_config` (Object({alias, attach, comment, gateway, gs_groups, hw_address, ip_address, ip_mask, ip_type, mtu, netflow_exporters, tags}), computed)
  * `alias` (String, computed) - ip interface name
  * `attach` (List(String), computed) - network ports ,tool ports or circuit ports
  * `comment` (String, computed)
  * `gateway` (String, computed) - gateway ipv4 or ipv6 address
  * `gs_groups` (List(String), computed) - Gs Groups associated with the IP Interface
  * `hw_address` (String, computed)
  * `ip_address` (String, computed) - ipv4/ipv6 address
  * `ip_mask` (String, computed) - ipAddress netmask required with ipAddress
  * `ip_type` (String, computed)
  * `mtu` (Number, computed)
  * `netflow_exporters` (List(String), computed) - Netflow Exporters associated with the IP Interface
  * `tags` (List(Object({tag_key, tag_values})), computed)
* `location` (Object({cluster_id, engine_ports}), computed) - Location of the engine port
  * `cluster_id` (String, computed)
  * `engine_ports` (List(String), computed)
* `site_name` (String, computed) - Name of the site where the node is deployed

