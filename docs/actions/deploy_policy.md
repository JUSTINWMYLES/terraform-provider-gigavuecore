---
page_title: "gigavuecore_deploy_policy Action - gigavuecore"
subcategory: ""
description: |-
  Deploy a new policy
---

# gigavuecore_deploy_policy Action

Deploy a new policy

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_deploy_policy" "example" {
  config {
    comment             = "example"
    deployed            = true
    deployment_error    = "example"
    deployment_percent  = "example"
    dest_port_timestamp = "example"
    health_state        = "green"
    health_state_reasons = [{
      message                               = "example"
      severity                              = "green"
      traffic_health_state_computation_type = "PORT_LOW_UTIL"
    }]
    name             = "example"
    policy_id        = "example"
    policy_timestamp = "example"
    priority         = true
    rules = [{
      comment = "example"
      gs_operations = {
        gs_apps = {
          apf = {
            enabled = "enabled"
          }
          dedup = {
            enabled = "enabled"
          }
          diameter_whitelist = {
            enabled = "enabled"
          }
          flow_filter = {
            type = "gtp"
          }
          flow_sampling = {
            type = "ip"
          }
          gseries_header_add = {
            types = ["srcid"]
          }
          gseries_header_remove = {
            enabled = "enabled"
          }
          gseries_load_balance = {
            fixed_offset = {
              hash   = "checksum"
              length = 1
              offset = 0
            }
            variable_offset = {
              end_delim   = "example"
              hash        = "checksum"
              start_delim = "example"
              start_field = "example"
            }
          }
          gseries_pattern_match = {
            fixed_offset = {
              length = 1
              offset = 0
            }
            variable_offset = {
              end_delim   = "example"
              start_delim = "example"
            }
          }
          gtp_whitelist = {
            enabled = "enabled"
          }
          header_add = {
            vlan = 0
          }
          header_remove = {
            ah1                = "none"
            ah2                = "none"
            custom_len         = 1
            erspan_flow_id     = 0
            fp_dst_switch_id   = 0
            fp_src_switch_id   = 0
            header_count       = 1
            offset             = "start"
            offset_range_value = 0
            protocol           = "gtp"
            timestamp_format   = "gigasmart"
            vlan_header        = "all"
            vxlan_id           = 0
          }
          icap = {
            icap_profile = "example"
          }
          inline_ssl = {
            inline_ssl_profile = "example"
          }
          load_balance = {
            enhanced = {
              elb_alias = "example"
            }
            stateful = {
              app_type               = "gtp"
              diameter_key_hash_type = "sessionId"
              diameter_key_multi_hash_type = [{
                avp_codevalue = 0
                key           = "sessionId"
              }]
              gtp_key_hash_type = "imsi"
              lb_type           = "leastBw"
              sip_key_hash_type = "callerId"
            }
            stateless = {
              field_location = "inner"
              hash_fields    = "ipOnly"
            }
          }
          masking = {
            content_type = "message_cpim"
            length       = 1
            offset       = 0
            pattern      = "a1"
            protocol     = "none"
          }
          metadata_export = {
            cache = "example"
          }
          netflow = {
            enabled = "enabled"
          }
          sa_apf = {
            enabled = "enabled"
          }
          sip_whitelist = {
            enabled = "enabled"
          }
          slicing = {
            enhanced = "example"
            offset   = 4
            protocol = "none"
          }
          ssl_decrypt = {
            in_port  = 0
            out_port = 0
          }
          trailer_add = {
            types = ["crc"]
          }
          trailer_remove = {
            enabled = "enabled"
          }
          tunnel_decap = {
            custom = {
              port_dst = 0
              port_src = 0
            }
            erspan_flow_id = 0
            gmip_port      = 0
            l2_gre_key     = 0
            tls_pcapng = {
              decap_key = "example"
              listener  = "example"
            }
            type = "gmip"
            vxlan = {
              port_dst = 1
              port_src = 0
              vni      = 0
            }
          }
          tunnel_encap = {
            gmip_config = {
              dscp       = 0
              dst_ip     = "example"
              dst_port   = 0
              flow_label = 0
              prec       = 0
              src_port   = 0
              ttl        = 1
            }
            l2_gre_config = {
              dscp          = 0
              dst_ip        = "example"
              flow_label    = 0
              key           = 0
              pg_dst        = "example"
              prec          = 0
              session_field = "fiveTupleIpv4"
              session_pos   = "inner"
              ttl           = 1
            }
            tls_pcapng = {
              exporter       = "example"
              exporter_group = "example"
            }
            type = "gmip"
            vxlan_config = {
              dscp     = 0
              dst_ip   = "example"
              dst_port = 4789
              src_port = 0
              ttl      = 1
              vni      = 1
            }
          }
        }
        gs_engines = ["example"]
        gs_group_params = {
          app_tcp = {
            application  = "broadcast"
            load_balance = true
            tcp_control  = "broadcast"
          }
          dedup = {
            action    = "count"
            ip_tclass = "include"
            ip_tos    = "include"
            tcp_seq   = "include"
            timer     = 10
            vlan      = "include"
          }
          diameter_packet = {
            timeout = 1
          }
          diameter_s6_a_session = {
            limit   = 1
            timeout = 30
          }
          diameter_whitelist = {
            whitelist = "example"
          }
          eflow = {
            enabled      = true
            interval     = 0
            log_enabled  = true
            packet_count = 0
            packet_ratio = 0
          }
          engine_watchdog_timer = {
            time = 0
          }
          erspan3 = {
            timestamp_format = "gigasmart"
          }
          flow_mask = {
            enabled = true
            length  = 1
            offset  = 0
          }
          flow_sampling = {
            ip_ranges = ["example"]
            rate      = 5
            timeout   = 1
            type      = "deviceIp"
          }
          generic_session_timeout = {
            time = 5
          }
          gpfcp_profiles = {
            g_pfcp_profiles = ["example"]
          }
          gs_group_system = {
            cpu_load_alarm_threshold = 20
          }
          gta_profiles = {
            gta_profiles = ["example"]
          }
          gtp_control_sampling = {
            enabled = true
          }
          gtp_flow = {
            timeout = 1
          }
          gtp_gpfcp_delay = {
            timeout = 0
          }
          gtp_persistence = {
            enabled          = true
            file_age_timeout = 10
            interval         = 10
            restart_age_time = 10
          }
          gtp_random_sampling = {
            enabled  = true
            interval = 12
          }
          gtp_whitelist = {
            multi_whitelists = ["example"]
            whitelist        = "example"
          }
          health_check = {
            action          = "pass"
            dst_port        = 1
            enabled         = true
            interval        = 5
            protocol        = "icmp"
            rcv_port        = 1
            retries         = 1
            round_trip_time = 1
            src_port        = 1
          }
          hsm_group = {
            hsm_group = "example"
          }
          ip_frag = {
            forward              = true
            head_session_timeout = 15
            timeout              = 5
          }
          load_balance = {
            failover = {
              enabled               = true
              threshold_lt_bw       = 50
              threshold_lt_pkt_rate = 500
            }
            link_weight_type = "speed"
            replicate_gtpc   = true
          }
          netflow = {
            monitor = "example"
          }
          node_role = {
            mob5_g_limit     = 1
            mob_lte_limit    = 1
            stand_alone_mode = true
            type             = "control"
          }
          port_throttle_sip = {
            port_throttle = "example"
          }
          resource = {
            buffer_asf_size = 0
            cpu = {
              overload_threshold = 0
            }
            hsm_ssl = {
              buffer        = 0
              packet_buffer = 20
              session_count = 0
            }
            inline_ssl = {
              standalone = true
            }
            metadata = 0
            packet_buffer = {
              overload_threshold = 0
            }
            session_overload = {
              overload_threshold = 0
            }
            tunnel_overload = {
              overload_threshold = 0
            }
            xpkt_match = {
              flows = 0
            }
          }
          rtp_ports = {
            range = {
              port     = 0
              port_max = 0
            }
          }
          sa_apf = {
            buffer_size = 0
          }
          session_logging = {
            interface          = "example"
            log_level          = "err"
            remote_syslog_ip   = "example"
            remote_syslog_port = 0
          }
          sffp_profiles = {
            sffp_profiles = ["example"]
          }
          sip_media = {
            timeout = 30
          }
          sip_ports = {
            ports = [1]
          }
          sip_session = {
            timeout = 30
          }
          sip_tcp_idle_timeout = {
            time = 20
          }
          sip_whitelist = {
            whitelist = "example"
          }
          ssl_decrypt = {
            decrypt_fail_action = "drop"
            enabled             = true
            hsm_pkcs11 = {
              debug_level    = 0
              dynamic_object = true
              load_sharing   = true
            }
            hsm_timeout             = 2
            key_cache_timeout       = 1
            key_map                 = "example"
            non_ssl_traffic         = "drop"
            pending_session_timeout = 30
            session_timeout         = 30
            tcp_syn_timeout         = 20
            ticket_cache_timeout    = 1
          }
          xpkt_match = {
            enabled = true
          }
        }
        templates = [{
          gs_apps = {
            apf = {
              enabled = "enabled"
            }
            dedup = {
              enabled = "enabled"
            }
            diameter_whitelist = {
              enabled = "enabled"
            }
            flow_filter = {
              type = "gtp"
            }
            flow_sampling = {
              type = "ip"
            }
            gseries_header_add = {
              types = ["srcid"]
            }
            gseries_header_remove = {
              enabled = "enabled"
            }
            gseries_load_balance = {
              fixed_offset = {
                hash   = "checksum"
                length = 1
                offset = 0
              }
              variable_offset = {
                end_delim   = "example"
                hash        = "checksum"
                start_delim = "example"
                start_field = "example"
              }
            }
            gseries_pattern_match = {
              fixed_offset = {
                length = 1
                offset = 0
              }
              variable_offset = {
                end_delim   = "example"
                start_delim = "example"
              }
            }
            gtp_whitelist = {
              enabled = "enabled"
            }
            header_add = {
              vlan = 0
            }
            header_remove = {
              ah1                = "none"
              ah2                = "none"
              custom_len         = 1
              erspan_flow_id     = 0
              fp_dst_switch_id   = 0
              fp_src_switch_id   = 0
              header_count       = 1
              offset             = "start"
              offset_range_value = 0
              protocol           = "gtp"
              timestamp_format   = "gigasmart"
              vlan_header        = "all"
              vxlan_id           = 0
            }
            icap = {
              icap_profile = "example"
            }
            inline_ssl = {
              inline_ssl_profile = "example"
            }
            load_balance = {
              enhanced = {
                elb_alias = "example"
              }
              stateful = {
                app_type               = "gtp"
                diameter_key_hash_type = "sessionId"
                diameter_key_multi_hash_type = [{
                  avp_codevalue = 0
                  key           = "sessionId"
                }]
                gtp_key_hash_type = "imsi"
                lb_type           = "leastBw"
                sip_key_hash_type = "callerId"
              }
              stateless = {
                field_location = "inner"
                hash_fields    = "ipOnly"
              }
            }
            masking = {
              content_type = "message_cpim"
              length       = 1
              offset       = 0
              pattern      = "a1"
              protocol     = "none"
            }
            metadata_export = {
              cache = "example"
            }
            netflow = {
              enabled = "enabled"
            }
            sa_apf = {
              enabled = "enabled"
            }
            sip_whitelist = {
              enabled = "enabled"
            }
            slicing = {
              enhanced = "example"
              offset   = 4
              protocol = "none"
            }
            ssl_decrypt = {
              in_port  = 0
              out_port = 0
            }
            trailer_add = {
              types = ["crc"]
            }
            trailer_remove = {
              enabled = "enabled"
            }
            tunnel_decap = {
              custom = {
                port_dst = 0
                port_src = 0
              }
              erspan_flow_id = 0
              gmip_port      = 0
              l2_gre_key     = 0
              tls_pcapng = {
                decap_key = "example"
                listener  = "example"
              }
              type = "gmip"
              vxlan = {
                port_dst = 1
                port_src = 0
                vni      = 0
              }
            }
            tunnel_encap = {
              gmip_config = {
                dscp       = 0
                dst_ip     = "example"
                dst_port   = 0
                flow_label = 0
                prec       = 0
                src_port   = 0
                ttl        = 1
              }
              l2_gre_config = {
                dscp          = 0
                dst_ip        = "example"
                flow_label    = 0
                key           = 0
                pg_dst        = "example"
                prec          = 0
                session_field = "fiveTupleIpv4"
                session_pos   = "inner"
                ttl           = 1
              }
              tls_pcapng = {
                exporter       = "example"
                exporter_group = "example"
              }
              type = "gmip"
              vxlan_config = {
                dscp     = 0
                dst_ip   = "example"
                dst_port = 4789
                src_port = 0
                ttl      = 1
                vni      = 1
              }
            }
          }
          gs_group_params = {
            app_tcp = {
              application  = "broadcast"
              load_balance = true
              tcp_control  = "broadcast"
            }
            dedup = {
              action    = "count"
              ip_tclass = "include"
              ip_tos    = "include"
              tcp_seq   = "include"
              timer     = 10
              vlan      = "include"
            }
            diameter_packet = {
              timeout = 1
            }
            diameter_s6_a_session = {
              limit   = 1
              timeout = 30
            }
            diameter_whitelist = {
              whitelist = "example"
            }
            eflow = {
              enabled      = true
              interval     = 0
              log_enabled  = true
              packet_count = 0
              packet_ratio = 0
            }
            engine_watchdog_timer = {
              time = 0
            }
            erspan3 = {
              timestamp_format = "gigasmart"
            }
            flow_mask = {
              enabled = true
              length  = 1
              offset  = 0
            }
            flow_sampling = {
              ip_ranges = ["example"]
              rate      = 5
              timeout   = 1
              type      = "deviceIp"
            }
            generic_session_timeout = {
              time = 5
            }
            gpfcp_profiles = {
              g_pfcp_profiles = ["example"]
            }
            gs_group_system = {
              cpu_load_alarm_threshold = 20
            }
            gta_profiles = {
              gta_profiles = ["example"]
            }
            gtp_control_sampling = {
              enabled = true
            }
            gtp_flow = {
              timeout = 1
            }
            gtp_gpfcp_delay = {
              timeout = 0
            }
            gtp_persistence = {
              enabled          = true
              file_age_timeout = 10
              interval         = 10
              restart_age_time = 10
            }
            gtp_random_sampling = {
              enabled  = true
              interval = 12
            }
            gtp_whitelist = {
              multi_whitelists = ["example"]
              whitelist        = "example"
            }
            health_check = {
              action          = "pass"
              dst_port        = 1
              enabled         = true
              interval        = 5
              protocol        = "icmp"
              rcv_port        = 1
              retries         = 1
              round_trip_time = 1
              src_port        = 1
            }
            hsm_group = {
              hsm_group = "example"
            }
            ip_frag = {
              forward              = true
              head_session_timeout = 15
              timeout              = 5
            }
            load_balance = {
              failover = {
                enabled               = true
                threshold_lt_bw       = 50
                threshold_lt_pkt_rate = 500
              }
              link_weight_type = "speed"
              replicate_gtpc   = true
            }
            netflow = {
              monitor = "example"
            }
            node_role = {
              mob5_g_limit     = 1
              mob_lte_limit    = 1
              stand_alone_mode = true
              type             = "control"
            }
            port_throttle_sip = {
              port_throttle = "example"
            }
            resource = {
              buffer_asf_size = 0
              cpu = {
                overload_threshold = 0
              }
              hsm_ssl = {
                buffer        = 0
                packet_buffer = 20
                session_count = 0
              }
              inline_ssl = {
                standalone = true
              }
              metadata = 0
              packet_buffer = {
                overload_threshold = 0
              }
              session_overload = {
                overload_threshold = 0
              }
              tunnel_overload = {
                overload_threshold = 0
              }
              xpkt_match = {
                flows = 0
              }
            }
            rtp_ports = {
              range = {
                port     = 0
                port_max = 0
              }
            }
            sa_apf = {
              buffer_size = 0
            }
            session_logging = {
              interface          = "example"
              log_level          = "err"
              remote_syslog_ip   = "example"
              remote_syslog_port = 0
            }
            sffp_profiles = {
              sffp_profiles = ["example"]
            }
            sip_media = {
              timeout = 30
            }
            sip_ports = {
              ports = [1]
            }
            sip_session = {
              timeout = 30
            }
            sip_tcp_idle_timeout = {
              time = 20
            }
            sip_whitelist = {
              whitelist = "example"
            }
            ssl_decrypt = {
              decrypt_fail_action = "drop"
              enabled             = true
              hsm_pkcs11 = {
                debug_level    = 0
                dynamic_object = true
                load_sharing   = true
              }
              hsm_timeout             = 2
              key_cache_timeout       = 1
              key_map                 = "example"
              non_ssl_traffic         = "drop"
              pending_session_timeout = 30
              session_timeout         = 30
              tcp_syn_timeout         = 20
              ticket_cache_timeout    = 1
            }
            xpkt_match = {
              enabled = true
            }
          }
          template_id = "example"
        }]
      }
      health_state = "green"
      health_state_reasons = [{
        message                               = "example"
        severity                              = "green"
        traffic_health_state_computation_type = "PORT_LOW_UTIL"
      }]
      high_priority_drop = true
      matches = [{
        aggregate_filters = [{
          name           = "example"
          operation_type = "example"
          simple_filters = {
            advanced_filter_properties = {
              key    = "example"
              values = ["example"]
            }
            type = "example"
            values = {
              mask      = "example"
              offset    = "example"
              subset    = "example"
              value     = "example"
              value_max = "example"
            }
          }
          template_id = "example"
        }]
        operation_type   = "example"
        rule_criteria_id = "example"
      }]
      no_expansion_tags = ["example"]
      rule_id           = "example"
      rule_name         = "example"
      tools = [{
        alias          = "example"
        is_drop        = true
        is_giga_stream = true
        is_port        = true
        via_gsop       = true
      }]
      type = "example"
    }]
    src_port_timestamp = "example"
    src_ports_info = {
      comment      = "example"
      ports        = ["example"]
      template_ids = ["example"]
    }
    tags = [{
      tag_key    = "example"
      tag_values = ["example"]
    }]
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `comment` (String, optional)
* `deployed` (Boolean, optional) - policy is deployed or not
* `deployment_error` (String, optional) - policy deployment error message
* `deployment_percent` (String, optional) - policy deployment percentage
* `dest_port_timestamp` (String, optional)
* `health_state` (String, optional) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List, optional) (see [below for nested schema](#nestedatt--health_state_reasons))
* `name` (String, required) - policy name
* `policy_id` (String, optional) - generated unique policy ID
* `policy_timestamp` (String, optional)
* `priority` (Boolean, optional)
* `rules` (Attributes List, optional) (see [below for nested schema](#nestedatt--rules))
* `src_port_timestamp` (String, optional)
* `src_ports_info` (Attributes, optional) (see [below for nested schema](#nestedatt--src_ports_info))
* `tags` (Attributes List, optional) (see [below for nested schema](#nestedatt--tags))

<a id="nestedatt--health_state_reasons"></a>
### Nested Schema for `health_state_reasons`

Optional:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type

<a id="nestedatt--rules"></a>
### Nested Schema for `rules`

Optional:

* `comment` (String)
* `gs_operations` (Attributes) (see [below for nested schema](#nestedatt--rules--gs_operations))
* `health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--rules--health_state_reasons))
* `high_priority_drop` (Boolean)
* `matches` (Attributes List) (see [below for nested schema](#nestedatt--rules--matches))
* `no_expansion_tags` (List of String)
* `rule_id` (String)
* `rule_name` (String)
* `tools` (Attributes List) (see [below for nested schema](#nestedatt--rules--tools))
* `type` (String)

<a id="nestedatt--rules--gs_operations"></a>
### Nested Schema for `rules.gs_operations`

Optional:

* `gs_apps` (Attributes) - GigaSMART Applications for a GSOP. At least one GsApp must be defined for the config to be valid (see [below for nested schema](#nestedatt--rules--gs_operations--gs_apps))
* `gs_engines` (List of String)
* `gs_group_params` (Attributes) - GsGroup Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--gs_group_params))
* `templates` (Attributes List) (see [below for nested schema](#nestedatt--rules--gs_operations--templates))

<a id="nestedatt--rules--gs_operations--gs_apps"></a>
### Nested Schema for `rules.gs_operations.gs_apps`

Optional:

* `apf` (Attributes) - Only applicable for H-series (see [below for nested schema](#nestedatt--rules--gs_operations--gs_apps--apf))
* `dedup` (Attributes) (see [below for nested schema](#nestedatt--rules--gs_operations--gs_apps--dedup))
* `diameter_whitelist` (Attributes) - Only applicable for H-series (see [below for nested schema](#nestedatt--rules--gs_operations--gs_apps--diameter_whitelist))
* `flow_filter` (Attributes) - GigaSMART 'Flow Filter' Application Configuration (see [below for nested schema](#nestedatt--rules--gs_operations--gs_apps--flow_filter))
* `flow_sampling` (Attributes) - GigaSMART 'Flow Sampling' Application Configuration (see [below for nested schema](#nestedatt--rules--gs_operations--gs_apps--flow_sampling))
* `gseries_header_add` (Attributes) - Only applicable for G-series (see [below for nested schema](#nestedatt--rules--gs_operations--gs_apps--gseries_header_add))
* `gseries_header_remove` (Attributes) - Only applicable for G-series (see [below for nested schema](#nestedatt--rules--gs_operations--gs_apps--gseries_header_remove))
* `gseries_load_balance` (Attributes) - Only applicable for G-series per-rule GSOP (see [below for nested schema](#nestedatt--rules--gs_operations--gs_apps--gseries_load_balance))
* `gseries_pattern_match` (Attributes) - Only applicable for G-series per-rule GSOP (see [below for nested schema](#nestedatt--rules--gs_operations--gs_apps--gseries_pattern_match))
* `gtp_whitelist` (Attributes) - Only applicable for H-series (see [below for nested schema](#nestedatt--rules--gs_operations--gs_apps--gtp_whitelist))
* `header_add` (Attributes) - GigaSMART 'Add Header' Application Configuration (see [below for nested schema](#nestedatt--rules--gs_operations--gs_apps--header_add))
* `header_remove` (Attributes) - GigaSMART 'Remove Header' Application Configuration (see [below for nested schema](#nestedatt--rules--gs_operations--gs_apps--header_remove))
* `icap` (Attributes) - GigaSMART ICAP Configuration (see [below for nested schema](#nestedatt--rules--gs_operations--gs_apps--icap))
* `inline_ssl` (Attributes) - GigaSMART Inline SSL Profile Configuration (see [below for nested schema](#nestedatt--rules--gs_operations--gs_apps--inline_ssl))
* `load_balance` (Attributes) - GigaSMART 'Load Balancing' Application Configuration (see [below for nested schema](#nestedatt--rules--gs_operations--gs_apps--load_balance))
* `masking` (Attributes) - GigaSMART 'Masking' Application Configuration (see [below for nested schema](#nestedatt--rules--gs_operations--gs_apps--masking))
* `metadata_export` (Attributes) - Only applicable for H-series (see [below for nested schema](#nestedatt--rules--gs_operations--gs_apps--metadata_export))
* `netflow` (Attributes) - Only applicable for H-series (see [below for nested schema](#nestedatt--rules--gs_operations--gs_apps--netflow))
* `sa_apf` (Attributes) - Deprecated (functionality will not be supported from 5.4 release onwards); Only applicable for H-series (see [below for nested schema](#nestedatt--rules--gs_operations--gs_apps--sa_apf))
* `sip_whitelist` (Attributes) - Only applicable for H-series (see [below for nested schema](#nestedatt--rules--gs_operations--gs_apps--sip_whitelist))
* `slicing` (Attributes) - GigaSMART 'Slicing' Application Configuration (see [below for nested schema](#nestedatt--rules--gs_operations--gs_apps--slicing))
* `ssl_decrypt` (Attributes) - GigaSMART 'SSL Decrypt' Application Configuration (see [below for nested schema](#nestedatt--rules--gs_operations--gs_apps--ssl_decrypt))
* `trailer_add` (Attributes) - GigaSMART 'Add Trailer' Application Configuration (see [below for nested schema](#nestedatt--rules--gs_operations--gs_apps--trailer_add))
* `trailer_remove` (Attributes) (see [below for nested schema](#nestedatt--rules--gs_operations--gs_apps--trailer_remove))
* `tunnel_decap` (Attributes) - GigaSMART 'Decapsulate Tunnel' Application Configuration (see [below for nested schema](#nestedatt--rules--gs_operations--gs_apps--tunnel_decap))
* `tunnel_encap` (Attributes) - GigaSMART 'Encapsulate Tunnel' Application Configuration (see [below for nested schema](#nestedatt--rules--gs_operations--gs_apps--tunnel_encap))

<a id="nestedatt--rules--gs_operations--gs_apps--apf"></a>
### Nested Schema for `rules.gs_operations.gs_apps.apf`

Optional:

* `enabled` (String)

<a id="nestedatt--rules--gs_operations--gs_apps--dedup"></a>
### Nested Schema for `rules.gs_operations.gs_apps.dedup`

Optional:

* `enabled` (String)

<a id="nestedatt--rules--gs_operations--gs_apps--diameter_whitelist"></a>
### Nested Schema for `rules.gs_operations.gs_apps.diameter_whitelist`

Optional:

* `enabled` (String)

<a id="nestedatt--rules--gs_operations--gs_apps--flow_filter"></a>
### Nested Schema for `rules.gs_operations.gs_apps.flow_filter`

Required:

* `type` (String)

<a id="nestedatt--rules--gs_operations--gs_apps--flow_sampling"></a>
### Nested Schema for `rules.gs_operations.gs_apps.flow_sampling`

Required:

* `type` (String)

<a id="nestedatt--rules--gs_operations--gs_apps--gseries_header_add"></a>
### Nested Schema for `rules.gs_operations.gs_apps.gseries_header_add`

Required:

* `types` (Set of String)

<a id="nestedatt--rules--gs_operations--gs_apps--gseries_header_remove"></a>
### Nested Schema for `rules.gs_operations.gs_apps.gseries_header_remove`

Optional:

* `enabled` (String)

<a id="nestedatt--rules--gs_operations--gs_apps--gseries_load_balance"></a>
### Nested Schema for `rules.gs_operations.gs_apps.gseries_load_balance`

Optional:

* `fixed_offset` (Attributes) - Per-rule Fixed Offset Load Balancing config for G-seres devices (see [below for nested schema](#nestedatt--rules--gs_operations--gs_apps--gseries_load_balance--fixed_offset))
* `variable_offset` (Attributes) - Per-rule Variable Offset Load Balancing config for G-seres devices (see [below for nested schema](#nestedatt--rules--gs_operations--gs_apps--gseries_load_balance--variable_offset))

<a id="nestedatt--rules--gs_operations--gs_apps--gseries_load_balance--fixed_offset"></a>
### Nested Schema for `rules.gs_operations.gs_apps.gseries_load_balance.fixed_offset`

Required:

* `hash` (String)
* `length` (Number)
* `offset` (Number)

<a id="nestedatt--rules--gs_operations--gs_apps--gseries_load_balance--variable_offset"></a>
### Nested Schema for `rules.gs_operations.gs_apps.gseries_load_balance.variable_offset`

Required:

* `end_delim` (String)
* `hash` (String)
* `start_delim` (String)
* `start_field` (String)

<a id="nestedatt--rules--gs_operations--gs_apps--gseries_pattern_match"></a>
### Nested Schema for `rules.gs_operations.gs_apps.gseries_pattern_match`

Optional:

* `fixed_offset` (Attributes) - Per-rule Fixed Offset Pattern Match config for G-seres devices (see [below for nested schema](#nestedatt--rules--gs_operations--gs_apps--gseries_pattern_match--fixed_offset))
* `variable_offset` (Attributes) - Per-rule Variable Offset Pattern Match config for G-seres devices (see [below for nested schema](#nestedatt--rules--gs_operations--gs_apps--gseries_pattern_match--variable_offset))

<a id="nestedatt--rules--gs_operations--gs_apps--gseries_pattern_match--fixed_offset"></a>
### Nested Schema for `rules.gs_operations.gs_apps.gseries_pattern_match.fixed_offset`

Required:

* `length` (Number)
* `offset` (Number)

<a id="nestedatt--rules--gs_operations--gs_apps--gseries_pattern_match--variable_offset"></a>
### Nested Schema for `rules.gs_operations.gs_apps.gseries_pattern_match.variable_offset`

Required:

* `end_delim` (String)
* `start_delim` (String)

<a id="nestedatt--rules--gs_operations--gs_apps--gtp_whitelist"></a>
### Nested Schema for `rules.gs_operations.gs_apps.gtp_whitelist`

Optional:

* `enabled` (String)

<a id="nestedatt--rules--gs_operations--gs_apps--header_add"></a>
### Nested Schema for `rules.gs_operations.gs_apps.header_add`

Required:

* `vlan` (Number)

<a id="nestedatt--rules--gs_operations--gs_apps--header_remove"></a>
### Nested Schema for `rules.gs_operations.gs_apps.header_remove`

Required:

* `protocol` (String) - 'gre' and 'fabricPath' are only applicable for H-series

Optional:

* `ah1` (String) - only valid and required for 'generic'. First anchor header, from which the header to be stripped is occurred.
* `ah2` (String) - only valid and required for 'generic'. next anchor header.
* `custom_len` (Number) - only valid for 'generic'. length of unknown header.
* `erspan_flow_id` (Number) - only applicable for 'erspan', specifies the flow id to strip. Value of 0 will strip all flow ids
* `fp_dst_switch_id` (Number) - Only valid and required for 'fabricPath', 12 bit destination switch id
* `fp_src_switch_id` (Number) - Only valid and required for 'fabricPath', 12 bit source switch id
* `header_count` (Number) - only valid for 'generic'. Number of headers to be stripped.
* `offset` (String) - only valid and required for 'generic'.'start': strip from start of ah1, 'end': strip from end of ah1 or any other offset in the range of length of ah1, 'offsetRangeValue' is required if offset is 'offsetRange'.
* `offset_range_value` (Number) - only valid and required when offset is 'offsetRange', integer within range of size of header
* `timestamp_format` (String) - Timestamp format. Only valid and required for 'fm6000Ts'
* `vlan_header` (String) - Only valid when protocol is 'vlan'. Specifies the target vlan header to strip. Defaults to 'outer'
* `vxlan_id` (Number) - 24-bit value. Only valid when protocol is 'vxlan'. Specifies the vxlan id to strip. Value of '0' will strip all vxlan ids

<a id="nestedatt--rules--gs_operations--gs_apps--icap"></a>
### Nested Schema for `rules.gs_operations.gs_apps.icap`

Required:

* `icap_profile` (String) - Alias of referenced ICAP Profile

<a id="nestedatt--rules--gs_operations--gs_apps--inline_ssl"></a>
### Nested Schema for `rules.gs_operations.gs_apps.inline_ssl`

Required:

* `inline_ssl_profile` (String) - Alias of referenced Inline SSL Profile

<a id="nestedatt--rules--gs_operations--gs_apps--load_balance"></a>
### Nested Schema for `rules.gs_operations.gs_apps.load_balance`

Optional:

* `enhanced` (Attributes) - Private class. Enhanced part of the GigaSMART 'Load Balancing' Application Configuration (see [below for nested schema](#nestedatt--rules--gs_operations--gs_apps--load_balance--enhanced))
* `stateful` (Attributes) - Private class. Stateful part of the GigaSMART 'Load Balancing' Application Configuration (see [below for nested schema](#nestedatt--rules--gs_operations--gs_apps--load_balance--stateful))
* `stateless` (Attributes) - Private class. Stateless part of the GigaSMART 'Load Balancing' Application Configuration (see [below for nested schema](#nestedatt--rules--gs_operations--gs_apps--load_balance--stateless))

<a id="nestedatt--rules--gs_operations--gs_apps--load_balance--enhanced"></a>
### Nested Schema for `rules.gs_operations.gs_apps.load_balance.enhanced`

Required:

* `elb_alias` (String) - elb app alias

<a id="nestedatt--rules--gs_operations--gs_apps--load_balance--stateful"></a>
### Nested Schema for `rules.gs_operations.gs_apps.load_balance.stateful`

Required:

* `app_type` (String) - 'sapf' option is added in release H 4.3.01. 'sip' and 'tunnel' added in release H 5.1. 'diameter' is added in release H 5.5
* `lb_type` (String) - 'gtpKeyHash' is only applicable when 'appType' == 'gtp'. 'sipKeyHash' is only applicable when 'appType' == 'sip'

Optional:

* `diameter_key_hash_type` (String) - required when 'appType' == 'diameter' and 'lbType' == 'diameterKeyHash'
* `diameter_key_multi_hash_type` (Attributes List) - required when 'appType' == 'diameter' and 'lbType' == 'diameterKeyMultiHash' (see [below for nested schema](#nestedatt--rules--gs_operations--gs_apps--load_balance--stateful--diameter_key_multi_hash_type))
* `gtp_key_hash_type` (String) - required when 'appType' == 'gtp' and 'lbType' == 'gtpKeyHash'. ignored otherwise
* `sip_key_hash_type` (String) - required when 'appType' == 'sip' and 'lbType' == 'sipKeyHash'. ignored otherwise

<a id="nestedatt--rules--gs_operations--gs_apps--load_balance--stateful--diameter_key_multi_hash_type"></a>
### Nested Schema for `rules.gs_operations.gs_apps.load_balance.stateful.diameter_key_multi_hash_type`

Required:

* `key` (String)

Optional:

* `avp_codevalue` (Number) - required when 'key' == 'avpCode'

<a id="nestedatt--rules--gs_operations--gs_apps--load_balance--stateless"></a>
### Nested Schema for `rules.gs_operations.gs_apps.load_balance.stateless`

Required:

* `hash_fields` (String)

Optional:

* `field_location` (String) - Ignored when 'hashFields' == 'gtpuTeid'. Required otherwise

<a id="nestedatt--rules--gs_operations--gs_apps--masking"></a>
### Nested Schema for `rules.gs_operations.gs_apps.masking`

Required:

* `protocol` (String)

Optional:

* `content_type` (String) - content type that will trigger masking, only valid and required for protocol 'sip'
* `length` (Number) - max length is 9600 on H-series; 9000 on G-series, not applicable for protocol 'sip' valid and required otherwise
* `offset` (Number) - min offset is 0 for 'none'; 1 for protocol-referenced offsets, not applicable for protocol 'sip' valid and required otherwise
* `pattern` (String) - 1-byte hex mask, not applicable for protocol 'sip' valid and required otherwise

<a id="nestedatt--rules--gs_operations--gs_apps--metadata_export"></a>
### Nested Schema for `rules.gs_operations.gs_apps.metadata_export`

Optional:

* `cache` (String) - metadata cache alias. cache should have exporters defined, and the exporters should have applicationProfiles defined

<a id="nestedatt--rules--gs_operations--gs_apps--netflow"></a>
### Nested Schema for `rules.gs_operations.gs_apps.netflow`

Optional:

* `enabled` (String)

<a id="nestedatt--rules--gs_operations--gs_apps--sa_apf"></a>
### Nested Schema for `rules.gs_operations.gs_apps.sa_apf`

Optional:

* `enabled` (String)

<a id="nestedatt--rules--gs_operations--gs_apps--sip_whitelist"></a>
### Nested Schema for `rules.gs_operations.gs_apps.sip_whitelist`

Optional:

* `enabled` (String)

<a id="nestedatt--rules--gs_operations--gs_apps--slicing"></a>
### Nested Schema for `rules.gs_operations.gs_apps.slicing`

Required:

* `offset` (Number) - min offset is 64 for 'none'; 4 for protocol-referenced offsets. required property till H 5.6
* `protocol` (String) - required property till H 5.6

Optional:

* `enhanced` (String) - enhanced-slicing apps alias

<a id="nestedatt--rules--gs_operations--gs_apps--ssl_decrypt"></a>
### Nested Schema for `rules.gs_operations.gs_apps.ssl_decrypt`

Optional:

* `in_port` (Number) - Port number of 0 represents 'any' port
* `out_port` (Number) - Port number of 0 represents 'auto' port

<a id="nestedatt--rules--gs_operations--gs_apps--trailer_add"></a>
### Nested Schema for `rules.gs_operations.gs_apps.trailer_add`

Required:

* `types` (Set of String) - 'crc' is not applicable for G-series

<a id="nestedatt--rules--gs_operations--gs_apps--trailer_remove"></a>
### Nested Schema for `rules.gs_operations.gs_apps.trailer_remove`

Optional:

* `enabled` (String)

<a id="nestedatt--rules--gs_operations--gs_apps--tunnel_decap"></a>
### Nested Schema for `rules.gs_operations.gs_apps.tunnel_decap`

Required:

* `type` (String)

Optional:

* `custom` (Attributes) - only applicable for 'custom', in which case it is required. (see [below for nested schema](#nestedatt--rules--gs_operations--gs_apps--tunnel_decap--custom))
* `erspan_flow_id` (Number) - only applicable for 'erspan', A Flow ID of 0 decapsulates all ERSPAN tunnel traffic regardless of Flow ID
* `gmip_port` (Number) - only applicable for 'gmip', in which case it is required. Specifies the UDP port on which the Tunnel Network port on the receiving GigaVUE H Series is listening. Must match the configuration of the portdst configured on the sending end of the tunnel
* `l2_gre_key` (Number) - only applicable for 'l2gre', in which case it is required.
* `tls_pcapng` (Attributes) - only applicable for 'tls-pcapng', in which case it is required (see [below for nested schema](#nestedatt--rules--gs_operations--gs_apps--tunnel_decap--tls_pcapng))
* `vxlan` (Attributes) - only applicable for 'vxlan', in which case it is required. (see [below for nested schema](#nestedatt--rules--gs_operations--gs_apps--tunnel_decap--vxlan))

<a id="nestedatt--rules--gs_operations--gs_apps--tunnel_decap--custom"></a>
### Nested Schema for `rules.gs_operations.gs_apps.tunnel_decap.custom`

Required:

* `port_dst` (Number) - when specified as 0, no validation will be done in the packet.
* `port_src` (Number) - when specified as 0, no validation will be done in the packet.

<a id="nestedatt--rules--gs_operations--gs_apps--tunnel_decap--tls_pcapng"></a>
### Nested Schema for `rules.gs_operations.gs_apps.tunnel_decap.tls_pcapng`

Optional:

* `decap_key` (String)
* `listener` (String)

<a id="nestedatt--rules--gs_operations--gs_apps--tunnel_decap--vxlan"></a>
### Nested Schema for `rules.gs_operations.gs_apps.tunnel_decap.vxlan`

Required:

* `port_dst` (Number)
* `port_src` (Number) - when specified as 0, no validation will be done in the packet.
* `vni` (Number)

<a id="nestedatt--rules--gs_operations--gs_apps--tunnel_encap"></a>
### Nested Schema for `rules.gs_operations.gs_apps.tunnel_encap`

Required:

* `type` (String)

Optional:

* `gmip_config` (Attributes) - Configuration for GMIP Tunnel Encapsulate GigaSmaprt App (see [below for nested schema](#nestedatt--rules--gs_operations--gs_apps--tunnel_encap--gmip_config))
* `l2_gre_config` (Attributes) - Configuration for l2GRE Tunnel Encapsulate GigaSmaprt App (see [below for nested schema](#nestedatt--rules--gs_operations--gs_apps--tunnel_encap--l2_gre_config))
* `tls_pcapng` (Attributes) - only applicable for 'tls-pcapng', in which case it is required (see [below for nested schema](#nestedatt--rules--gs_operations--gs_apps--tunnel_encap--tls_pcapng))
* `vxlan_config` (Attributes) - Configuration for Vxlan Tunnel Encapsulate GigaSmaprt App (see [below for nested schema](#nestedatt--rules--gs_operations--gs_apps--tunnel_encap--vxlan_config))

<a id="nestedatt--rules--gs_operations--gs_apps--tunnel_encap--gmip_config"></a>
### Nested Schema for `rules.gs_operations.gs_apps.tunnel_encap.gmip_config`

Required:

* `dst_ip` (String) - IP Destination. IPv4 or IPv6.
* `dst_port` (Number)
* `src_port` (Number)

Optional:

* `dscp` (Number)
* `flow_label` (Number)
* `prec` (Number) - decimal Precedence value from 0-7 to be used in the ToS byte of the outer headers of tunneled packets
* `ttl` (Number)

<a id="nestedatt--rules--gs_operations--gs_apps--tunnel_encap--l2_gre_config"></a>
### Nested Schema for `rules.gs_operations.gs_apps.tunnel_encap.l2_gre_config`

Required:

* `key` (Number)

Optional:

* `dscp` (Number)
* `dst_ip` (String) - ip destination, mutually exclusive with 'pgDst'. IPv4 or IPv6.
* `flow_label` (Number)
* `pg_dst` (String) - port group destination alias, mutually exclusive with 'dstIp'
* `prec` (Number) - decimal Precedence value from 0-7 to be used in the ToS byte of the outer headers of tunneled packets
* `session_field` (String) - required with stateful loadBalance when 'appType' is 'tunnel'
* `session_pos` (String) - required if 'sessionField' is specified
* `ttl` (Number)

<a id="nestedatt--rules--gs_operations--gs_apps--tunnel_encap--tls_pcapng"></a>
### Nested Schema for `rules.gs_operations.gs_apps.tunnel_encap.tls_pcapng`

Optional:

* `exporter` (String)
* `exporter_group` (String)

<a id="nestedatt--rules--gs_operations--gs_apps--tunnel_encap--vxlan_config"></a>
### Nested Schema for `rules.gs_operations.gs_apps.tunnel_encap.vxlan_config`

Required:

* `dst_ip` (String) - IP Destination. IPv4 or IPv6.
* `dst_port` (Number)
* `src_port` (Number)
* `vni` (Number)

Optional:

* `dscp` (Number)
* `ttl` (Number)

<a id="nestedatt--rules--gs_operations--gs_group_params"></a>
### Nested Schema for `rules.gs_operations.gs_group_params`

Optional:

* `app_tcp` (Attributes) - GsGroup TCP Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--gs_group_params--app_tcp))
* `dedup` (Attributes) - GsGroup Dedup Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--gs_group_params--dedup))
* `diameter_packet` (Attributes) - GsGroup Diameter Packet Timeout Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--gs_group_params--diameter_packet))
* `diameter_s6_a_session` (Attributes) - GsGroup Diameter s6a Session Timeout Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--gs_group_params--diameter_s6_a_session))
* `diameter_whitelist` (Attributes) - GsGroup Diameter Whitelist Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--gs_group_params--diameter_whitelist))
* `eflow` (Attributes) - GsGroup Eflow Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--gs_group_params--eflow))
* `engine_watchdog_timer` (Attributes) - GsGroup EngineWatchdogTimer Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--gs_group_params--engine_watchdog_timer))
* `erspan3` (Attributes) - GsGroup ERSPAN III Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--gs_group_params--erspan3))
* `flow_mask` (Attributes) - GsGroup Flow Mask Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--gs_group_params--flow_mask))
* `flow_sampling` (Attributes) - GsGroup Flow Sampling Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--gs_group_params--flow_sampling))
* `generic_session_timeout` (Attributes) - GsGroup Generic Session Timeout Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--gs_group_params--generic_session_timeout))
* `gpfcp_profiles` (Attributes) - Enriched CUPS Gpfcp profile aliases (see [below for nested schema](#nestedatt--rules--gs_operations--gs_group_params--gpfcp_profiles))
* `gs_group_system` (Attributes) - GsGroup System Monitoring Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--gs_group_params--gs_group_system))
* `gta_profiles` (Attributes) - 3GPP CUPS gta profile aliases (see [below for nested schema](#nestedatt--rules--gs_operations--gs_group_params--gta_profiles))
* `gtp_control_sampling` (Attributes) - GsGroup Gtp Control Sampling Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--gs_group_params--gtp_control_sampling))
* `gtp_flow` (Attributes) - GsGroup Gtp Flow Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--gs_group_params--gtp_flow))
* `gtp_gpfcp_delay` (Attributes) - GsGroup Gtp GPFCP Delay time (see [below for nested schema](#nestedatt--rules--gs_operations--gs_group_params--gtp_gpfcp_delay))
* `gtp_persistence` (Attributes) - GsGroup Gtp Persistence Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--gs_group_params--gtp_persistence))
* `gtp_random_sampling` (Attributes) - GsGroup Gtp Random Sampling Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--gs_group_params--gtp_random_sampling))
* `gtp_whitelist` (Attributes) - GsGroup GTP Whitelist Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--gs_group_params--gtp_whitelist))
* `health_check` (Attributes) - health check configuration (see [below for nested schema](#nestedatt--rules--gs_operations--gs_group_params--health_check))
* `hsm_group` (Attributes) - GsGroup Hsm Group Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--gs_group_params--hsm_group))
* `ip_frag` (Attributes) - GsGroup IP Fragmentation Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--gs_group_params--ip_frag))
* `load_balance` (Attributes) - GsGroup Load Balancing Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--gs_group_params--load_balance))
* `netflow` (Attributes) - GsGroup Netflow Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--gs_group_params--netflow))
* `node_role` (Attributes) (see [below for nested schema](#nestedatt--rules--gs_operations--gs_group_params--node_role))
* `port_throttle_sip` (Attributes) - GsGroup SIP Port Throttle Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--gs_group_params--port_throttle_sip))
* `resource` (Attributes) - GsGroup Resource Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--gs_group_params--resource))
* `rtp_ports` (Attributes) - GsGroup Rtp Ports Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--gs_group_params--rtp_ports))
* `sa_apf` (Attributes) - GsGroup Session Aware APF Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--gs_group_params--sa_apf))
* `session_logging` (Attributes) - GsGroup Session Logging Configuration (see [below for nested schema](#nestedatt--rules--gs_operations--gs_group_params--session_logging))
* `sffp_profiles` (Attributes) - 3GPP CUPS sffp profile aliases (see [below for nested schema](#nestedatt--rules--gs_operations--gs_group_params--sffp_profiles))
* `sip_media` (Attributes) - GsGroup Sip Media Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--gs_group_params--sip_media))
* `sip_ports` (Attributes) - GsGroup Sip Ports Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--gs_group_params--sip_ports))
* `sip_session` (Attributes) - GsGroup Sip Session Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--gs_group_params--sip_session))
* `sip_tcp_idle_timeout` (Attributes) - GsGroup Sip Tcp Idle Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--gs_group_params--sip_tcp_idle_timeout))
* `sip_whitelist` (Attributes) - GsGroup SIP Whitelist Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--gs_group_params--sip_whitelist))
* `ssl_decrypt` (Attributes) - GsGroup SSL Decrypt Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--gs_group_params--ssl_decrypt))
* `xpkt_match` (Attributes) - cross packet match configuration (see [below for nested schema](#nestedatt--rules--gs_operations--gs_group_params--xpkt_match))

<a id="nestedatt--rules--gs_operations--gs_group_params--app_tcp"></a>
### Nested Schema for `rules.gs_operations.gs_group_params.app_tcp`

Optional:

* `application` (String) - To choose the action on Unknown Application Data
* `load_balance` (Boolean) - Enables TCP loadbalancing for the Tool Ports. Supports SIP & N11 only
* `tcp_control` (String) - To choose the action on TCP Control messages

<a id="nestedatt--rules--gs_operations--gs_group_params--dedup"></a>
### Nested Schema for `rules.gs_operations.gs_group_params.dedup`

Optional:

* `action` (String)
* `ip_tclass` (String)
* `ip_tos` (String)
* `tcp_seq` (String)
* `timer` (Number) - in microseconds
* `vlan` (String)

<a id="nestedatt--rules--gs_operations--gs_group_params--diameter_packet"></a>
### Nested Schema for `rules.gs_operations.gs_group_params.diameter_packet`

Optional:

* `timeout` (Number) - timeout in seconds to remove OOO or fragmented packet

<a id="nestedatt--rules--gs_operations--gs_group_params--diameter_s6_a_session"></a>
### Nested Schema for `rules.gs_operations.gs_group_params.diameter_s6_a_session`

Optional:

* `limit` (Number) - Number of sessions to allocate for Diameter S6A
* `timeout` (Number) - timeout in seconds used to clean inactive sessions

<a id="nestedatt--rules--gs_operations--gs_group_params--diameter_whitelist"></a>
### Nested Schema for `rules.gs_operations.gs_group_params.diameter_whitelist`

Required:

* `whitelist` (String) - Alias of referenced diameter Whitelist

<a id="nestedatt--rules--gs_operations--gs_group_params--eflow"></a>
### Nested Schema for `rules.gs_operations.gs_group_params.eflow`

Optional:

* `enabled` (Boolean) - Enable/Disable elephant flow detection and handling
* `interval` (Number) - time interval in seconds
* `log_enabled` (Boolean) - Enable/Disable logging of elephant flow parameters into gs logs
* `packet_count` (Number) - Number of packets to be received by a flow
* `packet_ratio` (Number) - Percentage of packets in a flow vs overall packet count

<a id="nestedatt--rules--gs_operations--gs_group_params--engine_watchdog_timer"></a>
### Nested Schema for `rules.gs_operations.gs_group_params.engine_watchdog_timer`

Optional:

* `time` (Number) - Engine Watchdog Timer value, time to restart in seconds .Valid values: 60-600. 0 is disable

<a id="nestedatt--rules--gs_operations--gs_group_params--erspan3"></a>
### Nested Schema for `rules.gs_operations.gs_group_params.erspan3`

Optional:

* `timestamp_format` (String) - erspan III tunnelDecap timestamp format

<a id="nestedatt--rules--gs_operations--gs_group_params--flow_mask"></a>
### Nested Schema for `rules.gs_operations.gs_group_params.flow_mask`

Optional:

* `enabled` (Boolean)
* `length` (Number)
* `offset` (Number)

<a id="nestedatt--rules--gs_operations--gs_group_params--flow_sampling"></a>
### Nested Schema for `rules.gs_operations.gs_group_params.flow_sampling`

Optional:

* `ip_ranges` (List of Dynamic)
* `rate` (Number) - in percent
* `timeout` (Number) - in minutes
* `type` (String)

<a id="nestedatt--rules--gs_operations--gs_group_params--generic_session_timeout"></a>
### Nested Schema for `rules.gs_operations.gs_group_params.generic_session_timeout`

Optional:

* `time` (Number) - Maximum timeout for session entry

<a id="nestedatt--rules--gs_operations--gs_group_params--gpfcp_profiles"></a>
### Nested Schema for `rules.gs_operations.gs_group_params.gpfcp_profiles`

Optional:

* `g_pfcp_profiles` (List of String)

<a id="nestedatt--rules--gs_operations--gs_group_params--gs_group_system"></a>
### Nested Schema for `rules.gs_operations.gs_group_params.gs_group_system`

Optional:

* `cpu_load_alarm_threshold` (Number) - CPU rising threshold percentage 20-99%. Once this threshold is crossed for 5 seconds, an alarm is generated

<a id="nestedatt--rules--gs_operations--gs_group_params--gta_profiles"></a>
### Nested Schema for `rules.gs_operations.gs_group_params.gta_profiles`

Optional:

* `gta_profiles` (List of String)

<a id="nestedatt--rules--gs_operations--gs_group_params--gtp_control_sampling"></a>
### Nested Schema for `rules.gs_operations.gs_group_params.gtp_control_sampling`

Optional:

* `enabled` (Boolean) - When enabled, GTP Control plane traffic for the subscribers contained within the subscriber sample defined in GTP flowsample maps will be sent to tools. When disabled, all GTP Control plane traffic for all subscribers in GTP flowsample maps will be sent to tools.

<a id="nestedatt--rules--gs_operations--gs_group_params--gtp_flow"></a>
### Nested Schema for `rules.gs_operations.gs_group_params.gtp_flow`

Optional:

* `timeout` (Number) - Session Timeout. in units of 10 minutes. Default of 48 is 8 hours

<a id="nestedatt--rules--gs_operations--gs_group_params--gtp_gpfcp_delay"></a>
### Nested Schema for `rules.gs_operations.gs_group_params.gtp_gpfcp_delay`

Optional:

* `timeout` (Number)

<a id="nestedatt--rules--gs_operations--gs_group_params--gtp_persistence"></a>
### Nested Schema for `rules.gs_operations.gs_group_params.gtp_persistence`

Optional:

* `enabled` (Boolean) - GTP Persistence Status
* `file_age_timeout` (Number) - GTP Persistence File Age Timeout(mins)
* `interval` (Number) - GTP Persistence Interval(mins)
* `restart_age_time` (Number) - GTP Persistence Restart Age Time(mins)

<a id="nestedatt--rules--gs_operations--gs_group_params--gtp_random_sampling"></a>
### Nested Schema for `rules.gs_operations.gs_group_params.gtp_random_sampling`

Optional:

* `enabled` (Boolean) - When enabled, sampling of subscriber's sessions happens in random fashion
* `interval` (Number) - Rotation Interval in multiples of 12 (hrs)

<a id="nestedatt--rules--gs_operations--gs_group_params--gtp_whitelist"></a>
### Nested Schema for `rules.gs_operations.gs_group_params.gtp_whitelist`

Required:

* `whitelist` (String) - Alias of referenced GTP Whitelist.Deprecated since H 5.12

Optional:

* `multi_whitelists` (List of String) - Alias/Aliases of referenced GTP Whitelists.

<a id="nestedatt--rules--gs_operations--gs_group_params--health_check"></a>
### Nested Schema for `rules.gs_operations.gs_group_params.health_check`

Optional:

* `action` (String)
* `dst_port` (Number)
* `enabled` (Boolean)
* `interval` (Number)
* `protocol` (String)
* `rcv_port` (Number)
* `retries` (Number)
* `round_trip_time` (Number)
* `src_port` (Number)

<a id="nestedatt--rules--gs_operations--gs_group_params--hsm_group"></a>
### Nested Schema for `rules.gs_operations.gs_group_params.hsm_group`

Optional:

* `hsm_group` (String) - alias of Hsm Group

<a id="nestedatt--rules--gs_operations--gs_group_params--ip_frag"></a>
### Nested Schema for `rules.gs_operations.gs_group_params.ip_frag`

Optional:

* `forward` (Boolean)
* `head_session_timeout` (Number) - in seconds
* `timeout` (Number) - in seconds

<a id="nestedatt--rules--gs_operations--gs_group_params--load_balance"></a>
### Nested Schema for `rules.gs_operations.gs_group_params.load_balance`

Optional:

* `failover` (Attributes) - Private class. Failover part of the GsGroup Load Balancing Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--gs_group_params--load_balance--failover))
* `link_weight_type` (String) - Controls LB port weight resolution for 'weighted' load balancing (wtRoundRobin, wtLeastBw, wtLeastPktRate, wtLeastConn, wtLeastTotalTraffic,wtImsi, wtSupi). For 'assigned', explicit user-configured port weights on the port group are used. For 'speed', link speed is used and user-assigned port weights are ignored. Ex: if a port group consists of four tool ports, and one of them is 100G and the others are 10G, the 100G link will be selected about 10 times more than the 10G links
* `replicate_gtpc` (Boolean) - Enables replication of GTP control packets (GTP-c)

<a id="nestedatt--rules--gs_operations--gs_group_params--load_balance--failover"></a>
### Nested Schema for `rules.gs_operations.gs_group_params.load_balance.failover`

Optional:

* `enabled` (Boolean) - Enables or disables failover when tool ports are down or thresholds to other tool ports in the load balancing port group are exceeded. A GigaSMART application failover will occur no more than once in 30 seconds
* `threshold_lt_bw` (Number) - Mutually exclusive with 'thresholdLtPktRate'. Failover threshold for Least Bandwidth. In percent's of the maximum bandwidth of a tool port. Ex: for a 1G port, a failover threshold of 90% means that failover to another tool port occurs when the bandwidth reaches 900Mbps. The default is 80%
* `threshold_lt_pkt_rate` (Number) - Mutually exclusive with 'thresholdLtBw'. Failover threshold for Least Packet Rate. In In kilo-packets-per-second: \[500k..5M\] pps. A tool port will failover to another tool port when the packet rate is over the specified threshold, in packets per second. The default is 1M

<a id="nestedatt--rules--gs_operations--gs_group_params--netflow"></a>
### Nested Schema for `rules.gs_operations.gs_group_params.netflow`

Optional:

* `monitor` (String) - Alias of referenced Netflow Monitor

<a id="nestedatt--rules--gs_operations--gs_group_params--node_role"></a>
### Nested Schema for `rules.gs_operations.gs_group_params.node_role`

Optional:

* `mob5_g_limit` (Number) - Number of sessions to allocate for Control 5G Node
* `mob_lte_limit` (Number) - Number of sessions to allocate for LTE CPN / UPN Node
* `stand_alone_mode` (Boolean) - Enables UPN Stand-alone mode
* `type` (String)

<a id="nestedatt--rules--gs_operations--gs_group_params--port_throttle_sip"></a>
### Nested Schema for `rules.gs_operations.gs_group_params.port_throttle_sip`

Optional:

* `port_throttle` (String) - Alias of referenced Port Throttle

<a id="nestedatt--rules--gs_operations--gs_group_params--resource"></a>
### Nested Schema for `rules.gs_operations.gs_group_params.resource`

Optional:

* `buffer_asf_size` (Number) - Session aware APF buffer. Valid values: 2-5: in million sessions (0 = disabled). Changes take effect after card or system reboot
* `cpu` (Attributes) (see [below for nested schema](#nestedatt--rules--gs_operations--gs_group_params--resource--cpu))
* `hsm_ssl` (Attributes) - GsGroup Resource Hsm Ssl Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--gs_group_params--resource--hsm_ssl))
* `inline_ssl` (Attributes) - Used to configure other GS apps in addition to Inline SSL on a HC1 box (see [below for nested schema](#nestedatt--rules--gs_operations--gs_group_params--resource--inline_ssl))
* `metadata` (Number) - flows in millions, how many flows to support for metadata. 0 to disable
* `packet_buffer` (Attributes) (see [below for nested schema](#nestedatt--rules--gs_operations--gs_group_params--resource--packet_buffer))
* `session_overload` (Attributes) (see [below for nested schema](#nestedatt--rules--gs_operations--gs_group_params--resource--session_overload))
* `tunnel_overload` (Attributes) (see [below for nested schema](#nestedatt--rules--gs_operations--gs_group_params--resource--tunnel_overload))
* `xpkt_match` (Attributes) - GsGroup Resource Cross Packet Match Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--gs_group_params--resource--xpkt_match))

<a id="nestedatt--rules--gs_operations--gs_group_params--resource--cpu"></a>
### Nested Schema for `rules.gs_operations.gs_group_params.resource.cpu`

Optional:

* `overload_threshold` (Number) - Valid values 50 - 90%. 0 is disabled. Changes take effect after card or system reboot

<a id="nestedatt--rules--gs_operations--gs_group_params--resource--hsm_ssl"></a>
### Nested Schema for `rules.gs_operations.gs_group_params.resource.hsm_ssl`

Optional:

* `buffer` (Number) - resource for application hsm-ssl buffer in MB. 0 to disable
* `packet_buffer` (Number) - resource for application hsm-ssl packet-buffer per connection
* `session_count` (Number) - resource for application hsm-ssl buffer session count in million, 0 to disable

<a id="nestedatt--rules--gs_operations--gs_group_params--resource--inline_ssl"></a>
### Nested Schema for `rules.gs_operations.gs_group_params.resource.inline_ssl`

Optional:

* `standalone` (Boolean) - If enabled , behaves in a normal way. If disabled, can configure other GS apps along side Inline ssl with Inline ssl occupying only 50% of the memory

<a id="nestedatt--rules--gs_operations--gs_group_params--resource--packet_buffer"></a>
### Nested Schema for `rules.gs_operations.gs_group_params.resource.packet_buffer`

Optional:

* `overload_threshold` (Number) - Valid values 50 - 80%. 0 is disabled. Changes take effect after card or system reboot

<a id="nestedatt--rules--gs_operations--gs_group_params--resource--session_overload"></a>
### Nested Schema for `rules.gs_operations.gs_group_params.resource.session_overload`

Optional:

* `overload_threshold` (Number) - Session overload threshold value , Default value is 90 and 0 is disabled.

<a id="nestedatt--rules--gs_operations--gs_group_params--resource--tunnel_overload"></a>
### Nested Schema for `rules.gs_operations.gs_group_params.resource.tunnel_overload`

Optional:

* `overload_threshold` (Number) - Tunnel overload threshold value , Default value is 90 and 0 is disabled.

<a id="nestedatt--rules--gs_operations--gs_group_params--resource--xpkt_match"></a>
### Nested Schema for `rules.gs_operations.gs_group_params.resource.xpkt_match`

Optional:

* `flows` (Number) - num in 100K flows. 0 is disable

<a id="nestedatt--rules--gs_operations--gs_group_params--rtp_ports"></a>
### Nested Schema for `rules.gs_operations.gs_group_params.rtp_ports`

Optional:

* `range` (Attributes) - GsGroup Rtp Port Range Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--gs_group_params--rtp_ports--range))

<a id="nestedatt--rules--gs_operations--gs_group_params--rtp_ports--range"></a>
### Nested Schema for `rules.gs_operations.gs_group_params.rtp_ports.range`

Required:

* `port` (Number)

Optional:

* `port_max` (Number) - If specified should be greater than 'port'

<a id="nestedatt--rules--gs_operations--gs_group_params--sa_apf"></a>
### Nested Schema for `rules.gs_operations.gs_group_params.sa_apf`

Optional:

* `buffer_size` (Number) - Session aware APF buffer. Valid values: 2-5: in million sessions (0 = disabled). Changes take effect after card or system reboot

<a id="nestedatt--rules--gs_operations--gs_group_params--session_logging"></a>
### Nested Schema for `rules.gs_operations.gs_group_params.session_logging`

Optional:

* `interface` (String) - Associated IP Interface
* `log_level` (String) - Log Level
* `remote_syslog_ip` (String) - Remote Syslog IP
* `remote_syslog_port` (Number) - Remote Syslog Port Number

<a id="nestedatt--rules--gs_operations--gs_group_params--sffp_profiles"></a>
### Nested Schema for `rules.gs_operations.gs_group_params.sffp_profiles`

Optional:

* `sffp_profiles` (List of String)

<a id="nestedatt--rules--gs_operations--gs_group_params--sip_media"></a>
### Nested Schema for `rules.gs_operations.gs_group_params.sip_media`

Optional:

* `timeout` (Number) - Sip media timeout value in seconds .Valid values: 30-300.

<a id="nestedatt--rules--gs_operations--gs_group_params--sip_ports"></a>
### Nested Schema for `rules.gs_operations.gs_group_params.sip_ports`

Optional:

* `ports` (List of Number) - list of TCP/UDP ports. Valid ports 1 - 65535. maximum 10 ports

<a id="nestedatt--rules--gs_operations--gs_group_params--sip_session"></a>
### Nested Schema for `rules.gs_operations.gs_group_params.sip_session`

Optional:

* `timeout` (Number) - Sip session inactivity timer, value in seconds .Valid values: 30-300.

<a id="nestedatt--rules--gs_operations--gs_group_params--sip_tcp_idle_timeout"></a>
### Nested Schema for `rules.gs_operations.gs_group_params.sip_tcp_idle_timeout`

Optional:

* `time` (Number) - Sip tcp idle timeout value in seconds .Valid values: 20-600.

<a id="nestedatt--rules--gs_operations--gs_group_params--sip_whitelist"></a>
### Nested Schema for `rules.gs_operations.gs_group_params.sip_whitelist`

Required:

* `whitelist` (String) - Alias of referenced SIP Whitelist

<a id="nestedatt--rules--gs_operations--gs_group_params--ssl_decrypt"></a>
### Nested Schema for `rules.gs_operations.gs_group_params.ssl_decrypt`

Required:

* `key_map` (String) - references one of the pre-defined 'SslDecryptionKeyMap' groups

Optional:

* `decrypt_fail_action` (String)
* `enabled` (Boolean)
* `hsm_pkcs11` (Attributes) - GsGroup Ssl Decrypt Hsm Pkcs11 Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--gs_group_params--ssl_decrypt--hsm_pkcs11))
* `hsm_timeout` (Number) - in milliseconds
* `key_cache_timeout` (Number) - in seconds
* `non_ssl_traffic` (String)
* `pending_session_timeout` (Number) - in seconds
* `session_timeout` (Number) - in seconds
* `tcp_syn_timeout` (Number) - in seconds
* `ticket_cache_timeout` (Number) - in seconds

<a id="nestedatt--rules--gs_operations--gs_group_params--ssl_decrypt--hsm_pkcs11"></a>
### Nested Schema for `rules.gs_operations.gs_group_params.ssl_decrypt.hsm_pkcs11`

Optional:

* `debug_level` (Number) - hsm pkcs11 debug level
* `dynamic_object` (Boolean) - hsm pkcs11 dynamic object
* `load_sharing` (Boolean) - hsm pkcs11 load sharing

<a id="nestedatt--rules--gs_operations--gs_group_params--xpkt_match"></a>
### Nested Schema for `rules.gs_operations.gs_group_params.xpkt_match`

Optional:

* `enabled` (Boolean)

<a id="nestedatt--rules--gs_operations--templates"></a>
### Nested Schema for `rules.gs_operations.templates`

Optional:

* `gs_apps` (Attributes) - GigaSMART Applications for a GSOP. At least one GsApp must be defined for the config to be valid (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_apps))
* `gs_group_params` (Attributes) - GsGroup Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_group_params))
* `template_id` (String)

<a id="nestedatt--rules--gs_operations--templates--gs_apps"></a>
### Nested Schema for `rules.gs_operations.templates.gs_apps`

Optional:

* `apf` (Attributes) - Only applicable for H-series (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_apps--apf))
* `dedup` (Attributes) (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_apps--dedup))
* `diameter_whitelist` (Attributes) - Only applicable for H-series (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_apps--diameter_whitelist))
* `flow_filter` (Attributes) - GigaSMART 'Flow Filter' Application Configuration (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_apps--flow_filter))
* `flow_sampling` (Attributes) - GigaSMART 'Flow Sampling' Application Configuration (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_apps--flow_sampling))
* `gseries_header_add` (Attributes) - Only applicable for G-series (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_apps--gseries_header_add))
* `gseries_header_remove` (Attributes) - Only applicable for G-series (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_apps--gseries_header_remove))
* `gseries_load_balance` (Attributes) - Only applicable for G-series per-rule GSOP (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_apps--gseries_load_balance))
* `gseries_pattern_match` (Attributes) - Only applicable for G-series per-rule GSOP (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_apps--gseries_pattern_match))
* `gtp_whitelist` (Attributes) - Only applicable for H-series (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_apps--gtp_whitelist))
* `header_add` (Attributes) - GigaSMART 'Add Header' Application Configuration (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_apps--header_add))
* `header_remove` (Attributes) - GigaSMART 'Remove Header' Application Configuration (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_apps--header_remove))
* `icap` (Attributes) - GigaSMART ICAP Configuration (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_apps--icap))
* `inline_ssl` (Attributes) - GigaSMART Inline SSL Profile Configuration (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_apps--inline_ssl))
* `load_balance` (Attributes) - GigaSMART 'Load Balancing' Application Configuration (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_apps--load_balance))
* `masking` (Attributes) - GigaSMART 'Masking' Application Configuration (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_apps--masking))
* `metadata_export` (Attributes) - Only applicable for H-series (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_apps--metadata_export))
* `netflow` (Attributes) - Only applicable for H-series (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_apps--netflow))
* `sa_apf` (Attributes) - Deprecated (functionality will not be supported from 5.4 release onwards); Only applicable for H-series (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_apps--sa_apf))
* `sip_whitelist` (Attributes) - Only applicable for H-series (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_apps--sip_whitelist))
* `slicing` (Attributes) - GigaSMART 'Slicing' Application Configuration (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_apps--slicing))
* `ssl_decrypt` (Attributes) - GigaSMART 'SSL Decrypt' Application Configuration (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_apps--ssl_decrypt))
* `trailer_add` (Attributes) - GigaSMART 'Add Trailer' Application Configuration (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_apps--trailer_add))
* `trailer_remove` (Attributes) (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_apps--trailer_remove))
* `tunnel_decap` (Attributes) - GigaSMART 'Decapsulate Tunnel' Application Configuration (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_apps--tunnel_decap))
* `tunnel_encap` (Attributes) - GigaSMART 'Encapsulate Tunnel' Application Configuration (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_apps--tunnel_encap))

<a id="nestedatt--rules--gs_operations--templates--gs_apps--apf"></a>
### Nested Schema for `rules.gs_operations.templates.gs_apps.apf`

Optional:

* `enabled` (String)

<a id="nestedatt--rules--gs_operations--templates--gs_apps--dedup"></a>
### Nested Schema for `rules.gs_operations.templates.gs_apps.dedup`

Optional:

* `enabled` (String)

<a id="nestedatt--rules--gs_operations--templates--gs_apps--diameter_whitelist"></a>
### Nested Schema for `rules.gs_operations.templates.gs_apps.diameter_whitelist`

Optional:

* `enabled` (String)

<a id="nestedatt--rules--gs_operations--templates--gs_apps--flow_filter"></a>
### Nested Schema for `rules.gs_operations.templates.gs_apps.flow_filter`

Required:

* `type` (String)

<a id="nestedatt--rules--gs_operations--templates--gs_apps--flow_sampling"></a>
### Nested Schema for `rules.gs_operations.templates.gs_apps.flow_sampling`

Required:

* `type` (String)

<a id="nestedatt--rules--gs_operations--templates--gs_apps--gseries_header_add"></a>
### Nested Schema for `rules.gs_operations.templates.gs_apps.gseries_header_add`

Required:

* `types` (Set of String)

<a id="nestedatt--rules--gs_operations--templates--gs_apps--gseries_header_remove"></a>
### Nested Schema for `rules.gs_operations.templates.gs_apps.gseries_header_remove`

Optional:

* `enabled` (String)

<a id="nestedatt--rules--gs_operations--templates--gs_apps--gseries_load_balance"></a>
### Nested Schema for `rules.gs_operations.templates.gs_apps.gseries_load_balance`

Optional:

* `fixed_offset` (Attributes) - Per-rule Fixed Offset Load Balancing config for G-seres devices (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_apps--gseries_load_balance--fixed_offset))
* `variable_offset` (Attributes) - Per-rule Variable Offset Load Balancing config for G-seres devices (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_apps--gseries_load_balance--variable_offset))

<a id="nestedatt--rules--gs_operations--templates--gs_apps--gseries_load_balance--fixed_offset"></a>
### Nested Schema for `rules.gs_operations.templates.gs_apps.gseries_load_balance.fixed_offset`

Required:

* `hash` (String)
* `length` (Number)
* `offset` (Number)

<a id="nestedatt--rules--gs_operations--templates--gs_apps--gseries_load_balance--variable_offset"></a>
### Nested Schema for `rules.gs_operations.templates.gs_apps.gseries_load_balance.variable_offset`

Required:

* `end_delim` (String)
* `hash` (String)
* `start_delim` (String)
* `start_field` (String)

<a id="nestedatt--rules--gs_operations--templates--gs_apps--gseries_pattern_match"></a>
### Nested Schema for `rules.gs_operations.templates.gs_apps.gseries_pattern_match`

Optional:

* `fixed_offset` (Attributes) - Per-rule Fixed Offset Pattern Match config for G-seres devices (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_apps--gseries_pattern_match--fixed_offset))
* `variable_offset` (Attributes) - Per-rule Variable Offset Pattern Match config for G-seres devices (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_apps--gseries_pattern_match--variable_offset))

<a id="nestedatt--rules--gs_operations--templates--gs_apps--gseries_pattern_match--fixed_offset"></a>
### Nested Schema for `rules.gs_operations.templates.gs_apps.gseries_pattern_match.fixed_offset`

Required:

* `length` (Number)
* `offset` (Number)

<a id="nestedatt--rules--gs_operations--templates--gs_apps--gseries_pattern_match--variable_offset"></a>
### Nested Schema for `rules.gs_operations.templates.gs_apps.gseries_pattern_match.variable_offset`

Required:

* `end_delim` (String)
* `start_delim` (String)

<a id="nestedatt--rules--gs_operations--templates--gs_apps--gtp_whitelist"></a>
### Nested Schema for `rules.gs_operations.templates.gs_apps.gtp_whitelist`

Optional:

* `enabled` (String)

<a id="nestedatt--rules--gs_operations--templates--gs_apps--header_add"></a>
### Nested Schema for `rules.gs_operations.templates.gs_apps.header_add`

Required:

* `vlan` (Number)

<a id="nestedatt--rules--gs_operations--templates--gs_apps--header_remove"></a>
### Nested Schema for `rules.gs_operations.templates.gs_apps.header_remove`

Required:

* `protocol` (String) - 'gre' and 'fabricPath' are only applicable for H-series

Optional:

* `ah1` (String) - only valid and required for 'generic'. First anchor header, from which the header to be stripped is occurred.
* `ah2` (String) - only valid and required for 'generic'. next anchor header.
* `custom_len` (Number) - only valid for 'generic'. length of unknown header.
* `erspan_flow_id` (Number) - only applicable for 'erspan', specifies the flow id to strip. Value of 0 will strip all flow ids
* `fp_dst_switch_id` (Number) - Only valid and required for 'fabricPath', 12 bit destination switch id
* `fp_src_switch_id` (Number) - Only valid and required for 'fabricPath', 12 bit source switch id
* `header_count` (Number) - only valid for 'generic'. Number of headers to be stripped.
* `offset` (String) - only valid and required for 'generic'.'start': strip from start of ah1, 'end': strip from end of ah1 or any other offset in the range of length of ah1, 'offsetRangeValue' is required if offset is 'offsetRange'.
* `offset_range_value` (Number) - only valid and required when offset is 'offsetRange', integer within range of size of header
* `timestamp_format` (String) - Timestamp format. Only valid and required for 'fm6000Ts'
* `vlan_header` (String) - Only valid when protocol is 'vlan'. Specifies the target vlan header to strip. Defaults to 'outer'
* `vxlan_id` (Number) - 24-bit value. Only valid when protocol is 'vxlan'. Specifies the vxlan id to strip. Value of '0' will strip all vxlan ids

<a id="nestedatt--rules--gs_operations--templates--gs_apps--icap"></a>
### Nested Schema for `rules.gs_operations.templates.gs_apps.icap`

Required:

* `icap_profile` (String) - Alias of referenced ICAP Profile

<a id="nestedatt--rules--gs_operations--templates--gs_apps--inline_ssl"></a>
### Nested Schema for `rules.gs_operations.templates.gs_apps.inline_ssl`

Required:

* `inline_ssl_profile` (String) - Alias of referenced Inline SSL Profile

<a id="nestedatt--rules--gs_operations--templates--gs_apps--load_balance"></a>
### Nested Schema for `rules.gs_operations.templates.gs_apps.load_balance`

Optional:

* `enhanced` (Attributes) - Private class. Enhanced part of the GigaSMART 'Load Balancing' Application Configuration (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_apps--load_balance--enhanced))
* `stateful` (Attributes) - Private class. Stateful part of the GigaSMART 'Load Balancing' Application Configuration (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_apps--load_balance--stateful))
* `stateless` (Attributes) - Private class. Stateless part of the GigaSMART 'Load Balancing' Application Configuration (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_apps--load_balance--stateless))

<a id="nestedatt--rules--gs_operations--templates--gs_apps--load_balance--enhanced"></a>
### Nested Schema for `rules.gs_operations.templates.gs_apps.load_balance.enhanced`

Required:

* `elb_alias` (String) - elb app alias

<a id="nestedatt--rules--gs_operations--templates--gs_apps--load_balance--stateful"></a>
### Nested Schema for `rules.gs_operations.templates.gs_apps.load_balance.stateful`

Required:

* `app_type` (String) - 'sapf' option is added in release H 4.3.01. 'sip' and 'tunnel' added in release H 5.1. 'diameter' is added in release H 5.5
* `lb_type` (String) - 'gtpKeyHash' is only applicable when 'appType' == 'gtp'. 'sipKeyHash' is only applicable when 'appType' == 'sip'

Optional:

* `diameter_key_hash_type` (String) - required when 'appType' == 'diameter' and 'lbType' == 'diameterKeyHash'
* `diameter_key_multi_hash_type` (Attributes List) - required when 'appType' == 'diameter' and 'lbType' == 'diameterKeyMultiHash' (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_apps--load_balance--stateful--diameter_key_multi_hash_type))
* `gtp_key_hash_type` (String) - required when 'appType' == 'gtp' and 'lbType' == 'gtpKeyHash'. ignored otherwise
* `sip_key_hash_type` (String) - required when 'appType' == 'sip' and 'lbType' == 'sipKeyHash'. ignored otherwise

<a id="nestedatt--rules--gs_operations--templates--gs_apps--load_balance--stateful--diameter_key_multi_hash_type"></a>
### Nested Schema for `rules.gs_operations.templates.gs_apps.load_balance.stateful.diameter_key_multi_hash_type`

Required:

* `key` (String)

Optional:

* `avp_codevalue` (Number) - required when 'key' == 'avpCode'

<a id="nestedatt--rules--gs_operations--templates--gs_apps--load_balance--stateless"></a>
### Nested Schema for `rules.gs_operations.templates.gs_apps.load_balance.stateless`

Required:

* `hash_fields` (String)

Optional:

* `field_location` (String) - Ignored when 'hashFields' == 'gtpuTeid'. Required otherwise

<a id="nestedatt--rules--gs_operations--templates--gs_apps--masking"></a>
### Nested Schema for `rules.gs_operations.templates.gs_apps.masking`

Required:

* `protocol` (String)

Optional:

* `content_type` (String) - content type that will trigger masking, only valid and required for protocol 'sip'
* `length` (Number) - max length is 9600 on H-series; 9000 on G-series, not applicable for protocol 'sip' valid and required otherwise
* `offset` (Number) - min offset is 0 for 'none'; 1 for protocol-referenced offsets, not applicable for protocol 'sip' valid and required otherwise
* `pattern` (String) - 1-byte hex mask, not applicable for protocol 'sip' valid and required otherwise

<a id="nestedatt--rules--gs_operations--templates--gs_apps--metadata_export"></a>
### Nested Schema for `rules.gs_operations.templates.gs_apps.metadata_export`

Optional:

* `cache` (String) - metadata cache alias. cache should have exporters defined, and the exporters should have applicationProfiles defined

<a id="nestedatt--rules--gs_operations--templates--gs_apps--netflow"></a>
### Nested Schema for `rules.gs_operations.templates.gs_apps.netflow`

Optional:

* `enabled` (String)

<a id="nestedatt--rules--gs_operations--templates--gs_apps--sa_apf"></a>
### Nested Schema for `rules.gs_operations.templates.gs_apps.sa_apf`

Optional:

* `enabled` (String)

<a id="nestedatt--rules--gs_operations--templates--gs_apps--sip_whitelist"></a>
### Nested Schema for `rules.gs_operations.templates.gs_apps.sip_whitelist`

Optional:

* `enabled` (String)

<a id="nestedatt--rules--gs_operations--templates--gs_apps--slicing"></a>
### Nested Schema for `rules.gs_operations.templates.gs_apps.slicing`

Required:

* `offset` (Number) - min offset is 64 for 'none'; 4 for protocol-referenced offsets. required property till H 5.6
* `protocol` (String) - required property till H 5.6

Optional:

* `enhanced` (String) - enhanced-slicing apps alias

<a id="nestedatt--rules--gs_operations--templates--gs_apps--ssl_decrypt"></a>
### Nested Schema for `rules.gs_operations.templates.gs_apps.ssl_decrypt`

Optional:

* `in_port` (Number) - Port number of 0 represents 'any' port
* `out_port` (Number) - Port number of 0 represents 'auto' port

<a id="nestedatt--rules--gs_operations--templates--gs_apps--trailer_add"></a>
### Nested Schema for `rules.gs_operations.templates.gs_apps.trailer_add`

Required:

* `types` (Set of String) - 'crc' is not applicable for G-series

<a id="nestedatt--rules--gs_operations--templates--gs_apps--trailer_remove"></a>
### Nested Schema for `rules.gs_operations.templates.gs_apps.trailer_remove`

Optional:

* `enabled` (String)

<a id="nestedatt--rules--gs_operations--templates--gs_apps--tunnel_decap"></a>
### Nested Schema for `rules.gs_operations.templates.gs_apps.tunnel_decap`

Required:

* `type` (String)

Optional:

* `custom` (Attributes) - only applicable for 'custom', in which case it is required. (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_apps--tunnel_decap--custom))
* `erspan_flow_id` (Number) - only applicable for 'erspan', A Flow ID of 0 decapsulates all ERSPAN tunnel traffic regardless of Flow ID
* `gmip_port` (Number) - only applicable for 'gmip', in which case it is required. Specifies the UDP port on which the Tunnel Network port on the receiving GigaVUE H Series is listening. Must match the configuration of the portdst configured on the sending end of the tunnel
* `l2_gre_key` (Number) - only applicable for 'l2gre', in which case it is required.
* `tls_pcapng` (Attributes) - only applicable for 'tls-pcapng', in which case it is required (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_apps--tunnel_decap--tls_pcapng))
* `vxlan` (Attributes) - only applicable for 'vxlan', in which case it is required. (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_apps--tunnel_decap--vxlan))

<a id="nestedatt--rules--gs_operations--templates--gs_apps--tunnel_decap--custom"></a>
### Nested Schema for `rules.gs_operations.templates.gs_apps.tunnel_decap.custom`

Required:

* `port_dst` (Number) - when specified as 0, no validation will be done in the packet.
* `port_src` (Number) - when specified as 0, no validation will be done in the packet.

<a id="nestedatt--rules--gs_operations--templates--gs_apps--tunnel_decap--tls_pcapng"></a>
### Nested Schema for `rules.gs_operations.templates.gs_apps.tunnel_decap.tls_pcapng`

Optional:

* `decap_key` (String)
* `listener` (String)

<a id="nestedatt--rules--gs_operations--templates--gs_apps--tunnel_decap--vxlan"></a>
### Nested Schema for `rules.gs_operations.templates.gs_apps.tunnel_decap.vxlan`

Required:

* `port_dst` (Number)
* `port_src` (Number) - when specified as 0, no validation will be done in the packet.
* `vni` (Number)

<a id="nestedatt--rules--gs_operations--templates--gs_apps--tunnel_encap"></a>
### Nested Schema for `rules.gs_operations.templates.gs_apps.tunnel_encap`

Required:

* `type` (String)

Optional:

* `gmip_config` (Attributes) - Configuration for GMIP Tunnel Encapsulate GigaSmaprt App (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_apps--tunnel_encap--gmip_config))
* `l2_gre_config` (Attributes) - Configuration for l2GRE Tunnel Encapsulate GigaSmaprt App (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_apps--tunnel_encap--l2_gre_config))
* `tls_pcapng` (Attributes) - only applicable for 'tls-pcapng', in which case it is required (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_apps--tunnel_encap--tls_pcapng))
* `vxlan_config` (Attributes) - Configuration for Vxlan Tunnel Encapsulate GigaSmaprt App (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_apps--tunnel_encap--vxlan_config))

<a id="nestedatt--rules--gs_operations--templates--gs_apps--tunnel_encap--gmip_config"></a>
### Nested Schema for `rules.gs_operations.templates.gs_apps.tunnel_encap.gmip_config`

Required:

* `dst_ip` (String) - IP Destination. IPv4 or IPv6.
* `dst_port` (Number)
* `src_port` (Number)

Optional:

* `dscp` (Number)
* `flow_label` (Number)
* `prec` (Number) - decimal Precedence value from 0-7 to be used in the ToS byte of the outer headers of tunneled packets
* `ttl` (Number)

<a id="nestedatt--rules--gs_operations--templates--gs_apps--tunnel_encap--l2_gre_config"></a>
### Nested Schema for `rules.gs_operations.templates.gs_apps.tunnel_encap.l2_gre_config`

Required:

* `key` (Number)

Optional:

* `dscp` (Number)
* `dst_ip` (String) - ip destination, mutually exclusive with 'pgDst'. IPv4 or IPv6.
* `flow_label` (Number)
* `pg_dst` (String) - port group destination alias, mutually exclusive with 'dstIp'
* `prec` (Number) - decimal Precedence value from 0-7 to be used in the ToS byte of the outer headers of tunneled packets
* `session_field` (String) - required with stateful loadBalance when 'appType' is 'tunnel'
* `session_pos` (String) - required if 'sessionField' is specified
* `ttl` (Number)

<a id="nestedatt--rules--gs_operations--templates--gs_apps--tunnel_encap--tls_pcapng"></a>
### Nested Schema for `rules.gs_operations.templates.gs_apps.tunnel_encap.tls_pcapng`

Optional:

* `exporter` (String)
* `exporter_group` (String)

<a id="nestedatt--rules--gs_operations--templates--gs_apps--tunnel_encap--vxlan_config"></a>
### Nested Schema for `rules.gs_operations.templates.gs_apps.tunnel_encap.vxlan_config`

Required:

* `dst_ip` (String) - IP Destination. IPv4 or IPv6.
* `dst_port` (Number)
* `src_port` (Number)
* `vni` (Number)

Optional:

* `dscp` (Number)
* `ttl` (Number)

<a id="nestedatt--rules--gs_operations--templates--gs_group_params"></a>
### Nested Schema for `rules.gs_operations.templates.gs_group_params`

Optional:

* `app_tcp` (Attributes) - GsGroup TCP Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_group_params--app_tcp))
* `dedup` (Attributes) - GsGroup Dedup Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_group_params--dedup))
* `diameter_packet` (Attributes) - GsGroup Diameter Packet Timeout Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_group_params--diameter_packet))
* `diameter_s6_a_session` (Attributes) - GsGroup Diameter s6a Session Timeout Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_group_params--diameter_s6_a_session))
* `diameter_whitelist` (Attributes) - GsGroup Diameter Whitelist Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_group_params--diameter_whitelist))
* `eflow` (Attributes) - GsGroup Eflow Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_group_params--eflow))
* `engine_watchdog_timer` (Attributes) - GsGroup EngineWatchdogTimer Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_group_params--engine_watchdog_timer))
* `erspan3` (Attributes) - GsGroup ERSPAN III Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_group_params--erspan3))
* `flow_mask` (Attributes) - GsGroup Flow Mask Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_group_params--flow_mask))
* `flow_sampling` (Attributes) - GsGroup Flow Sampling Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_group_params--flow_sampling))
* `generic_session_timeout` (Attributes) - GsGroup Generic Session Timeout Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_group_params--generic_session_timeout))
* `gpfcp_profiles` (Attributes) - Enriched CUPS Gpfcp profile aliases (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_group_params--gpfcp_profiles))
* `gs_group_system` (Attributes) - GsGroup System Monitoring Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_group_params--gs_group_system))
* `gta_profiles` (Attributes) - 3GPP CUPS gta profile aliases (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_group_params--gta_profiles))
* `gtp_control_sampling` (Attributes) - GsGroup Gtp Control Sampling Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_group_params--gtp_control_sampling))
* `gtp_flow` (Attributes) - GsGroup Gtp Flow Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_group_params--gtp_flow))
* `gtp_gpfcp_delay` (Attributes) - GsGroup Gtp GPFCP Delay time (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_group_params--gtp_gpfcp_delay))
* `gtp_persistence` (Attributes) - GsGroup Gtp Persistence Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_group_params--gtp_persistence))
* `gtp_random_sampling` (Attributes) - GsGroup Gtp Random Sampling Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_group_params--gtp_random_sampling))
* `gtp_whitelist` (Attributes) - GsGroup GTP Whitelist Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_group_params--gtp_whitelist))
* `health_check` (Attributes) - health check configuration (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_group_params--health_check))
* `hsm_group` (Attributes) - GsGroup Hsm Group Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_group_params--hsm_group))
* `ip_frag` (Attributes) - GsGroup IP Fragmentation Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_group_params--ip_frag))
* `load_balance` (Attributes) - GsGroup Load Balancing Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_group_params--load_balance))
* `netflow` (Attributes) - GsGroup Netflow Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_group_params--netflow))
* `node_role` (Attributes) (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_group_params--node_role))
* `port_throttle_sip` (Attributes) - GsGroup SIP Port Throttle Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_group_params--port_throttle_sip))
* `resource` (Attributes) - GsGroup Resource Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_group_params--resource))
* `rtp_ports` (Attributes) - GsGroup Rtp Ports Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_group_params--rtp_ports))
* `sa_apf` (Attributes) - GsGroup Session Aware APF Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_group_params--sa_apf))
* `session_logging` (Attributes) - GsGroup Session Logging Configuration (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_group_params--session_logging))
* `sffp_profiles` (Attributes) - 3GPP CUPS sffp profile aliases (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_group_params--sffp_profiles))
* `sip_media` (Attributes) - GsGroup Sip Media Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_group_params--sip_media))
* `sip_ports` (Attributes) - GsGroup Sip Ports Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_group_params--sip_ports))
* `sip_session` (Attributes) - GsGroup Sip Session Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_group_params--sip_session))
* `sip_tcp_idle_timeout` (Attributes) - GsGroup Sip Tcp Idle Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_group_params--sip_tcp_idle_timeout))
* `sip_whitelist` (Attributes) - GsGroup SIP Whitelist Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_group_params--sip_whitelist))
* `ssl_decrypt` (Attributes) - GsGroup SSL Decrypt Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_group_params--ssl_decrypt))
* `xpkt_match` (Attributes) - cross packet match configuration (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_group_params--xpkt_match))

<a id="nestedatt--rules--gs_operations--templates--gs_group_params--app_tcp"></a>
### Nested Schema for `rules.gs_operations.templates.gs_group_params.app_tcp`

Optional:

* `application` (String) - To choose the action on Unknown Application Data
* `load_balance` (Boolean) - Enables TCP loadbalancing for the Tool Ports. Supports SIP & N11 only
* `tcp_control` (String) - To choose the action on TCP Control messages

<a id="nestedatt--rules--gs_operations--templates--gs_group_params--dedup"></a>
### Nested Schema for `rules.gs_operations.templates.gs_group_params.dedup`

Optional:

* `action` (String)
* `ip_tclass` (String)
* `ip_tos` (String)
* `tcp_seq` (String)
* `timer` (Number) - in microseconds
* `vlan` (String)

<a id="nestedatt--rules--gs_operations--templates--gs_group_params--diameter_packet"></a>
### Nested Schema for `rules.gs_operations.templates.gs_group_params.diameter_packet`

Optional:

* `timeout` (Number) - timeout in seconds to remove OOO or fragmented packet

<a id="nestedatt--rules--gs_operations--templates--gs_group_params--diameter_s6_a_session"></a>
### Nested Schema for `rules.gs_operations.templates.gs_group_params.diameter_s6_a_session`

Optional:

* `limit` (Number) - Number of sessions to allocate for Diameter S6A
* `timeout` (Number) - timeout in seconds used to clean inactive sessions

<a id="nestedatt--rules--gs_operations--templates--gs_group_params--diameter_whitelist"></a>
### Nested Schema for `rules.gs_operations.templates.gs_group_params.diameter_whitelist`

Required:

* `whitelist` (String) - Alias of referenced diameter Whitelist

<a id="nestedatt--rules--gs_operations--templates--gs_group_params--eflow"></a>
### Nested Schema for `rules.gs_operations.templates.gs_group_params.eflow`

Optional:

* `enabled` (Boolean) - Enable/Disable elephant flow detection and handling
* `interval` (Number) - time interval in seconds
* `log_enabled` (Boolean) - Enable/Disable logging of elephant flow parameters into gs logs
* `packet_count` (Number) - Number of packets to be received by a flow
* `packet_ratio` (Number) - Percentage of packets in a flow vs overall packet count

<a id="nestedatt--rules--gs_operations--templates--gs_group_params--engine_watchdog_timer"></a>
### Nested Schema for `rules.gs_operations.templates.gs_group_params.engine_watchdog_timer`

Optional:

* `time` (Number) - Engine Watchdog Timer value, time to restart in seconds .Valid values: 60-600. 0 is disable

<a id="nestedatt--rules--gs_operations--templates--gs_group_params--erspan3"></a>
### Nested Schema for `rules.gs_operations.templates.gs_group_params.erspan3`

Optional:

* `timestamp_format` (String) - erspan III tunnelDecap timestamp format

<a id="nestedatt--rules--gs_operations--templates--gs_group_params--flow_mask"></a>
### Nested Schema for `rules.gs_operations.templates.gs_group_params.flow_mask`

Optional:

* `enabled` (Boolean)
* `length` (Number)
* `offset` (Number)

<a id="nestedatt--rules--gs_operations--templates--gs_group_params--flow_sampling"></a>
### Nested Schema for `rules.gs_operations.templates.gs_group_params.flow_sampling`

Optional:

* `ip_ranges` (List of Dynamic)
* `rate` (Number) - in percent
* `timeout` (Number) - in minutes
* `type` (String)

<a id="nestedatt--rules--gs_operations--templates--gs_group_params--generic_session_timeout"></a>
### Nested Schema for `rules.gs_operations.templates.gs_group_params.generic_session_timeout`

Optional:

* `time` (Number) - Maximum timeout for session entry

<a id="nestedatt--rules--gs_operations--templates--gs_group_params--gpfcp_profiles"></a>
### Nested Schema for `rules.gs_operations.templates.gs_group_params.gpfcp_profiles`

Optional:

* `g_pfcp_profiles` (List of String)

<a id="nestedatt--rules--gs_operations--templates--gs_group_params--gs_group_system"></a>
### Nested Schema for `rules.gs_operations.templates.gs_group_params.gs_group_system`

Optional:

* `cpu_load_alarm_threshold` (Number) - CPU rising threshold percentage 20-99%. Once this threshold is crossed for 5 seconds, an alarm is generated

<a id="nestedatt--rules--gs_operations--templates--gs_group_params--gta_profiles"></a>
### Nested Schema for `rules.gs_operations.templates.gs_group_params.gta_profiles`

Optional:

* `gta_profiles` (List of String)

<a id="nestedatt--rules--gs_operations--templates--gs_group_params--gtp_control_sampling"></a>
### Nested Schema for `rules.gs_operations.templates.gs_group_params.gtp_control_sampling`

Optional:

* `enabled` (Boolean) - When enabled, GTP Control plane traffic for the subscribers contained within the subscriber sample defined in GTP flowsample maps will be sent to tools. When disabled, all GTP Control plane traffic for all subscribers in GTP flowsample maps will be sent to tools.

<a id="nestedatt--rules--gs_operations--templates--gs_group_params--gtp_flow"></a>
### Nested Schema for `rules.gs_operations.templates.gs_group_params.gtp_flow`

Optional:

* `timeout` (Number) - Session Timeout. in units of 10 minutes. Default of 48 is 8 hours

<a id="nestedatt--rules--gs_operations--templates--gs_group_params--gtp_gpfcp_delay"></a>
### Nested Schema for `rules.gs_operations.templates.gs_group_params.gtp_gpfcp_delay`

Optional:

* `timeout` (Number)

<a id="nestedatt--rules--gs_operations--templates--gs_group_params--gtp_persistence"></a>
### Nested Schema for `rules.gs_operations.templates.gs_group_params.gtp_persistence`

Optional:

* `enabled` (Boolean) - GTP Persistence Status
* `file_age_timeout` (Number) - GTP Persistence File Age Timeout(mins)
* `interval` (Number) - GTP Persistence Interval(mins)
* `restart_age_time` (Number) - GTP Persistence Restart Age Time(mins)

<a id="nestedatt--rules--gs_operations--templates--gs_group_params--gtp_random_sampling"></a>
### Nested Schema for `rules.gs_operations.templates.gs_group_params.gtp_random_sampling`

Optional:

* `enabled` (Boolean) - When enabled, sampling of subscriber's sessions happens in random fashion
* `interval` (Number) - Rotation Interval in multiples of 12 (hrs)

<a id="nestedatt--rules--gs_operations--templates--gs_group_params--gtp_whitelist"></a>
### Nested Schema for `rules.gs_operations.templates.gs_group_params.gtp_whitelist`

Required:

* `whitelist` (String) - Alias of referenced GTP Whitelist.Deprecated since H 5.12

Optional:

* `multi_whitelists` (List of String) - Alias/Aliases of referenced GTP Whitelists.

<a id="nestedatt--rules--gs_operations--templates--gs_group_params--health_check"></a>
### Nested Schema for `rules.gs_operations.templates.gs_group_params.health_check`

Optional:

* `action` (String)
* `dst_port` (Number)
* `enabled` (Boolean)
* `interval` (Number)
* `protocol` (String)
* `rcv_port` (Number)
* `retries` (Number)
* `round_trip_time` (Number)
* `src_port` (Number)

<a id="nestedatt--rules--gs_operations--templates--gs_group_params--hsm_group"></a>
### Nested Schema for `rules.gs_operations.templates.gs_group_params.hsm_group`

Optional:

* `hsm_group` (String) - alias of Hsm Group

<a id="nestedatt--rules--gs_operations--templates--gs_group_params--ip_frag"></a>
### Nested Schema for `rules.gs_operations.templates.gs_group_params.ip_frag`

Optional:

* `forward` (Boolean)
* `head_session_timeout` (Number) - in seconds
* `timeout` (Number) - in seconds

<a id="nestedatt--rules--gs_operations--templates--gs_group_params--load_balance"></a>
### Nested Schema for `rules.gs_operations.templates.gs_group_params.load_balance`

Optional:

* `failover` (Attributes) - Private class. Failover part of the GsGroup Load Balancing Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_group_params--load_balance--failover))
* `link_weight_type` (String) - Controls LB port weight resolution for 'weighted' load balancing (wtRoundRobin, wtLeastBw, wtLeastPktRate, wtLeastConn, wtLeastTotalTraffic,wtImsi, wtSupi). For 'assigned', explicit user-configured port weights on the port group are used. For 'speed', link speed is used and user-assigned port weights are ignored. Ex: if a port group consists of four tool ports, and one of them is 100G and the others are 10G, the 100G link will be selected about 10 times more than the 10G links
* `replicate_gtpc` (Boolean) - Enables replication of GTP control packets (GTP-c)

<a id="nestedatt--rules--gs_operations--templates--gs_group_params--load_balance--failover"></a>
### Nested Schema for `rules.gs_operations.templates.gs_group_params.load_balance.failover`

Optional:

* `enabled` (Boolean) - Enables or disables failover when tool ports are down or thresholds to other tool ports in the load balancing port group are exceeded. A GigaSMART application failover will occur no more than once in 30 seconds
* `threshold_lt_bw` (Number) - Mutually exclusive with 'thresholdLtPktRate'. Failover threshold for Least Bandwidth. In percent's of the maximum bandwidth of a tool port. Ex: for a 1G port, a failover threshold of 90% means that failover to another tool port occurs when the bandwidth reaches 900Mbps. The default is 80%
* `threshold_lt_pkt_rate` (Number) - Mutually exclusive with 'thresholdLtBw'. Failover threshold for Least Packet Rate. In In kilo-packets-per-second: \[500k..5M\] pps. A tool port will failover to another tool port when the packet rate is over the specified threshold, in packets per second. The default is 1M

<a id="nestedatt--rules--gs_operations--templates--gs_group_params--netflow"></a>
### Nested Schema for `rules.gs_operations.templates.gs_group_params.netflow`

Optional:

* `monitor` (String) - Alias of referenced Netflow Monitor

<a id="nestedatt--rules--gs_operations--templates--gs_group_params--node_role"></a>
### Nested Schema for `rules.gs_operations.templates.gs_group_params.node_role`

Optional:

* `mob5_g_limit` (Number) - Number of sessions to allocate for Control 5G Node
* `mob_lte_limit` (Number) - Number of sessions to allocate for LTE CPN / UPN Node
* `stand_alone_mode` (Boolean) - Enables UPN Stand-alone mode
* `type` (String)

<a id="nestedatt--rules--gs_operations--templates--gs_group_params--port_throttle_sip"></a>
### Nested Schema for `rules.gs_operations.templates.gs_group_params.port_throttle_sip`

Optional:

* `port_throttle` (String) - Alias of referenced Port Throttle

<a id="nestedatt--rules--gs_operations--templates--gs_group_params--resource"></a>
### Nested Schema for `rules.gs_operations.templates.gs_group_params.resource`

Optional:

* `buffer_asf_size` (Number) - Session aware APF buffer. Valid values: 2-5: in million sessions (0 = disabled). Changes take effect after card or system reboot
* `cpu` (Attributes) (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_group_params--resource--cpu))
* `hsm_ssl` (Attributes) - GsGroup Resource Hsm Ssl Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_group_params--resource--hsm_ssl))
* `inline_ssl` (Attributes) - Used to configure other GS apps in addition to Inline SSL on a HC1 box (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_group_params--resource--inline_ssl))
* `metadata` (Number) - flows in millions, how many flows to support for metadata. 0 to disable
* `packet_buffer` (Attributes) (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_group_params--resource--packet_buffer))
* `session_overload` (Attributes) (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_group_params--resource--session_overload))
* `tunnel_overload` (Attributes) (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_group_params--resource--tunnel_overload))
* `xpkt_match` (Attributes) - GsGroup Resource Cross Packet Match Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_group_params--resource--xpkt_match))

<a id="nestedatt--rules--gs_operations--templates--gs_group_params--resource--cpu"></a>
### Nested Schema for `rules.gs_operations.templates.gs_group_params.resource.cpu`

Optional:

* `overload_threshold` (Number) - Valid values 50 - 90%. 0 is disabled. Changes take effect after card or system reboot

<a id="nestedatt--rules--gs_operations--templates--gs_group_params--resource--hsm_ssl"></a>
### Nested Schema for `rules.gs_operations.templates.gs_group_params.resource.hsm_ssl`

Optional:

* `buffer` (Number) - resource for application hsm-ssl buffer in MB. 0 to disable
* `packet_buffer` (Number) - resource for application hsm-ssl packet-buffer per connection
* `session_count` (Number) - resource for application hsm-ssl buffer session count in million, 0 to disable

<a id="nestedatt--rules--gs_operations--templates--gs_group_params--resource--inline_ssl"></a>
### Nested Schema for `rules.gs_operations.templates.gs_group_params.resource.inline_ssl`

Optional:

* `standalone` (Boolean) - If enabled , behaves in a normal way. If disabled, can configure other GS apps along side Inline ssl with Inline ssl occupying only 50% of the memory

<a id="nestedatt--rules--gs_operations--templates--gs_group_params--resource--packet_buffer"></a>
### Nested Schema for `rules.gs_operations.templates.gs_group_params.resource.packet_buffer`

Optional:

* `overload_threshold` (Number) - Valid values 50 - 80%. 0 is disabled. Changes take effect after card or system reboot

<a id="nestedatt--rules--gs_operations--templates--gs_group_params--resource--session_overload"></a>
### Nested Schema for `rules.gs_operations.templates.gs_group_params.resource.session_overload`

Optional:

* `overload_threshold` (Number) - Session overload threshold value , Default value is 90 and 0 is disabled.

<a id="nestedatt--rules--gs_operations--templates--gs_group_params--resource--tunnel_overload"></a>
### Nested Schema for `rules.gs_operations.templates.gs_group_params.resource.tunnel_overload`

Optional:

* `overload_threshold` (Number) - Tunnel overload threshold value , Default value is 90 and 0 is disabled.

<a id="nestedatt--rules--gs_operations--templates--gs_group_params--resource--xpkt_match"></a>
### Nested Schema for `rules.gs_operations.templates.gs_group_params.resource.xpkt_match`

Optional:

* `flows` (Number) - num in 100K flows. 0 is disable

<a id="nestedatt--rules--gs_operations--templates--gs_group_params--rtp_ports"></a>
### Nested Schema for `rules.gs_operations.templates.gs_group_params.rtp_ports`

Optional:

* `range` (Attributes) - GsGroup Rtp Port Range Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_group_params--rtp_ports--range))

<a id="nestedatt--rules--gs_operations--templates--gs_group_params--rtp_ports--range"></a>
### Nested Schema for `rules.gs_operations.templates.gs_group_params.rtp_ports.range`

Required:

* `port` (Number)

Optional:

* `port_max` (Number) - If specified should be greater than 'port'

<a id="nestedatt--rules--gs_operations--templates--gs_group_params--sa_apf"></a>
### Nested Schema for `rules.gs_operations.templates.gs_group_params.sa_apf`

Optional:

* `buffer_size` (Number) - Session aware APF buffer. Valid values: 2-5: in million sessions (0 = disabled). Changes take effect after card or system reboot

<a id="nestedatt--rules--gs_operations--templates--gs_group_params--session_logging"></a>
### Nested Schema for `rules.gs_operations.templates.gs_group_params.session_logging`

Optional:

* `interface` (String) - Associated IP Interface
* `log_level` (String) - Log Level
* `remote_syslog_ip` (String) - Remote Syslog IP
* `remote_syslog_port` (Number) - Remote Syslog Port Number

<a id="nestedatt--rules--gs_operations--templates--gs_group_params--sffp_profiles"></a>
### Nested Schema for `rules.gs_operations.templates.gs_group_params.sffp_profiles`

Optional:

* `sffp_profiles` (List of String)

<a id="nestedatt--rules--gs_operations--templates--gs_group_params--sip_media"></a>
### Nested Schema for `rules.gs_operations.templates.gs_group_params.sip_media`

Optional:

* `timeout` (Number) - Sip media timeout value in seconds .Valid values: 30-300.

<a id="nestedatt--rules--gs_operations--templates--gs_group_params--sip_ports"></a>
### Nested Schema for `rules.gs_operations.templates.gs_group_params.sip_ports`

Optional:

* `ports` (List of Number) - list of TCP/UDP ports. Valid ports 1 - 65535. maximum 10 ports

<a id="nestedatt--rules--gs_operations--templates--gs_group_params--sip_session"></a>
### Nested Schema for `rules.gs_operations.templates.gs_group_params.sip_session`

Optional:

* `timeout` (Number) - Sip session inactivity timer, value in seconds .Valid values: 30-300.

<a id="nestedatt--rules--gs_operations--templates--gs_group_params--sip_tcp_idle_timeout"></a>
### Nested Schema for `rules.gs_operations.templates.gs_group_params.sip_tcp_idle_timeout`

Optional:

* `time` (Number) - Sip tcp idle timeout value in seconds .Valid values: 20-600.

<a id="nestedatt--rules--gs_operations--templates--gs_group_params--sip_whitelist"></a>
### Nested Schema for `rules.gs_operations.templates.gs_group_params.sip_whitelist`

Required:

* `whitelist` (String) - Alias of referenced SIP Whitelist

<a id="nestedatt--rules--gs_operations--templates--gs_group_params--ssl_decrypt"></a>
### Nested Schema for `rules.gs_operations.templates.gs_group_params.ssl_decrypt`

Required:

* `key_map` (String) - references one of the pre-defined 'SslDecryptionKeyMap' groups

Optional:

* `decrypt_fail_action` (String)
* `enabled` (Boolean)
* `hsm_pkcs11` (Attributes) - GsGroup Ssl Decrypt Hsm Pkcs11 Parameters (see [below for nested schema](#nestedatt--rules--gs_operations--templates--gs_group_params--ssl_decrypt--hsm_pkcs11))
* `hsm_timeout` (Number) - in milliseconds
* `key_cache_timeout` (Number) - in seconds
* `non_ssl_traffic` (String)
* `pending_session_timeout` (Number) - in seconds
* `session_timeout` (Number) - in seconds
* `tcp_syn_timeout` (Number) - in seconds
* `ticket_cache_timeout` (Number) - in seconds

<a id="nestedatt--rules--gs_operations--templates--gs_group_params--ssl_decrypt--hsm_pkcs11"></a>
### Nested Schema for `rules.gs_operations.templates.gs_group_params.ssl_decrypt.hsm_pkcs11`

Optional:

* `debug_level` (Number) - hsm pkcs11 debug level
* `dynamic_object` (Boolean) - hsm pkcs11 dynamic object
* `load_sharing` (Boolean) - hsm pkcs11 load sharing

<a id="nestedatt--rules--gs_operations--templates--gs_group_params--xpkt_match"></a>
### Nested Schema for `rules.gs_operations.templates.gs_group_params.xpkt_match`

Optional:

* `enabled` (Boolean)

<a id="nestedatt--rules--health_state_reasons"></a>
### Nested Schema for `rules.health_state_reasons`

Optional:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type

<a id="nestedatt--rules--matches"></a>
### Nested Schema for `rules.matches`

Optional:

* `aggregate_filters` (Attributes List) (see [below for nested schema](#nestedatt--rules--matches--aggregate_filters))
* `operation_type` (String)
* `rule_criteria_id` (String)

<a id="nestedatt--rules--matches--aggregate_filters"></a>
### Nested Schema for `rules.matches.aggregate_filters`

Optional:

* `name` (String)
* `operation_type` (String)
* `simple_filters` (Attributes) (see [below for nested schema](#nestedatt--rules--matches--aggregate_filters--simple_filters))
* `template_id` (String)

<a id="nestedatt--rules--matches--aggregate_filters--simple_filters"></a>
### Nested Schema for `rules.matches.aggregate_filters.simple_filters`

Optional:

* `advanced_filter_properties` (Attributes) (see [below for nested schema](#nestedatt--rules--matches--aggregate_filters--simple_filters--advanced_filter_properties))
* `type` (String)
* `values` (Attributes) (see [below for nested schema](#nestedatt--rules--matches--aggregate_filters--simple_filters--values))

<a id="nestedatt--rules--matches--aggregate_filters--simple_filters--advanced_filter_properties"></a>
### Nested Schema for `rules.matches.aggregate_filters.simple_filters.advanced_filter_properties`

Optional:

* `key` (String)
* `values` (List of Dynamic)

<a id="nestedatt--rules--matches--aggregate_filters--simple_filters--values"></a>
### Nested Schema for `rules.matches.aggregate_filters.simple_filters.values`

Optional:

* `mask` (String)
* `offset` (String)
* `subset` (String)
* `value` (String)
* `value_max` (String)

<a id="nestedatt--rules--tools"></a>
### Nested Schema for `rules.tools`

Optional:

* `alias` (String) - alias of destination. Either port or gigastream, in form <clusterId:portId> or <clusterId:gigastreamAlias>
* `is_drop` (Boolean) - destination is drop
* `is_giga_stream` (Boolean) - destination alias is a gigastream (tool or hybrid)
* `is_port` (Boolean) - destination alias is a port (tool or hybrid)
* `via_gsop` (Boolean) - Policy has GigaSMART operation

<a id="nestedatt--src_ports_info"></a>
### Nested Schema for `src_ports_info`

Optional:

* `comment` (String)
* `ports` (List of String)
* `template_ids` (List of String)

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Required:

* `tag_key` (String) - Name of the tag
* `tag_values` (List of String) - All possible values of the tag

