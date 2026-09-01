---
page_title: "gigavuecore_load_all_fm_templates_hierarchial_config Data Source - gigavuecore"
subcategory: ""
description: |-
  new in FM 5.7
---

# gigavuecore_load_all_fm_templates_hierarchial_config Data Source

new in FM 5.7

## Example Usage

```terraform
data "gigavuecore_load_all_fm_templates_hierarchial_config" "example" {
  config_type = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `config_type` (String, optional) - configType of the fm template

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) - All FM Template  Configurations (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `config` (Attributes) (see [below for nested schema](#nestedatt--items--config))
* `config_level` (String) - Scope of the applied FM template
* `config_level_value` (List of String)
* `config_resource` (Dynamic) - For a particular config type, there may be multiple templates with different levels. For example, cluster level the user can have more than one template and in same way in tag level there may be more than one template for each tag combination. In such cases, the configResource property can be used to contain the list of clusters associated to the template or the tag key values combination of the template. The data structure changes based on the level hence it is kept as type Object.
* `config_type` (String) - Configuration Type of the FM template
* `modifiable` (Boolean)
* `ref_count` (Number)
* `ref_object` (Attributes) (see [below for nested schema](#nestedatt--items--ref_object))
* `template_name` (String)
* `update_time` (String)

<a id="nestedatt--items--config"></a>
### Nested Schema for `items.config`

Read-Only:

* `aaa_auth_config` (Attributes) - FM Global AAA Authentication config to the device's (see [below for nested schema](#nestedatt--items--config--aaa_auth_config))
* `acme_certificate` (Attributes List) - Available when ConfigType is ACME\_TEMPLATE (see [below for nested schema](#nestedatt--items--config--acme_certificate))
* `device_ssl_certificate_configs` (Attributes List) - Available when ConfigType is SSL\_CERTIFICATE\_TEMPLATE (see [below for nested schema](#nestedatt--items--config--device_ssl_certificate_configs))
* `export_metadata_app_profile` (Attributes) (see [below for nested schema](#nestedatt--items--config--export_metadata_app_profile))
* `giga_port_neighbors_discovery_config` (Attributes) - To contain the port neighbour discovery template config (see [below for nested schema](#nestedatt--items--config--giga_port_neighbors_discovery_config))
* `giga_stream_threshold_config` (Attributes) - Gigastream threshold configs (see [below for nested schema](#nestedatt--items--config--giga_stream_threshold_config))
* `giga_user_defined_application_config` (Attributes) (see [below for nested schema](#nestedatt--items--config--giga_user_defined_application_config))
* `ldap_server_system_config` (Dynamic) - Available for ConfigType LDAP\_SYSTEM\_CONFIG\_TEMPLATE
* `ldap_servers` (Attributes List) - List of LDAP Servers. Available for ConfigType LDAP\_SERVERS\_TEMPLATE ChildType ldapServers (see [below for nested schema](#nestedatt--items--config--ldap_servers))
* `metadata_exporter` (Attributes) (see [below for nested schema](#nestedatt--items--config--metadata_exporter))
* `port_packet_threshold_config` (Attributes) - Contains port threshold configs (see [below for nested schema](#nestedatt--items--config--port_packet_threshold_config))
* `proxy_server_profile` (Attributes) (see [below for nested schema](#nestedatt--items--config--proxy_server_profile))
* `snmp_trap_event_configs` (Attributes List) - Available when ConfigType is SNMPTRAPS (see [below for nested schema](#nestedatt--items--config--snmp_trap_event_configs))
* `snmp_v3_users_config` (Attributes) - Contains list of FM's snmpv3 users (see [below for nested schema](#nestedatt--items--config--snmp_v3_users_config))
* `ssh_ciphers_config` (Attributes) - System SSH Cipher (see [below for nested schema](#nestedatt--items--config--ssh_ciphers_config))

<a id="nestedatt--items--config--aaa_auth_config"></a>
### Nested Schema for `items.config.aaa_auth_config`

Read-Only:

* `auth_sequence` (Set of String) - Authentication methods order for user login. Valid values are 'local', 'ldap', 'radius', 'tacacs'
* `external_login_mapping` (Attributes) - Specifies how externally authenticated logins (RADIUS, TACACS+, or LDAP) are mapped to local accounts. Private class (see [below for nested schema](#nestedatt--items--config--aaa_auth_config--external_login_mapping))

<a id="nestedatt--items--config--aaa_auth_config--external_login_mapping"></a>
### Nested Schema for `items.config.aaa_auth_config.external_login_mapping`

Read-Only:

* `default_local_user` (String) - Specifies the account to which externally authenticated logins are mapped when map order is set to remote-first (if there is no matching local account) or local-only
* `user_map_order` (String) - Specifies how externally authenticated logins (RADIUS, TACACS+, or LDAP) are mapped to local accounts

<a id="nestedatt--items--config--acme_certificate"></a>
### Nested Schema for `items.config.acme_certificate`

Read-Only:

* `acme_server_url` (String)
* `algorithm` (String)
* `operation_type` (String)
* `renew_days` (Number) - default will be 1/3rd of certificate validity period

<a id="nestedatt--items--config--device_ssl_certificate_configs"></a>
### Nested Schema for `items.config.device_ssl_certificate_configs`

Read-Only:

* `issuer` (String) - issuer details of the certificate
* `not_after` (String) - date and time when certificate stops being valid (\[rfc3339\](https://xml2rfc.tools.ietf.org/public/rfc/html/rfc3339.html\#anchor14))
* `not_before` (String) - date and time when certificate starts being valid (\[rfc3339\](https://xml2rfc.tools.ietf.org/public/rfc/html/rfc3339.html\#anchor14))
* `operation_type` (String)
* `signature_algorithm` (String)
* `subject` (String) - subject name of the certificate
* `trusted_ca` (Attributes) (see [below for nested schema](#nestedatt--items--config--device_ssl_certificate_configs--trusted_ca))
* `upload_spec` (Attributes) (see [below for nested schema](#nestedatt--items--config--device_ssl_certificate_configs--upload_spec))

<a id="nestedatt--items--config--device_ssl_certificate_configs--trusted_ca"></a>
### Nested Schema for `items.config.device_ssl_certificate_configs.trusted_ca`

Read-Only:

* `name` (String) - name of the certificate

<a id="nestedatt--items--config--device_ssl_certificate_configs--upload_spec"></a>
### Nested Schema for `items.config.device_ssl_certificate_configs.upload_spec`

Read-Only:

* `info` (Attributes) - Certificate info (see [below for nested schema](#nestedatt--items--config--device_ssl_certificate_configs--upload_spec--info))
* `pem` (String) - contents of the certificate in pem format

<a id="nestedatt--items--config--device_ssl_certificate_configs--upload_spec--info"></a>
### Nested Schema for `items.config.device_ssl_certificate_configs.upload_spec.info`

Read-Only:

* `comment` (String) - a short description of the certificate
* `name` (String) - name of the certificate
* `passphrase` (String) - used to decrypt pkcs12 and private keys
* `type` (String) - type of the certificate

<a id="nestedatt--items--config--export_metadata_app_profile"></a>
### Nested Schema for `items.config.export_metadata_app_profile`

Read-Only:

* `alias` (String) - application profile alias
* `application_id` (Boolean) - only valid with 'export' type
* `applications` (Attributes List) - application and attributes. (see [below for nested schema](#nestedatt--items--config--export_metadata_app_profile--applications))
* `counter` (Attributes) (see [below for nested schema](#nestedatt--items--config--export_metadata_app_profile--counter))
* `datalink` (Attributes) (see [below for nested schema](#nestedatt--items--config--export_metadata_app_profile--datalink))
* `description` (String)
* `flow` (Attributes) (see [below for nested schema](#nestedatt--items--config--export_metadata_app_profile--flow))
* `gtpu` (Attributes) (see [below for nested schema](#nestedatt--items--config--export_metadata_app_profile--gtpu))
* `interface` (Attributes) (see [below for nested schema](#nestedatt--items--config--export_metadata_app_profile--interface))
* `ip` (Attributes) (see [below for nested schema](#nestedatt--items--config--export_metadata_app_profile--ip))
* `ipv4` (Attributes) (see [below for nested schema](#nestedatt--items--config--export_metadata_app_profile--ipv4))
* `ipv6` (Attributes) (see [below for nested schema](#nestedatt--items--config--export_metadata_app_profile--ipv6))
* `outer_ipv4` (Attributes) (see [below for nested schema](#nestedatt--items--config--export_metadata_app_profile--outer_ipv4))
* `outer_ipv6` (Attributes) (see [below for nested schema](#nestedatt--items--config--export_metadata_app_profile--outer_ipv6))
* `timestamp` (Attributes) (see [below for nested schema](#nestedatt--items--config--export_metadata_app_profile--timestamp))
* `transport` (Attributes) (see [below for nested schema](#nestedatt--items--config--export_metadata_app_profile--transport))
* `type` (String)

<a id="nestedatt--items--config--export_metadata_app_profile--applications"></a>
### Nested Schema for `items.config.export_metadata_app_profile.applications`

Read-Only:

* `attributes` (Attributes List) (see [below for nested schema](#nestedatt--items--config--export_metadata_app_profile--applications--attributes))
* `is_user_defined` (Boolean) - Default value is false. Must be set to true only for App Intel solution while configuring user defined apps
* `name` (String) - application name

<a id="nestedatt--items--config--export_metadata_app_profile--applications--attributes"></a>
### Nested Schema for `items.config.export_metadata_app_profile.applications.attributes`

Read-Only:

* `name` (String) - attribute name
* `value` (String) - application's attribute value

<a id="nestedatt--items--config--export_metadata_app_profile--counter"></a>
### Nested Schema for `items.config.export_metadata_app_profile.counter`

Read-Only:

* `bytes` (Boolean)
* `bytes_long` (Boolean)
* `inner_byte` (Boolean)
* `inner_byte_long` (Boolean)
* `packets` (Boolean)
* `packets_long` (Boolean)

<a id="nestedatt--items--config--export_metadata_app_profile--datalink"></a>
### Nested Schema for `items.config.export_metadata_app_profile.datalink`

Read-Only:

* `mac_dst` (Boolean)
* `mac_src` (Boolean)
* `vlan` (Boolean)

<a id="nestedatt--items--config--export_metadata_app_profile--flow"></a>
### Nested Schema for `items.config.export_metadata_app_profile.flow`

Read-Only:

* `end_reason` (Boolean)

<a id="nestedatt--items--config--export_metadata_app_profile--gtpu"></a>
### Nested Schema for `items.config.export_metadata_app_profile.gtpu`

Read-Only:

* `qfi` (Boolean)
* `teid` (Boolean)

<a id="nestedatt--items--config--export_metadata_app_profile--interface"></a>
### Nested Schema for `items.config.export_metadata_app_profile.interface`

Read-Only:

* `in_name_width` (Number)
* `in_physical_width` (Number)
* `out_physical_width` (Number)

<a id="nestedatt--items--config--export_metadata_app_profile--ip"></a>
### Nested Schema for `items.config.export_metadata_app_profile.ip`

Read-Only:

* `version` (Boolean)

<a id="nestedatt--items--config--export_metadata_app_profile--ipv4"></a>
### Nested Schema for `items.config.export_metadata_app_profile.ipv4`

Read-Only:

* `destination` (Attributes) (see [below for nested schema](#nestedatt--items--config--export_metadata_app_profile--ipv4--destination))
* `dscp` (Boolean)
* `fragmentation` (Attributes) (see [below for nested schema](#nestedatt--items--config--export_metadata_app_profile--ipv4--fragmentation))
* `header_len` (Boolean)
* `option_map` (Boolean)
* `precedence` (Boolean)
* `protocol` (Boolean)
* `section` (Attributes) (see [below for nested schema](#nestedatt--items--config--export_metadata_app_profile--ipv4--section))
* `source` (Attributes) (see [below for nested schema](#nestedatt--items--config--export_metadata_app_profile--ipv4--source))
* `tos` (Boolean)
* `total_length` (Boolean)
* `ttl` (Boolean)

<a id="nestedatt--items--config--export_metadata_app_profile--ipv4--destination"></a>
### Nested Schema for `items.config.export_metadata_app_profile.ipv4.destination`

Read-Only:

* `prefix_min_mask` (String)

<a id="nestedatt--items--config--export_metadata_app_profile--ipv4--fragmentation"></a>
### Nested Schema for `items.config.export_metadata_app_profile.ipv4.fragmentation`

Read-Only:

* `flags` (Boolean)
* `offset` (Boolean)

<a id="nestedatt--items--config--export_metadata_app_profile--ipv4--section"></a>
### Nested Schema for `items.config.export_metadata_app_profile.ipv4.section`

Read-Only:

* `header_size` (Number)
* `payload_size` (Number)

<a id="nestedatt--items--config--export_metadata_app_profile--ipv4--source"></a>
### Nested Schema for `items.config.export_metadata_app_profile.ipv4.source`

Read-Only:

* `prefix_min_mask` (String) - ipv4 source prefix minimum-mask - netmask or mask length

<a id="nestedatt--items--config--export_metadata_app_profile--ipv6"></a>
### Nested Schema for `items.config.export_metadata_app_profile.ipv6`

Read-Only:

* `destination` (Attributes) (see [below for nested schema](#nestedatt--items--config--export_metadata_app_profile--ipv6--destination))
* `dscp` (Boolean)
* `extension_map` (Boolean)
* `flow_label` (Boolean)
* `fragmentation` (Attributes) (see [below for nested schema](#nestedatt--items--config--export_metadata_app_profile--ipv6--fragmentation))
* `hop_limit` (Boolean)
* `length` (Attributes) (see [below for nested schema](#nestedatt--items--config--export_metadata_app_profile--ipv6--length))
* `next_header` (Boolean)
* `precedence` (Boolean)
* `section` (Attributes) (see [below for nested schema](#nestedatt--items--config--export_metadata_app_profile--ipv6--section))
* `source` (Attributes) (see [below for nested schema](#nestedatt--items--config--export_metadata_app_profile--ipv6--source))
* `traffic_class` (Boolean)

<a id="nestedatt--items--config--export_metadata_app_profile--ipv6--destination"></a>
### Nested Schema for `items.config.export_metadata_app_profile.ipv6.destination`

Read-Only:

* `prefix_min_mask` (String)

<a id="nestedatt--items--config--export_metadata_app_profile--ipv6--fragmentation"></a>
### Nested Schema for `items.config.export_metadata_app_profile.ipv6.fragmentation`

Read-Only:

* `flags` (Boolean)
* `offset` (Boolean)

<a id="nestedatt--items--config--export_metadata_app_profile--ipv6--length"></a>
### Nested Schema for `items.config.export_metadata_app_profile.ipv6.length`

Read-Only:

* `header` (Boolean)
* `payload` (Boolean)
* `total` (Boolean)

<a id="nestedatt--items--config--export_metadata_app_profile--ipv6--section"></a>
### Nested Schema for `items.config.export_metadata_app_profile.ipv6.section`

Read-Only:

* `header_size` (Number)
* `payload_size` (Number)

<a id="nestedatt--items--config--export_metadata_app_profile--ipv6--source"></a>
### Nested Schema for `items.config.export_metadata_app_profile.ipv6.source`

Read-Only:

* `prefix_min_mask` (String)

<a id="nestedatt--items--config--export_metadata_app_profile--outer_ipv4"></a>
### Nested Schema for `items.config.export_metadata_app_profile.outer_ipv4`

Read-Only:

* `destination` (Boolean)
* `source` (Boolean)

<a id="nestedatt--items--config--export_metadata_app_profile--outer_ipv6"></a>
### Nested Schema for `items.config.export_metadata_app_profile.outer_ipv6`

Read-Only:

* `destination` (Boolean)
* `source` (Boolean)

<a id="nestedatt--items--config--export_metadata_app_profile--timestamp"></a>
### Nested Schema for `items.config.export_metadata_app_profile.timestamp`

Read-Only:

* `flow_end_msec` (Boolean)
* `flow_endsec` (Boolean)
* `flow_start_msec` (Boolean)
* `flow_startsec` (Boolean)
* `sys_up_time_first` (Boolean)
* `sys_up_time_last` (Boolean)

<a id="nestedatt--items--config--export_metadata_app_profile--transport"></a>
### Nested Schema for `items.config.export_metadata_app_profile.transport`

Read-Only:

* `dst_port` (Boolean)
* `icmp` (Attributes) (see [below for nested schema](#nestedatt--items--config--export_metadata_app_profile--transport--icmp))
* `src_port` (Boolean)
* `tcp` (Attributes) (see [below for nested schema](#nestedatt--items--config--export_metadata_app_profile--transport--tcp))
* `udp` (Attributes) (see [below for nested schema](#nestedatt--items--config--export_metadata_app_profile--transport--udp))

<a id="nestedatt--items--config--export_metadata_app_profile--transport--icmp"></a>
### Nested Schema for `items.config.export_metadata_app_profile.transport.icmp`

Read-Only:

* `ipv4_code` (Boolean)
* `ipv4_type` (Boolean)
* `ipv6_code` (Boolean)
* `ipv6_type` (Boolean)

<a id="nestedatt--items--config--export_metadata_app_profile--transport--tcp"></a>
### Nested Schema for `items.config.export_metadata_app_profile.transport.tcp`

Read-Only:

* `ack_number` (Boolean)
* `dst_port` (Boolean)
* `flags` (Boolean)
* `header_len` (Boolean)
* `seq_number` (Boolean)
* `src_port` (Boolean)
* `urgent_ptr` (Boolean)
* `window_size` (Boolean)

<a id="nestedatt--items--config--export_metadata_app_profile--transport--udp"></a>
### Nested Schema for `items.config.export_metadata_app_profile.transport.udp`

Read-Only:

* `dst_port` (Boolean)
* `msg_len` (Boolean)
* `src_port` (Boolean)

<a id="nestedatt--items--config--giga_port_neighbors_discovery_config"></a>
### Nested Schema for `items.config.giga_port_neighbors_discovery_config`

Read-Only:

* `resource_configs` (Attributes List) - List of config for each port type (see [below for nested schema](#nestedatt--items--config--giga_port_neighbors_discovery_config--resource_configs))

<a id="nestedatt--items--config--giga_port_neighbors_discovery_config--resource_configs"></a>
### Nested Schema for `items.config.giga_port_neighbors_discovery_config.resource_configs`

Read-Only:

* `cdp` (Boolean)
* `gdp` (Boolean)
* `lldp` (Boolean)
* `port_type` (Dynamic) - Type of the port

<a id="nestedatt--items--config--giga_stream_threshold_config"></a>
### Nested Schema for `items.config.giga_stream_threshold_config`

Read-Only:

* `giga_stream_type_thresholds` (Attributes List) - GigaStream threshold config for port types (see [below for nested schema](#nestedatt--items--config--giga_stream_threshold_config--giga_stream_type_thresholds))

<a id="nestedatt--items--config--giga_stream_threshold_config--giga_stream_type_thresholds"></a>
### Nested Schema for `items.config.giga_stream_threshold_config.giga_stream_type_thresholds`

Read-Only:

* `type` (String) - Port type
* `variance_threshold` (Number) - Threshold value in percentage. '-1' will disable threshold feature.

<a id="nestedatt--items--config--giga_user_defined_application_config"></a>
### Nested Schema for `items.config.giga_user_defined_application_config`

Read-Only:

* `alias` (String) - pattern: ^\[a-zA-Z0-9&'+\_=\|\-\]+$
* `app_id` (Number) - minimum: 16777216 maximum: 16777343   Ranges from 2^24 to 2^24 + 127. It needs to be unique
* `priority` (Number) - minimum: 1 maximum: 120  Lower the priority higher the precedence
* `rules` (Attributes) (see [below for nested schema](#nestedatt--items--config--giga_user_defined_application_config--rules))

<a id="nestedatt--items--config--giga_user_defined_application_config--rules"></a>
### Nested Schema for `items.config.giga_user_defined_application_config.rules`

Read-Only:

* `rule` (Attributes List) (see [below for nested schema](#nestedatt--items--config--giga_user_defined_application_config--rules--rule))

<a id="nestedatt--items--config--giga_user_defined_application_config--rules--rule"></a>
### Nested Schema for `items.config.giga_user_defined_application_config.rules.rule`

Read-Only:

* `address` (String)
* `code` (String)
* `common_name` (String)
* `content` (String)
* `cts_cookie` (String)
* `cts_page_url` (String)
* `cts_referer` (String)
* `cts_server` (String)
* `cts_uri` (String)
* `cts_user_agent` (String)
* `dscp` (String)
* `mime_type` (String)
* `mindata` (Number)
* `port` (String)
* `resolv_name` (String)
* `stc_location` (String)
* `stc_server_agent` (String)
* `stc_subject_alt_name` (String)
* `stream` (String)
* `typeval` (String)
* `user_agent` (String)

<a id="nestedatt--items--config--ldap_servers"></a>
### Nested Schema for `items.config.ldap_servers`

Read-Only:

* `order` (String) - The order in which the server is to be reached. 1 means server will be contacted first
* `server_address` (String) - ipv4 or ipv6 or hostname. Specifies address of the LDAP server where authentication requests will be sent

<a id="nestedatt--items--config--metadata_exporter"></a>
### Nested Schema for `items.config.metadata_exporter`

Read-Only:

* `alias` (String)
* `application_profiles` (List of String) - application profile aliases to attach to the exporter
* `cef` (Attributes) - cef attributes (see [below for nested schema](#nestedatt--items--config--metadata_exporter--cef))
* `description` (String)
* `destination` (Attributes) - destination attributes (see [below for nested schema](#nestedatt--items--config--metadata_exporter--destination))
* `max_pkt_size` (Number)
* `mobility_sam` (Attributes) (see [below for nested schema](#nestedatt--items--config--metadata_exporter--mobility_sam))
* `monitor` (Attributes) - monitor attributes (see [below for nested schema](#nestedatt--items--config--metadata_exporter--monitor))
* `netflow` (Attributes) - netflow attributes (see [below for nested schema](#nestedatt--items--config--metadata_exporter--netflow))
* `snmp` (Attributes) (see [below for nested schema](#nestedatt--items--config--metadata_exporter--snmp))
* `source` (Attributes) - source tunnel port (see [below for nested schema](#nestedatt--items--config--metadata_exporter--source))
* `type` (String)

<a id="nestedatt--items--config--metadata_exporter--cef"></a>
### Nested Schema for `items.config.metadata_exporter.cef`

Read-Only:

* `active_timeout` (Number) - in seconds
* `inactive_timeout` (Number) - in seconds

<a id="nestedatt--items--config--metadata_exporter--destination"></a>
### Nested Schema for `items.config.metadata_exporter.destination`

Read-Only:

* `dscp` (Number)
* `ipv4_address` (String) - ipv4 address
* `l4_port_dst` (Number)
* `l4_port_src` (Number)
* `l4_protocol` (String)
* `ttl` (Number)

<a id="nestedatt--items--config--metadata_exporter--mobility_sam"></a>
### Nested Schema for `items.config.metadata_exporter.mobility_sam`

Read-Only:

* `encoding` (String)
* `encoding_format` (String)
* `event_enable` (Attributes) (see [below for nested schema](#nestedatt--items--config--metadata_exporter--mobility_sam--event_enable))
* `trigger` (String)

<a id="nestedatt--items--config--metadata_exporter--mobility_sam--event_enable"></a>
### Nested Schema for `items.config.metadata_exporter.mobility_sam.event_enable`

Read-Only:

* `modify` (Boolean)
* `update` (Boolean)

<a id="nestedatt--items--config--metadata_exporter--monitor"></a>
### Nested Schema for `items.config.metadata_exporter.monitor`

Read-Only:

* `timeout` (Number) - how often to export in seconds

<a id="nestedatt--items--config--metadata_exporter--netflow"></a>
### Nested Schema for `items.config.metadata_exporter.netflow`

Read-Only:

* `active_timeout` (Number) - in seconds
* `inactive_timeout` (Number) - in seconds
* `template_refresh` (Number) - template refresh interval in seconds
* `template_type` (String)
* `version` (String)

<a id="nestedatt--items--config--metadata_exporter--snmp"></a>
### Nested Schema for `items.config.metadata_exporter.snmp`

Read-Only:

* `enabled` (Boolean) - snmp reverse lookup enable/disable

<a id="nestedatt--items--config--metadata_exporter--source"></a>
### Nested Schema for `items.config.metadata_exporter.source`

Read-Only:

* `ip_interface` (String)

<a id="nestedatt--items--config--port_packet_threshold_config"></a>
### Nested Schema for `items.config.port_packet_threshold_config`

Read-Only:

* `drop_threshold` (Attributes) - Drop threshold config for Rx and Tx (see [below for nested schema](#nestedatt--items--config--port_packet_threshold_config--drop_threshold))
* `error_threshold` (Attributes) - Error threshold config for Rx and Tx (see [below for nested schema](#nestedatt--items--config--port_packet_threshold_config--error_threshold))

<a id="nestedatt--items--config--port_packet_threshold_config--drop_threshold"></a>
### Nested Schema for `items.config.port_packet_threshold_config.drop_threshold`

Read-Only:

* `rx` (Attributes List) - Threshold values for Rx packet drop (see [below for nested schema](#nestedatt--items--config--port_packet_threshold_config--drop_threshold--rx))
* `tx` (Attributes List) - Threshold values for Tx packet drop (see [below for nested schema](#nestedatt--items--config--port_packet_threshold_config--drop_threshold--tx))

<a id="nestedatt--items--config--port_packet_threshold_config--drop_threshold--rx"></a>
### Nested Schema for `items.config.port_packet_threshold_config.drop_threshold.rx`

Read-Only:

* `count_` (Number) - Threshold value in number of packets out of total packets in a given interval. '-1' will disable threshold feature.
* `percent` (Number) - Threshold value in percent of packets out of total packets in a given interval. '-1' will disable threshold feature.
* `port_type` (String) - Port type

<a id="nestedatt--items--config--port_packet_threshold_config--drop_threshold--tx"></a>
### Nested Schema for `items.config.port_packet_threshold_config.drop_threshold.tx`

Read-Only:

* `count_` (Number) - Threshold value in number of packets out of total packets in a given interval. '-1' will disable threshold feature.
* `percent` (Number) - Threshold value in percent of packets out of total packets in a given interval. '-1' will disable threshold feature.
* `port_type` (String) - Port type

<a id="nestedatt--items--config--port_packet_threshold_config--error_threshold"></a>
### Nested Schema for `items.config.port_packet_threshold_config.error_threshold`

Read-Only:

* `rx` (Attributes List) - Threshold values for Rx packet error (see [below for nested schema](#nestedatt--items--config--port_packet_threshold_config--error_threshold--rx))
* `tx` (Attributes List) - Threshold values for Tx packet error (see [below for nested schema](#nestedatt--items--config--port_packet_threshold_config--error_threshold--tx))

<a id="nestedatt--items--config--port_packet_threshold_config--error_threshold--rx"></a>
### Nested Schema for `items.config.port_packet_threshold_config.error_threshold.rx`

Read-Only:

* `count_` (Number) - Threshold value in number of packets out of total packets in a given interval. '-1' will disable threshold feature.
* `percent` (Number) - Threshold value in percent of packets out of total packets in a given interval. '-1' will disable threshold feature.
* `port_type` (String) - Port type

<a id="nestedatt--items--config--port_packet_threshold_config--error_threshold--tx"></a>
### Nested Schema for `items.config.port_packet_threshold_config.error_threshold.tx`

Read-Only:

* `count_` (Number) - Threshold value in number of packets out of total packets in a given interval. '-1' will disable threshold feature.
* `percent` (Number) - Threshold value in percent of packets out of total packets in a given interval. '-1' will disable threshold feature.
* `port_type` (String) - Port type

<a id="nestedatt--items--config--proxy_server_profile"></a>
### Nested Schema for `items.config.proxy_server_profile`

Read-Only:

* `alias` (String)
* `auth_type` (String)
* `comment` (String)
* `password` (String)
* `periodic_ping` (String)
* `periodic_ping_failure_retry` (Number)
* `periodic_ping_interval` (Number)
* `periodic_ping_type` (String)
* `port` (Number)
* `protocol` (String)
* `proxy_address` (String)
* `ssl_apps` (Attributes) (see [below for nested schema](#nestedatt--items--config--proxy_server_profile--ssl_apps))
* `username` (String)

<a id="nestedatt--items--config--proxy_server_profile--ssl_apps"></a>
### Nested Schema for `items.config.proxy_server_profile.ssl_apps`

Read-Only:

* `cluster_name` (List of String) - Cluster Name where the proxy deployed

<a id="nestedatt--items--config--snmp_trap_event_configs"></a>
### Nested Schema for `items.config.snmp_trap_event_configs`

Read-Only:

* `enabled` (Boolean)
* `notify_event` (String)

<a id="nestedatt--items--config--snmp_v3_users_config"></a>
### Nested Schema for `items.config.snmp_v3_users_config`

Read-Only:

* `snmp_v3_user` (Attributes List) (see [below for nested schema](#nestedatt--items--config--snmp_v3_users_config--snmp_v3_user))

<a id="nestedatt--items--config--snmp_v3_users_config--snmp_v3_user"></a>
### Nested Schema for `items.config.snmp_v3_users_config.snmp_v3_user`

Read-Only:

* `auth_key` (String) - auth key/passphrase
* `auth_protocol` (String)
* `min_sw_version` (String) - minimum software version of the device
* `previous_username` (String) - When the username is changed, the username prior changing is sent in this property. Based on this application handles the username update logic. Not found error is returned when this property is not sent during the username change
* `priv_key` (String) - privacy key/passphrase
* `priv_protocol` (String)
* `username` (String) - snmpv3 username

<a id="nestedatt--items--config--ssh_ciphers_config"></a>
### Nested Schema for `items.config.ssh_ciphers_config`

Read-Only:

* `classic` (Attributes) - System SSH Cipher (see [below for nested schema](#nestedatt--items--config--ssh_ciphers_config--classic))
* `crypto` (Attributes) - System SSH Cipher (see [below for nested schema](#nestedatt--items--config--ssh_ciphers_config--crypto))
* `fips` (Attributes) - System SSH Cipher (see [below for nested schema](#nestedatt--items--config--ssh_ciphers_config--fips))

<a id="nestedatt--items--config--ssh_ciphers_config--classic"></a>
### Nested Schema for `items.config.ssh_ciphers_config.classic`

Read-Only:

* `client_ciphers` (List of String)
* `client_hostkey` (List of String)
* `client_kex` (List of String)
* `client_macs` (List of String)
* `server_ciphers` (List of String)
* `server_hostkey` (List of String)
* `server_kex` (List of String)
* `server_macs` (List of String)

<a id="nestedatt--items--config--ssh_ciphers_config--crypto"></a>
### Nested Schema for `items.config.ssh_ciphers_config.crypto`

Read-Only:

* `client_ciphers` (List of String)
* `client_hostkey` (List of String)
* `client_kex` (List of String)
* `client_macs` (List of String)
* `server_ciphers` (List of String)
* `server_hostkey` (List of String)
* `server_kex` (List of String)
* `server_macs` (List of String)

<a id="nestedatt--items--config--ssh_ciphers_config--fips"></a>
### Nested Schema for `items.config.ssh_ciphers_config.fips`

Read-Only:

* `client_ciphers` (List of String)
* `client_hostkey` (List of String)
* `client_kex` (List of String)
* `client_macs` (List of String)
* `server_ciphers` (List of String)
* `server_hostkey` (List of String)
* `server_kex` (List of String)
* `server_macs` (List of String)

<a id="nestedatt--items--ref_object"></a>
### Nested Schema for `items.ref_object`

Read-Only:

* `ref_object_type` (String) - Type of the reference object

