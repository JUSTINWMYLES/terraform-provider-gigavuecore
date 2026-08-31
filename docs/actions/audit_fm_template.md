---
page_title: "gigavuecore_audit_fm_template Action - gigavuecore"
subcategory: ""
description: |-
  new in FM 5.7
---

# gigavuecore_audit_fm_template Action

new in FM 5.7

## Example Usage

```terraform
action "gigavuecore_audit_fm_template" "example" {
  config {
    async              = true
    body_config_type   = "SNMPTRAPS"
    body_template_name = "example"
    config = {
      aaa_auth_config = {
        auth_sequence = [ "local" ]
        external_login_mapping = {
          default_local_user = "example"
          user_map_order     = "localOnly"
        }
      }
      acme_certificate = [{
        acme_server_url = "example"
        algorithm       = "rsa-2048"
        operation_type  = "issue"
        renew_days      = 0
      }]
      device_ssl_certificate_configs = [{
        issuer              = "example"
        not_after           = "example"
        not_before          = "example"
        operation_type      = "add"
        signature_algorithm = "example"
        subject             = "example"
        trusted_ca = {
          name = "example"
        }
        upload_spec = {
          info = {
            comment    = "example"
            name       = "example"
            passphrase = "example"
            type       = "privateKey"
          }
          pem = "example"
        }
      }]
      export_metadata_app_profile = {
        alias          = "example"
        application_id = true
        applications = [{
          attributes = [{
            name  = "example"
            value = "example"
          }]
          is_user_defined = true
          name            = "example"
        }]
        counter = {
          bytes           = true
          bytes_long      = true
          inner_byte      = true
          inner_byte_long = true
          packets         = true
          packets_long    = true
        }
        datalink = {
          mac_dst = true
          mac_src = true
          vlan    = true
        }
        description = "example"
        flow = {
          end_reason = true
        }
        gtpu = {
          qfi  = true
          teid = true
        }
        interface = {
          in_name_width      = 1
          in_physical_width  = 2
          out_physical_width = 2
        }
        ip = {
          version = true
        }
        ipv4 = {
          destination = {
            prefix_min_mask = "example"
          }
          dscp = true
          fragmentation = {
            flags  = true
            offset = true
          }
          header_len = true
          option_map = true
          precedence = true
          protocol   = true
          section = {
            header_size  = 1
            payload_size = 1
          }
          source = {
            prefix_min_mask = "example"
          }
          tos          = true
          total_length = true
          ttl          = true
        }
        ipv6 = {
          destination = {
            prefix_min_mask = "example"
          }
          dscp          = true
          extension_map = true
          flow_label    = true
          fragmentation = {
            flags  = true
            offset = true
          }
          hop_limit = true
          length = {
            header  = true
            payload = true
            total   = true
          }
          next_header = true
          precedence  = true
          section = {
            header_size  = 1
            payload_size = 1
          }
          source = {
            prefix_min_mask = "example"
          }
          traffic_class = true
        }
        outer_ipv4 = {
          destination = true
          source      = true
        }
        outer_ipv6 = {
          destination = true
          source      = true
        }
        timestamp = {
          flow_end_msec     = true
          flow_endsec       = true
          flow_start_msec   = true
          flow_startsec     = true
          sys_up_time_first = true
          sys_up_time_last  = true
        }
        transport = {
          dst_port = true
          icmp = {
            ipv4_code = true
            ipv4_type = true
            ipv6_code = true
            ipv6_type = true
          }
          src_port = true
          tcp = {
            ack_number  = true
            dst_port    = true
            flags       = true
            header_len  = true
            seq_number  = true
            src_port    = true
            urgent_ptr  = true
            window_size = true
          }
          udp = {
            dst_port = true
            msg_len  = true
            src_port = true
          }
        }
        type = "export"
      }
      giga_port_neighbors_discovery_config = {
        resource_configs = [{
          cdp       = true
          gdp       = true
          lldp      = true
          port_type = "example"
        }]
      }
      giga_stream_threshold_config = {
        giga_stream_type_thresholds = [{
          type               = "all"
          variance_threshold = 1.0
        }]
      }
      giga_user_defined_application_config = {
        alias    = "example"
        app_id   = 0
        priority = 0
        rules = {
          rule = [{
            address              = "example"
            code                 = "example"
            common_name          = "example"
            content              = "example"
            cts_cookie           = "example"
            cts_page_url         = "example"
            cts_referer          = "example"
            cts_server           = "example"
            cts_uri              = "example"
            cts_user_agent       = "example"
            dscp                 = "example"
            mime_type            = "example"
            mindata              = 0
            port                 = "example"
            resolv_name          = "example"
            stc_location         = "example"
            stc_server_agent     = "example"
            stc_subject_alt_name = "example"
            stream               = "example"
            typeval              = "example"
            user_agent           = "example"
          }]
        }
      }
      ldap_server_system_config = "example"
      ldap_servers = [{
        order          = "example"
        server_address = "example"
      }]
      metadata_exporter = {
        alias                = "example"
        application_profiles = [ "example" ]
        cef = {
          active_timeout   = 1
          inactive_timeout = 1
        }
        description = "example"
        destination = {
          dscp         = 0
          ipv4_address = "example"
          l4_port_dst  = 1
          l4_port_src  = 1
          l4_protocol  = "udp"
          ttl          = 1
        }
        max_pkt_size = 0
        mobility_sam = {
          encoding        = "example"
          encoding_format = "hierarchy"
          event_enable = {
            modify = true
            update = true
          }
          trigger = "example"
        }
        monitor = {
          timeout = 60
        }
        netflow = {
          active_timeout   = 1
          inactive_timeout = 1
          template_refresh = 1
          template_type    = "cohesive"
          version          = "v5"
        }
        snmp = {
          enabled = true
        }
        source = {
          ip_interface = "example"
        }
        type = "cef"
      }
      port_packet_threshold_config = {
        drop_threshold = {
          rx = [{
            count_    = 0
            percent   = 1.0
            port_type = "all"
          }]
          tx = [{
            count_    = 0
            percent   = 1.0
            port_type = "all"
          }]
        }
        error_threshold = {
          rx = [{
            count_    = 0
            percent   = 1.0
            port_type = "all"
          }]
          tx = [{
            count_    = 0
            percent   = 1.0
            port_type = "all"
          }]
        }
      }
      proxy_server_profile = {
        alias                       = "example"
        auth_type                   = "none"
        comment                     = "example"
        password                    = "example"
        periodic_ping               = "enable"
        periodic_ping_failure_retry = 1
        periodic_ping_interval      = 1
        periodic_ping_type          = "http-connect"
        port                        = 1
        protocol                    = "http"
        proxy_address               = "example"
        ssl_apps = {
          cluster_name = [ "example" ]
        }
        username = "example"
      }
      snmp_trap_event_configs = [{
        enabled      = true
        notify_event = "example"
      }]
      snmp_v3_users_config = {
        snmp_v3_user = [{
          auth_key          = "example"
          auth_protocol     = "md5"
          min_sw_version    = "example"
          previous_username = "example"
          priv_key          = "example"
          priv_protocol     = "des"
          username          = "example"
        }]
      }
      ssh_ciphers_config = {
        classic = {
          client_ciphers = [ "default" ]
          client_hostkey = [ "default" ]
          client_kex     = [ "default" ]
          client_macs    = [ "default" ]
          server_ciphers = [ "default" ]
          server_hostkey = [ "default" ]
          server_kex     = [ "default" ]
          server_macs    = [ "default" ]
        }
        crypto = {
          client_ciphers = [ "default" ]
          client_hostkey = [ "default" ]
          client_kex     = [ "default" ]
          client_macs    = [ "default" ]
          server_ciphers = [ "default" ]
          server_hostkey = [ "default" ]
          server_kex     = [ "default" ]
          server_macs    = [ "default" ]
        }
        fips = {
          client_ciphers = [ "default" ]
          client_hostkey = [ "default" ]
          client_kex     = [ "default" ]
          client_macs    = [ "default" ]
          server_ciphers = [ "default" ]
          server_hostkey = [ "default" ]
          server_kex     = [ "default" ]
          server_macs    = [ "default" ]
        }
      }
    }
    config_level       = "GLOBAL"
    config_level_value = [ "example" ]
    config_resource    = "example"
    config_type        = "example"
    modifiable         = true
    ref_count          = 0
    ref_object = {
      ref_object_type = "PHYSICAL"
    }
    template_name = "example"
    update_time   = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `async` (Boolean, optional) - audit operation is sync or async
* `body_config_type` (String, optional) - Configuration Type of the FM template
* `body_template_name` (String, optional)
* `config` (Attributes, optional) (see [below for nested schema](#nestedatt--config))
* `config_level` (String, optional) - Scope of the applied FM template
* `config_level_value` (List of String, optional)
* `config_resource` (Dynamic, optional) - For a particular config type, there may be multiple templates with different levels. For example, cluster level the user can have more than one template and in same way in tag level there may be more than one template for each tag combination. In such cases, the configResource property can be used to contain the list of clusters associated to the template or the tag key values combination of the template. The data structure changes based on the level hence it is kept as type Object.
* `config_type` (String, required) - configType of the fm template
* `modifiable` (Boolean, optional)
* `ref_count` (Number, optional)
* `ref_object` (Attributes, optional) (see [below for nested schema](#nestedatt--ref_object))
* `template_name` (String, optional) - Name of the template
* `update_time` (String, optional)

<a id="nestedatt--config"></a>
### Nested Schema for `config`

Optional:

* `aaa_auth_config` (Attributes) - FM Global AAA Authentication config to the device's (see [below for nested schema](#nestedatt--config--aaa_auth_config))
* `acme_certificate` (Attributes List) - Available when ConfigType is ACME\_TEMPLATE (see [below for nested schema](#nestedatt--config--acme_certificate))
* `device_ssl_certificate_configs` (Attributes List) - Available when ConfigType is SSL\_CERTIFICATE\_TEMPLATE (see [below for nested schema](#nestedatt--config--device_ssl_certificate_configs))
* `export_metadata_app_profile` (Attributes) (see [below for nested schema](#nestedatt--config--export_metadata_app_profile))
* `giga_port_neighbors_discovery_config` (Attributes) - To contain the port neighbour discovery template config (see [below for nested schema](#nestedatt--config--giga_port_neighbors_discovery_config))
* `giga_stream_threshold_config` (Attributes) - Gigastream threshold configs (see [below for nested schema](#nestedatt--config--giga_stream_threshold_config))
* `giga_user_defined_application_config` (Attributes) (see [below for nested schema](#nestedatt--config--giga_user_defined_application_config))
* `ldap_server_system_config` (Dynamic) - Available for ConfigType LDAP\_SYSTEM\_CONFIG\_TEMPLATE
* `ldap_servers` (Attributes List) - List of LDAP Servers. Available for ConfigType LDAP\_SERVERS\_TEMPLATE ChildType ldapServers (see [below for nested schema](#nestedatt--config--ldap_servers))
* `metadata_exporter` (Attributes) (see [below for nested schema](#nestedatt--config--metadata_exporter))
* `port_packet_threshold_config` (Attributes) - Contains port threshold configs (see [below for nested schema](#nestedatt--config--port_packet_threshold_config))
* `proxy_server_profile` (Attributes) (see [below for nested schema](#nestedatt--config--proxy_server_profile))
* `snmp_trap_event_configs` (Attributes List) - Available when ConfigType is SNMPTRAPS (see [below for nested schema](#nestedatt--config--snmp_trap_event_configs))
* `snmp_v3_users_config` (Attributes) - Contains list of FM's snmpv3 users (see [below for nested schema](#nestedatt--config--snmp_v3_users_config))
* `ssh_ciphers_config` (Attributes) - System SSH Cipher (see [below for nested schema](#nestedatt--config--ssh_ciphers_config))

<a id="nestedatt--config--aaa_auth_config"></a>
### Nested Schema for `config.aaa_auth_config`

Required:

* `auth_sequence` (Set of String) - Authentication methods order for user login. Valid values are 'local', 'ldap', 'radius', 'tacacs'

Optional:

* `external_login_mapping` (Attributes) - Specifies how externally authenticated logins (RADIUS, TACACS+, or LDAP) are mapped to local accounts. Private class (see [below for nested schema](#nestedatt--config--aaa_auth_config--external_login_mapping))

<a id="nestedatt--config--aaa_auth_config--external_login_mapping"></a>
### Nested Schema for `config.aaa_auth_config.external_login_mapping`

Optional:

* `default_local_user` (String) - Specifies the account to which externally authenticated logins are mapped when map order is set to remote-first (if there is no matching local account) or local-only
* `user_map_order` (String) - Specifies how externally authenticated logins (RADIUS, TACACS+, or LDAP) are mapped to local accounts

<a id="nestedatt--config--acme_certificate"></a>
### Nested Schema for `config.acme_certificate`

Required:

* `acme_server_url` (String)
* `operation_type` (String)

Optional:

* `algorithm` (String)
* `renew_days` (Number) - default will be 1/3rd of certificate validity period

<a id="nestedatt--config--device_ssl_certificate_configs"></a>
### Nested Schema for `config.device_ssl_certificate_configs`

Required:

* `operation_type` (String)
* `trusted_ca` (Attributes) (see [below for nested schema](#nestedatt--config--device_ssl_certificate_configs--trusted_ca))
* `upload_spec` (Attributes) (see [below for nested schema](#nestedatt--config--device_ssl_certificate_configs--upload_spec))

Optional:

* `issuer` (String) - issuer details of the certificate
* `not_after` (String) - date and time when certificate stops being valid (\[rfc3339\](https://xml2rfc.tools.ietf.org/public/rfc/html/rfc3339.html\#anchor14))
* `not_before` (String) - date and time when certificate starts being valid (\[rfc3339\](https://xml2rfc.tools.ietf.org/public/rfc/html/rfc3339.html\#anchor14))
* `signature_algorithm` (String)
* `subject` (String) - subject name of the certificate

<a id="nestedatt--config--device_ssl_certificate_configs--trusted_ca"></a>
### Nested Schema for `config.device_ssl_certificate_configs.trusted_ca`

Required:

* `name` (String) - name of the certificate

<a id="nestedatt--config--device_ssl_certificate_configs--upload_spec"></a>
### Nested Schema for `config.device_ssl_certificate_configs.upload_spec`

Optional:

* `info` (Attributes) - Certificate info (see [below for nested schema](#nestedatt--config--device_ssl_certificate_configs--upload_spec--info))
* `pem` (String) - contents of the certificate in pem format

<a id="nestedatt--config--device_ssl_certificate_configs--upload_spec--info"></a>
### Nested Schema for `config.device_ssl_certificate_configs.upload_spec.info`

Required:

* `name` (String) - name of the certificate
* `type` (String) - type of the certificate

Optional:

* `comment` (String) - a short description of the certificate
* `passphrase` (String) - used to decrypt pkcs12 and private keys

<a id="nestedatt--config--export_metadata_app_profile"></a>
### Nested Schema for `config.export_metadata_app_profile`

Required:

* `alias` (String) - application profile alias

Optional:

* `application_id` (Boolean) - only valid with 'export' type
* `applications` (Attributes List) - application and attributes. (see [below for nested schema](#nestedatt--config--export_metadata_app_profile--applications))
* `counter` (Attributes) (see [below for nested schema](#nestedatt--config--export_metadata_app_profile--counter))
* `datalink` (Attributes) (see [below for nested schema](#nestedatt--config--export_metadata_app_profile--datalink))
* `description` (String)
* `flow` (Attributes) (see [below for nested schema](#nestedatt--config--export_metadata_app_profile--flow))
* `gtpu` (Attributes) (see [below for nested schema](#nestedatt--config--export_metadata_app_profile--gtpu))
* `interface` (Attributes) (see [below for nested schema](#nestedatt--config--export_metadata_app_profile--interface))
* `ip` (Attributes) (see [below for nested schema](#nestedatt--config--export_metadata_app_profile--ip))
* `ipv4` (Attributes) (see [below for nested schema](#nestedatt--config--export_metadata_app_profile--ipv4))
* `ipv6` (Attributes) (see [below for nested schema](#nestedatt--config--export_metadata_app_profile--ipv6))
* `outer_ipv4` (Attributes) (see [below for nested schema](#nestedatt--config--export_metadata_app_profile--outer_ipv4))
* `outer_ipv6` (Attributes) (see [below for nested schema](#nestedatt--config--export_metadata_app_profile--outer_ipv6))
* `timestamp` (Attributes) (see [below for nested schema](#nestedatt--config--export_metadata_app_profile--timestamp))
* `transport` (Attributes) (see [below for nested schema](#nestedatt--config--export_metadata_app_profile--transport))
* `type` (String)

<a id="nestedatt--config--export_metadata_app_profile--applications"></a>
### Nested Schema for `config.export_metadata_app_profile.applications`

Required:

* `name` (String) - application name

Optional:

* `attributes` (Attributes List) (see [below for nested schema](#nestedatt--config--export_metadata_app_profile--applications--attributes))
* `is_user_defined` (Boolean) - Default value is false. Must be set to true only for App Intel solution while configuring user defined apps

<a id="nestedatt--config--export_metadata_app_profile--applications--attributes"></a>
### Nested Schema for `config.export_metadata_app_profile.applications.attributes`

Required:

* `name` (String) - attribute name

Optional:

* `value` (String) - application's attribute value

<a id="nestedatt--config--export_metadata_app_profile--counter"></a>
### Nested Schema for `config.export_metadata_app_profile.counter`

Optional:

* `bytes` (Boolean)
* `bytes_long` (Boolean)
* `inner_byte` (Boolean)
* `inner_byte_long` (Boolean)
* `packets` (Boolean)
* `packets_long` (Boolean)

<a id="nestedatt--config--export_metadata_app_profile--datalink"></a>
### Nested Schema for `config.export_metadata_app_profile.datalink`

Optional:

* `mac_dst` (Boolean)
* `mac_src` (Boolean)
* `vlan` (Boolean)

<a id="nestedatt--config--export_metadata_app_profile--flow"></a>
### Nested Schema for `config.export_metadata_app_profile.flow`

Optional:

* `end_reason` (Boolean)

<a id="nestedatt--config--export_metadata_app_profile--gtpu"></a>
### Nested Schema for `config.export_metadata_app_profile.gtpu`

Optional:

* `qfi` (Boolean)
* `teid` (Boolean)

<a id="nestedatt--config--export_metadata_app_profile--interface"></a>
### Nested Schema for `config.export_metadata_app_profile.interface`

Optional:

* `in_name_width` (Number)
* `in_physical_width` (Number)
* `out_physical_width` (Number)

<a id="nestedatt--config--export_metadata_app_profile--ip"></a>
### Nested Schema for `config.export_metadata_app_profile.ip`

Optional:

* `version` (Boolean)

<a id="nestedatt--config--export_metadata_app_profile--ipv4"></a>
### Nested Schema for `config.export_metadata_app_profile.ipv4`

Optional:

* `destination` (Attributes) (see [below for nested schema](#nestedatt--config--export_metadata_app_profile--ipv4--destination))
* `dscp` (Boolean)
* `fragmentation` (Attributes) (see [below for nested schema](#nestedatt--config--export_metadata_app_profile--ipv4--fragmentation))
* `header_len` (Boolean)
* `option_map` (Boolean)
* `precedence` (Boolean)
* `protocol` (Boolean)
* `section` (Attributes) (see [below for nested schema](#nestedatt--config--export_metadata_app_profile--ipv4--section))
* `source` (Attributes) (see [below for nested schema](#nestedatt--config--export_metadata_app_profile--ipv4--source))
* `tos` (Boolean)
* `total_length` (Boolean)
* `ttl` (Boolean)

<a id="nestedatt--config--export_metadata_app_profile--ipv4--destination"></a>
### Nested Schema for `config.export_metadata_app_profile.ipv4.destination`

Optional:

* `prefix_min_mask` (String)

<a id="nestedatt--config--export_metadata_app_profile--ipv4--fragmentation"></a>
### Nested Schema for `config.export_metadata_app_profile.ipv4.fragmentation`

Optional:

* `flags` (Boolean)
* `offset` (Boolean)

<a id="nestedatt--config--export_metadata_app_profile--ipv4--section"></a>
### Nested Schema for `config.export_metadata_app_profile.ipv4.section`

Optional:

* `header_size` (Number)
* `payload_size` (Number)

<a id="nestedatt--config--export_metadata_app_profile--ipv4--source"></a>
### Nested Schema for `config.export_metadata_app_profile.ipv4.source`

Optional:

* `prefix_min_mask` (String) - ipv4 source prefix minimum-mask - netmask or mask length

<a id="nestedatt--config--export_metadata_app_profile--ipv6"></a>
### Nested Schema for `config.export_metadata_app_profile.ipv6`

Optional:

* `destination` (Attributes) (see [below for nested schema](#nestedatt--config--export_metadata_app_profile--ipv6--destination))
* `dscp` (Boolean)
* `extension_map` (Boolean)
* `flow_label` (Boolean)
* `fragmentation` (Attributes) (see [below for nested schema](#nestedatt--config--export_metadata_app_profile--ipv6--fragmentation))
* `hop_limit` (Boolean)
* `length` (Attributes) (see [below for nested schema](#nestedatt--config--export_metadata_app_profile--ipv6--length))
* `next_header` (Boolean)
* `precedence` (Boolean)
* `section` (Attributes) (see [below for nested schema](#nestedatt--config--export_metadata_app_profile--ipv6--section))
* `source` (Attributes) (see [below for nested schema](#nestedatt--config--export_metadata_app_profile--ipv6--source))
* `traffic_class` (Boolean)

<a id="nestedatt--config--export_metadata_app_profile--ipv6--destination"></a>
### Nested Schema for `config.export_metadata_app_profile.ipv6.destination`

Optional:

* `prefix_min_mask` (String)

<a id="nestedatt--config--export_metadata_app_profile--ipv6--fragmentation"></a>
### Nested Schema for `config.export_metadata_app_profile.ipv6.fragmentation`

Optional:

* `flags` (Boolean)
* `offset` (Boolean)

<a id="nestedatt--config--export_metadata_app_profile--ipv6--length"></a>
### Nested Schema for `config.export_metadata_app_profile.ipv6.length`

Optional:

* `header` (Boolean)
* `payload` (Boolean)
* `total` (Boolean)

<a id="nestedatt--config--export_metadata_app_profile--ipv6--section"></a>
### Nested Schema for `config.export_metadata_app_profile.ipv6.section`

Optional:

* `header_size` (Number)
* `payload_size` (Number)

<a id="nestedatt--config--export_metadata_app_profile--ipv6--source"></a>
### Nested Schema for `config.export_metadata_app_profile.ipv6.source`

Optional:

* `prefix_min_mask` (String)

<a id="nestedatt--config--export_metadata_app_profile--outer_ipv4"></a>
### Nested Schema for `config.export_metadata_app_profile.outer_ipv4`

Optional:

* `destination` (Boolean)
* `source` (Boolean)

<a id="nestedatt--config--export_metadata_app_profile--outer_ipv6"></a>
### Nested Schema for `config.export_metadata_app_profile.outer_ipv6`

Optional:

* `destination` (Boolean)
* `source` (Boolean)

<a id="nestedatt--config--export_metadata_app_profile--timestamp"></a>
### Nested Schema for `config.export_metadata_app_profile.timestamp`

Optional:

* `flow_end_msec` (Boolean)
* `flow_endsec` (Boolean)
* `flow_start_msec` (Boolean)
* `flow_startsec` (Boolean)
* `sys_up_time_first` (Boolean)
* `sys_up_time_last` (Boolean)

<a id="nestedatt--config--export_metadata_app_profile--transport"></a>
### Nested Schema for `config.export_metadata_app_profile.transport`

Optional:

* `dst_port` (Boolean)
* `icmp` (Attributes) (see [below for nested schema](#nestedatt--config--export_metadata_app_profile--transport--icmp))
* `src_port` (Boolean)
* `tcp` (Attributes) (see [below for nested schema](#nestedatt--config--export_metadata_app_profile--transport--tcp))
* `udp` (Attributes) (see [below for nested schema](#nestedatt--config--export_metadata_app_profile--transport--udp))

<a id="nestedatt--config--export_metadata_app_profile--transport--icmp"></a>
### Nested Schema for `config.export_metadata_app_profile.transport.icmp`

Optional:

* `ipv4_code` (Boolean)
* `ipv4_type` (Boolean)
* `ipv6_code` (Boolean)
* `ipv6_type` (Boolean)

<a id="nestedatt--config--export_metadata_app_profile--transport--tcp"></a>
### Nested Schema for `config.export_metadata_app_profile.transport.tcp`

Optional:

* `ack_number` (Boolean)
* `dst_port` (Boolean)
* `flags` (Boolean)
* `header_len` (Boolean)
* `seq_number` (Boolean)
* `src_port` (Boolean)
* `urgent_ptr` (Boolean)
* `window_size` (Boolean)

<a id="nestedatt--config--export_metadata_app_profile--transport--udp"></a>
### Nested Schema for `config.export_metadata_app_profile.transport.udp`

Optional:

* `dst_port` (Boolean)
* `msg_len` (Boolean)
* `src_port` (Boolean)

<a id="nestedatt--config--giga_port_neighbors_discovery_config"></a>
### Nested Schema for `config.giga_port_neighbors_discovery_config`

Optional:

* `resource_configs` (Attributes List) - List of config for each port type (see [below for nested schema](#nestedatt--config--giga_port_neighbors_discovery_config--resource_configs))

<a id="nestedatt--config--giga_port_neighbors_discovery_config--resource_configs"></a>
### Nested Schema for `config.giga_port_neighbors_discovery_config.resource_configs`

Optional:

* `cdp` (Boolean)
* `gdp` (Boolean)
* `lldp` (Boolean)
* `port_type` (Dynamic) - Type of the port

<a id="nestedatt--config--giga_stream_threshold_config"></a>
### Nested Schema for `config.giga_stream_threshold_config`

Optional:

* `giga_stream_type_thresholds` (Attributes List) - GigaStream threshold config for port types (see [below for nested schema](#nestedatt--config--giga_stream_threshold_config--giga_stream_type_thresholds))

<a id="nestedatt--config--giga_stream_threshold_config--giga_stream_type_thresholds"></a>
### Nested Schema for `config.giga_stream_threshold_config.giga_stream_type_thresholds`

Optional:

* `type` (String) - Port type
* `variance_threshold` (Number) - Threshold value in percentage. '-1' will disable threshold feature.

<a id="nestedatt--config--giga_user_defined_application_config"></a>
### Nested Schema for `config.giga_user_defined_application_config`

Optional:

* `alias` (String) - pattern: ^\[a-zA-Z0-9&'+\_=\|\-\]+$
* `app_id` (Number) - minimum: 16777216 maximum: 16777343   Ranges from 2^24 to 2^24 + 127. It needs to be unique
* `priority` (Number) - minimum: 1 maximum: 120  Lower the priority higher the precedence
* `rules` (Attributes) (see [below for nested schema](#nestedatt--config--giga_user_defined_application_config--rules))

<a id="nestedatt--config--giga_user_defined_application_config--rules"></a>
### Nested Schema for `config.giga_user_defined_application_config.rules`

Optional:

* `rule` (Attributes List) (see [below for nested schema](#nestedatt--config--giga_user_defined_application_config--rules--rule))

<a id="nestedatt--config--giga_user_defined_application_config--rules--rule"></a>
### Nested Schema for `config.giga_user_defined_application_config.rules.rule`

Optional:

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

<a id="nestedatt--config--ldap_servers"></a>
### Nested Schema for `config.ldap_servers`

Required:

* `server_address` (String) - ipv4 or ipv6 or hostname. Specifies address of the LDAP server where authentication requests will be sent

Optional:

* `order` (String) - The order in which the server is to be reached. 1 means server will be contacted first

<a id="nestedatt--config--metadata_exporter"></a>
### Nested Schema for `config.metadata_exporter`

Required:

* `alias` (String)

Optional:

* `application_profiles` (List of String) - application profile aliases to attach to the exporter
* `cef` (Attributes) - cef attributes (see [below for nested schema](#nestedatt--config--metadata_exporter--cef))
* `description` (String)
* `destination` (Attributes) - destination attributes (see [below for nested schema](#nestedatt--config--metadata_exporter--destination))
* `max_pkt_size` (Number)
* `mobility_sam` (Attributes) (see [below for nested schema](#nestedatt--config--metadata_exporter--mobility_sam))
* `monitor` (Attributes) - monitor attributes (see [below for nested schema](#nestedatt--config--metadata_exporter--monitor))
* `netflow` (Attributes) - netflow attributes (see [below for nested schema](#nestedatt--config--metadata_exporter--netflow))
* `snmp` (Attributes) (see [below for nested schema](#nestedatt--config--metadata_exporter--snmp))
* `source` (Attributes) - source tunnel port (see [below for nested schema](#nestedatt--config--metadata_exporter--source))
* `type` (String)

<a id="nestedatt--config--metadata_exporter--cef"></a>
### Nested Schema for `config.metadata_exporter.cef`

Optional:

* `active_timeout` (Number) - in seconds
* `inactive_timeout` (Number) - in seconds

<a id="nestedatt--config--metadata_exporter--destination"></a>
### Nested Schema for `config.metadata_exporter.destination`

Optional:

* `dscp` (Number)
* `ipv4_address` (String) - ipv4 address
* `l4_port_dst` (Number)
* `l4_port_src` (Number)
* `l4_protocol` (String)
* `ttl` (Number)

<a id="nestedatt--config--metadata_exporter--mobility_sam"></a>
### Nested Schema for `config.metadata_exporter.mobility_sam`

Optional:

* `encoding` (String)
* `encoding_format` (String)
* `event_enable` (Attributes) (see [below for nested schema](#nestedatt--config--metadata_exporter--mobility_sam--event_enable))
* `trigger` (String)

<a id="nestedatt--config--metadata_exporter--mobility_sam--event_enable"></a>
### Nested Schema for `config.metadata_exporter.mobility_sam.event_enable`

Optional:

* `modify` (Boolean)
* `update` (Boolean)

<a id="nestedatt--config--metadata_exporter--monitor"></a>
### Nested Schema for `config.metadata_exporter.monitor`

Optional:

* `timeout` (Number) - how often to export in seconds

<a id="nestedatt--config--metadata_exporter--netflow"></a>
### Nested Schema for `config.metadata_exporter.netflow`

Optional:

* `active_timeout` (Number) - in seconds
* `inactive_timeout` (Number) - in seconds
* `template_refresh` (Number) - template refresh interval in seconds
* `template_type` (String)
* `version` (String)

<a id="nestedatt--config--metadata_exporter--snmp"></a>
### Nested Schema for `config.metadata_exporter.snmp`

Optional:

* `enabled` (Boolean) - snmp reverse lookup enable/disable

<a id="nestedatt--config--metadata_exporter--source"></a>
### Nested Schema for `config.metadata_exporter.source`

Optional:

* `ip_interface` (String)

<a id="nestedatt--config--port_packet_threshold_config"></a>
### Nested Schema for `config.port_packet_threshold_config`

Optional:

* `drop_threshold` (Attributes) - Drop threshold config for Rx and Tx (see [below for nested schema](#nestedatt--config--port_packet_threshold_config--drop_threshold))
* `error_threshold` (Attributes) - Error threshold config for Rx and Tx (see [below for nested schema](#nestedatt--config--port_packet_threshold_config--error_threshold))

<a id="nestedatt--config--port_packet_threshold_config--drop_threshold"></a>
### Nested Schema for `config.port_packet_threshold_config.drop_threshold`

Optional:

* `rx` (Attributes List) - Threshold values for Rx packet drop (see [below for nested schema](#nestedatt--config--port_packet_threshold_config--drop_threshold--rx))
* `tx` (Attributes List) - Threshold values for Tx packet drop (see [below for nested schema](#nestedatt--config--port_packet_threshold_config--drop_threshold--tx))

<a id="nestedatt--config--port_packet_threshold_config--drop_threshold--rx"></a>
### Nested Schema for `config.port_packet_threshold_config.drop_threshold.rx`

Optional:

* `count_` (Number) - Threshold value in number of packets out of total packets in a given interval. '-1' will disable threshold feature.
* `percent` (Number) - Threshold value in percent of packets out of total packets in a given interval. '-1' will disable threshold feature.
* `port_type` (String) - Port type

<a id="nestedatt--config--port_packet_threshold_config--drop_threshold--tx"></a>
### Nested Schema for `config.port_packet_threshold_config.drop_threshold.tx`

Optional:

* `count_` (Number) - Threshold value in number of packets out of total packets in a given interval. '-1' will disable threshold feature.
* `percent` (Number) - Threshold value in percent of packets out of total packets in a given interval. '-1' will disable threshold feature.
* `port_type` (String) - Port type

<a id="nestedatt--config--port_packet_threshold_config--error_threshold"></a>
### Nested Schema for `config.port_packet_threshold_config.error_threshold`

Optional:

* `rx` (Attributes List) - Threshold values for Rx packet error (see [below for nested schema](#nestedatt--config--port_packet_threshold_config--error_threshold--rx))
* `tx` (Attributes List) - Threshold values for Tx packet error (see [below for nested schema](#nestedatt--config--port_packet_threshold_config--error_threshold--tx))

<a id="nestedatt--config--port_packet_threshold_config--error_threshold--rx"></a>
### Nested Schema for `config.port_packet_threshold_config.error_threshold.rx`

Optional:

* `count_` (Number) - Threshold value in number of packets out of total packets in a given interval. '-1' will disable threshold feature.
* `percent` (Number) - Threshold value in percent of packets out of total packets in a given interval. '-1' will disable threshold feature.
* `port_type` (String) - Port type

<a id="nestedatt--config--port_packet_threshold_config--error_threshold--tx"></a>
### Nested Schema for `config.port_packet_threshold_config.error_threshold.tx`

Optional:

* `count_` (Number) - Threshold value in number of packets out of total packets in a given interval. '-1' will disable threshold feature.
* `percent` (Number) - Threshold value in percent of packets out of total packets in a given interval. '-1' will disable threshold feature.
* `port_type` (String) - Port type

<a id="nestedatt--config--proxy_server_profile"></a>
### Nested Schema for `config.proxy_server_profile`

Required:

* `alias` (String)
* `auth_type` (String)
* `periodic_ping` (String)
* `port` (Number)
* `protocol` (String)
* `proxy_address` (String)

Optional:

* `comment` (String)
* `password` (String)
* `periodic_ping_failure_retry` (Number)
* `periodic_ping_interval` (Number)
* `periodic_ping_type` (String)
* `ssl_apps` (Attributes) (see [below for nested schema](#nestedatt--config--proxy_server_profile--ssl_apps))
* `username` (String)

<a id="nestedatt--config--proxy_server_profile--ssl_apps"></a>
### Nested Schema for `config.proxy_server_profile.ssl_apps`

Optional:

* `cluster_name` (List of String) - Cluster Name where the proxy deployed

<a id="nestedatt--config--snmp_trap_event_configs"></a>
### Nested Schema for `config.snmp_trap_event_configs`

Optional:

* `enabled` (Boolean)
* `notify_event` (String)

<a id="nestedatt--config--snmp_v3_users_config"></a>
### Nested Schema for `config.snmp_v3_users_config`

Required:

* `snmp_v3_user` (Attributes List) (see [below for nested schema](#nestedatt--config--snmp_v3_users_config--snmp_v3_user))

<a id="nestedatt--config--snmp_v3_users_config--snmp_v3_user"></a>
### Nested Schema for `config.snmp_v3_users_config.snmp_v3_user`

Required:

* `auth_key` (String) - auth key/passphrase
* `auth_protocol` (String)
* `priv_key` (String) - privacy key/passphrase
* `priv_protocol` (String)
* `username` (String) - snmpv3 username

Optional:

* `min_sw_version` (String) - minimum software version of the device
* `previous_username` (String) - When the username is changed, the username prior changing is sent in this property. Based on this application handles the username update logic. Not found error is returned when this property is not sent during the username change

<a id="nestedatt--config--ssh_ciphers_config"></a>
### Nested Schema for `config.ssh_ciphers_config`

Optional:

* `classic` (Attributes) - System SSH Cipher (see [below for nested schema](#nestedatt--config--ssh_ciphers_config--classic))
* `crypto` (Attributes) - System SSH Cipher (see [below for nested schema](#nestedatt--config--ssh_ciphers_config--crypto))
* `fips` (Attributes) - System SSH Cipher (see [below for nested schema](#nestedatt--config--ssh_ciphers_config--fips))

<a id="nestedatt--config--ssh_ciphers_config--classic"></a>
### Nested Schema for `config.ssh_ciphers_config.classic`

Optional:

* `client_ciphers` (List of String)
* `client_hostkey` (List of String)
* `client_kex` (List of String)
* `client_macs` (List of String)
* `server_ciphers` (List of String)
* `server_hostkey` (List of String)
* `server_kex` (List of String)
* `server_macs` (List of String)

<a id="nestedatt--config--ssh_ciphers_config--crypto"></a>
### Nested Schema for `config.ssh_ciphers_config.crypto`

Optional:

* `client_ciphers` (List of String)
* `client_hostkey` (List of String)
* `client_kex` (List of String)
* `client_macs` (List of String)
* `server_ciphers` (List of String)
* `server_hostkey` (List of String)
* `server_kex` (List of String)
* `server_macs` (List of String)

<a id="nestedatt--config--ssh_ciphers_config--fips"></a>
### Nested Schema for `config.ssh_ciphers_config.fips`

Optional:

* `client_ciphers` (List of String)
* `client_hostkey` (List of String)
* `client_kex` (List of String)
* `client_macs` (List of String)
* `server_ciphers` (List of String)
* `server_hostkey` (List of String)
* `server_kex` (List of String)
* `server_macs` (List of String)

<a id="nestedatt--ref_object"></a>
### Nested Schema for `ref_object`

Optional:

* `ref_object_type` (String) - Type of the reference object

