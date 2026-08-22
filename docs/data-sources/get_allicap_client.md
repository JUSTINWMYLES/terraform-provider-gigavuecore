---
page_title: "gigavuecore_get_allicap_client Data Source - gigavuecore"
subcategory: ""
description: |-
  Get All icap client apps across FM
---

# gigavuecore_get_allicap_client Data Source

Get All icap client apps across FM

## Example Usage

```terraform
data "gigavuecore_get_allicap_client" "example" {
  cluster_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, optional) - If provided, icap client only for that cluster are returned

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({alias, cluster_id, config_status, config_status_reasons, gs_engines, gs_grp_alias, gsop_alias, health_state, health_state_reasons, icap_map_alias, icap_profile_config, icap_server_grp_alias, icap_servers, ing_alias, inline_networks, ip_interface_alias})), computed)

