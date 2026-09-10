---
page_title: "gigavuecore_get_all_flex_inline_ssl_apps Data Source - gigavuecore"
subcategory: ""
description: |-
  Get All Inline ssl apps across FM
---

# gigavuecore_get_all_flex_inline_ssl_apps Data Source

Get All Inline ssl apps across FM

## Example Usage

```terraform
data "gigavuecore_get_all_flex_inline_ssl_apps" "example" {
  cluster_id = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, optional) - If provided, ssl apps only for that cluster are returned

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `alias` (String) - Alias of the inline ssl app
* `app_intent_configs` (Attributes) (see [below for nested schema](#nestedatt--items--app_intent_configs))
* `cluster_name` (String)
* `config_status` (String) - Configuration status of this SSL app
* `config_status_reasons` (List of String) - In case of configuration failure, this message provides details about the possible cause of the failure
* `health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--items--health_state_reasons))
* `m_tls` (String) - Enable mTLS to do client Authentication
* `ria_configs` (Attributes List) - Configs applicable if RIA is enabled in ssl app, the listed properties here will not be applicable in appIntent configs if RIA is enabled (see [below for nested schema](#nestedatt--items--ria_configs))
* `ria_enabled` (String) - Enable RIA to configure ssl app in two nodes which is needed for RIA SSL

<a id="nestedatt--items--app_intent_configs"></a>
### Nested Schema for `items.app_intent_configs`

Read-Only:

* `black_list_config` (Attributes) (see [below for nested schema](#nestedatt--items--app_intent_configs--black_list_config))
* `global_default_configs` (Attributes) - inline SSL configuration parameters (see [below for nested schema](#nestedatt--items--app_intent_configs--global_default_configs))
* `gs_engines` (List of String)
* `gs_group_alias` (String) - Alias of the gs group auto created by ssl app
* `gs_group_param_configs` (Attributes) (see [below for nested schema](#nestedatt--items--app_intent_configs--gs_group_param_configs))
* `gsop_alias` (String) - Alias of the gsop auto created by ssl app
* `inline_ssl` (Attributes) - Used to configure other GS apps in addition to Inline SSL on a HC1 box (see [below for nested schema](#nestedatt--items--app_intent_configs--inline_ssl))
* `key_store_configs` (Attributes) (see [below for nested schema](#nestedatt--items--app_intent_configs--key_store_configs))
* `m_tls_configs` (Attributes) (see [below for nested schema](#nestedatt--items--app_intent_configs--m_tls_configs))
* `network_access_configs` (Attributes List) (see [below for nested schema](#nestedatt--items--app_intent_configs--network_access_configs))
* `ssl_path` (Attributes List) (see [below for nested schema](#nestedatt--items--app_intent_configs--ssl_path))
* `ssl_profile_alias` (String) - Alias of the ssl profile auto created by ssl app
* `ssl_profile_config` (Attributes) - Flexible Inline SSL Profile (see [below for nested schema](#nestedatt--items--app_intent_configs--ssl_profile_config))
* `tag_protocol_id` (String) - When tool VLAN tag is added , this protocol Id will be added which egress out the traffic
* `trust_store_configs` (Attributes) (see [below for nested schema](#nestedatt--items--app_intent_configs--trust_store_configs))
* `vlan_id` (Number) - Vlan id for proxy maps
* `vport_alias` (String) - Alias of the vport auto created by ssl app
* `vport_config` (Attributes) - GigaSMART vPort (see [below for nested schema](#nestedatt--items--app_intent_configs--vport_config))
* `white_list_config` (Attributes) (see [below for nested schema](#nestedatt--items--app_intent_configs--white_list_config))

<a id="nestedatt--items--app_intent_configs--black_list_config"></a>
### Nested Schema for `items.app_intent_configs.black_list_config`

Read-Only:

* `operation` (String)
* `profile_list` (Attributes) (see [below for nested schema](#nestedatt--items--app_intent_configs--black_list_config--profile_list))

<a id="nestedatt--items--app_intent_configs--black_list_config--profile_list"></a>
### Nested Schema for `items.app_intent_configs.black_list_config.profile_list`

Read-Only:

* `file_source` (Attributes) - Remote file source or destination (see [below for nested schema](#nestedatt--items--app_intent_configs--black_list_config--profile_list--file_source))
* `list` (String) - The nodecryptlist or decryptlist . Mutually exclusive with 'fileSource'

<a id="nestedatt--items--app_intent_configs--black_list_config--profile_list--file_source"></a>
### Nested Schema for `items.app_intent_configs.black_list_config.profile_list.file_source`

Read-Only:

* `hostname` (String) - server address
* `password` (String) - password to use for server login
* `path` (String) - file path on server
* `protocol` (String) - Protocol to access server. HTTP and HTTPS only applicable for retrieving file
* `username` (String) - user name to use for server login

<a id="nestedatt--items--app_intent_configs--global_default_configs"></a>
### Nested Schema for `items.app_intent_configs.global_default_configs`

Read-Only:

* `caching` (Attributes) (see [below for nested schema](#nestedatt--items--app_intent_configs--global_default_configs--caching))
* `cluster_id` (String) - id of the defining cluster
* `dhe_ciphersuit` (String)
* `monitor` (Attributes) - This property is moved to ssl profile for device version >=5.7 (see [below for nested schema](#nestedatt--items--app_intent_configs--global_default_configs--monitor))
* `resumption` (Attributes) (see [below for nested schema](#nestedatt--items--app_intent_configs--global_default_configs--resumption))
* `ssl_versions` (Attributes) (see [below for nested schema](#nestedatt--items--app_intent_configs--global_default_configs--ssl_versions))
* `start_tls` (Attributes) (see [below for nested schema](#nestedatt--items--app_intent_configs--global_default_configs--start_tls))

<a id="nestedatt--items--app_intent_configs--global_default_configs--caching"></a>
### Nested Schema for `items.app_intent_configs.global_default_configs.caching`

Read-Only:

* `persistence` (Attributes) (see [below for nested schema](#nestedatt--items--app_intent_configs--global_default_configs--caching--persistence))

<a id="nestedatt--items--app_intent_configs--global_default_configs--caching--persistence"></a>
### Nested Schema for `items.app_intent_configs.global_default_configs.caching.persistence`

Read-Only:

* `enable` (Boolean)

<a id="nestedatt--items--app_intent_configs--global_default_configs--monitor"></a>
### Nested Schema for `items.app_intent_configs.global_default_configs.monitor`

Read-Only:

* `enable` (Boolean)

<a id="nestedatt--items--app_intent_configs--global_default_configs--resumption"></a>
### Nested Schema for `items.app_intent_configs.global_default_configs.resumption`

Read-Only:

* `client` (Attributes) - enable client initiated resumption (for debug purposes only) (see [below for nested schema](#nestedatt--items--app_intent_configs--global_default_configs--resumption--client))

<a id="nestedatt--items--app_intent_configs--global_default_configs--resumption--client"></a>
### Nested Schema for `items.app_intent_configs.global_default_configs.resumption.client`

Read-Only:

* `enable` (Boolean)

<a id="nestedatt--items--app_intent_configs--global_default_configs--ssl_versions"></a>
### Nested Schema for `items.app_intent_configs.global_default_configs.ssl_versions`

Read-Only:

* `connection_reset_action_for_max_version` (String) - Action to take to reset connection if TLS version is higher than configured
* `connection_reset_action_for_min_version` (String) - Action to take to reset connection if TLS version is lower than configured
* `max_version` (String) - maxVersion must be greater than minVersion
* `min_version` (String)

<a id="nestedatt--items--app_intent_configs--global_default_configs--start_tls"></a>
### Nested Schema for `items.app_intent_configs.global_default_configs.start_tls`

Read-Only:

* `enable` (Boolean)

<a id="nestedatt--items--app_intent_configs--gs_group_param_configs"></a>
### Nested Schema for `items.app_intent_configs.gs_group_param_configs`

Read-Only:

* `hsm_group` (String) - Alias of the hsmGroup
* `session_logging` (Attributes) - GsGroup Session Logging Configuration (see [below for nested schema](#nestedatt--items--app_intent_configs--gs_group_param_configs--session_logging))

<a id="nestedatt--items--app_intent_configs--gs_group_param_configs--session_logging"></a>
### Nested Schema for `items.app_intent_configs.gs_group_param_configs.session_logging`

Read-Only:

* `interface` (String) - Associated IP Interface
* `log_level` (String) - Log Level
* `remote_syslog_ip` (String) - Remote Syslog IP
* `remote_syslog_port` (Number) - Remote Syslog Port Number

<a id="nestedatt--items--app_intent_configs--inline_ssl"></a>
### Nested Schema for `items.app_intent_configs.inline_ssl`

Read-Only:

* `standalone` (Boolean) - If enabled , behaves in a normal way. If disabled, can configure other GS apps along side Inline ssl with Inline ssl occupying only 50% of the memory

<a id="nestedatt--items--app_intent_configs--key_store_configs"></a>
### Nested Schema for `items.app_intent_configs.key_store_configs`

Read-Only:

* `deployment_type` (String)
* `inbound_keys` (Attributes List) (see [below for nested schema](#nestedatt--items--app_intent_configs--key_store_configs--inbound_keys))
* `outboundkeys` (Attributes List) (see [below for nested schema](#nestedatt--items--app_intent_configs--key_store_configs--outboundkeys))

<a id="nestedatt--items--app_intent_configs--key_store_configs--inbound_keys"></a>
### Nested Schema for `items.app_intent_configs.key_store_configs.inbound_keys`

Read-Only:

* `key_alias` (String) - Key to decrypt the traffic for the associated server/domain address
* `server_domain_alias` (String)

<a id="nestedatt--items--app_intent_configs--key_store_configs--outboundkeys"></a>
### Nested Schema for `items.app_intent_configs.key_store_configs.outboundkeys`

Read-Only:

* `key_alias` (String)
* `signing_for` (String)

<a id="nestedatt--items--app_intent_configs--m_tls_configs"></a>
### Nested Schema for `items.app_intent_configs.m_tls_configs`

Read-Only:

* `primary_signing` (String) - Alias of the key from keyStore used for primary signing for mTLS
* `secondary_signing` (String) - Alias of the key from keyStore used for secondary signing for mTLS
* `trust_store` (String) - Alias of the client trust store

<a id="nestedatt--items--app_intent_configs--network_access_configs"></a>
### Nested Schema for `items.app_intent_configs.network_access_configs`

Read-Only:

* `network_access` (Attributes) (see [below for nested schema](#nestedatt--items--app_intent_configs--network_access_configs--network_access))
* `operation` (String) - Adds or deletes(None) engine interface connectivity

<a id="nestedatt--items--app_intent_configs--network_access_configs--network_access"></a>
### Nested Schema for `items.app_intent_configs.network_access_configs.network_access`

Read-Only:

* `cluster_id` (String) - id of the defining cluster
* `dhcp` (Boolean)
* `dns` (String) - Only valid and required when 'dhcp' is false
* `eport` (String) - GigaSMART engine port
* `gateway` (String) - Only valid and required when 'dhcp' is false
* `hw_address` (String)
* `interface` (String)
* `ip_address` (String) - Only valid and required when 'dhcp' is false
* `ip_mask` (String) - Only valid and required when 'dhcp' is false
* `mtu` (Number) - Only valid when 'dhcp' is false
* `proxy_server_profile` (String)
* `status` (String)
* `vlan` (Number)

<a id="nestedatt--items--app_intent_configs--ssl_path"></a>
### Nested Schema for `items.app_intent_configs.ssl_path`

Read-Only:

* `alias` (String)
* `flex_inline_map` (Attributes) - When 'oobCopy' is configured at least one of 'aToB' or 'bToA' must be configured (see [below for nested schema](#nestedatt--items--app_intent_configs--ssl_path--flex_inline_map))

<a id="nestedatt--items--app_intent_configs--ssl_path--flex_inline_map"></a>
### Nested Schema for `items.app_intent_configs.ssl_path.flex_inline_map`

Read-Only:

* `a_to_b` (Attributes) (see [below for nested schema](#nestedatt--items--app_intent_configs--ssl_path--flex_inline_map--a_to_b))
* `b_to_a` (Attributes) (see [below for nested schema](#nestedatt--items--app_intent_configs--ssl_path--flex_inline_map--b_to_a))
* `oob_copy` (Attributes List) (see [below for nested schema](#nestedatt--items--app_intent_configs--ssl_path--flex_inline_map--oob_copy))
* `svt_mode` (Boolean)
* `svt_tag` (Number) - only applicable when svtMode is enabled
* `tag` (Attributes) (see [below for nested schema](#nestedatt--items--app_intent_configs--ssl_path--flex_inline_map--tag))

<a id="nestedatt--items--app_intent_configs--ssl_path--flex_inline_map--a_to_b"></a>
### Nested Schema for `items.app_intent_configs.ssl_path.flex_inline_map.a_to_b`

Read-Only:

* `ib_pathway` (String) - ibPathway alias. Only applicable when type is 'ibPathway'
* `tools` (List of String) - ordered list of inline tools or vports. Only applicable when 'type' is 'tools'
* `type` (String) - when set to 'tools', the list of the processing inline tools have to be provided. when set to 'ibPathway', ibPathway alias should be provided

<a id="nestedatt--items--app_intent_configs--ssl_path--flex_inline_map--b_to_a"></a>
### Nested Schema for `items.app_intent_configs.ssl_path.flex_inline_map.b_to_a`

Read-Only:

* `ib_pathway` (String) - ibPathway alias. Only applicable when type is 'ibPathway'
* `tools` (List of String) - ordered list of inline tools or vports. Only applicable when 'type' is 'tools'
* `type` (String) - when set to 'tools', the list of the processing inline tools have to be provided. when set to 'ibPathway', ibPathway alias should be provided

<a id="nestedatt--items--app_intent_configs--ssl_path--flex_inline_map--oob_copy"></a>
### Nested Schema for `items.app_intent_configs.ssl_path.flex_inline_map.oob_copy`

Read-Only:

* `direction` (String)
* `dst_ports` (List of String) - list of destination tool ports
* `src_ports` (List of String) - inline network or an item from a-to-b and b-to-a lists
* `tag` (Attributes) (see [below for nested schema](#nestedatt--items--app_intent_configs--ssl_path--flex_inline_map--oob_copy--tag))

<a id="nestedatt--items--app_intent_configs--ssl_path--flex_inline_map--oob_copy--tag"></a>
### Nested Schema for `items.app_intent_configs.ssl_path.flex_inline_map.oob_copy.tag`

Read-Only:

* `type` (String)

<a id="nestedatt--items--app_intent_configs--ssl_path--flex_inline_map--tag"></a>
### Nested Schema for `items.app_intent_configs.ssl_path.flex_inline_map.tag`

Read-Only:

* `tag_protocol_id` (String) - When tool VLAN tag is added , this protocol Id will be added which egress out the traffic
* `type` (String)
* `vlan_id` (Number) - only applicable when type is 'vlan'

<a id="nestedatt--items--app_intent_configs--ssl_profile_config"></a>
### Nested Schema for `items.app_intent_configs.ssl_profile_config`

Read-Only:

* `alias` (String) - Alias of the inline SSL profile
* `certificate` (Attributes) - SSL profile certificate configuration (see [below for nested schema](#nestedatt--items--app_intent_configs--ssl_profile_config--certificate))
* `client_auth` (Attributes) - SSL profile certificate configuration (see [below for nested schema](#nestedatt--items--app_intent_configs--ssl_profile_config--client_auth))
* `cluster_id` (String) - id of the defining cluster
* `decrypt` (Attributes) - SSL profile configuration on decrypt action (see [below for nested schema](#nestedatt--items--app_intent_configs--ssl_profile_config--decrypt))
* `default_action` (String) - Action to take if none of the profile rules match
* `high_avail` (Attributes) - SSL profile configuration on high availability (see [below for nested schema](#nestedatt--items--app_intent_configs--ssl_profile_config--high_avail))
* `key_map` (Attributes List) (see [below for nested schema](#nestedatt--items--app_intent_configs--ssl_profile_config--key_map))
* `monitor` (String)
* `network_group` (Attributes) - SSL Profile configuration for multiple entry in network groups (see [below for nested schema](#nestedatt--items--app_intent_configs--ssl_profile_config--network_group))
* `no_decrypt` (Attributes) - SSL profile configuration on no-decrypt action (see [below for nested schema](#nestedatt--items--app_intent_configs--ssl_profile_config--no_decrypt))
* `non_ssl_tcp` (Attributes) - SSL profile configuration on TCP proxy action (see [below for nested schema](#nestedatt--items--app_intent_configs--ssl_profile_config--non_ssl_tcp))
* `rules` (List of Dynamic) - inline SSL profile rules
* `split_proxy` (Attributes) - SSL profile configuration for Split Proxy (see [below for nested schema](#nestedatt--items--app_intent_configs--ssl_profile_config--split_proxy))
* `start_tls` (Attributes) - SSL profile configuration on start TLS action (see [below for nested schema](#nestedatt--items--app_intent_configs--ssl_profile_config--start_tls))
* `tcp` (Attributes) - SSL profile configuration on TCP (see [below for nested schema](#nestedatt--items--app_intent_configs--ssl_profile_config--tcp))
* `tool` (Attributes) - SSL Profile configuration for Tools (see [below for nested schema](#nestedatt--items--app_intent_configs--ssl_profile_config--tool))
* `url_cache` (Attributes) - SSL profile configuration on url-cache (see [below for nested schema](#nestedatt--items--app_intent_configs--ssl_profile_config--url_cache))

<a id="nestedatt--items--app_intent_configs--ssl_profile_config--certificate"></a>
### Nested Schema for `items.app_intent_configs.ssl_profile_config.certificate`

Read-Only:

* `expired` (String) - SSL profile on expired certificate
* `invalid` (String) - SSL profile on invalid certificate
* `revocation` (Attributes) - SSL profile certificate revocation configuration (see [below for nested schema](#nestedatt--items--app_intent_configs--ssl_profile_config--certificate--revocation))
* `self_signed` (String) - SSL profile on self-signed certificate
* `unknown_ca` (String) - SSL profile on unknown CA certificate

<a id="nestedatt--items--app_intent_configs--ssl_profile_config--certificate--revocation"></a>
### Nested Schema for `items.app_intent_configs.ssl_profile_config.certificate.revocation`

Read-Only:

* `crl` (Attributes) - SSL profile certificate revocation configuration (see [below for nested schema](#nestedatt--items--app_intent_configs--ssl_profile_config--certificate--revocation--crl))
* `ocsp` (Attributes) - SSL profile certificate revocation configuration (see [below for nested schema](#nestedatt--items--app_intent_configs--ssl_profile_config--certificate--revocation--ocsp))

<a id="nestedatt--items--app_intent_configs--ssl_profile_config--certificate--revocation--crl"></a>
### Nested Schema for `items.app_intent_configs.ssl_profile_config.certificate.revocation.crl`

Read-Only:

* `defer` (Number) - timeout in seconds
* `enabled` (Boolean)
* `fail` (String) - only applicable when enabled

<a id="nestedatt--items--app_intent_configs--ssl_profile_config--certificate--revocation--ocsp"></a>
### Nested Schema for `items.app_intent_configs.ssl_profile_config.certificate.revocation.ocsp`

Read-Only:

* `defer` (Number) - timeout in seconds
* `enabled` (Boolean)
* `fail` (String) - only applicable when enabled

<a id="nestedatt--items--app_intent_configs--ssl_profile_config--client_auth"></a>
### Nested Schema for `items.app_intent_configs.ssl_profile_config.client_auth`

Read-Only:

* `expired` (String) - SSL profile on expired certificate
* `invalid` (String) - SSL profile on invalid certificate
* `revocation` (Attributes) - SSL profile certificate revocation configuration (see [below for nested schema](#nestedatt--items--app_intent_configs--ssl_profile_config--client_auth--revocation))
* `self_signed` (String) - SSL profile on self-signed certificate
* `unknown_ca` (String) - SSL profile on unknown CA certificate

<a id="nestedatt--items--app_intent_configs--ssl_profile_config--client_auth--revocation"></a>
### Nested Schema for `items.app_intent_configs.ssl_profile_config.client_auth.revocation`

Read-Only:

* `crl` (Attributes) - SSL profile certificate revocation configuration (see [below for nested schema](#nestedatt--items--app_intent_configs--ssl_profile_config--client_auth--revocation--crl))
* `ocsp` (Attributes) - SSL profile certificate revocation configuration (see [below for nested schema](#nestedatt--items--app_intent_configs--ssl_profile_config--client_auth--revocation--ocsp))

<a id="nestedatt--items--app_intent_configs--ssl_profile_config--client_auth--revocation--crl"></a>
### Nested Schema for `items.app_intent_configs.ssl_profile_config.client_auth.revocation.crl`

Read-Only:

* `defer` (Number) - timeout in seconds
* `enabled` (Boolean)
* `fail` (String) - only applicable when enabled

<a id="nestedatt--items--app_intent_configs--ssl_profile_config--client_auth--revocation--ocsp"></a>
### Nested Schema for `items.app_intent_configs.ssl_profile_config.client_auth.revocation.ocsp`

Read-Only:

* `defer` (Number) - timeout in seconds
* `enabled` (Boolean)
* `fail` (String) - only applicable when enabled

<a id="nestedatt--items--app_intent_configs--ssl_profile_config--decrypt"></a>
### Nested Schema for `items.app_intent_configs.ssl_profile_config.decrypt`

Read-Only:

* `tcp` (Attributes) - SSL decryption TCP control (see [below for nested schema](#nestedatt--items--app_intent_configs--ssl_profile_config--decrypt--tcp))
* `tool_bypass` (Attributes) (see [below for nested schema](#nestedatt--items--app_intent_configs--ssl_profile_config--decrypt--tool_bypass))

<a id="nestedatt--items--app_intent_configs--ssl_profile_config--decrypt--tcp"></a>
### Nested Schema for `items.app_intent_configs.ssl_profile_config.decrypt.tcp`

Read-Only:

* `inactive_timeout` (Number) - SSL decryption TCP inactive timeout (in minutes)
* `port_map` (Attributes) - SSL decryption port map (see [below for nested schema](#nestedatt--items--app_intent_configs--ssl_profile_config--decrypt--tcp--port_map))

<a id="nestedatt--items--app_intent_configs--ssl_profile_config--decrypt--tcp--port_map"></a>
### Nested Schema for `items.app_intent_configs.ssl_profile_config.decrypt.tcp.port_map`

Read-Only:

* `default_out_port` (Number) - egress port for decryption port map. 0 is disabled.
* `ports` (Attributes List) (see [below for nested schema](#nestedatt--items--app_intent_configs--ssl_profile_config--decrypt--tcp--port_map--ports))

<a id="nestedatt--items--app_intent_configs--ssl_profile_config--decrypt--tcp--port_map--ports"></a>
### Nested Schema for `items.app_intent_configs.ssl_profile_config.decrypt.tcp.port_map.ports`

Read-Only:

* `in_port` (Number) - ingress port for decryption port map
* `out_port` (Number) - egress port for decryption port map
* `rule_id` (Number)

<a id="nestedatt--items--app_intent_configs--ssl_profile_config--decrypt--tool_bypass"></a>
### Nested Schema for `items.app_intent_configs.ssl_profile_config.decrypt.tool_bypass`

Read-Only:

* `enable` (Boolean)

<a id="nestedatt--items--app_intent_configs--ssl_profile_config--high_avail"></a>
### Nested Schema for `items.app_intent_configs.ssl_profile_config.high_avail`

Read-Only:

* `active_standby` (Attributes) (see [below for nested schema](#nestedatt--items--app_intent_configs--ssl_profile_config--high_avail--active_standby))

<a id="nestedatt--items--app_intent_configs--ssl_profile_config--high_avail--active_standby"></a>
### Nested Schema for `items.app_intent_configs.ssl_profile_config.high_avail.active_standby`

Read-Only:

* `enable` (Boolean)

<a id="nestedatt--items--app_intent_configs--ssl_profile_config--key_map"></a>
### Nested Schema for `items.app_intent_configs.ssl_profile_config.key_map`

Read-Only:

* `hostname` (String) - hostname or IP address
* `key` (String) - SSL key alias
* `rule_id` (Number)

<a id="nestedatt--items--app_intent_configs--ssl_profile_config--network_group"></a>
### Nested Schema for `items.app_intent_configs.ssl_profile_config.network_group`

Read-Only:

* `multiple_entry` (Attributes) (see [below for nested schema](#nestedatt--items--app_intent_configs--ssl_profile_config--network_group--multiple_entry))

<a id="nestedatt--items--app_intent_configs--ssl_profile_config--network_group--multiple_entry"></a>
### Nested Schema for `items.app_intent_configs.ssl_profile_config.network_group.multiple_entry`

Read-Only:

* `enable` (Boolean)

<a id="nestedatt--items--app_intent_configs--ssl_profile_config--no_decrypt"></a>
### Nested Schema for `items.app_intent_configs.ssl_profile_config.no_decrypt`

Read-Only:

* `tool_bypass` (Attributes) (see [below for nested schema](#nestedatt--items--app_intent_configs--ssl_profile_config--no_decrypt--tool_bypass))

<a id="nestedatt--items--app_intent_configs--ssl_profile_config--no_decrypt--tool_bypass"></a>
### Nested Schema for `items.app_intent_configs.ssl_profile_config.no_decrypt.tool_bypass`

Read-Only:

* `enable` (Boolean)

<a id="nestedatt--items--app_intent_configs--ssl_profile_config--non_ssl_tcp"></a>
### Nested Schema for `items.app_intent_configs.ssl_profile_config.non_ssl_tcp`

Read-Only:

* `tool_bypass` (Attributes) (see [below for nested schema](#nestedatt--items--app_intent_configs--ssl_profile_config--non_ssl_tcp--tool_bypass))

<a id="nestedatt--items--app_intent_configs--ssl_profile_config--non_ssl_tcp--tool_bypass"></a>
### Nested Schema for `items.app_intent_configs.ssl_profile_config.non_ssl_tcp.tool_bypass`

Read-Only:

* `enable` (Boolean)

<a id="nestedatt--items--app_intent_configs--ssl_profile_config--split_proxy"></a>
### Nested Schema for `items.app_intent_configs.ssl_profile_config.split_proxy`

Read-Only:

* `mode` (Attributes) (see [below for nested schema](#nestedatt--items--app_intent_configs--ssl_profile_config--split_proxy--mode))
* `server_non_pfs_ciphers` (Attributes) (see [below for nested schema](#nestedatt--items--app_intent_configs--ssl_profile_config--split_proxy--server_non_pfs_ciphers))

<a id="nestedatt--items--app_intent_configs--ssl_profile_config--split_proxy--mode"></a>
### Nested Schema for `items.app_intent_configs.ssl_profile_config.split_proxy.mode`

Read-Only:

* `enable` (Boolean)

<a id="nestedatt--items--app_intent_configs--ssl_profile_config--split_proxy--server_non_pfs_ciphers"></a>
### Nested Schema for `items.app_intent_configs.ssl_profile_config.split_proxy.server_non_pfs_ciphers`

Read-Only:

* `enable` (Boolean)

<a id="nestedatt--items--app_intent_configs--ssl_profile_config--start_tls"></a>
### Nested Schema for `items.app_intent_configs.ssl_profile_config.start_tls`

Read-Only:

* `l4_port` (List of Number)

<a id="nestedatt--items--app_intent_configs--ssl_profile_config--tcp"></a>
### Nested Schema for `items.app_intent_configs.ssl_profile_config.tcp`

Read-Only:

* `delayed_ack` (Boolean) - enable/disable TCP delayed ACK
* `syn_retries` (Number) - TCP Sync retries
* `timewait_timeout` (Number) - TCP Wait Timeout value

<a id="nestedatt--items--app_intent_configs--ssl_profile_config--tool"></a>
### Nested Schema for `items.app_intent_configs.ssl_profile_config.tool`

Read-Only:

* `early_engage` (Boolean) - enable/disable tool early engage
* `fail_action` (String) - Action to take if the tool fails

<a id="nestedatt--items--app_intent_configs--ssl_profile_config--url_cache"></a>
### Nested Schema for `items.app_intent_configs.ssl_profile_config.url_cache`

Read-Only:

* `miss_action` (String) - The action to take if local URL category resolution misses
* `timeout` (Number) - defer timeout in seconds. Only applicable for missAction 'defer'

<a id="nestedatt--items--app_intent_configs--trust_store_configs"></a>
### Nested Schema for `items.app_intent_configs.trust_store_configs`

Read-Only:

* `trust_store_append_configs` (Attributes) (see [below for nested schema](#nestedatt--items--app_intent_configs--trust_store_configs--trust_store_append_configs))
* `trust_store_replace_configs` (Attributes) (see [below for nested schema](#nestedatt--items--app_intent_configs--trust_store_configs--trust_store_replace_configs))

<a id="nestedatt--items--app_intent_configs--trust_store_configs--trust_store_append_configs"></a>
### Nested Schema for `items.app_intent_configs.trust_store_configs.trust_store_append_configs`

Read-Only:

* `file` (String) - The contents of the file. Mutually exclusive with 'fileSource'
* `file_source` (Attributes) - Remote file source or destination (see [below for nested schema](#nestedatt--items--app_intent_configs--trust_store_configs--trust_store_append_configs--file_source))

<a id="nestedatt--items--app_intent_configs--trust_store_configs--trust_store_append_configs--file_source"></a>
### Nested Schema for `items.app_intent_configs.trust_store_configs.trust_store_append_configs.file_source`

Read-Only:

* `hostname` (String) - server address
* `password` (String) - password to use for server login
* `path` (String) - file path on server
* `protocol` (String) - Protocol to access server. HTTP and HTTPS only applicable for retrieving file
* `username` (String) - user name to use for server login

<a id="nestedatt--items--app_intent_configs--trust_store_configs--trust_store_replace_configs"></a>
### Nested Schema for `items.app_intent_configs.trust_store_configs.trust_store_replace_configs`

Read-Only:

* `file` (String) - The contents of the file. Mutually exclusive with 'fileSource'
* `file_source` (Attributes) - Remote file source or destination (see [below for nested schema](#nestedatt--items--app_intent_configs--trust_store_configs--trust_store_replace_configs--file_source))

<a id="nestedatt--items--app_intent_configs--trust_store_configs--trust_store_replace_configs--file_source"></a>
### Nested Schema for `items.app_intent_configs.trust_store_configs.trust_store_replace_configs.file_source`

Read-Only:

* `hostname` (String) - server address
* `password` (String) - password to use for server login
* `path` (String) - file path on server
* `protocol` (String) - Protocol to access server. HTTP and HTTPS only applicable for retrieving file
* `username` (String) - user name to use for server login

<a id="nestedatt--items--app_intent_configs--vport_config"></a>
### Nested Schema for `items.app_intent_configs.vport_config`

Read-Only:

* `alias` (String)
* `deferred_binding` (Boolean) - enable/disable deferred-binding
* `fail_over_action` (String)
* `gs_group` (String) - Alias of referenced managing GsGroup
* `health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--items--app_intent_configs--vport_config--health_state_reasons))
* `inline_status` (String)
* `inner_traffic_path` (String) - Similar to inline-network traffic-path, applicable for inner map
* `metadata_monitoring` (Attributes) (see [below for nested schema](#nestedatt--items--app_intent_configs--vport_config--metadata_monitoring))
* `mode` (String)
* `outer_traffic_path` (String) - Similar to inline-tool/inline-tool-group flex-traffic path, applicable for outer map
* `sa_apf_profile` (String) - ASF session profile

<a id="nestedatt--items--app_intent_configs--vport_config--health_state_reasons"></a>
### Nested Schema for `items.app_intent_configs.vport_config.health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type

<a id="nestedatt--items--app_intent_configs--vport_config--metadata_monitoring"></a>
### Nested Schema for `items.app_intent_configs.vport_config.metadata_monitoring`

Read-Only:

* `action` (String) - metadata monitoring action
* `exporters` (List of String)

<a id="nestedatt--items--app_intent_configs--white_list_config"></a>
### Nested Schema for `items.app_intent_configs.white_list_config`

Read-Only:

* `operation` (String)
* `profile_list` (Attributes) (see [below for nested schema](#nestedatt--items--app_intent_configs--white_list_config--profile_list))

<a id="nestedatt--items--app_intent_configs--white_list_config--profile_list"></a>
### Nested Schema for `items.app_intent_configs.white_list_config.profile_list`

Read-Only:

* `file_source` (Attributes) - Remote file source or destination (see [below for nested schema](#nestedatt--items--app_intent_configs--white_list_config--profile_list--file_source))
* `list` (String) - The nodecryptlist or decryptlist . Mutually exclusive with 'fileSource'

<a id="nestedatt--items--app_intent_configs--white_list_config--profile_list--file_source"></a>
### Nested Schema for `items.app_intent_configs.white_list_config.profile_list.file_source`

Read-Only:

* `hostname` (String) - server address
* `password` (String) - password to use for server login
* `path` (String) - file path on server
* `protocol` (String) - Protocol to access server. HTTP and HTTPS only applicable for retrieving file
* `username` (String) - user name to use for server login

<a id="nestedatt--items--health_state_reasons"></a>
### Nested Schema for `items.health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type

<a id="nestedatt--items--ria_configs"></a>
### Nested Schema for `items.ria_configs`

Read-Only:

* `cluster_name` (String)
* `gs_engines` (List of String)
* `gs_group_alias` (String) - Alias of the gs group auto created by ssl app
* `gsop_alias` (String) - Alias of the gsop auto created by ssl app
* `network_access_configs` (Attributes List) (see [below for nested schema](#nestedatt--items--ria_configs--network_access_configs))
* `ssl_profile_alias` (String) - Alias of the ssl profile auto created by ssl app
* `vport_alias` (String) - Alias of the vport auto created by ssl app

<a id="nestedatt--items--ria_configs--network_access_configs"></a>
### Nested Schema for `items.ria_configs.network_access_configs`

Read-Only:

* `network_access` (Attributes) (see [below for nested schema](#nestedatt--items--ria_configs--network_access_configs--network_access))
* `operation` (String) - Adds or deletes(None) engine interface connectivity

<a id="nestedatt--items--ria_configs--network_access_configs--network_access"></a>
### Nested Schema for `items.ria_configs.network_access_configs.network_access`

Read-Only:

* `cluster_id` (String) - id of the defining cluster
* `dhcp` (Boolean)
* `dns` (String) - Only valid and required when 'dhcp' is false
* `eport` (String) - GigaSMART engine port
* `gateway` (String) - Only valid and required when 'dhcp' is false
* `hw_address` (String)
* `interface` (String)
* `ip_address` (String) - Only valid and required when 'dhcp' is false
* `ip_mask` (String) - Only valid and required when 'dhcp' is false
* `mtu` (Number) - Only valid when 'dhcp' is false
* `proxy_server_profile` (String)
* `status` (String)
* `vlan` (Number)

