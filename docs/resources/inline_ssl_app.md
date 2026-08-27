---
page_title: "gigavuecore_inline_ssl_app Resource - gigavuecore"
subcategory: ""
description: |-
  Get inline ssl app details by alias
---

# gigavuecore_inline_ssl_app Resource

Get inline ssl app details by alias

## Example Usage

```terraform
resource "gigavuecore_inline_ssl_app" "example" {
  alias                 = null
  app_intent_configs    = {}
  cluster_id            = null
  cluster_name          = null
  config_status         = null
  config_status_reasons = []
  health_state          = null
  health_state_reasons  = []
  m_tls                 = null
  ria_configs           = []
  ria_enabled           = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, optional) - Alias of the inline ssl app
* `app_intent_configs` (Attributes, optional) (see [below for nested schema](#nestedatt--app_intent_configs))
* `cluster_id` (String, optional) - Inline SSL app only for that cluster is returned
* `cluster_name` (String, optional)
* `config_status` (String, optional) - Configuration status of this SSL app
* `config_status_reasons` (List of String, optional) - In case of configuration failure, this message provides details about the possible cause of the failure
* `health_state` (String, optional) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List, optional) (see [below for nested schema](#nestedatt--health_state_reasons))
* `m_tls` (String, optional) - Enable mTLS to do client Authentication
* `ria_configs` (Attributes List, optional) - Configs applicable if RIA is enabled in ssl app, the listed properties here will not be applicable in appIntent configs if RIA is enabled (see [below for nested schema](#nestedatt--ria_configs))
* `ria_enabled` (String, optional) - Enable RIA to configure ssl app in two nodes which is needed for RIA SSL

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `alias` (String, computed) - Alias of the inline ssl app
* `app_intent_configs` (Attributes, computed) (see [below for nested schema](#nestedatt--app_intent_configs))
* `cluster_name` (String, computed)
* `config_status` (String, computed) - Configuration status of this SSL app
* `config_status_reasons` (List of String, computed) - In case of configuration failure, this message provides details about the possible cause of the failure
* `health_state` (String, computed) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List, computed) (see [below for nested schema](#nestedatt--health_state_reasons))
* `m_tls` (String, computed) - Enable mTLS to do client Authentication
* `ria_configs` (Attributes List, computed) - Configs applicable if RIA is enabled in ssl app, the listed properties here will not be applicable in appIntent configs if RIA is enabled (see [below for nested schema](#nestedatt--ria_configs))
* `ria_enabled` (String, computed) - Enable RIA to configure ssl app in two nodes which is needed for RIA SSL

<a id="nestedatt--app_intent_configs"></a>
### Nested Schema for `app_intent_configs`

Optional:

* `black_list_config` (Attributes) (see [below for nested schema](#nestedatt--app_intent_configs--black_list_config))
* `global_default_configs` (Attributes) - inline SSL configuration parameters (see [below for nested schema](#nestedatt--app_intent_configs--global_default_configs))
* `gs_engines` (List of String)
* `gs_group_alias` (String) - Alias of the gs group auto created by ssl app
* `gs_group_param_configs` (Attributes) (see [below for nested schema](#nestedatt--app_intent_configs--gs_group_param_configs))
* `gsop_alias` (String) - Alias of the gsop auto created by ssl app
* `inline_ssl` (Attributes) - Used to configure other GS apps in addition to Inline SSL on a HC1 box (see [below for nested schema](#nestedatt--app_intent_configs--inline_ssl))
* `key_store_configs` (Attributes) (see [below for nested schema](#nestedatt--app_intent_configs--key_store_configs))
* `m_tls_configs` (Attributes) (see [below for nested schema](#nestedatt--app_intent_configs--m_tls_configs))
* `network_access_configs` (Attributes List) (see [below for nested schema](#nestedatt--app_intent_configs--network_access_configs))
* `ssl_path` (Attributes List) (see [below for nested schema](#nestedatt--app_intent_configs--ssl_path))
* `ssl_profile_alias` (String) - Alias of the ssl profile auto created by ssl app
* `ssl_profile_config` (Attributes) - Flexible Inline SSL Profile (see [below for nested schema](#nestedatt--app_intent_configs--ssl_profile_config))
* `tag_protocol_id` (String) - When tool VLAN tag is added , this protocol Id will be added which egress out the traffic
* `trust_store_configs` (Attributes) (see [below for nested schema](#nestedatt--app_intent_configs--trust_store_configs))
* `vlan_id` (Number) - Vlan id for proxy maps
* `vport_alias` (String) - Alias of the vport auto created by ssl app
* `vport_config` (Attributes) - GigaSMART vPort (see [below for nested schema](#nestedatt--app_intent_configs--vport_config))
* `white_list_config` (Attributes) (see [below for nested schema](#nestedatt--app_intent_configs--white_list_config))
<a id="nestedatt--app_intent_configs--black_list_config"></a>
### Nested Schema for `app_intent_configs.black_list_config`

Optional:

* `operation` (String)
* `profile_list` (Attributes) (see [below for nested schema](#nestedatt--app_intent_configs--black_list_config--profile_list))
<a id="nestedatt--app_intent_configs--black_list_config--profile_list"></a>
### Nested Schema for `app_intent_configs.black_list_config.profile_list`

Optional:

* `file_source` (Attributes) - Remote file source or destination (see [below for nested schema](#nestedatt--app_intent_configs--black_list_config--profile_list--file_source))
* `list` (String) - The nodecryptlist or decryptlist . Mutually exclusive with 'fileSource'
<a id="nestedatt--app_intent_configs--black_list_config--profile_list--file_source"></a>
### Nested Schema for `app_intent_configs.black_list_config.profile_list.file_source`

Required:

* `hostname` (String) - server address
* `path` (String) - file path on server
* `protocol` (String) - Protocol to access server. HTTP and HTTPS only applicable for retrieving file
Optional:

* `password` (String) - password to use for server login
* `username` (String) - user name to use for server login
<a id="nestedatt--app_intent_configs--global_default_configs"></a>
### Nested Schema for `app_intent_configs.global_default_configs`

Optional:

* `caching` (Attributes) (see [below for nested schema](#nestedatt--app_intent_configs--global_default_configs--caching))
* `cluster_id` (String) - id of the defining cluster
* `dhe_ciphersuit` (String)
* `monitor` (Attributes) (see [below for nested schema](#nestedatt--app_intent_configs--global_default_configs--monitor))
* `resumption` (Attributes) (see [below for nested schema](#nestedatt--app_intent_configs--global_default_configs--resumption))
* `ssl_versions` (Attributes) (see [below for nested schema](#nestedatt--app_intent_configs--global_default_configs--ssl_versions))
* `start_tls` (Attributes) (see [below for nested schema](#nestedatt--app_intent_configs--global_default_configs--start_tls))
<a id="nestedatt--app_intent_configs--global_default_configs--caching"></a>
### Nested Schema for `app_intent_configs.global_default_configs.caching`

Optional:

* `persistence` (Attributes) (see [below for nested schema](#nestedatt--app_intent_configs--global_default_configs--caching--persistence))
<a id="nestedatt--app_intent_configs--global_default_configs--caching--persistence"></a>
### Nested Schema for `app_intent_configs.global_default_configs.caching.persistence`

Optional:

* `enable` (Boolean)
<a id="nestedatt--app_intent_configs--global_default_configs--monitor"></a>
### Nested Schema for `app_intent_configs.global_default_configs.monitor`

Optional:

* `enable` (Boolean)
<a id="nestedatt--app_intent_configs--global_default_configs--resumption"></a>
### Nested Schema for `app_intent_configs.global_default_configs.resumption`

Optional:

* `client` (Attributes) (see [below for nested schema](#nestedatt--app_intent_configs--global_default_configs--resumption--client))
<a id="nestedatt--app_intent_configs--global_default_configs--resumption--client"></a>
### Nested Schema for `app_intent_configs.global_default_configs.resumption.client`

Optional:

* `enable` (Boolean)
<a id="nestedatt--app_intent_configs--global_default_configs--ssl_versions"></a>
### Nested Schema for `app_intent_configs.global_default_configs.ssl_versions`

Required:

* `max_version` (String) - maxVersion must be greater than minVersion
* `min_version` (String)
Optional:

* `connection_reset_action_for_max_version` (String) - Action to take to reset connection if TLS version is higher than configured
* `connection_reset_action_for_min_version` (String) - Action to take to reset connection if TLS version is lower than configured
<a id="nestedatt--app_intent_configs--global_default_configs--start_tls"></a>
### Nested Schema for `app_intent_configs.global_default_configs.start_tls`

Optional:

* `enable` (Boolean)
<a id="nestedatt--app_intent_configs--gs_group_param_configs"></a>
### Nested Schema for `app_intent_configs.gs_group_param_configs`

Optional:

* `hsm_group` (String) - Alias of the hsmGroup
* `session_logging` (Attributes) - GsGroup Session Logging Configuration (see [below for nested schema](#nestedatt--app_intent_configs--gs_group_param_configs--session_logging))
<a id="nestedatt--app_intent_configs--gs_group_param_configs--session_logging"></a>
### Nested Schema for `app_intent_configs.gs_group_param_configs.session_logging`

Optional:

* `interface` (String) - Associated IP Interface
* `log_level` (String) - Log Level
* `remote_syslog_ip` (String) - Remote Syslog IP
* `remote_syslog_port` (Number) - Remote Syslog Port Number
<a id="nestedatt--app_intent_configs--inline_ssl"></a>
### Nested Schema for `app_intent_configs.inline_ssl`

Optional:

* `standalone` (Boolean) - If enabled , behaves in a normal way. If disabled, can configure other GS apps along side Inline ssl with Inline ssl occupying only 50% of the memory
<a id="nestedatt--app_intent_configs--key_store_configs"></a>
### Nested Schema for `app_intent_configs.key_store_configs`

Optional:

* `deployment_type` (String)
* `inbound_keys` (Attributes List) (see [below for nested schema](#nestedatt--app_intent_configs--key_store_configs--inbound_keys))
* `outboundkeys` (Attributes List) (see [below for nested schema](#nestedatt--app_intent_configs--key_store_configs--outboundkeys))
<a id="nestedatt--app_intent_configs--key_store_configs--inbound_keys"></a>
### Nested Schema for `app_intent_configs.key_store_configs.inbound_keys`

Optional:

* `key_alias` (String) - Key to decrypt the traffic for the associated server/domain address
* `server_domain_alias` (String)
<a id="nestedatt--app_intent_configs--key_store_configs--outboundkeys"></a>
### Nested Schema for `app_intent_configs.key_store_configs.outboundkeys`

Optional:

* `key_alias` (String)
* `signing_for` (String)
<a id="nestedatt--app_intent_configs--m_tls_configs"></a>
### Nested Schema for `app_intent_configs.m_tls_configs`

Optional:

* `primary_signing` (String) - Alias of the key from keyStore used for primary signing for mTLS
* `secondary_signing` (String) - Alias of the key from keyStore used for secondary signing for mTLS
* `trust_store` (String) - Alias of the client trust store
<a id="nestedatt--app_intent_configs--network_access_configs"></a>
### Nested Schema for `app_intent_configs.network_access_configs`

Optional:

* `network_access` (Attributes) (see [below for nested schema](#nestedatt--app_intent_configs--network_access_configs--network_access))
* `operation` (String) - Adds or deletes(None) engine interface connectivity
<a id="nestedatt--app_intent_configs--network_access_configs--network_access"></a>
### Nested Schema for `app_intent_configs.network_access_configs.network_access`

Required:

* `eport` (String) - GigaSMART engine port
Optional:

* `cluster_id` (String) - id of the defining cluster
* `dhcp` (Boolean)
* `dns` (String) - Only valid and required when 'dhcp' is false
* `gateway` (String) - Only valid and required when 'dhcp' is false
* `hw_address` (String)
* `interface` (String)
* `ip_address` (String) - Only valid and required when 'dhcp' is false
* `ip_mask` (String) - Only valid and required when 'dhcp' is false
* `mtu` (Number) - Only valid when 'dhcp' is false
* `proxy_server_profile` (String)
* `status` (String)
* `vlan` (Number)
<a id="nestedatt--app_intent_configs--ssl_path"></a>
### Nested Schema for `app_intent_configs.ssl_path`

Optional:

* `alias` (String)
* `flex_inline_map` (Attributes) - When 'oobCopy' is configured at least one of 'aToB' or 'bToA' must be configured (see [below for nested schema](#nestedatt--app_intent_configs--ssl_path--flex_inline_map))
<a id="nestedatt--app_intent_configs--ssl_path--flex_inline_map"></a>
### Nested Schema for `app_intent_configs.ssl_path.flex_inline_map`

Optional:

* `a_to_b` (Attributes) (see [below for nested schema](#nestedatt--app_intent_configs--ssl_path--flex_inline_map--a_to_b))
* `b_to_a` (Attributes) (see [below for nested schema](#nestedatt--app_intent_configs--ssl_path--flex_inline_map--b_to_a))
* `oob_copy` (Attributes List) (see [below for nested schema](#nestedatt--app_intent_configs--ssl_path--flex_inline_map--oob_copy))
* `svt_mode` (Boolean)
* `svt_tag` (Number) - only applicable when svtMode is enabled
* `tag` (Attributes) (see [below for nested schema](#nestedatt--app_intent_configs--ssl_path--flex_inline_map--tag))
<a id="nestedatt--app_intent_configs--ssl_path--flex_inline_map--a_to_b"></a>
### Nested Schema for `app_intent_configs.ssl_path.flex_inline_map.a_to_b`

Required:

* `type` (String) - when set to 'tools', the list of the processing inline tools have to be provided. when set to 'ibPathway', ibPathway alias should be provided
Optional:

* `ib_pathway` (String) - ibPathway alias. Only applicable when type is 'ibPathway'
* `tools` (List of String) - ordered list of inline tools or vports. Only applicable when 'type' is 'tools'
<a id="nestedatt--app_intent_configs--ssl_path--flex_inline_map--b_to_a"></a>
### Nested Schema for `app_intent_configs.ssl_path.flex_inline_map.b_to_a`

Required:

* `type` (String) - when set to 'tools', the list of the processing inline tools have to be provided. when set to 'ibPathway', ibPathway alias should be provided
Optional:

* `ib_pathway` (String) - ibPathway alias. Only applicable when type is 'ibPathway'
* `tools` (List of String) - ordered list of inline tools or vports. Only applicable when 'type' is 'tools'
<a id="nestedatt--app_intent_configs--ssl_path--flex_inline_map--oob_copy"></a>
### Nested Schema for `app_intent_configs.ssl_path.flex_inline_map.oob_copy`

Required:

* `dst_ports` (List of String) - list of destination tool ports
* `src_ports` (List of String) - inline network or an item from a-to-b and b-to-a lists
Optional:

* `direction` (String)
* `tag` (Attributes) (see [below for nested schema](#nestedatt--app_intent_configs--ssl_path--flex_inline_map--oob_copy--tag))
<a id="nestedatt--app_intent_configs--ssl_path--flex_inline_map--oob_copy--tag"></a>
### Nested Schema for `app_intent_configs.ssl_path.flex_inline_map.oob_copy.tag`

Required:

* `type` (String)
<a id="nestedatt--app_intent_configs--ssl_path--flex_inline_map--tag"></a>
### Nested Schema for `app_intent_configs.ssl_path.flex_inline_map.tag`

Required:

* `type` (String)
Optional:

* `tag_protocol_id` (String) - When tool VLAN tag is added , this protocol Id will be added which egress out the traffic
* `vlan_id` (Number) - only applicable when type is 'vlan'
<a id="nestedatt--app_intent_configs--ssl_profile_config"></a>
### Nested Schema for `app_intent_configs.ssl_profile_config`

Required:

* `alias` (String) - Alias of the inline SSL profile
Optional:

* `certificate` (Attributes) (see [below for nested schema](#nestedatt--app_intent_configs--ssl_profile_config--certificate))
* `client_auth` (Attributes) (see [below for nested schema](#nestedatt--app_intent_configs--ssl_profile_config--client_auth))
* `cluster_id` (String) - id of the defining cluster
* `decrypt` (Attributes) (see [below for nested schema](#nestedatt--app_intent_configs--ssl_profile_config--decrypt))
* `default_action` (String) - Action to take if none of the profile rules match
* `high_avail` (Attributes) (see [below for nested schema](#nestedatt--app_intent_configs--ssl_profile_config--high_avail))
* `key_map` (Attributes List) (see [below for nested schema](#nestedatt--app_intent_configs--ssl_profile_config--key_map))
* `monitor` (String)
* `network_group` (Attributes) (see [below for nested schema](#nestedatt--app_intent_configs--ssl_profile_config--network_group))
* `no_decrypt` (Attributes) (see [below for nested schema](#nestedatt--app_intent_configs--ssl_profile_config--no_decrypt))
* `non_ssl_tcp` (Attributes) (see [below for nested schema](#nestedatt--app_intent_configs--ssl_profile_config--non_ssl_tcp))
* `rules` (List of Dynamic) - inline SSL profile rules
* `split_proxy` (Attributes) (see [below for nested schema](#nestedatt--app_intent_configs--ssl_profile_config--split_proxy))
* `start_tls` (Attributes) (see [below for nested schema](#nestedatt--app_intent_configs--ssl_profile_config--start_tls))
* `tcp` (Attributes) (see [below for nested schema](#nestedatt--app_intent_configs--ssl_profile_config--tcp))
* `tool` (Attributes) (see [below for nested schema](#nestedatt--app_intent_configs--ssl_profile_config--tool))
* `url_cache` (Attributes) (see [below for nested schema](#nestedatt--app_intent_configs--ssl_profile_config--url_cache))
<a id="nestedatt--app_intent_configs--ssl_profile_config--certificate"></a>
### Nested Schema for `app_intent_configs.ssl_profile_config.certificate`

Optional:

* `expired` (String) - SSL profile on expired certificate
* `invalid` (String) - SSL profile on invalid certificate
* `revocation` (Attributes) (see [below for nested schema](#nestedatt--app_intent_configs--ssl_profile_config--certificate--revocation))
* `self_signed` (String) - SSL profile on self-signed certificate
* `unknown_ca` (String) - SSL profile on unknown CA certificate
<a id="nestedatt--app_intent_configs--ssl_profile_config--certificate--revocation"></a>
### Nested Schema for `app_intent_configs.ssl_profile_config.certificate.revocation`

Optional:

* `crl` (Attributes) (see [below for nested schema](#nestedatt--app_intent_configs--ssl_profile_config--certificate--revocation--crl))
* `ocsp` (Attributes) (see [below for nested schema](#nestedatt--app_intent_configs--ssl_profile_config--certificate--revocation--ocsp))
<a id="nestedatt--app_intent_configs--ssl_profile_config--certificate--revocation--crl"></a>
### Nested Schema for `app_intent_configs.ssl_profile_config.certificate.revocation.crl`

Optional:

* `defer` (Number) - timeout in seconds
* `enabled` (Boolean)
* `fail` (String) - only applicable when enabled
<a id="nestedatt--app_intent_configs--ssl_profile_config--certificate--revocation--ocsp"></a>
### Nested Schema for `app_intent_configs.ssl_profile_config.certificate.revocation.ocsp`

Optional:

* `defer` (Number) - timeout in seconds
* `enabled` (Boolean)
* `fail` (String) - only applicable when enabled
<a id="nestedatt--app_intent_configs--ssl_profile_config--client_auth"></a>
### Nested Schema for `app_intent_configs.ssl_profile_config.client_auth`

Optional:

* `expired` (String) - SSL profile on expired certificate
* `invalid` (String) - SSL profile on invalid certificate
* `revocation` (Attributes) (see [below for nested schema](#nestedatt--app_intent_configs--ssl_profile_config--client_auth--revocation))
* `self_signed` (String) - SSL profile on self-signed certificate
* `unknown_ca` (String) - SSL profile on unknown CA certificate
<a id="nestedatt--app_intent_configs--ssl_profile_config--client_auth--revocation"></a>
### Nested Schema for `app_intent_configs.ssl_profile_config.client_auth.revocation`

Optional:

* `crl` (Attributes) (see [below for nested schema](#nestedatt--app_intent_configs--ssl_profile_config--client_auth--revocation--crl))
* `ocsp` (Attributes) (see [below for nested schema](#nestedatt--app_intent_configs--ssl_profile_config--client_auth--revocation--ocsp))
<a id="nestedatt--app_intent_configs--ssl_profile_config--client_auth--revocation--crl"></a>
### Nested Schema for `app_intent_configs.ssl_profile_config.client_auth.revocation.crl`

Optional:

* `defer` (Number) - timeout in seconds
* `enabled` (Boolean)
* `fail` (String) - only applicable when enabled
<a id="nestedatt--app_intent_configs--ssl_profile_config--client_auth--revocation--ocsp"></a>
### Nested Schema for `app_intent_configs.ssl_profile_config.client_auth.revocation.ocsp`

Optional:

* `defer` (Number) - timeout in seconds
* `enabled` (Boolean)
* `fail` (String) - only applicable when enabled
<a id="nestedatt--app_intent_configs--ssl_profile_config--decrypt"></a>
### Nested Schema for `app_intent_configs.ssl_profile_config.decrypt`

Optional:

* `tcp` (Attributes) (see [below for nested schema](#nestedatt--app_intent_configs--ssl_profile_config--decrypt--tcp))
* `tool_bypass` (Attributes) (see [below for nested schema](#nestedatt--app_intent_configs--ssl_profile_config--decrypt--tool_bypass))
<a id="nestedatt--app_intent_configs--ssl_profile_config--decrypt--tcp"></a>
### Nested Schema for `app_intent_configs.ssl_profile_config.decrypt.tcp`

Optional:

* `inactive_timeout` (Number) - SSL decryption TCP inactive timeout (in minutes)
* `port_map` (Attributes) (see [below for nested schema](#nestedatt--app_intent_configs--ssl_profile_config--decrypt--tcp--port_map))
<a id="nestedatt--app_intent_configs--ssl_profile_config--decrypt--tcp--port_map"></a>
### Nested Schema for `app_intent_configs.ssl_profile_config.decrypt.tcp.port_map`

Optional:

* `default_out_port` (Number) - egress port for decryption port map. 0 is disabled.
* `ports` (Attributes List) (see [below for nested schema](#nestedatt--app_intent_configs--ssl_profile_config--decrypt--tcp--port_map--ports))
<a id="nestedatt--app_intent_configs--ssl_profile_config--decrypt--tcp--port_map--ports"></a>
### Nested Schema for `app_intent_configs.ssl_profile_config.decrypt.tcp.port_map.ports`

Required:

* `in_port` (Number) - ingress port for decryption port map
* `out_port` (Number) - egress port for decryption port map
Optional:

* `rule_id` (Number)
<a id="nestedatt--app_intent_configs--ssl_profile_config--decrypt--tool_bypass"></a>
### Nested Schema for `app_intent_configs.ssl_profile_config.decrypt.tool_bypass`

Optional:

* `enable` (Boolean)
<a id="nestedatt--app_intent_configs--ssl_profile_config--high_avail"></a>
### Nested Schema for `app_intent_configs.ssl_profile_config.high_avail`

Optional:

* `active_standby` (Attributes) (see [below for nested schema](#nestedatt--app_intent_configs--ssl_profile_config--high_avail--active_standby))
<a id="nestedatt--app_intent_configs--ssl_profile_config--high_avail--active_standby"></a>
### Nested Schema for `app_intent_configs.ssl_profile_config.high_avail.active_standby`

Optional:

* `enable` (Boolean)
<a id="nestedatt--app_intent_configs--ssl_profile_config--key_map"></a>
### Nested Schema for `app_intent_configs.ssl_profile_config.key_map`

Required:

* `hostname` (String) - hostname or IP address
* `key` (String) - SSL key alias
Optional:

* `rule_id` (Number)
<a id="nestedatt--app_intent_configs--ssl_profile_config--network_group"></a>
### Nested Schema for `app_intent_configs.ssl_profile_config.network_group`

Optional:

* `multiple_entry` (Attributes) (see [below for nested schema](#nestedatt--app_intent_configs--ssl_profile_config--network_group--multiple_entry))
<a id="nestedatt--app_intent_configs--ssl_profile_config--network_group--multiple_entry"></a>
### Nested Schema for `app_intent_configs.ssl_profile_config.network_group.multiple_entry`

Optional:

* `enable` (Boolean)
<a id="nestedatt--app_intent_configs--ssl_profile_config--no_decrypt"></a>
### Nested Schema for `app_intent_configs.ssl_profile_config.no_decrypt`

Optional:

* `tool_bypass` (Attributes) (see [below for nested schema](#nestedatt--app_intent_configs--ssl_profile_config--no_decrypt--tool_bypass))
<a id="nestedatt--app_intent_configs--ssl_profile_config--no_decrypt--tool_bypass"></a>
### Nested Schema for `app_intent_configs.ssl_profile_config.no_decrypt.tool_bypass`

Optional:

* `enable` (Boolean)
<a id="nestedatt--app_intent_configs--ssl_profile_config--non_ssl_tcp"></a>
### Nested Schema for `app_intent_configs.ssl_profile_config.non_ssl_tcp`

Optional:

* `tool_bypass` (Attributes) (see [below for nested schema](#nestedatt--app_intent_configs--ssl_profile_config--non_ssl_tcp--tool_bypass))
<a id="nestedatt--app_intent_configs--ssl_profile_config--non_ssl_tcp--tool_bypass"></a>
### Nested Schema for `app_intent_configs.ssl_profile_config.non_ssl_tcp.tool_bypass`

Optional:

* `enable` (Boolean)
<a id="nestedatt--app_intent_configs--ssl_profile_config--split_proxy"></a>
### Nested Schema for `app_intent_configs.ssl_profile_config.split_proxy`

Optional:

* `mode` (Attributes) (see [below for nested schema](#nestedatt--app_intent_configs--ssl_profile_config--split_proxy--mode))
* `server_non_pfs_ciphers` (Attributes) (see [below for nested schema](#nestedatt--app_intent_configs--ssl_profile_config--split_proxy--server_non_pfs_ciphers))
<a id="nestedatt--app_intent_configs--ssl_profile_config--split_proxy--mode"></a>
### Nested Schema for `app_intent_configs.ssl_profile_config.split_proxy.mode`

Optional:

* `enable` (Boolean)
<a id="nestedatt--app_intent_configs--ssl_profile_config--split_proxy--server_non_pfs_ciphers"></a>
### Nested Schema for `app_intent_configs.ssl_profile_config.split_proxy.server_non_pfs_ciphers`

Optional:

* `enable` (Boolean)
<a id="nestedatt--app_intent_configs--ssl_profile_config--start_tls"></a>
### Nested Schema for `app_intent_configs.ssl_profile_config.start_tls`

Optional:

* `l4_port` (List of Number)
<a id="nestedatt--app_intent_configs--ssl_profile_config--tcp"></a>
### Nested Schema for `app_intent_configs.ssl_profile_config.tcp`

Optional:

* `delayed_ack` (Boolean) - enable/disable TCP delayed ACK
* `syn_retries` (Number) - TCP Sync retries
* `timewait_timeout` (Number) - TCP Wait Timeout value
<a id="nestedatt--app_intent_configs--ssl_profile_config--tool"></a>
### Nested Schema for `app_intent_configs.ssl_profile_config.tool`

Optional:

* `early_engage` (Boolean) - enable/disable tool early engage
* `fail_action` (String) - Action to take if the tool fails
<a id="nestedatt--app_intent_configs--ssl_profile_config--url_cache"></a>
### Nested Schema for `app_intent_configs.ssl_profile_config.url_cache`

Optional:

* `miss_action` (String) - The action to take if local URL category resolution misses
* `timeout` (Number) - defer timeout in seconds. Only applicable for missAction 'defer'
<a id="nestedatt--app_intent_configs--trust_store_configs"></a>
### Nested Schema for `app_intent_configs.trust_store_configs`

Optional:

* `trust_store_append_configs` (Attributes) (see [below for nested schema](#nestedatt--app_intent_configs--trust_store_configs--trust_store_append_configs))
* `trust_store_replace_configs` (Attributes) (see [below for nested schema](#nestedatt--app_intent_configs--trust_store_configs--trust_store_replace_configs))
<a id="nestedatt--app_intent_configs--trust_store_configs--trust_store_append_configs"></a>
### Nested Schema for `app_intent_configs.trust_store_configs.trust_store_append_configs`

Optional:

* `file` (String) - The contents of the file. Mutually exclusive with 'fileSource'
* `file_source` (Attributes) - Remote file source or destination (see [below for nested schema](#nestedatt--app_intent_configs--trust_store_configs--trust_store_append_configs--file_source))
<a id="nestedatt--app_intent_configs--trust_store_configs--trust_store_append_configs--file_source"></a>
### Nested Schema for `app_intent_configs.trust_store_configs.trust_store_append_configs.file_source`

Required:

* `hostname` (String) - server address
* `path` (String) - file path on server
* `protocol` (String) - Protocol to access server. HTTP and HTTPS only applicable for retrieving file
Optional:

* `password` (String) - password to use for server login
* `username` (String) - user name to use for server login
<a id="nestedatt--app_intent_configs--trust_store_configs--trust_store_replace_configs"></a>
### Nested Schema for `app_intent_configs.trust_store_configs.trust_store_replace_configs`

Optional:

* `file` (String) - The contents of the file. Mutually exclusive with 'fileSource'
* `file_source` (Attributes) - Remote file source or destination (see [below for nested schema](#nestedatt--app_intent_configs--trust_store_configs--trust_store_replace_configs--file_source))
<a id="nestedatt--app_intent_configs--trust_store_configs--trust_store_replace_configs--file_source"></a>
### Nested Schema for `app_intent_configs.trust_store_configs.trust_store_replace_configs.file_source`

Required:

* `hostname` (String) - server address
* `path` (String) - file path on server
* `protocol` (String) - Protocol to access server. HTTP and HTTPS only applicable for retrieving file
Optional:

* `password` (String) - password to use for server login
* `username` (String) - user name to use for server login
<a id="nestedatt--app_intent_configs--vport_config"></a>
### Nested Schema for `app_intent_configs.vport_config`

Required:

* `alias` (String)
* `gs_group` (String) - Alias of referenced managing GsGroup
Optional:

* `deferred_binding` (Boolean) - enable/disable deferred-binding
* `fail_over_action` (String)
* `health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--app_intent_configs--vport_config--health_state_reasons))
* `inline_status` (String)
* `inner_traffic_path` (String) - Similar to inline-network traffic-path, applicable for inner map
* `metadata_monitoring` (Attributes) (see [below for nested schema](#nestedatt--app_intent_configs--vport_config--metadata_monitoring))
* `mode` (String)
* `outer_traffic_path` (String) - Similar to inline-tool/inline-tool-group flex-traffic path, applicable for outer map
* `sa_apf_profile` (String) - ASF session profile
<a id="nestedatt--app_intent_configs--vport_config--health_state_reasons"></a>
### Nested Schema for `app_intent_configs.vport_config.health_state_reasons`

Optional:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type
<a id="nestedatt--app_intent_configs--vport_config--metadata_monitoring"></a>
### Nested Schema for `app_intent_configs.vport_config.metadata_monitoring`

Optional:

* `action` (String) - metadata monitoring action
* `exporters` (List of String)
<a id="nestedatt--app_intent_configs--white_list_config"></a>
### Nested Schema for `app_intent_configs.white_list_config`

Optional:

* `operation` (String)
* `profile_list` (Attributes) (see [below for nested schema](#nestedatt--app_intent_configs--white_list_config--profile_list))
<a id="nestedatt--app_intent_configs--white_list_config--profile_list"></a>
### Nested Schema for `app_intent_configs.white_list_config.profile_list`

Optional:

* `file_source` (Attributes) - Remote file source or destination (see [below for nested schema](#nestedatt--app_intent_configs--white_list_config--profile_list--file_source))
* `list` (String) - The nodecryptlist or decryptlist . Mutually exclusive with 'fileSource'
<a id="nestedatt--app_intent_configs--white_list_config--profile_list--file_source"></a>
### Nested Schema for `app_intent_configs.white_list_config.profile_list.file_source`

Required:

* `hostname` (String) - server address
* `path` (String) - file path on server
* `protocol` (String) - Protocol to access server. HTTP and HTTPS only applicable for retrieving file
Optional:

* `password` (String) - password to use for server login
* `username` (String) - user name to use for server login
<a id="nestedatt--health_state_reasons"></a>
### Nested Schema for `health_state_reasons`

Optional:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type
<a id="nestedatt--ria_configs"></a>
### Nested Schema for `ria_configs`

Optional:

* `cluster_name` (String)
* `gs_engines` (List of String)
* `gs_group_alias` (String) - Alias of the gs group auto created by ssl app
* `gsop_alias` (String) - Alias of the gsop auto created by ssl app
* `network_access_configs` (Attributes List) (see [below for nested schema](#nestedatt--ria_configs--network_access_configs))
* `ssl_profile_alias` (String) - Alias of the ssl profile auto created by ssl app
* `vport_alias` (String) - Alias of the vport auto created by ssl app
<a id="nestedatt--ria_configs--network_access_configs"></a>
### Nested Schema for `ria_configs.network_access_configs`

Optional:

* `network_access` (Attributes) (see [below for nested schema](#nestedatt--ria_configs--network_access_configs--network_access))
* `operation` (String) - Adds or deletes(None) engine interface connectivity
<a id="nestedatt--ria_configs--network_access_configs--network_access"></a>
### Nested Schema for `ria_configs.network_access_configs.network_access`

Required:

* `eport` (String) - GigaSMART engine port
Optional:

* `cluster_id` (String) - id of the defining cluster
* `dhcp` (Boolean)
* `dns` (String) - Only valid and required when 'dhcp' is false
* `gateway` (String) - Only valid and required when 'dhcp' is false
* `hw_address` (String)
* `interface` (String)
* `ip_address` (String) - Only valid and required when 'dhcp' is false
* `ip_mask` (String) - Only valid and required when 'dhcp' is false
* `mtu` (Number) - Only valid when 'dhcp' is false
* `proxy_server_profile` (String)
* `status` (String)
* `vlan` (Number)

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_inline_ssl_app.example {alias}
```
