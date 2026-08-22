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
  alias = null
  app_intent_configs = {}
  cluster_name = null
  config_status = null
  config_status_reasons = []
  health_state = null
  health_state_reasons = []
  m_tls = null
  ria_configs = []
  ria_enabled = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, optional) - Alias of the inline ssl app
* `app_intent_configs` (Object({black_list_config, global_default_configs, gs_engines, gs_group_alias, gs_group_param_configs, gsop_alias, inline_ssl, key_store_configs, m_tls_configs, network_access_configs, ssl_path, ssl_profile_alias, ssl_profile_config, tag_protocol_id, trust_store_configs, vlan_id, vport_alias, vport_config, white_list_config}), optional)
  * `black_list_config` (Object({operation, profile_list}), optional)
    * `operation` (String, optional)
    * `profile_list` (Object({file_source, list}), optional)
      * `file_source` (Object({hostname, password, path, protocol, username}), optional) - Remote file source or destination
        * `hostname` (String, required) - server address
        * `password` (String, optional) - password to use for server login
        * `path` (String, required) - file path on server
        * `protocol` (String, required) - Protocol to access server. HTTP and HTTPS only applicable for retrieving file
        * `username` (String, optional) - user name to use for server login
      * `list` (String, optional) - The nodecryptlist or decryptlist . Mutually exclusive with 'fileSource'
  * `global_default_configs` (Object({caching, cluster_id, dhe_ciphersuit, monitor, resumption, ssl_versions, start_tls}), optional) - inline SSL configuration parameters
    * `caching` (Object({persistence}), optional)
      * `persistence` (Object({enable}), optional)
        * `enable` (Bool, optional)
    * `cluster_id` (String, optional) - id of the defining cluster
    * `dhe_ciphersuit` (String, optional)
    * `monitor` (Object({enable}), optional)
      * `enable` (Bool, optional)
    * `resumption` (Object({client}), optional)
      * `client` (Object({enable}), optional)
        * `enable` (Bool, optional)
    * `ssl_versions` (Object({connection_reset_action_for_max_version, connection_reset_action_for_min_version, max_version, min_version}), optional)
      * `connection_reset_action_for_max_version` (String, optional) - Action to take to reset connection if TLS version is higher than configured
      * `connection_reset_action_for_min_version` (String, optional) - Action to take to reset connection if TLS version is lower than configured
      * `max_version` (String, required) - maxVersion must be greater than minVersion
      * `min_version` (String, required)
    * `start_tls` (Object({enable}), optional)
      * `enable` (Bool, optional)
  * `gs_engines` (List(String), optional)
  * `gs_group_alias` (String, optional) - Alias of the gs group auto created by ssl app
  * `gs_group_param_configs` (Object({hsm_group, session_logging}), optional)
    * `hsm_group` (String, optional) - Alias of the hsmGroup
    * `session_logging` (Object({interface, log_level, remote_syslog_ip, remote_syslog_port}), optional) - GsGroup Session Logging Configuration
      * `interface` (String, optional) - Associated IP Interface
      * `log_level` (String, optional) - Log Level
      * `remote_syslog_ip` (String, optional) - Remote Syslog IP
      * `remote_syslog_port` (Number, optional) - Remote Syslog Port Number
  * `gsop_alias` (String, optional) - Alias of the gsop auto created by ssl app
  * `inline_ssl` (Object({standalone}), optional) - Used to configure other GS apps in addition to Inline SSL on a HC1 box
    * `standalone` (Bool, optional) - If enabled , behaves in a normal way. If disabled, can configure other GS apps along side Inline ssl with Inline ssl occupying only 50% of the memory
  * `key_store_configs` (Object({deployment_type, inbound_keys, outboundkeys}), optional)
    * `deployment_type` (String, optional)
    * `inbound_keys` (List(Object({key_alias, server_domain_alias})), optional)
    * `outboundkeys` (List(Object({key_alias, signing_for})), optional)
  * `m_tls_configs` (Object({primary_signing, secondary_signing, trust_store}), optional)
    * `primary_signing` (String, optional) - Alias of the key from keyStore used for primary signing for mTLS
    * `secondary_signing` (String, optional) - Alias of the key from keyStore used for secondary signing for mTLS
    * `trust_store` (String, optional) - Alias of the client trust store
  * `network_access_configs` (List(Object({network_access, operation})), optional)
  * `ssl_path` (List(Object({alias, flex_inline_map})), optional)
  * `ssl_profile_alias` (String, optional) - Alias of the ssl profile auto created by ssl app
  * `ssl_profile_config` (Object({alias, certificate, client_auth, cluster_id, decrypt, default_action, high_avail, key_map, monitor, network_group, no_decrypt, non_ssl_tcp, rules, split_proxy, start_tls, tcp, tool, url_cache}), optional) - Flexible Inline SSL Profile
    * `alias` (String, required) - Alias of the inline SSL profile
    * `certificate` (Object({expired, invalid, revocation, self_signed, unknown_ca}), optional)
      * `expired` (String, optional) - SSL profile on expired certificate
      * `invalid` (String, optional) - SSL profile on invalid certificate
      * `revocation` (Object({crl, ocsp}), optional)
        * `crl` (Object({defer, enabled, fail}), optional)
          * `defer` (Number, optional) - timeout in seconds
          * `enabled` (Bool, optional)
          * `fail` (String, optional) - only applicable when enabled
        * `ocsp` (Object({defer, enabled, fail}), optional)
          * `defer` (Number, optional) - timeout in seconds
          * `enabled` (Bool, optional)
          * `fail` (String, optional) - only applicable when enabled
      * `self_signed` (String, optional) - SSL profile on self-signed certificate
      * `unknown_ca` (String, optional) - SSL profile on unknown CA certificate
    * `client_auth` (Object({expired, invalid, revocation, self_signed, unknown_ca}), optional)
      * `expired` (String, optional) - SSL profile on expired certificate
      * `invalid` (String, optional) - SSL profile on invalid certificate
      * `revocation` (Object({crl, ocsp}), optional)
        * `crl` (Object({defer, enabled, fail}), optional)
          * `defer` (Number, optional) - timeout in seconds
          * `enabled` (Bool, optional)
          * `fail` (String, optional) - only applicable when enabled
        * `ocsp` (Object({defer, enabled, fail}), optional)
          * `defer` (Number, optional) - timeout in seconds
          * `enabled` (Bool, optional)
          * `fail` (String, optional) - only applicable when enabled
      * `self_signed` (String, optional) - SSL profile on self-signed certificate
      * `unknown_ca` (String, optional) - SSL profile on unknown CA certificate
    * `cluster_id` (String, optional) - id of the defining cluster
    * `decrypt` (Object({tcp, tool_bypass}), optional)
      * `tcp` (Object({inactive_timeout, port_map}), optional)
        * `inactive_timeout` (Number, optional) - SSL decryption TCP inactive timeout (in minutes)
        * `port_map` (Object({default_out_port, ports}), optional)
          * `default_out_port` (Number, optional) - egress port for decryption port map. 0 is disabled.
          * `ports` (List(Object({in_port, out_port, rule_id})), optional)
      * `tool_bypass` (Object({enable}), optional)
        * `enable` (Bool, optional)
    * `default_action` (String, optional) - Action to take if none of the profile rules match
    * `high_avail` (Object({active_standby}), optional)
      * `active_standby` (Object({enable}), optional)
        * `enable` (Bool, optional)
    * `key_map` (List(Object({hostname, key, rule_id})), optional)
    * `monitor` (String, optional)
    * `network_group` (Object({multiple_entry}), optional)
      * `multiple_entry` (Object({enable}), optional)
        * `enable` (Bool, optional)
    * `no_decrypt` (Object({tool_bypass}), optional)
      * `tool_bypass` (Object({enable}), optional)
        * `enable` (Bool, optional)
    * `non_ssl_tcp` (Object({tool_bypass}), optional)
      * `tool_bypass` (Object({enable}), optional)
        * `enable` (Bool, optional)
    * `rules` (List(Dynamic), optional) - inline SSL profile rules
    * `split_proxy` (Object({mode, server_non_pfs_ciphers}), optional)
      * `mode` (Object({enable}), optional)
        * `enable` (Bool, optional)
      * `server_non_pfs_ciphers` (Object({enable}), optional)
        * `enable` (Bool, optional)
    * `start_tls` (Object({l4_port}), optional)
      * `l4_port` (List(Number), optional)
    * `tcp` (Object({delayed_ack, syn_retries, timewait_timeout}), optional)
      * `delayed_ack` (Bool, optional) - enable/disable TCP delayed ACK
      * `syn_retries` (Number, optional) - TCP Sync retries
      * `timewait_timeout` (Number, optional) - TCP Wait Timeout value
    * `tool` (Object({early_engage, fail_action}), optional)
      * `early_engage` (Bool, optional) - enable/disable tool early engage
      * `fail_action` (String, optional) - Action to take if the tool fails
    * `url_cache` (Object({miss_action, timeout}), optional)
      * `miss_action` (String, optional) - The action to take if local URL category resolution misses
      * `timeout` (Number, optional) - defer timeout in seconds. Only applicable for missAction 'defer'
  * `tag_protocol_id` (String, optional) - When tool VLAN tag is added , this protocol Id will be added which egress out the traffic
  * `trust_store_configs` (Object({trust_store_append_configs, trust_store_replace_configs}), optional)
    * `trust_store_append_configs` (Object({file, file_source}), optional)
      * `file` (String, optional) - The contents of the file. Mutually exclusive with 'fileSource'
      * `file_source` (Object({hostname, password, path, protocol, username}), optional) - Remote file source or destination
        * `hostname` (String, required) - server address
        * `password` (String, optional) - password to use for server login
        * `path` (String, required) - file path on server
        * `protocol` (String, required) - Protocol to access server. HTTP and HTTPS only applicable for retrieving file
        * `username` (String, optional) - user name to use for server login
    * `trust_store_replace_configs` (Object({file, file_source}), optional)
      * `file` (String, optional) - The contents of the file. Mutually exclusive with 'fileSource'
      * `file_source` (Object({hostname, password, path, protocol, username}), optional) - Remote file source or destination
        * `hostname` (String, required) - server address
        * `password` (String, optional) - password to use for server login
        * `path` (String, required) - file path on server
        * `protocol` (String, required) - Protocol to access server. HTTP and HTTPS only applicable for retrieving file
        * `username` (String, optional) - user name to use for server login
  * `vlan_id` (Number, optional) - Vlan id for proxy maps
  * `vport_alias` (String, optional) - Alias of the vport auto created by ssl app
  * `vport_config` (Object({alias, deferred_binding, fail_over_action, gs_group, health_state, health_state_reasons, inline_status, inner_traffic_path, metadata_monitoring, mode, outer_traffic_path, sa_apf_profile}), optional) - GigaSMART vPort
    * `alias` (String, required)
    * `deferred_binding` (Bool, optional) - enable/disable deferred-binding
    * `fail_over_action` (String, optional)
    * `gs_group` (String, required) - Alias of referenced managing GsGroup
    * `health_state` (String, optional) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
    * `health_state_reasons` (List(Object({message, severity, traffic_health_state_computation_type})), optional)
    * `inline_status` (String, optional)
    * `inner_traffic_path` (String, optional) - Similar to inline-network traffic-path, applicable for inner map
    * `metadata_monitoring` (Object({action, exporters}), optional)
      * `action` (String, optional) - metadata monitoring action
      * `exporters` (List(String), optional)
    * `mode` (String, optional)
    * `outer_traffic_path` (String, optional) - Similar to inline-tool/inline-tool-group flex-traffic path, applicable for outer map
    * `sa_apf_profile` (String, optional) - ASF session profile
  * `white_list_config` (Object({operation, profile_list}), optional)
    * `operation` (String, optional)
    * `profile_list` (Object({file_source, list}), optional)
      * `file_source` (Object({hostname, password, path, protocol, username}), optional) - Remote file source or destination
        * `hostname` (String, required) - server address
        * `password` (String, optional) - password to use for server login
        * `path` (String, required) - file path on server
        * `protocol` (String, required) - Protocol to access server. HTTP and HTTPS only applicable for retrieving file
        * `username` (String, optional) - user name to use for server login
      * `list` (String, optional) - The nodecryptlist or decryptlist . Mutually exclusive with 'fileSource'
