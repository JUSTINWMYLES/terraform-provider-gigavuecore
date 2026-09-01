---
page_title: "gigavuecore_create_ha_group Action - gigavuecore"
subcategory: ""
description: |-
  create HA group
---

# gigavuecore_create_ha_group Action

create HA group

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

~> **Warning:** This action accepts attributes whose names indicate secrets (password), but action schemas cannot mark attributes Sensitive, so their values are stored and displayed in plain text. Avoid passing real secrets through these attributes.

## Example Usage

```terraform
action "gigavuecore_create_ha_group" "example" {
  config {
    fm_ha_tunnel = {
      tunnel_auth_mode = "PSK"
    }
    name = "example"
    nodes = {
      cluster_ip_address    = "example"
      entity_id             = "example"
      hostname              = "example"
      idp_meta_data_url     = "example"
      management_ip_address = "example"
      password              = "example"
      public_ip_address     = "example"
      reachable             = true
      seed_node             = true
      username              = "example"
    }
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `fm_ha_tunnel` (Attributes, required) - Auth mode for creating tunnels in new FMHA cluster (see [below for nested schema](#nestedatt--fm_ha_tunnel))
* `name` (String, required) - Name of the new FMHA cluster
* `nodes` (Attributes, required) - Nodes of the new FMHA cluster (see [below for nested schema](#nestedatt--nodes))

<a id="nestedatt--fm_ha_tunnel"></a>
### Nested Schema for `fm_ha_tunnel`

Optional:

* `tunnel_auth_mode` (String) - Auth mode for setting up FMHA tunnel. PSK is Pre Shared Key. PKI is Public Key Infrastructure(Certificate based)

<a id="nestedatt--nodes"></a>
### Nested Schema for `nodes`

Required:

* `cluster_ip_address` (String) - Cluster IP address for FM HA node
* `entity_id` (String) - Entity Id for FM HA node
* `hostname` (String) - DNS Name or IP Address for FM HA node
* `management_ip_address` (String) - Management IP address for FM HA node
* `password` (String) - password for FM HA node
* `public_ip_address` (String) - Public IP address for FM HA node
* `reachable` (Boolean) - Reachable for FM HA node
* `seed_node` (Boolean) - Seed node for FM HA node
* `username` (String) - username for FM HA node

Optional:

* `idp_meta_data_url` (String) - IDP Meta data URL for FM HA node

