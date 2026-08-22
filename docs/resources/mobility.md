---
page_title: "gigavuecore_mobility Resource - gigavuecore"
subcategory: ""
description: |-
  Load mobility solutions by alias
---

# gigavuecore_mobility Resource

Load mobility solutions by alias

## Example Usage

```terraform
resource "gigavuecore_mobility" "example" {
  sites = []
  solution_alias = null
  solution_type = null
  tags = []
  traffic_policies = {}
}
```

## Schema

### Arguments

The following arguments are supported:

* `sites` (List(Object({alias, cp_nodes, gtp_nodes, network_ports, sam_exporter_nodes, site_override_of_policy_arrangements, skip_deployment, tags, tool_bindings, up_nodes})), required)
* `solution_alias` (String, required) - Alias of the solution
* `solution_type` (String, required) - Type of the solution
* `tags` (List(Object({tag_key, tag_values})), optional) - RBAC Tags
* `traffic_policies` (Object({for5_g, for_lte, for_non_cups_lte}), optional) - Forwarding policies for the processing nodes
  * `for5_g` (Object({gtp_flow_timeout, gtp_persistence, load_balancing, overlap_mode, sampling, whitelisting}), optional)
    * `gtp_flow_timeout` (Number, optional)
    * `gtp_persistence` (Object({enabled, file_age_timeout, interval, restart_age_time}), optional) - GsGroup Gtp Persistence Parameters
      * `enabled` (Bool, optional) - GTP Persistence Status
      * `file_age_timeout` (Number, optional) - GTP Persistence File Age Timeout(mins)
      * `interval` (Number, optional) - GTP Persistence Interval(mins)
      * `restart_age_time` (Number, optional) - GTP Persistence Restart Age Time(mins)
    * `load_balancing` (Object({app_type, hashing_key}), optional)
      * `app_type` (String, required)
      * `hashing_key` (String, required)
    * `overlap_mode` (Bool, optional) - When enabled flow-filtering cannot be configured
    * `sampling` (Object({flow_maps}), optional)
      * `flow_maps` (List(Object({alias, comment, rules, source_group_id, tags, tool})), optional)
    * `whitelisting` (Object({flow_maps, multi_whitelists, white_list_alias}), optional)
      * `flow_maps` (List(Object({alias, comment, rules, source_group_id, tags, tool})), required)
      * `multi_whitelists` (List(String), optional) - Alias/Aliases of referenced GTP Whitelists.Use either whiteListAlias or multiWhitelists.
      * `white_list_alias` (String, optional) - Alias of referenced GTP Whitelist.Use either whiteListAlias or multiWhitelists.
  * `for_lte` (Object({gtp_flow_timeout, gtp_persistence, load_balancing_lte, overlap_mode, sampling, whitelisting}), optional)
    * `gtp_flow_timeout` (Number, optional)
    * `gtp_persistence` (Object({enabled, file_age_timeout, interval, restart_age_time}), optional) - GsGroup Gtp Persistence Parameters
      * `enabled` (Bool, optional) - GTP Persistence Status
      * `file_age_timeout` (Number, optional) - GTP Persistence File Age Timeout(mins)
      * `interval` (Number, optional) - GTP Persistence Interval(mins)
      * `restart_age_time` (Number, optional) - GTP Persistence Restart Age Time(mins)
    * `load_balancing_lte` (Object({app_type, hashing_key}), computed)
      * `app_type` (String, computed)
      * `hashing_key` (String, computed)
    * `overlap_mode` (Bool, optional) - When enabled flow-filtering cannot be configured
    * `sampling` (Object({flow_maps}), optional)
      * `flow_maps` (List(Object({alias, comment, rules, source_group_id, tags, tool})), optional)
    * `whitelisting` (Object({flow_maps, multi_whitelists, white_list_alias}), optional)
      * `flow_maps` (List(Object({alias, comment, rules, source_group_id, tool})), required)
      * `multi_whitelists` (List(String), optional) - Alias/Aliases of referenced GTP Whitelists.Use either whiteListAlias or multiWhitelists.
      * `white_list_alias` (String, optional) - Alias of referenced GTP Whitelist.Use either whiteListAlias or multiWhitelists.
  * `for_non_cups_lte` (Object({flowfiltering, gtp_flow_timeout, gtp_persistence, load_balancing, overlap_mode, sampling, whitelisting}), optional)
    * `flowfiltering` (Object({flow_maps}), optional)
      * `flow_maps` (List(Object({alias, comment, drop_rules, pass_rules, source_group_id, tags, tool})), optional)
    * `gtp_flow_timeout` (Number, optional)
    * `gtp_persistence` (Object({enabled, file_age_timeout, interval, restart_age_time}), optional) - GsGroup Gtp Persistence Parameters
      * `enabled` (Bool, optional) - GTP Persistence Status
      * `file_age_timeout` (Number, optional) - GTP Persistence File Age Timeout(mins)
      * `interval` (Number, optional) - GTP Persistence Interval(mins)
      * `restart_age_time` (Number, optional) - GTP Persistence Restart Age Time(mins)
    * `load_balancing` (Object({app_type, hashing_key}), optional)
      * `app_type` (String, required)
      * `hashing_key` (String, required)
    * `overlap_mode` (Bool, optional) - When enabled flow-filtering cannot be configured
    * `sampling` (Object({flow_maps}), optional)
      * `flow_maps` (List(Object({alias, comment, rules, source_group_id, tags, tool})), optional)
    * `whitelisting` (Object({flow_maps, multi_whitelists, white_list_alias}), optional)
      * `flow_maps` (List(Object({alias, comment, rules, source_group_id, tags, tool})), required)
      * `multi_whitelists` (List(String), optional) - Alias/Aliases of referenced GTP Whitelists.Use either whiteListAlias or multiWhitelists.
      * `white_list_alias` (String, optional) - Alias of referenced GTP Whitelist.Use either whiteListAlias or multiWhitelists.

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `health_state` (String, computed) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (List(Object({message, severity, traffic_health_state_computation_type})), computed)
* `tags` (List(Object({tag_key, tag_values})), computed) - RBAC Tags
* `traffic_policies` (Object({for5_g, for_lte, for_non_cups_lte}), computed) - Forwarding policies for the processing nodes
  * `for5_g` (Object({gtp_flow_timeout, gtp_persistence, load_balancing, overlap_mode, sampling, whitelisting}), optional)
    * `gtp_flow_timeout` (Number, optional)
    * `gtp_persistence` (Object({enabled, file_age_timeout, interval, restart_age_time}), optional) - GsGroup Gtp Persistence Parameters
      * `enabled` (Bool, optional) - GTP Persistence Status
      * `file_age_timeout` (Number, optional) - GTP Persistence File Age Timeout(mins)
      * `interval` (Number, optional) - GTP Persistence Interval(mins)
      * `restart_age_time` (Number, optional) - GTP Persistence Restart Age Time(mins)
    * `load_balancing` (Object({app_type, hashing_key}), optional)
      * `app_type` (String, required)
      * `hashing_key` (String, required)
    * `overlap_mode` (Bool, optional) - When enabled flow-filtering cannot be configured
    * `sampling` (Object({flow_maps}), optional)
      * `flow_maps` (List(Object({alias, comment, rules, source_group_id, tags, tool})), optional)
    * `whitelisting` (Object({flow_maps, multi_whitelists, white_list_alias}), optional)
      * `flow_maps` (List(Object({alias, comment, rules, source_group_id, tags, tool})), required)
      * `multi_whitelists` (List(String), optional) - Alias/Aliases of referenced GTP Whitelists.Use either whiteListAlias or multiWhitelists.
      * `white_list_alias` (String, optional) - Alias of referenced GTP Whitelist.Use either whiteListAlias or multiWhitelists.
  * `for_lte` (Object({gtp_flow_timeout, gtp_persistence, load_balancing_lte, overlap_mode, sampling, whitelisting}), optional)
    * `gtp_flow_timeout` (Number, optional)
    * `gtp_persistence` (Object({enabled, file_age_timeout, interval, restart_age_time}), optional) - GsGroup Gtp Persistence Parameters
      * `enabled` (Bool, optional) - GTP Persistence Status
      * `file_age_timeout` (Number, optional) - GTP Persistence File Age Timeout(mins)
      * `interval` (Number, optional) - GTP Persistence Interval(mins)
      * `restart_age_time` (Number, optional) - GTP Persistence Restart Age Time(mins)
    * `load_balancing_lte` (Object({app_type, hashing_key}), computed)
      * `app_type` (String, computed)
      * `hashing_key` (String, computed)
    * `overlap_mode` (Bool, optional) - When enabled flow-filtering cannot be configured
    * `sampling` (Object({flow_maps}), optional)
      * `flow_maps` (List(Object({alias, comment, rules, source_group_id, tags, tool})), optional)
    * `whitelisting` (Object({flow_maps, multi_whitelists, white_list_alias}), optional)
      * `flow_maps` (List(Object({alias, comment, rules, source_group_id, tool})), required)
      * `multi_whitelists` (List(String), optional) - Alias/Aliases of referenced GTP Whitelists.Use either whiteListAlias or multiWhitelists.
      * `white_list_alias` (String, optional) - Alias of referenced GTP Whitelist.Use either whiteListAlias or multiWhitelists.
  * `for_non_cups_lte` (Object({flowfiltering, gtp_flow_timeout, gtp_persistence, load_balancing, overlap_mode, sampling, whitelisting}), optional)
    * `flowfiltering` (Object({flow_maps}), optional)
      * `flow_maps` (List(Object({alias, comment, drop_rules, pass_rules, source_group_id, tags, tool})), optional)
    * `gtp_flow_timeout` (Number, optional)
    * `gtp_persistence` (Object({enabled, file_age_timeout, interval, restart_age_time}), optional) - GsGroup Gtp Persistence Parameters
      * `enabled` (Bool, optional) - GTP Persistence Status
      * `file_age_timeout` (Number, optional) - GTP Persistence File Age Timeout(mins)
      * `interval` (Number, optional) - GTP Persistence Interval(mins)
      * `restart_age_time` (Number, optional) - GTP Persistence Restart Age Time(mins)
    * `load_balancing` (Object({app_type, hashing_key}), optional)
      * `app_type` (String, required)
      * `hashing_key` (String, required)
    * `overlap_mode` (Bool, optional) - When enabled flow-filtering cannot be configured
    * `sampling` (Object({flow_maps}), optional)
      * `flow_maps` (List(Object({alias, comment, rules, source_group_id, tags, tool})), optional)
    * `whitelisting` (Object({flow_maps, multi_whitelists, white_list_alias}), optional)
      * `flow_maps` (List(Object({alias, comment, rules, source_group_id, tags, tool})), required)
      * `multi_whitelists` (List(String), optional) - Alias/Aliases of referenced GTP Whitelists.Use either whiteListAlias or multiWhitelists.
      * `white_list_alias` (String, optional) - Alias of referenced GTP Whitelist.Use either whiteListAlias or multiWhitelists.

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_mobility.example {solution_alias}
```
