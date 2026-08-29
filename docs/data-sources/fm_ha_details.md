---
page_title: "gigavuecore_fm_ha_details Data Source - gigavuecore"
subcategory: ""
description: |-
  Get fmHa group details
---

# gigavuecore_fm_ha_details Data Source

Get fmHa group details

## Example Usage

```terraform
data "gigavuecore_fm_ha_details" "example" {
  load_system_details = true
}
```

## Schema

### Arguments

The following arguments are supported:

* `load_system_details` (Boolean, required) - Load system details

### Attributes

In addition to all arguments above, the following attributes are exported:

* `fips_enabled` (Boolean, computed) - fips enabled
* `fm_ha_tunnel` (Attributes, computed) (see [below for nested schema](#nestedatt--fm_ha_tunnel))
* `ha_status` (Attributes, computed) (see [below for nested schema](#nestedatt--ha_status))
* `hostname` (String, computed) - DNS Name or IP Address of HA group
* `name` (String, computed) - name of the HA group
* `nodes` (Attributes, computed) (see [below for nested schema](#nestedatt--nodes))

<a id="nestedatt--fm_ha_tunnel"></a>
### Nested Schema for `fm_ha_tunnel`

Read-Only:

* `tunnel_auth_mode` (String) - Auth mode for setting up FMHA tunnel. PSK is Pre Shared Key. PKI is Public Key Infrastructure(Certificate based)
* `tunnel_status` (Attributes) - Includes details of the configured tunnels and their health status in FMHA cluster (see [below for nested schema](#nestedatt--fm_ha_tunnel--tunnel_status))
<a id="nestedatt--fm_ha_tunnel--tunnel_status"></a>
### Nested Schema for `fm_ha_tunnel.tunnel_status`

Read-Only:

* `total_tunnels` (Attributes) - Lists all the configured tunnels in each node in a FMHA cluster (see [below for nested schema](#nestedatt--fm_ha_tunnel--tunnel_status--total_tunnels))
* `tunnel_health` (String) - Overall health status of all tunnels configured in FMHA cluster
* `unreachable_tunnels` (Attributes) - Lists the unreachable tunnels from a particular FMHA node (see [below for nested schema](#nestedatt--fm_ha_tunnel--tunnel_status--unreachable_tunnels))
<a id="nestedatt--fm_ha_tunnel--tunnel_status--total_tunnels"></a>
### Nested Schema for `fm_ha_tunnel.tunnel_status.total_tunnels`

Read-Only:

* `x198_51_100_42` (List of String)
* `x198_51_100_43` (List of String)
* `x198_51_100_44` (List of String)
<a id="nestedatt--fm_ha_tunnel--tunnel_status--unreachable_tunnels"></a>
### Nested Schema for `fm_ha_tunnel.tunnel_status.unreachable_tunnels`

Read-Only:

* `x198_51_100_42` (List of String)
<a id="nestedatt--ha_status"></a>
### Nested Schema for `ha_status`

Read-Only:

* `active_eligible` (String) - active eligible for ha status
* `support` (String) - support for ha status
<a id="nestedatt--nodes"></a>
### Nested Schema for `nodes`

Read-Only:

* `cluster_ip_address` (String) - Cluster IP address for FM HA node
* `entity_id` (String) - Entity Id for FM HA node
* `hostname` (String) - DNS Name or IP Address for FM HA node
* `idp_meta_data_url` (String) - IDP Meta data URL for FM HA node
* `management_ip_address` (String) - Management IP address for FM HA node
* `password` (String) - password for FM HA node
* `public_ip_address` (String) - Public IP address for FM HA node
* `reachable` (Boolean) - Reachable for FM HA node
* `seed_node` (Boolean) - Seed node for FM HA node
* `username` (String) - username for FM HA node

