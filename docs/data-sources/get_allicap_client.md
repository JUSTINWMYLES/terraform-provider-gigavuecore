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
  cluster_id = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, optional) - If provided, icap client only for that cluster are returned

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `alias` (String) - Icap Alias
* `cluster_id` (String) - id of the defining cluster
* `config_status` (String) - Configuration status of this ICAP Client app
* `config_status_reasons` (List of String) - In case of configuration failure, this message provides details about the possible cause of the failure
* `gs_engines` (List of String)
* `gs_grp_alias` (String) - GSGroup Alias // User provided (or) Auto-generated
* `gsop_alias` (String) - GSOP Alias
* `health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--items--health_state_reasons))
* `icap_map_alias` (String) - Icap MAP Alias
* `icap_profile_config` (Attributes) - ICAP Profile (see [below for nested schema](#nestedatt--items--icap_profile_config))
* `icap_server_grp_alias` (String) - Icap Server Group Alias // Auto-generated
* `icap_servers` (Attributes List) (see [below for nested schema](#nestedatt--items--icap_servers))
* `ing_alias` (String) - Inline network group Alias //Auto-generated
* `inline_networks` (List of String)
* `ip_interface_alias` (String) - IpInterface Alias

<a id="nestedatt--items--health_state_reasons"></a>
### Nested Schema for `items.health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type

<a id="nestedatt--items--icap_profile_config"></a>
### Nested Schema for `items.icap_profile_config`

Read-Only:

* `alias` (String) - Icap Alias
* `cluster_id` (String) - id of the defining cluster
* `exceed_action` (String) - Icap Profile action incase of Http request buffer exceeded
* `http_req_buf` (Number) - Icap Profile Http request buffer in KB
* `inactivity_timeout` (Number) - Icap Inactivity timeout in minutes
* `preview` (Number) - Icap Preview bytes in KB
* `resp_mod` (String) - Icap Response Modification Enable\|Disable
* `resp_timeout` (Number) - Icap Server Response Timeout value in seconds
* `resp_timeout_action` (String) - Response Timeout action Drop\|Bypass
* `server_group` (String) - Icap Server Group Alias
* `src_max_l4_port` (Number) - Icap Service Source l4 port maximum
* `src_min_l4_port` (Number) - Icap Service Source l4 port minimum

<a id="nestedatt--items--icap_servers"></a>
### Nested Schema for `items.icap_servers`

Read-Only:

* `alias` (String) - Icap Server Alias
* `cluster_id` (String) - id of the defining cluster
* `comment` (String) - Icap Server Comment
* `l3_address` (String) - Icap Server IP Address
* `l4_port` (Number) - Icap Server l4 Port
* `options_service_url` (String) - Options Service URL
* `reqmod_service_url` (String) - Request Modification Service URL
* `respmod_service_url` (String) - Response Modification Service URL

