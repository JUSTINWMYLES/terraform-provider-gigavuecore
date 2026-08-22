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
  load_system_details = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `load_system_details` (Bool, required) - Load system details

### Attributes

In addition to all arguments above, the following attributes are exported:

* `fips_enabled` (Bool, computed) - fips enabled
* `fm_ha_tunnel` (Object({tunnel_auth_mode, tunnel_status}), computed)
  * `tunnel_auth_mode` (String, computed) - Auth mode for setting up FMHA tunnel. PSK is Pre Shared Key. PKI is Public Key Infrastructure(Certificate based)
  * `tunnel_status` (Object({total_tunnels, tunnel_health, unreachable_tunnels}), computed) - Includes details of the configured tunnels and their health status in FMHA cluster
    * `total_tunnels` (Object({x198_51_100_42, x198_51_100_43, x198_51_100_44}), computed) - Lists all the configured tunnels in each node in a FMHA cluster
      * `x198_51_100_42` (List(String), computed)
      * `x198_51_100_43` (List(String), computed)
      * `x198_51_100_44` (List(String), computed)
    * `tunnel_health` (String, computed) - Overall health status of all tunnels configured in FMHA cluster
    * `unreachable_tunnels` (Object({x198_51_100_42}), computed) - Lists the unreachable tunnels from a particular FMHA node
      * `x198_51_100_42` (List(String), computed)
* `ha_status` (Object({active_eligible, support}), computed)
  * `active_eligible` (String, computed) - active eligible for ha status
  * `support` (String, computed) - support for ha status
* `hostname` (String, computed) - DNS Name or IP Address of HA group
* `name` (String, computed) - name of the HA group
* `nodes` (Object({cluster_ip_address, entity_id, hostname, idp_meta_data_url, management_ip_address, password, public_ip_address, reachable, seed_node, username}), computed)
  * `cluster_ip_address` (String, computed) - Cluster IP address for FM HA node
  * `entity_id` (String, computed) - Entity Id for FM HA node
  * `hostname` (String, computed) - DNS Name or IP Address for FM HA node
  * `idp_meta_data_url` (String, computed) - IDP Meta data URL for FM HA node
  * `management_ip_address` (String, computed) - Management IP address for FM HA node
  * `password` (String, computed) - password for FM HA node
  * `public_ip_address` (String, computed) - Public IP address for FM HA node
  * `reachable` (Bool, computed) - Reachable for FM HA node
  * `seed_node` (Bool, computed) - Seed node for FM HA node
  * `username` (String, computed) - username for FM HA node

