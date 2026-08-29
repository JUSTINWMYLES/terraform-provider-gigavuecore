---
page_title: "gigavuecore_fm_build_info Action - gigavuecore"
subcategory: ""
description: |-
  Get FM build info
---

# gigavuecore_fm_build_info Action

Get FM build info

## Example Usage

```terraform
action "gigavuecore_fm_build_info" "example" {
  config {
    eligible_node_role    = "example"
    entity_id             = "example"
    host_name             = "example"
    management_ip_address = "example"
    name                  = "example"
    node_roles = {
      es_node_role    = "example"
      mongo_node_role = "example"
    }
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
    password          = "example"
    public_ip_address = "example"
    username          = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `eligible_node_role` (String, required) - Eligible Node Role of the HA group
* `entity_id` (String, required) - Entity ID of the HA group
* `host_name` (String, required) - DNS Name or IP Address of the HA group
* `management_ip_address` (String, required) - Management Ip Address of the HA group
* `name` (String, required) - name of the HA group
* `node_roles` (Attributes, required) (see [below for nested schema](#nestedatt--node_roles))
* `nodes` (Attributes, required) (see [below for nested schema](#nestedatt--nodes))
* `password` (String, required) - Password of the HA group
* `public_ip_address` (String, required) - Public Ip Address of the HA group
* `username` (String, required) - Username of the HA group

<a id="nestedatt--node_roles"></a>
### Nested Schema for `node_roles`

Optional:

* `es_node_role` (String) - ES node role
* `mongo_node_role` (String) - Mongo node role
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