* `cluster_name` (String, optional)
* `config_status` (String, optional) - Configuration status of this SSL app
* `config_status_reasons` (List(String), optional) - In case of configuration failure, this message provides details about the possible cause of the failure
* `health_state` (String, optional) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (List(Object({message, severity, traffic_health_state_computation_type})), optional)
* `m_tls` (String, optional) - Enable mTLS to do client Authentication
* `ria_configs` (List(Object({cluster_name, gs_engines, gs_group_alias, gsop_alias, network_access_configs, ssl_profile_alias, vport_alias})), optional) - Configs applicable if RIA is enabled in ssl app, the listed properties here will not be applicable in appIntent configs if RIA is enabled
* `ria_enabled` (String, optional) - Enable RIA to configure ssl app in two nodes which is needed for RIA SSL

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `alias` (String, computed) - Alias of the inline ssl app
* `app_intent_configs` (Object({black_list_config, global_default_configs, gs_engines, gs_group_alias, gs_group_param_configs, gsop_alias, inline_ssl, key_store_configs, m_tls_configs, network_access_configs, ssl_path, ssl_profile_alias, ssl_profile_config, tag_protocol_id, trust_store_configs, vlan_id, vport_alias, vport_config, white_list_config}), computed)
  * `black_list_config` (Object({operation, profile_list}), optional)
    * `operation` (String, optional)
    * `profile_list` (Object({file_source, list}), optional)
      * `file_source` (Object({hostname, password, path, protocol, username}), optional) - Remote file source or destination
        * `hostname` (String, required) - server address
        * `password` (String, optional) - password to use for server login
        * `path` (String, required) - file path on server
        * `protocol` (String, required) - Protocol to access server. HTTP and HTTPS only applicable for retrieving file
        * `username` (String, optional) - user name to use for server login
      * `list` (String, optional) - The nodecryptlist or decryptlist . Mutually exclusive with 'fileSource'
  * `global_default_configs` (Object({caching, cluster_id, dhe_ciphersuit, monitor, resumption, ssl_versions, start_tls}), optional) - inline SSL configuration parameters
    * `caching` (Object({persistence}), optional)
      * `persistence` (Object({enable}), optional)
        * `enable` (Bool, optional)
    * `cluster_id` (String, optional) - id of the defining cluster
    * `dhe_ciphersuit` (String, optional)
    * `monitor` (Object({enable}), optional)
      * `enable` (Bool, optional)
    * `resumption` (Object({client}), optional)
      * `client` (Object({enable}), optional)
        * `enable` (Bool, optional)
    * `ssl_versions` (Object({connection_reset_action_for_max_version, connection_reset_action_for_min_version, max_version, min_version}), optional)
      * `connection_reset_action_for_max_version` (String, optional) - Action to take to reset connection if TLS version is higher than configured
      * `connection_reset_action_for_min_version` (String, optional) - Action to take to reset connection if TLS version is lower than configured
      * `max_version` (String, required) - maxVersion must be greater than minVersion
      * `min_version` (String, required)
    * `start_tls` (Object({enable}), optional)
      * `enable` (Bool, optional)
  * `gs_engines` (List(String), optional)
  * `gs_group_alias` (String, optional) - Alias of the gs group auto created by ssl app
  * `gs_group_param_configs` (Object({hsm_group, session_logging}), optional)
    * `hsm_group` (String, optional) - Alias of the hsmGroup
    * `session_logging` (Object({interface, log_level, remote_syslog_ip, remote_syslog_port}), optional) - GsGroup Session Logging Configuration
      * `interface` (String, optional) - Associated IP Interface
      * `log_level` (String, optional) - Log Level
      * `remote_syslog_ip` (String, optional) - Remote Syslog IP
      * `remote_syslog_port` (Number, optional) - Remote Syslog Port Number
  * `gsop_alias` (String, optional) - Alias of the gsop auto created by ssl app
  * `inline_ssl` (Object({standalone}), optional) - Used to configure other GS apps in addition to Inline SSL on a HC1 box
    * `standalone` (Bool, optional) - If enabled , behaves in a normal way. If disabled, can configure other GS apps along side Inline ssl with Inline ssl occupying only 50% of the memory
  * `key_store_configs` (Object({deployment_type, inbound_keys, outboundkeys}), optional)
    * `deployment_type` (String, optional)
    * `inbound_keys` (List(Object({key_alias, server_domain_alias})), optional)
    * `outboundkeys` (List(Object({key_alias, signing_for})), optional)
  * `m_tls_configs` (Object({primary_signing, secondary_signing, trust_store}), optional)
    * `primary_signing` (String, optional) - Alias of the key from keyStore used for primary signing for mTLS
    * `secondary_signing` (String, optional) - Alias of the key from keyStore used for secondary signing for mTLS
    * `trust_store` (String, optional) - Alias of the client trust store
  * `network_access_configs` (List(Object({network_access, operation})), optional)
  * `ssl_path` (List(Object({alias, flex_inline_map})), optional)
  * `ssl_profile_alias` (String, optional) - Alias of the ssl profile auto created by ssl app
  * `ssl_profile_config` (Object({alias, certificate, client_auth, cluster_id, decrypt, default_action, high_avail, key_map, monitor, network_group, no_decrypt, non_ssl_tcp, rules, split_proxy, start_tls, tcp, tool, url_cache}), optional) - Flexible Inline SSL Profile
    * `alias` (String, required) - Alias of the inline SSL profile
    * `certificate` (Object({expired, invalid, revocation, self_signed, unknown_ca}), optional)
      * `expired` (String, optional) - SSL profile on expired certificate
      * `invalid` (String, optional) - SSL profile on invalid certificate
      * `revocation` (Object({crl, ocsp}), optional)
        * `crl` (Object({defer, enabled, fail}), optional)
          * `defer` (Number, optional) - timeout in seconds
          * `enabled` (Bool, optional)
          * `fail` (String, optional) - only applicable when enabled
        * `ocsp` (Object({defer, enabled, fail}), optional)
          * `defer` (Number, optional) - timeout in seconds
          * `enabled` (Bool, optional)
          * `fail` (String, optional) - only applicable when enabled
      * `self_signed` (String, optional) - SSL profile on self-signed certificate
      * `unknown_ca` (String, optional) - SSL profile on unknown CA certificate
    * `client_auth` (Object({expired, invalid, revocation, self_signed, unknown_ca}), optional)
      * `expired` (String, optional) - SSL profile on expired certificate
      * `invalid` (String, optional) - SSL profile on invalid certificate
      * `revocation` (Object({crl, ocsp}), optional)
        * `crl` (Object({defer, enabled, fail}), optional)
          * `defer` (Number, optional) - timeout in seconds
          * `enabled` (Bool, optional)
          * `fail` (String, optional) - only applicable when enabled
        * `ocsp` (Object({defer, enabled, fail}), optional)
          * `defer` (Number, optional) - timeout in seconds
          * `enabled` (Bool, optional)
          * `fail` (String, optional) - only applicable when enabled
      * `self_signed` (String, optional) - SSL profile on self-signed certificate
      * `unknown_ca` (String, optional) - SSL profile on unknown CA certificate
    * `cluster_id` (String, optional) - id of the defining cluster
    * `decrypt` (Object({tcp, tool_bypass}), optional)
      * `tcp` (Object({inactive_timeout, port_map}), optional)
        * `inactive_timeout` (Number, optional) - SSL decryption TCP inactive timeout (in minutes)
        * `port_map` (Object({default_out_port, ports}), optional)
          * `default_out_port` (Number, optional) - egress port for decryption port map. 0 is disabled.
          * `ports` (List(Object({in_port, out_port, rule_id})), optional)
      * `tool_bypass` (Object({enable}), optional)
        * `enable` (Bool, optional)
    * `default_action` (String, optional) - Action to take if none of the profile rules match
    * `high_avail` (Object({active_standby}), optional)
      * `active_standby` (Object({enable}), optional)
        * `enable` (Bool, optional)
    * `key_map` (List(Object({hostname, key, rule_id})), optional)
    * `monitor` (String, optional)
    * `network_group` (Object({multiple_entry}), optional)
      * `multiple_entry` (Object({enable}), optional)
        * `enable` (Bool, optional)
    * `no_decrypt` (Object({tool_bypass}), optional)
      * `tool_bypass` (Object({enable}), optional)
        * `enable` (Bool, optional)
    * `non_ssl_tcp` (Object({tool_bypass}), optional)
      * `tool_bypass` (Object({enable}), optional)
        * `enable` (Bool, optional)
    * `rules` (List(Dynamic), optional) - inline SSL profile rules
    * `split_proxy` (Object({mode, server_non_pfs_ciphers}), optional)
      * `mode` (Object({enable}), optional)
        * `enable` (Bool, optional)
      * `server_non_pfs_ciphers` (Object({enable}), optional)
        * `enable` (Bool, optional)
    * `start_tls` (Object({l4_port}), optional)
      * `l4_port` (List(Number), optional)
    * `tcp` (Object({delayed_ack, syn_retries, timewait_timeout}), optional)
      * `delayed_ack` (Bool, optional) - enable/disable TCP delayed ACK
      * `syn_retries` (Number, optional) - TCP Sync retries
      * `timewait_timeout` (Number, optional) - TCP Wait Timeout value
    * `tool` (Object({early_engage, fail_action}), optional)
      * `early_engage` (Bool, optional) - enable/disable tool early engage
      * `fail_action` (String, optional) - Action to take if the tool fails
    * `url_cache` (Object({miss_action, timeout}), optional)
      * `miss_action` (String, optional) - The action to take if local URL category resolution misses
      * `timeout` (Number, optional) - defer timeout in seconds. Only applicable for missAction 'defer'
  * `tag_protocol_id` (String, optional) - When tool VLAN tag is added , this protocol Id will be added which egress out the traffic
  * `trust_store_configs` (Object({trust_store_append_configs, trust_store_replace_configs}), optional)
    * `trust_store_append_configs` (Object({file, file_source}), optional)
      * `file` (String, optional) - The contents of the file. Mutually exclusive with 'fileSource'
      * `file_source` (Object({hostname, password, path, protocol, username}), optional) - Remote file source or destination
        * `hostname` (String, required) - server address
        * `password` (String, optional) - password to use for server login
        * `path` (String, required) - file path on server
        * `protocol` (String, required) - Protocol to access server. HTTP and HTTPS only applicable for retrieving file
        * `username` (String, optional) - user name to use for server login
    * `trust_store_replace_configs` (Object({file, file_source}), optional)
      * `file` (String, optional) - The contents of the file. Mutually exclusive with 'fileSource'
      * `file_source` (Object({hostname, password, path, protocol, username}), optional) - Remote file source or destination
        * `hostname` (String, required) - server address
        * `password` (String, optional) - password to use for server login
        * `path` (String, required) - file path on server
        * `protocol` (String, required) - Protocol to access server. HTTP and HTTPS only applicable for retrieving file
        * `username` (String, optional) - user name to use for server login
  * `vlan_id` (Number, optional) - Vlan id for proxy maps
  * `vport_alias` (String, optional) - Alias of the vport auto created by ssl app
  * `vport_config` (Object({alias, deferred_binding, fail_over_action, gs_group, health_state, health_state_reasons, inline_status, inner_traffic_path, metadata_monitoring, mode, outer_traffic_path, sa_apf_profile}), optional) - GigaSMART vPort
    * `alias` (String, required)
    * `deferred_binding` (Bool, optional) - enable/disable deferred-binding
    * `fail_over_action` (String, optional)
    * `gs_group` (String, required) - Alias of referenced managing GsGroup
    * `health_state` (String, optional) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
    * `health_state_reasons` (List(Object({message, severity, traffic_health_state_computation_type})), optional)
    * `inline_status` (String, optional)
    * `inner_traffic_path` (String, optional) - Similar to inline-network traffic-path, applicable for inner map
    * `metadata_monitoring` (Object({action, exporters}), optional)
      * `action` (String, optional) - metadata monitoring action
      * `exporters` (List(String), optional)
    * `mode` (String, optional)
    * `outer_traffic_path` (String, optional) - Similar to inline-tool/inline-tool-group flex-traffic path, applicable for outer map
    * `sa_apf_profile` (String, optional) - ASF session profile
  * `white_list_config` (Object({operation, profile_list}), optional)
    * `operation` (String, optional)
    * `profile_list` (Object({file_source, list}), optional)
      * `file_source` (Object({hostname, password, path, protocol, username}), optional) - Remote file source or destination
        * `hostname` (String, required) - server address
        * `password` (String, optional) - password to use for server login
        * `path` (String, required) - file path on server
        * `protocol` (String, required) - Protocol to access server. HTTP and HTTPS only applicable for retrieving file
        * `username` (String, optional) - user name to use for server login
      * `list` (String, optional) - The nodecryptlist or decryptlist . Mutually exclusive with 'fileSource'
* `cluster_name` (String, computed)
* `config_status` (String, computed) - Configuration status of this SSL app
* `config_status_reasons` (List(String), computed) - In case of configuration failure, this message provides details about the possible cause of the failure
* `health_state` (String, computed) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (List(Object({message, severity, traffic_health_state_computation_type})), computed)
* `m_tls` (String, computed) - Enable mTLS to do client Authentication
* `ria_configs` (List(Object({cluster_name, gs_engines, gs_group_alias, gsop_alias, network_access_configs, ssl_profile_alias, vport_alias})), computed) - Configs applicable if RIA is enabled in ssl app, the listed properties here will not be applicable in appIntent configs if RIA is enabled
* `ria_enabled` (String, computed) - Enable RIA to configure ssl app in two nodes which is needed for RIA SSL

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_inline_ssl_app.example {alias}
```
