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
  alias                 = "example"
  cluster_id            = "example"
  config_status         = "SUCCESS"
  config_status_reasons = ["example"]
  gs_engines            = ["example"]
  gs_grp_alias          = "example"
  gsop_alias            = "example"
  health_state          = "green"
  health_state_reasons = [{
    message                               = "example"
    severity                              = "green"
    traffic_health_state_computation_type = "PORT_LOW_UTIL"
  }]
  icap_map_alias = "example"
  icap_profile_config = {
    alias               = "example"
    cluster_id          = "example"
    exceed_action       = "drop"
    http_req_buf        = 10
    inactivity_timeout  = 2
    preview             = 0
    resp_mod            = "enable"
    resp_timeout        = 5
    resp_timeout_action = "drop"
    server_group        = "example"
    src_max_l4_port     = 10000
    src_min_l4_port     = 10000
  }
  icap_server_grp_alias = "example"
  icap_servers = [{
    alias               = "example"
    cluster_id          = "example"
    comment             = "example"
    l3_address          = "example"
    l4_port             = 0
    options_service_url = "example"
    reqmod_service_url  = "example"
    respmod_service_url = "example"
  }]
  ing_alias          = "example"
  inline_networks    = ["example"]
  ip_interface_alias = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Icap Alias
* `cluster_id` (String, required) - id of the defining cluster
* `config_status` (String, optional) - Configuration status of this ICAP Client app
* `config_status_reasons` (List of String, optional) - In case of configuration failure, this message provides details about the possible cause of the failure
* `gs_engines` (List of String, optional)
* `gs_grp_alias` (String, optional) - GSGroup Alias // User provided (or) Auto-generated
* `gsop_alias` (String, optional) - GSOP Alias
* `health_state` (String, optional) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List, optional) (see [below for nested schema](#nestedatt--health_state_reasons))
* `icap_map_alias` (String, optional) - Icap MAP Alias
* `icap_profile_config` (Attributes, required) - ICAP Profile (see [below for nested schema](#nestedatt--icap_profile_config))
* `icap_server_grp_alias` (String, optional) - Icap Server Group Alias // Auto-generated
* `icap_servers` (Attributes List, optional) (see [below for nested schema](#nestedatt--icap_servers))
* `ing_alias` (String, optional) - Inline network group Alias //Auto-generated
* `inline_networks` (List of String, required)
* `ip_interface_alias` (String, required) - IpInterface Alias

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

<a id="nestedatt--health_state_reasons"></a>
### Nested Schema for `health_state_reasons`

Optional:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type

<a id="nestedatt--icap_profile_config"></a>
### Nested Schema for `icap_profile_config`

Required:

* `alias` (String) - Icap Alias
* `http_req_buf` (Number) - Icap Profile Http request buffer in KB
* `server_group` (String) - Icap Server Group Alias
* `src_max_l4_port` (Number) - Icap Service Source l4 port maximum
* `src_min_l4_port` (Number) - Icap Service Source l4 port minimum

Optional:

* `cluster_id` (String) - id of the defining cluster
* `exceed_action` (String) - Icap Profile action incase of Http request buffer exceeded
* `inactivity_timeout` (Number) - Icap Inactivity timeout in minutes
* `preview` (Number) - Icap Preview bytes in KB
* `resp_mod` (String) - Icap Response Modification Enable\|Disable
* `resp_timeout` (Number) - Icap Server Response Timeout value in seconds
* `resp_timeout_action` (String) - Response Timeout action Drop\|Bypass

<a id="nestedatt--icap_servers"></a>
### Nested Schema for `icap_servers`

Required:

* `alias` (String) - Icap Server Alias
* `l3_address` (String) - Icap Server IP Address
* `l4_port` (Number) - Icap Server l4 Port

Optional:

* `cluster_id` (String) - id of the defining cluster
* `comment` (String) - Icap Server Comment
* `options_service_url` (String) - Options Service URL
* `reqmod_service_url` (String) - Request Modification Service URL
* `respmod_service_url` (String) - Response Modification Service URL
<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

* `create` (String) - A create timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `read` (String) - A read timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).
* `update` (String) - An update timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `delete` (String) - A delete timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_icap.example {alias}
```
