---
page_title: "gigavuecore_icap Resource - gigavuecore"
subcategory: ""
description: |-
  Get icap client details by alias
---

# gigavuecore_icap Resource

Get icap client details by alias

## Example Usage

```terraform
resource "gigavuecore_icap" "example" {
  alias = null
  cluster_id = null
  config_status = null
  config_status_reasons = []
  gs_engines = []
  gs_grp_alias = null
  gsop_alias = null
  health_state = null
  health_state_reasons = []
  icap_map_alias = null
  icap_profile_config = {}
  icap_server_grp_alias = null
  icap_servers = []
  ing_alias = null
  inline_networks = []
  ip_interface_alias = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Icap Alias
* `cluster_id` (String, required) - id of the defining cluster
* `config_status` (String, optional) - Configuration status of this ICAP Client app
* `config_status_reasons` (List(String), optional) - In case of configuration failure, this message provides details about the possible cause of the failure
* `gs_engines` (List(String), optional)
* `gs_grp_alias` (String, optional) - GSGroup Alias // User provided (or) Auto-generated
* `gsop_alias` (String, optional) - GSOP Alias
* `health_state` (String, optional) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (List(Object({message, severity, traffic_health_state_computation_type})), optional)
* `icap_map_alias` (String, optional) - Icap MAP Alias
* `icap_profile_config` (Object({alias, cluster_id, exceed_action, http_req_buf, inactivity_timeout, preview, resp_mod, resp_timeout, resp_timeout_action, server_group, src_max_l4_port, src_min_l4_port}), required) - ICAP Profile
  * `alias` (String, required) - Icap Alias
  * `cluster_id` (String, optional) - id of the defining cluster
  * `exceed_action` (String, optional) - Icap Profile action incase of Http request buffer exceeded
  * `http_req_buf` (Number, required) - Icap Profile Http request buffer in KB
  * `inactivity_timeout` (Number, optional) - Icap Inactivity timeout in minutes
  * `preview` (Number, optional) - Icap Preview bytes in KB
  * `resp_mod` (String, optional) - Icap Response Modification Enable\|Disable
  * `resp_timeout` (Number, optional) - Icap Server Response Timeout value in seconds
  * `resp_timeout_action` (String, optional) - Response Timeout action Drop\|Bypass
  * `server_group` (String, required) - Icap Server Group Alias
  * `src_max_l4_port` (Number, required) - Icap Service Source l4 port maximum
  * `src_min_l4_port` (Number, required) - Icap Service Source l4 port minimum
* `icap_server_grp_alias` (String, optional) - Icap Server Group Alias // Auto-generated
* `icap_servers` (List(Object({alias, cluster_id, comment, l3_address, l4_port, options_service_url, reqmod_service_url, respmod_service_url})), optional)
* `ing_alias` (String, optional) - Inline network group Alias //Auto-generated
* `inline_networks` (List(String), required)
* `ip_interface_alias` (String, required) - IpInterface Alias

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `config_status` (String, computed) - Configuration status of this ICAP Client app
* `config_status_reasons` (List(String), computed) - In case of configuration failure, this message provides details about the possible cause of the failure
* `gs_engines` (List(String), computed)
* `gs_grp_alias` (String, computed) - GSGroup Alias // User provided (or) Auto-generated
* `gsop_alias` (String, computed) - GSOP Alias
* `health_state` (String, computed) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (List(Object({message, severity, traffic_health_state_computation_type})), computed)
* `icap_map_alias` (String, computed) - Icap MAP Alias
* `icap_server_grp_alias` (String, computed) - Icap Server Group Alias // Auto-generated
* `icap_servers` (List(Object({alias, cluster_id, comment, l3_address, l4_port, options_service_url, reqmod_service_url, respmod_service_url})), computed)
* `ing_alias` (String, computed) - Inline network group Alias //Auto-generated

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_icap.example {alias}
```
