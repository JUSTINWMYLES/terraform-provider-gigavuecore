---
page_title: "gigavuecore_fm_template Resource - gigavuecore"
subcategory: ""
description: |-
  Load FM Template by Template Name
---

# gigavuecore_fm_template Resource

Load FM Template by Template Name

## Example Usage

```terraform
resource "gigavuecore_fm_template" "example" {
  config = {}
  config_level = null
  config_level_value = []
  config_resource = null
  config_type = null
  modifiable = null
  ref_count = null
  ref_object = {}
  template_name = null
  update_time = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `config` (Object({aaa_auth_config, acme_certificate, device_ssl_certificate_configs, export_metadata_app_profile, giga_port_neighbors_discovery_config, giga_stream_threshold_config, giga_user_defined_application_config, ldap_server_system_config, ldap_servers, metadata_exporter, port_packet_threshold_config, proxy_server_profile, snmp_trap_event_configs, snmp_v3_users_config, ssh_ciphers_config}), optional)
  * `aaa_auth_config` (Object({auth_sequence, external_login_mapping}), optional) - FM Global AAA Authentication config to the device's
    * `auth_sequence` (Set(String), required) - Authentication methods order for user login. Valid values are 'local', 'ldap', 'radius', 'tacacs'
    * `external_login_mapping` (Object({default_local_user, user_map_order}), optional) - Specifies how externally authenticated logins (RADIUS, TACACS+, or LDAP) are mapped to local accounts. Private class
      * `default_local_user` (String, optional) - Specifies the account to which externally authenticated logins are mapped when map order is set to remote-first (if there is no matching local account) or local-only
      * `user_map_order` (String, optional) - Specifies how externally authenticated logins (RADIUS, TACACS+, or LDAP) are mapped to local accounts
  * `acme_certificate` (List(Object({acme_server_url, algorithm, operation_type, renew_days})), optional) - Available when ConfigType is ACME\_TEMPLATE
  * `device_ssl_certificate_configs` (List(Object({issuer, not_after, not_before, operation_type, signature_algorithm, subject, trusted_ca, upload_spec})), optional) - Available when ConfigType is SSL\_CERTIFICATE\_TEMPLATE
  * `export_metadata_app_profile` (Object({alias, application_id, applications, counter, datalink, description, flow, gtpu, interface, ip, ipv4, ipv6, outer_ipv4, outer_ipv6, timestamp, transport, type}), optional)
    * `alias` (String, required) - application profile alias
    * `application_id` (Bool, optional) - only valid with 'export' type
    * `applications` (List(Object({attributes, is_user_defined, name})), optional) - application and attributes.
    * `counter` (Object({bytes, bytes_long, inner_byte, inner_byte_long, packets, packets_long}), optional)
      * `bytes` (Bool, optional)
      * `bytes_long` (Bool, optional)
      * `inner_byte` (Bool, optional)
      * `inner_byte_long` (Bool, optional)
      * `packets` (Bool, optional)
      * `packets_long` (Bool, optional)
    * `datalink` (Object({mac_dst, mac_src, vlan}), optional)
      * `mac_dst` (Bool, optional)
      * `mac_src` (Bool, optional)
      * `vlan` (Bool, optional)
    * `description` (String, optional)
    * `flow` (Object({end_reason}), optional)
      * `end_reason` (Bool, optional)
    * `gtpu` (Object({qfi, teid}), optional)
      * `qfi` (Bool, optional)
      * `teid` (Bool, optional)
    * `interface` (Object({in_name_width, in_physical_width, out_physical_width}), optional)
      * `in_name_width` (Number, optional)
      * `in_physical_width` (Number, optional)
      * `out_physical_width` (Number, optional)
    * `ip` (Object({version}), optional)
      * `version` (Bool, optional)
    * `ipv4` (Object({destination, dscp, fragmentation, header_len, option_map, precedence, protocol, section, source, tos, total_length, ttl}), optional)
      * `destination` (Object({prefix_min_mask}), optional)
        * `prefix_min_mask` (String, optional)
      * `dscp` (Bool, optional)
      * `fragmentation` (Object({flags, offset}), optional)
        * `flags` (Bool, optional)
        * `offset` (Bool, optional)
      * `header_len` (Bool, optional)
      * `option_map` (Bool, optional)
      * `precedence` (Bool, optional)
      * `protocol` (Bool, optional)
      * `section` (Object({header_size, payload_size}), optional)
        * `header_size` (Number, optional)
        * `payload_size` (Number, optional)
      * `source` (Object({prefix_min_mask}), optional)
        * `prefix_min_mask` (String, optional) - ipv4 source prefix minimum-mask - netmask or mask length
      * `tos` (Bool, optional)
      * `total_length` (Bool, optional)
      * `ttl` (Bool, optional)
    * `ipv6` (Object({destination, dscp, extension_map, flow_label, fragmentation, hop_limit, length, next_header, precedence, section, source, traffic_class}), optional)
      * `destination` (Object({prefix_min_mask}), optional)
        * `prefix_min_mask` (String, optional)
      * `dscp` (Bool, optional)
      * `extension_map` (Bool, optional)
      * `flow_label` (Bool, optional)
      * `fragmentation` (Object({flags, offset}), optional)
        * `flags` (Bool, optional)
        * `offset` (Bool, optional)
      * `hop_limit` (Bool, optional)
      * `length` (Object({header, payload, total}), optional)
        * `header` (Bool, optional)
        * `payload` (Bool, optional)
        * `total` (Bool, optional)
      * `next_header` (Bool, optional)
      * `precedence` (Bool, optional)
      * `section` (Object({header_size, payload_size}), optional)
        * `header_size` (Number, optional)
        * `payload_size` (Number, optional)
      * `source` (Object({prefix_min_mask}), optional)
        * `prefix_min_mask` (String, optional)
      * `traffic_class` (Bool, optional)
    * `outer_ipv4` (Object({destination, source}), optional)
      * `destination` (Bool, optional)
      * `source` (Bool, optional)
    * `outer_ipv6` (Object({destination, source}), optional)
      * `destination` (Bool, optional)
      * `source` (Bool, optional)
    * `timestamp` (Object({flow_end_msec, flow_endsec, flow_start_msec, flow_startsec, sys_up_time_first, sys_up_time_last}), optional)
      * `flow_end_msec` (Bool, optional)
      * `flow_endsec` (Bool, optional)
      * `flow_start_msec` (Bool, optional)
      * `flow_startsec` (Bool, optional)
      * `sys_up_time_first` (Bool, optional)
      * `sys_up_time_last` (Bool, optional)
    * `transport` (Object({dst_port, icmp, src_port, tcp, udp}), optional)
      * `dst_port` (Bool, optional)
      * `icmp` (Object({ipv4_code, ipv4_type, ipv6_code, ipv6_type}), optional)
        * `ipv4_code` (Bool, optional)
        * `ipv4_type` (Bool, optional)
        * `ipv6_code` (Bool, optional)
        * `ipv6_type` (Bool, optional)
      * `src_port` (Bool, optional)
      * `tcp` (Object({ack_number, dst_port, flags, header_len, seq_number, src_port, urgent_ptr, window_size}), optional)
        * `ack_number` (Bool, optional)
        * `dst_port` (Bool, optional)
        * `flags` (Bool, optional)
        * `header_len` (Bool, optional)
        * `seq_number` (Bool, optional)
        * `src_port` (Bool, optional)
        * `urgent_ptr` (Bool, optional)
        * `window_size` (Bool, optional)
      * `udp` (Object({dst_port, msg_len, src_port}), optional)
        * `dst_port` (Bool, optional)
        * `msg_len` (Bool, optional)
        * `src_port` (Bool, optional)
    * `type` (String, optional)
  * `giga_port_neighbors_discovery_config` (Object({resource_configs}), optional) - To contain the port neighbour discovery template config
    * `resource_configs` (List(Object({cdp, gdp, lldp, port_type})), optional) - List of config for each port type
  * `giga_stream_threshold_config` (Object({giga_stream_type_thresholds}), optional) - Gigastream threshold configs
    * `giga_stream_type_thresholds` (List(Object({type, variance_threshold})), optional) - GigaStream threshold config for port types
  * `giga_user_defined_application_config` (Object({alias, app_id, priority, rules}), optional)
    * `alias` (String, optional) - pattern: ^\[a-zA-Z0-9&'+\_=\|\-\]+$
    * `app_id` (Number, optional) - minimum: 16777216 maximum: 16777343   Ranges from 2^24 to 2^24 + 127. It needs to be unique
    * `priority` (Number, optional) - minimum: 1 maximum: 120  Lower the priority higher the precedence
    * `rules` (Object({rule}), optional)
      * `rule` (List(Object({address, code, common_name, content, cts_cookie, cts_page_url, cts_referer, cts_server, cts_uri, cts_user_agent, dscp, mime_type, mindata, port, resolv_name, stc_location, stc_server_agent, stc_subject_alt_name, stream, typeval, user_agent})), optional)
  * `ldap_server_system_config` (Dynamic, optional) - Available for ConfigType LDAP\_SYSTEM\_CONFIG\_TEMPLATE
  * `ldap_servers` (List(Object({order, server_address})), optional) - List of LDAP Servers. Available for ConfigType LDAP\_SERVERS\_TEMPLATE ChildType ldapServers
  * `metadata_exporter` (Object({alias, application_profiles, cef, description, destination, max_pkt_size, mobility_sam, monitor, netflow, snmp, source, type}), optional)
    * `alias` (String, required)
    * `application_profiles` (List(String), optional) - application profile aliases to attach to the exporter
    * `cef` (Object({active_timeout, inactive_timeout}), optional)
      * `active_timeout` (Number, optional) - in seconds
      * `inactive_timeout` (Number, optional) - in seconds
    * `description` (String, optional)
    * `destination` (Object({dscp, ipv4_address, l4_port_dst, l4_port_src, l4_protocol, ttl}), optional)
      * `dscp` (Number, optional)
      * `ipv4_address` (String, optional) - ipv4 address
      * `l4_port_dst` (Number, optional)
      * `l4_port_src` (Number, optional)
      * `l4_protocol` (String, optional)
      * `ttl` (Number, optional)
    * `max_pkt_size` (Number, optional)
    * `mobility_sam` (Object({encoding, encoding_format, event_enable, trigger}), optional)
      * `encoding` (String, optional)
      * `encoding_format` (String, optional)
      * `event_enable` (Object({modify, update}), optional)
        * `modify` (Bool, optional)
        * `update` (Bool, optional)
      * `trigger` (String, optional)
    * `monitor` (Object({timeout}), optional)
      * `timeout` (Number, optional) - how often to export in seconds
    * `netflow` (Object({active_timeout, inactive_timeout, template_refresh, template_type, version}), optional)
      * `active_timeout` (Number, optional) - in seconds
      * `inactive_timeout` (Number, optional) - in seconds
      * `template_refresh` (Number, optional) - template refresh interval in seconds
      * `template_type` (String, optional)
      * `version` (String, optional)
    * `snmp` (Object({enabled}), optional)
      * `enabled` (Bool, optional) - snmp reverse lookup enable/disable
    * `source` (Object({ip_interface}), optional)
      * `ip_interface` (String, optional)
    * `type` (String, optional)
  * `port_packet_threshold_config` (Object({drop_threshold, error_threshold}), optional) - Contains port threshold configs
    * `drop_threshold` (Object({rx, tx}), optional) - Drop threshold config for Rx and Tx
      * `rx` (List(Object({count_, percent, port_type})), optional) - Threshold values for Rx packet drop
      * `tx` (List(Object({count_, percent, port_type})), optional) - Threshold values for Tx packet drop
    * `error_threshold` (Object({rx, tx}), optional) - Error threshold config for Rx and Tx
      * `rx` (List(Object({count_, percent, port_type})), optional) - Threshold values for Rx packet error
      * `tx` (List(Object({count_, percent, port_type})), optional) - Threshold values for Tx packet error
  * `proxy_server_profile` (Object({alias, auth_type, comment, password, periodic_ping, periodic_ping_failure_retry, periodic_ping_interval, periodic_ping_type, port, protocol, proxy_address, ssl_apps, username}), optional)
    * `alias` (String, required)
    * `auth_type` (String, required)
    * `comment` (String, optional)
    * `password` (String, optional)
    * `periodic_ping` (String, required)
    * `periodic_ping_failure_retry` (Number, optional)
    * `periodic_ping_interval` (Number, optional)
    * `periodic_ping_type` (String, optional)
    * `port` (Number, required)
    * `protocol` (String, required)
    * `proxy_address` (String, required)
    * `ssl_apps` (Object({cluster_name}), optional)
      * `cluster_name` (List(String), optional) - Cluster Name where the proxy deployed
    * `username` (String, optional)
  * `snmp_trap_event_configs` (List(Object({enabled, notify_event})), optional) - Available when ConfigType is SNMPTRAPS
  * `snmp_v3_users_config` (Object({snmp_v3_user}), optional) - Contains list of FM's snmpv3 users
    * `snmp_v3_user` (List(Object({auth_key, auth_protocol, min_sw_version, previous_username, priv_key, priv_protocol, username})), required)
  * `ssh_ciphers_config` (Object({classic, crypto, fips}), optional) - System SSH Cipher
    * `classic` (Object({client_ciphers, client_hostkey, client_kex, client_macs, server_ciphers, server_hostkey, server_kex, server_macs}), optional) - System SSH Cipher
      * `client_ciphers` (List(String), optional)
      * `client_hostkey` (List(String), optional)
      * `client_kex` (List(String), optional)
      * `client_macs` (List(String), optional)
      * `server_ciphers` (List(String), optional)
      * `server_hostkey` (List(String), optional)
      * `server_kex` (List(String), optional)
      * `server_macs` (List(String), optional)
    * `crypto` (Object({client_ciphers, client_hostkey, client_kex, client_macs, server_ciphers, server_hostkey, server_kex, server_macs}), optional) - System SSH Cipher
      * `client_ciphers` (List(String), optional)
      * `client_hostkey` (List(String), optional)
      * `client_kex` (List(String), optional)
      * `client_macs` (List(String), optional)
      * `server_ciphers` (List(String), optional)
      * `server_hostkey` (List(String), optional)
      * `server_kex` (List(String), optional)
      * `server_macs` (List(String), optional)
    * `fips` (Object({client_ciphers, client_hostkey, client_kex, client_macs, server_ciphers, server_hostkey, server_kex, server_macs}), optional) - System SSH Cipher
      * `client_ciphers` (List(String), optional)
      * `client_hostkey` (List(String), optional)
      * `client_kex` (List(String), optional)
      * `client_macs` (List(String), optional)
      * `server_ciphers` (List(String), optional)
      * `server_hostkey` (List(String), optional)
      * `server_kex` (List(String), optional)
      * `server_macs` (List(String), optional)
* `config_level` (String, optional) - Scope of the applied FM template
* `config_level_value` (List(String), optional)
* `config_resource` (Dynamic, optional) - For a particular config type, there may be multiple templates with different levels. For example, cluster level the user can have more than one template and in same way in tag level there may be more than one template for each tag combination. In such cases, the configResource property can be used to contain the list of clusters associated to the template or the tag key values combination of the template. The data structure changes based on the level hence it is kept as type Object.
* `config_type` (String, optional) - Configuration Type of the FM template
* `modifiable` (Bool, optional)
* `ref_count` (Number, optional)
* `ref_object` (Object({ref_object_type}), optional)
  * `ref_object_type` (String, optional) - Type of the reference object
* `template_name` (String, optional)
* `update_time` (String, optional)

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `config` (Object({aaa_auth_config, acme_certificate, device_ssl_certificate_configs, export_metadata_app_profile, giga_port_neighbors_discovery_config, giga_stream_threshold_config, giga_user_defined_application_config, ldap_server_system_config, ldap_servers, metadata_exporter, port_packet_threshold_config, proxy_server_profile, snmp_trap_event_configs, snmp_v3_users_config, ssh_ciphers_config}), computed)
  * `aaa_auth_config` (Object({auth_sequence, external_login_mapping}), optional) - FM Global AAA Authentication config to the device's
    * `auth_sequence` (Set(String), required) - Authentication methods order for user login. Valid values are 'local', 'ldap', 'radius', 'tacacs'
    * `external_login_mapping` (Object({default_local_user, user_map_order}), optional) - Specifies how externally authenticated logins (RADIUS, TACACS+, or LDAP) are mapped to local accounts. Private class
      * `default_local_user` (String, optional) - Specifies the account to which externally authenticated logins are mapped when map order is set to remote-first (if there is no matching local account) or local-only
      * `user_map_order` (String, optional) - Specifies how externally authenticated logins (RADIUS, TACACS+, or LDAP) are mapped to local accounts
  * `acme_certificate` (List(Object({acme_server_url, algorithm, operation_type, renew_days})), optional) - Available when ConfigType is ACME\_TEMPLATE
  * `device_ssl_certificate_configs` (List(Object({issuer, not_after, not_before, operation_type, signature_algorithm, subject, trusted_ca, upload_spec})), optional) - Available when ConfigType is SSL\_CERTIFICATE\_TEMPLATE
  * `export_metadata_app_profile` (Object({alias, application_id, applications, counter, datalink, description, flow, gtpu, interface, ip, ipv4, ipv6, outer_ipv4, outer_ipv6, timestamp, transport, type}), optional)
    * `alias` (String, required) - application profile alias
    * `application_id` (Bool, optional) - only valid with 'export' type
    * `applications` (List(Object({attributes, is_user_defined, name})), optional) - application and attributes.
    * `counter` (Object({bytes, bytes_long, inner_byte, inner_byte_long, packets, packets_long}), optional)
      * `bytes` (Bool, optional)
      * `bytes_long` (Bool, optional)
      * `inner_byte` (Bool, optional)
      * `inner_byte_long` (Bool, optional)
      * `packets` (Bool, optional)
      * `packets_long` (Bool, optional)
    * `datalink` (Object({mac_dst, mac_src, vlan}), optional)
      * `mac_dst` (Bool, optional)
      * `mac_src` (Bool, optional)
      * `vlan` (Bool, optional)
    * `description` (String, optional)
    * `flow` (Object({end_reason}), optional)
      * `end_reason` (Bool, optional)
    * `gtpu` (Object({qfi, teid}), optional)
      * `qfi` (Bool, optional)
      * `teid` (Bool, optional)
    * `interface` (Object({in_name_width, in_physical_width, out_physical_width}), optional)
      * `in_name_width` (Number, optional)
      * `in_physical_width` (Number, optional)
      * `out_physical_width` (Number, optional)
    * `ip` (Object({version}), optional)
      * `version` (Bool, optional)
    * `ipv4` (Object({destination, dscp, fragmentation, header_len, option_map, precedence, protocol, section, source, tos, total_length, ttl}), optional)
      * `destination` (Object({prefix_min_mask}), optional)
        * `prefix_min_mask` (String, optional)
      * `dscp` (Bool, optional)
      * `fragmentation` (Object({flags, offset}), optional)
        * `flags` (Bool, optional)
        * `offset` (Bool, optional)
      * `header_len` (Bool, optional)
      * `option_map` (Bool, optional)
      * `precedence` (Bool, optional)
      * `protocol` (Bool, optional)
      * `section` (Object({header_size, payload_size}), optional)
        * `header_size` (Number, optional)
        * `payload_size` (Number, optional)
      * `source` (Object({prefix_min_mask}), optional)
        * `prefix_min_mask` (String, optional) - ipv4 source prefix minimum-mask - netmask or mask length
      * `tos` (Bool, optional)
      * `total_length` (Bool, optional)
      * `ttl` (Bool, optional)
    * `ipv6` (Object({destination, dscp, extension_map, flow_label, fragmentation, hop_limit, length, next_header, precedence, section, source, traffic_class}), optional)
      * `destination` (Object({prefix_min_mask}), optional)
        * `prefix_min_mask` (String, optional)
      * `dscp` (Bool, optional)
      * `extension_map` (Bool, optional)
      * `flow_label` (Bool, optional)
      * `fragmentation` (Object({flags, offset}), optional)
        * `flags` (Bool, optional)
        * `offset` (Bool, optional)
      * `hop_limit` (Bool, optional)
      * `length` (Object({header, payload, total}), optional)
        * `header` (Bool, optional)
        * `payload` (Bool, optional)
        * `total` (Bool, optional)
      * `next_header` (Bool, optional)
      * `precedence` (Bool, optional)
      * `section` (Object({header_size, payload_size}), optional)
        * `header_size` (Number, optional)
        * `payload_size` (Number, optional)
      * `source` (Object({prefix_min_mask}), optional)
        * `prefix_min_mask` (String, optional)
      * `traffic_class` (Bool, optional)
    * `outer_ipv4` (Object({destination, source}), optional)
      * `destination` (Bool, optional)
      * `source` (Bool, optional)
    * `outer_ipv6` (Object({destination, source}), optional)
      * `destination` (Bool, optional)
      * `source` (Bool, optional)
    * `timestamp` (Object({flow_end_msec, flow_endsec, flow_start_msec, flow_startsec, sys_up_time_first, sys_up_time_last}), optional)
      * `flow_end_msec` (Bool, optional)
      * `flow_endsec` (Bool, optional)
      * `flow_start_msec` (Bool, optional)
      * `flow_startsec` (Bool, optional)
      * `sys_up_time_first` (Bool, optional)
      * `sys_up_time_last` (Bool, optional)
    * `transport` (Object({dst_port, icmp, src_port, tcp, udp}), optional)
      * `dst_port` (Bool, optional)
      * `icmp` (Object({ipv4_code, ipv4_type, ipv6_code, ipv6_type}), optional)
        * `ipv4_code` (Bool, optional)
        * `ipv4_type` (Bool, optional)
        * `ipv6_code` (Bool, optional)
        * `ipv6_type` (Bool, optional)
      * `src_port` (Bool, optional)
      * `tcp` (Object({ack_number, dst_port, flags, header_len, seq_number, src_port, urgent_ptr, window_size}), optional)
        * `ack_number` (Bool, optional)
        * `dst_port` (Bool, optional)
        * `flags` (Bool, optional)
        * `header_len` (Bool, optional)
        * `seq_number` (Bool, optional)
        * `src_port` (Bool, optional)
        * `urgent_ptr` (Bool, optional)
        * `window_size` (Bool, optional)
      * `udp` (Object({dst_port, msg_len, src_port}), optional)
        * `dst_port` (Bool, optional)
        * `msg_len` (Bool, optional)
        * `src_port` (Bool, optional)
    * `type` (String, optional)
  * `giga_port_neighbors_discovery_config` (Object({resource_configs}), optional) - To contain the port neighbour discovery template config
    * `resource_configs` (List(Object({cdp, gdp, lldp, port_type})), optional) - List of config for each port type
  * `giga_stream_threshold_config` (Object({giga_stream_type_thresholds}), optional) - Gigastream threshold configs
    * `giga_stream_type_thresholds` (List(Object({type, variance_threshold})), optional) - GigaStream threshold config for port types
  * `giga_user_defined_application_config` (Object({alias, app_id, priority, rules}), optional)
    * `alias` (String, optional) - pattern: ^\[a-zA-Z0-9&'+\_=\|\-\]+$
    * `app_id` (Number, optional) - minimum: 16777216 maximum: 16777343   Ranges from 2^24 to 2^24 + 127. It needs to be unique
    * `priority` (Number, optional) - minimum: 1 maximum: 120  Lower the priority higher the precedence
    * `rules` (Object({rule}), optional)
      * `rule` (List(Object({address, code, common_name, content, cts_cookie, cts_page_url, cts_referer, cts_server, cts_uri, cts_user_agent, dscp, mime_type, mindata, port, resolv_name, stc_location, stc_server_agent, stc_subject_alt_name, stream, typeval, user_agent})), optional)
  * `ldap_server_system_config` (Dynamic, optional) - Available for ConfigType LDAP\_SYSTEM\_CONFIG\_TEMPLATE
  * `ldap_servers` (List(Object({order, server_address})), optional) - List of LDAP Servers. Available for ConfigType LDAP\_SERVERS\_TEMPLATE ChildType ldapServers
  * `metadata_exporter` (Object({alias, application_profiles, cef, description, destination, max_pkt_size, mobility_sam, monitor, netflow, snmp, source, type}), optional)
    * `alias` (String, required)
    * `application_profiles` (List(String), optional) - application profile aliases to attach to the exporter
    * `cef` (Object({active_timeout, inactive_timeout}), optional)
      * `active_timeout` (Number, optional) - in seconds
      * `inactive_timeout` (Number, optional) - in seconds
    * `description` (String, optional)
    * `destination` (Object({dscp, ipv4_address, l4_port_dst, l4_port_src, l4_protocol, ttl}), optional)
      * `dscp` (Number, optional)
      * `ipv4_address` (String, optional) - ipv4 address
      * `l4_port_dst` (Number, optional)
      * `l4_port_src` (Number, optional)
      * `l4_protocol` (String, optional)
      * `ttl` (Number, optional)
    * `max_pkt_size` (Number, optional)
    * `mobility_sam` (Object({encoding, encoding_format, event_enable, trigger}), optional)
      * `encoding` (String, optional)
      * `encoding_format` (String, optional)
      * `event_enable` (Object({modify, update}), optional)
        * `modify` (Bool, optional)
        * `update` (Bool, optional)
      * `trigger` (String, optional)
    * `monitor` (Object({timeout}), optional)
      * `timeout` (Number, optional) - how often to export in seconds
    * `netflow` (Object({active_timeout, inactive_timeout, template_refresh, template_type, version}), optional)
      * `active_timeout` (Number, optional) - in seconds
      * `inactive_timeout` (Number, optional) - in seconds
      * `template_refresh` (Number, optional) - template refresh interval in seconds
      * `template_type` (String, optional)
      * `version` (String, optional)
    * `snmp` (Object({enabled}), optional)
      * `enabled` (Bool, optional) - snmp reverse lookup enable/disable
    * `source` (Object({ip_interface}), optional)
      * `ip_interface` (String, optional)
    * `type` (String, optional)
  * `port_packet_threshold_config` (Object({drop_threshold, error_threshold}), optional) - Contains port threshold configs
    * `drop_threshold` (Object({rx, tx}), optional) - Drop threshold config for Rx and Tx
      * `rx` (List(Object({count_, percent, port_type})), optional) - Threshold values for Rx packet drop
      * `tx` (List(Object({count_, percent, port_type})), optional) - Threshold values for Tx packet drop
    * `error_threshold` (Object({rx, tx}), optional) - Error threshold config for Rx and Tx
      * `rx` (List(Object({count_, percent, port_type})), optional) - Threshold values for Rx packet error
      * `tx` (List(Object({count_, percent, port_type})), optional) - Threshold values for Tx packet error
  * `proxy_server_profile` (Object({alias, auth_type, comment, password, periodic_ping, periodic_ping_failure_retry, periodic_ping_interval, periodic_ping_type, port, protocol, proxy_address, ssl_apps, username}), optional)
    * `alias` (String, required)
    * `auth_type` (String, required)
    * `comment` (String, optional)
    * `password` (String, optional)
    * `periodic_ping` (String, required)
    * `periodic_ping_failure_retry` (Number, optional)
    * `periodic_ping_interval` (Number, optional)
    * `periodic_ping_type` (String, optional)
    * `port` (Number, required)
    * `protocol` (String, required)
    * `proxy_address` (String, required)
    * `ssl_apps` (Object({cluster_name}), optional)
      * `cluster_name` (List(String), optional) - Cluster Name where the proxy deployed
    * `username` (String, optional)
  * `snmp_trap_event_configs` (List(Object({enabled, notify_event})), optional) - Available when ConfigType is SNMPTRAPS
  * `snmp_v3_users_config` (Object({snmp_v3_user}), optional) - Contains list of FM's snmpv3 users
    * `snmp_v3_user` (List(Object({auth_key, auth_protocol, min_sw_version, previous_username, priv_key, priv_protocol, username})), required)
  * `ssh_ciphers_config` (Object({classic, crypto, fips}), optional) - System SSH Cipher
    * `classic` (Object({client_ciphers, client_hostkey, client_kex, client_macs, server_ciphers, server_hostkey, server_kex, server_macs}), optional) - System SSH Cipher
      * `client_ciphers` (List(String), optional)
      * `client_hostkey` (List(String), optional)
      * `client_kex` (List(String), optional)
      * `client_macs` (List(String), optional)
      * `server_ciphers` (List(String), optional)
      * `server_hostkey` (List(String), optional)
      * `server_kex` (List(String), optional)
      * `server_macs` (List(String), optional)
    * `crypto` (Object({client_ciphers, client_hostkey, client_kex, client_macs, server_ciphers, server_hostkey, server_kex, server_macs}), optional) - System SSH Cipher
      * `client_ciphers` (List(String), optional)
      * `client_hostkey` (List(String), optional)
      * `client_kex` (List(String), optional)
      * `client_macs` (List(String), optional)
      * `server_ciphers` (List(String), optional)
      * `server_hostkey` (List(String), optional)
      * `server_kex` (List(String), optional)
      * `server_macs` (List(String), optional)
    * `fips` (Object({client_ciphers, client_hostkey, client_kex, client_macs, server_ciphers, server_hostkey, server_kex, server_macs}), optional) - System SSH Cipher
      * `client_ciphers` (List(String), optional)
      * `client_hostkey` (List(String), optional)
      * `client_kex` (List(String), optional)
      * `client_macs` (List(String), optional)
      * `server_ciphers` (List(String), optional)
      * `server_hostkey` (List(String), optional)
      * `server_kex` (List(String), optional)
      * `server_macs` (List(String), optional)
* `config_level` (String, computed) - Scope of the applied FM template
* `config_level_value` (List(String), computed)
* `config_resource` (Dynamic, computed) - For a particular config type, there may be multiple templates with different levels. For example, cluster level the user can have more than one template and in same way in tag level there may be more than one template for each tag combination. In such cases, the configResource property can be used to contain the list of clusters associated to the template or the tag key values combination of the template. The data structure changes based on the level hence it is kept as type Object.
* `config_type` (String, computed) - Configuration Type of the FM template
* `id` (String, computed)
* `modifiable` (Bool, computed)
* `ref_count` (Number, computed)
* `ref_object` (Object({ref_object_type}), computed)
  * `ref_object_type` (String, optional) - Type of the reference object
* `template_name` (String, computed)
* `update_time` (String, computed)

