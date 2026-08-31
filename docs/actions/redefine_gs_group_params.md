---
page_title: "gigavuecore_redefine_gs_group_params Action - gigavuecore"
subcategory: ""
description: |-
  Redefine GS Group's Params
---

# gigavuecore_redefine_gs_group_params Action

Redefine GS Group's Params

## Example Usage

```terraform
action "gigavuecore_redefine_gs_group_params" "example" {
  config {
    alias = "example"
    app_tcp = {
      application  = "broadcast"
      load_balance = true
      tcp_control  = "broadcast"
    }
    cluster_id = "example"
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
      ip_ranges = [ "example" ]
      rate      = 5
      timeout   = 1
      type      = "deviceIp"
    }
    generic_session_timeout = {
      time = 5
    }
    gpfcp_profiles = {
      g_pfcp_profiles = [ "example" ]
    }
    gs_group_system = {
      cpu_load_alarm_threshold = 20
    }
    gta_profiles = {
      gta_profiles = [ "example" ]
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
      multi_whitelists = [ "example" ]
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
      sffp_profiles = [ "example" ]
    }
    sip_media = {
      timeout = 30
    }
    sip_ports = {
      ports = [ 1 ]
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
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target GS Group
* `app_tcp` (Attributes, optional) - GsGroup TCP Parameters (see [below for nested schema](#nestedatt--app_tcp))
* `cluster_id` (String, required) - Target Cluster ID
* `dedup` (Attributes, optional) - GsGroup Dedup Parameters (see [below for nested schema](#nestedatt--dedup))
* `diameter_packet` (Attributes, optional) - GsGroup Diameter Packet Timeout Parameters (see [below for nested schema](#nestedatt--diameter_packet))
* `diameter_s6_a_session` (Attributes, optional) - GsGroup Diameter s6a Session Timeout Parameters (see [below for nested schema](#nestedatt--diameter_s6_a_session))
* `diameter_whitelist` (Attributes, optional) - GsGroup Diameter Whitelist Parameters (see [below for nested schema](#nestedatt--diameter_whitelist))
* `eflow` (Attributes, optional) - GsGroup Eflow Parameters (see [below for nested schema](#nestedatt--eflow))
* `engine_watchdog_timer` (Attributes, optional) - GsGroup EngineWatchdogTimer Parameters (see [below for nested schema](#nestedatt--engine_watchdog_timer))
* `erspan3` (Attributes, optional) - GsGroup ERSPAN III Parameters (see [below for nested schema](#nestedatt--erspan3))
* `flow_mask` (Attributes, optional) - GsGroup Flow Mask Parameters (see [below for nested schema](#nestedatt--flow_mask))
* `flow_sampling` (Attributes, optional) - GsGroup Flow Sampling Parameters (see [below for nested schema](#nestedatt--flow_sampling))
* `generic_session_timeout` (Attributes, optional) - GsGroup Generic Session Timeout Parameters (see [below for nested schema](#nestedatt--generic_session_timeout))
* `gpfcp_profiles` (Attributes, optional) - Enriched CUPS Gpfcp profile aliases (see [below for nested schema](#nestedatt--gpfcp_profiles))
* `gs_group_system` (Attributes, optional) - GsGroup System Monitoring Parameters (see [below for nested schema](#nestedatt--gs_group_system))
* `gta_profiles` (Attributes, optional) - 3GPP CUPS gta profile aliases (see [below for nested schema](#nestedatt--gta_profiles))
* `gtp_control_sampling` (Attributes, optional) - GsGroup Gtp Control Sampling Parameters (see [below for nested schema](#nestedatt--gtp_control_sampling))
* `gtp_flow` (Attributes, optional) - GsGroup Gtp Flow Parameters (see [below for nested schema](#nestedatt--gtp_flow))
* `gtp_gpfcp_delay` (Attributes, optional) - GsGroup Gtp GPFCP Delay time (see [below for nested schema](#nestedatt--gtp_gpfcp_delay))
* `gtp_persistence` (Attributes, optional) - GsGroup Gtp Persistence Parameters (see [below for nested schema](#nestedatt--gtp_persistence))
* `gtp_random_sampling` (Attributes, optional) - GsGroup Gtp Random Sampling Parameters (see [below for nested schema](#nestedatt--gtp_random_sampling))
* `gtp_whitelist` (Attributes, optional) - GsGroup GTP Whitelist Parameters (see [below for nested schema](#nestedatt--gtp_whitelist))
* `health_check` (Attributes, optional) - health check configuration (see [below for nested schema](#nestedatt--health_check))
* `hsm_group` (Attributes, optional) - GsGroup Hsm Group Parameters (see [below for nested schema](#nestedatt--hsm_group))
* `ip_frag` (Attributes, optional) - GsGroup IP Fragmentation Parameters (see [below for nested schema](#nestedatt--ip_frag))
* `load_balance` (Attributes, optional) - GsGroup Load Balancing Parameters (see [below for nested schema](#nestedatt--load_balance))
* `netflow` (Attributes, optional) - GsGroup Netflow Parameters (see [below for nested schema](#nestedatt--netflow))
* `node_role` (Attributes, optional) (see [below for nested schema](#nestedatt--node_role))
* `port_throttle_sip` (Attributes, optional) - GsGroup SIP Port Throttle Parameters (see [below for nested schema](#nestedatt--port_throttle_sip))
* `resource` (Attributes, optional) - GsGroup Resource Parameters (see [below for nested schema](#nestedatt--resource))
* `rtp_ports` (Attributes, optional) - GsGroup Rtp Ports Parameters (see [below for nested schema](#nestedatt--rtp_ports))
* `sa_apf` (Attributes, optional) - GsGroup Session Aware APF Parameters (see [below for nested schema](#nestedatt--sa_apf))
* `session_logging` (Attributes, optional) - GsGroup Session Logging Configuration (see [below for nested schema](#nestedatt--session_logging))
* `sffp_profiles` (Attributes, optional) - 3GPP CUPS sffp profile aliases (see [below for nested schema](#nestedatt--sffp_profiles))
* `sip_media` (Attributes, optional) - GsGroup Sip Media Parameters (see [below for nested schema](#nestedatt--sip_media))
* `sip_ports` (Attributes, optional) - GsGroup Sip Ports Parameters (see [below for nested schema](#nestedatt--sip_ports))
* `sip_session` (Attributes, optional) - GsGroup Sip Session Parameters (see [below for nested schema](#nestedatt--sip_session))
* `sip_tcp_idle_timeout` (Attributes, optional) - GsGroup Sip Tcp Idle Parameters (see [below for nested schema](#nestedatt--sip_tcp_idle_timeout))
* `sip_whitelist` (Attributes, optional) - GsGroup SIP Whitelist Parameters (see [below for nested schema](#nestedatt--sip_whitelist))
* `ssl_decrypt` (Attributes, optional) - GsGroup SSL Decrypt Parameters (see [below for nested schema](#nestedatt--ssl_decrypt))
* `xpkt_match` (Attributes, optional) - cross packet match configuration (see [below for nested schema](#nestedatt--xpkt_match))

<a id="nestedatt--app_tcp"></a>
### Nested Schema for `app_tcp`

Optional:

* `application` (String) - To choose the action on Unknown Application Data
* `load_balance` (Boolean) - Enables TCP loadbalancing for the Tool Ports. Supports SIP & N11 only
* `tcp_control` (String) - To choose the action on TCP Control messages

<a id="nestedatt--dedup"></a>
### Nested Schema for `dedup`

Optional:

* `action` (String)
* `ip_tclass` (String)
* `ip_tos` (String)
* `tcp_seq` (String)
* `timer` (Number) - in microseconds
* `vlan` (String)

<a id="nestedatt--diameter_packet"></a>
### Nested Schema for `diameter_packet`

Optional:

* `timeout` (Number) - timeout in seconds to remove OOO or fragmented packet

<a id="nestedatt--diameter_s6_a_session"></a>
### Nested Schema for `diameter_s6_a_session`

Optional:

* `limit` (Number) - Number of sessions to allocate for Diameter S6A
* `timeout` (Number) - timeout in seconds used to clean inactive sessions

<a id="nestedatt--diameter_whitelist"></a>
### Nested Schema for `diameter_whitelist`

Required:

* `whitelist` (String) - Alias of referenced diameter Whitelist

<a id="nestedatt--eflow"></a>
### Nested Schema for `eflow`

Optional:

* `enabled` (Boolean) - Enable/Disable elephant flow detection and handling
* `interval` (Number) - time interval in seconds
* `log_enabled` (Boolean) - Enable/Disable logging of elephant flow parameters into gs logs
* `packet_count` (Number) - Number of packets to be received by a flow
* `packet_ratio` (Number) - Percentage of packets in a flow vs overall packet count

<a id="nestedatt--engine_watchdog_timer"></a>
### Nested Schema for `engine_watchdog_timer`

Optional:

* `time` (Number) - Engine Watchdog Timer value, time to restart in seconds .Valid values: 60-600. 0 is disable

<a id="nestedatt--erspan3"></a>
### Nested Schema for `erspan3`

Optional:

* `timestamp_format` (String) - erspan III tunnelDecap timestamp format

<a id="nestedatt--flow_mask"></a>
### Nested Schema for `flow_mask`

Optional:

* `enabled` (Boolean)
* `length` (Number)
* `offset` (Number)

<a id="nestedatt--flow_sampling"></a>
### Nested Schema for `flow_sampling`

Optional:

* `ip_ranges` (List of Dynamic)
* `rate` (Number) - in percent
* `timeout` (Number) - in minutes
* `type` (String)

<a id="nestedatt--generic_session_timeout"></a>
### Nested Schema for `generic_session_timeout`

Optional:

* `time` (Number) - Maximum timeout for session entry

<a id="nestedatt--gpfcp_profiles"></a>
### Nested Schema for `gpfcp_profiles`

Optional:

* `g_pfcp_profiles` (List of String)

<a id="nestedatt--gs_group_system"></a>
### Nested Schema for `gs_group_system`

Optional:

* `cpu_load_alarm_threshold` (Number) - CPU rising threshold percentage 20-99%. Once this threshold is crossed for 5 seconds, an alarm is generated

<a id="nestedatt--gta_profiles"></a>
### Nested Schema for `gta_profiles`

Optional:

* `gta_profiles` (List of String)

<a id="nestedatt--gtp_control_sampling"></a>
### Nested Schema for `gtp_control_sampling`

Optional:

* `enabled` (Boolean) - When enabled, GTP Control plane traffic for the subscribers contained within the subscriber sample defined in GTP flowsample maps will be sent to tools. When disabled, all GTP Control plane traffic for all subscribers in GTP flowsample maps will be sent to tools.

<a id="nestedatt--gtp_flow"></a>
### Nested Schema for `gtp_flow`

Optional:

* `timeout` (Number) - Session Timeout. in units of 10 minutes. Default of 48 is 8 hours

<a id="nestedatt--gtp_gpfcp_delay"></a>
### Nested Schema for `gtp_gpfcp_delay`

Optional:

* `timeout` (Number)

<a id="nestedatt--gtp_persistence"></a>
### Nested Schema for `gtp_persistence`

Optional:

* `enabled` (Boolean) - GTP Persistence Status
* `file_age_timeout` (Number) - GTP Persistence File Age Timeout(mins)
* `interval` (Number) - GTP Persistence Interval(mins)
* `restart_age_time` (Number) - GTP Persistence Restart Age Time(mins)

<a id="nestedatt--gtp_random_sampling"></a>
### Nested Schema for `gtp_random_sampling`

Optional:

* `enabled` (Boolean) - When enabled, sampling of subscriber's sessions happens in random fashion
* `interval` (Number) - Rotation Interval in multiples of 12 (hrs)

<a id="nestedatt--gtp_whitelist"></a>
### Nested Schema for `gtp_whitelist`

Required:

* `whitelist` (String) - Alias of referenced GTP Whitelist.Deprecated since H 5.12

Optional:

* `multi_whitelists` (List of String) - Alias/Aliases of referenced GTP Whitelists.

<a id="nestedatt--health_check"></a>
### Nested Schema for `health_check`

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

<a id="nestedatt--hsm_group"></a>
### Nested Schema for `hsm_group`

Optional:

* `hsm_group` (String) - alias of Hsm Group

<a id="nestedatt--ip_frag"></a>
### Nested Schema for `ip_frag`

Optional:

* `forward` (Boolean)
* `head_session_timeout` (Number) - in seconds
* `timeout` (Number) - in seconds

<a id="nestedatt--load_balance"></a>
### Nested Schema for `load_balance`

Optional:

* `failover` (Attributes) - Private class. Failover part of the GsGroup Load Balancing Parameters (see [below for nested schema](#nestedatt--load_balance--failover))
* `link_weight_type` (String) - Controls LB port weight resolution for 'weighted' load balancing (wtRoundRobin, wtLeastBw, wtLeastPktRate, wtLeastConn, wtLeastTotalTraffic,wtImsi, wtSupi). For 'assigned', explicit user-configured port weights on the port group are used. For 'speed', link speed is used and user-assigned port weights are ignored. Ex: if a port group consists of four tool ports, and one of them is 100G and the others are 10G, the 100G link will be selected about 10 times more than the 10G links
* `replicate_gtpc` (Boolean) - Enables replication of GTP control packets (GTP-c)

<a id="nestedatt--load_balance--failover"></a>
### Nested Schema for `load_balance.failover`

Optional:

* `enabled` (Boolean) - Enables or disables failover when tool ports are down or thresholds to other tool ports in the load balancing port group are exceeded. A GigaSMART application failover will occur no more than once in 30 seconds
* `threshold_lt_bw` (Number) - Mutually exclusive with 'thresholdLtPktRate'. Failover threshold for Least Bandwidth. In percent's of the maximum bandwidth of a tool port. Ex: for a 1G port, a failover threshold of 90% means that failover to another tool port occurs when the bandwidth reaches 900Mbps. The default is 80%
* `threshold_lt_pkt_rate` (Number) - Mutually exclusive with 'thresholdLtBw'. Failover threshold for Least Packet Rate. In In kilo-packets-per-second: \[500k..5M\] pps. A tool port will failover to another tool port when the packet rate is over the specified threshold, in packets per second. The default is 1M

<a id="nestedatt--netflow"></a>
### Nested Schema for `netflow`

Optional:

* `monitor` (String) - Alias of referenced Netflow Monitor

<a id="nestedatt--node_role"></a>
### Nested Schema for `node_role`

Optional:

* `mob5_g_limit` (Number) - Number of sessions to allocate for Control 5G Node
* `mob_lte_limit` (Number) - Number of sessions to allocate for LTE CPN / UPN Node
* `stand_alone_mode` (Boolean) - Enables UPN Stand-alone mode
* `type` (String)

<a id="nestedatt--port_throttle_sip"></a>
### Nested Schema for `port_throttle_sip`

Optional:

* `port_throttle` (String) - Alias of referenced Port Throttle

<a id="nestedatt--resource"></a>
### Nested Schema for `resource`

Optional:

* `buffer_asf_size` (Number) - Session aware APF buffer. Valid values: 2-5: in million sessions (0 = disabled). Changes take effect after card or system reboot
* `cpu` (Attributes) (see [below for nested schema](#nestedatt--resource--cpu))
* `hsm_ssl` (Attributes) - GsGroup Resource Hsm Ssl Parameters (see [below for nested schema](#nestedatt--resource--hsm_ssl))
* `inline_ssl` (Attributes) - Used to configure other GS apps in addition to Inline SSL on a HC1 box (see [below for nested schema](#nestedatt--resource--inline_ssl))
* `metadata` (Number) - flows in millions, how many flows to support for metadata. 0 to disable
* `packet_buffer` (Attributes) (see [below for nested schema](#nestedatt--resource--packet_buffer))
* `session_overload` (Attributes) (see [below for nested schema](#nestedatt--resource--session_overload))
* `tunnel_overload` (Attributes) (see [below for nested schema](#nestedatt--resource--tunnel_overload))
* `xpkt_match` (Attributes) - GsGroup Resource Cross Packet Match Parameters (see [below for nested schema](#nestedatt--resource--xpkt_match))

<a id="nestedatt--resource--cpu"></a>
### Nested Schema for `resource.cpu`

Optional:

* `overload_threshold` (Number) - Valid values 50 - 90%. 0 is disabled. Changes take effect after card or system reboot

<a id="nestedatt--resource--hsm_ssl"></a>
### Nested Schema for `resource.hsm_ssl`

Optional:

* `buffer` (Number) - resource for application hsm-ssl buffer in MB. 0 to disable
* `packet_buffer` (Number) - resource for application hsm-ssl packet-buffer per connection
* `session_count` (Number) - resource for application hsm-ssl buffer session count in million, 0 to disable

<a id="nestedatt--resource--inline_ssl"></a>
### Nested Schema for `resource.inline_ssl`

Optional:

* `standalone` (Boolean) - If enabled , behaves in a normal way. If disabled, can configure other GS apps along side Inline ssl with Inline ssl occupying only 50% of the memory

<a id="nestedatt--resource--packet_buffer"></a>
### Nested Schema for `resource.packet_buffer`

Optional:

* `overload_threshold` (Number) - Valid values 50 - 80%. 0 is disabled. Changes take effect after card or system reboot

<a id="nestedatt--resource--session_overload"></a>
### Nested Schema for `resource.session_overload`

Optional:

* `overload_threshold` (Number) - Session overload threshold value , Default value is 90 and 0 is disabled.

<a id="nestedatt--resource--tunnel_overload"></a>
### Nested Schema for `resource.tunnel_overload`

Optional:

* `overload_threshold` (Number) - Tunnel overload threshold value , Default value is 90 and 0 is disabled.

<a id="nestedatt--resource--xpkt_match"></a>
### Nested Schema for `resource.xpkt_match`

Optional:

* `flows` (Number) - num in 100K flows. 0 is disable

<a id="nestedatt--rtp_ports"></a>
### Nested Schema for `rtp_ports`

Optional:

* `range` (Attributes) - GsGroup Rtp Port Range Parameters (see [below for nested schema](#nestedatt--rtp_ports--range))

<a id="nestedatt--rtp_ports--range"></a>
### Nested Schema for `rtp_ports.range`

Required:

* `port` (Number)

Optional:

* `port_max` (Number) - If specified should be greater than 'port'

<a id="nestedatt--sa_apf"></a>
### Nested Schema for `sa_apf`

Optional:

* `buffer_size` (Number) - Session aware APF buffer. Valid values: 2-5: in million sessions (0 = disabled). Changes take effect after card or system reboot

<a id="nestedatt--session_logging"></a>
### Nested Schema for `session_logging`

Optional:

* `interface` (String) - Associated IP Interface
* `log_level` (String) - Log Level
* `remote_syslog_ip` (String) - Remote Syslog IP
* `remote_syslog_port` (Number) - Remote Syslog Port Number

<a id="nestedatt--sffp_profiles"></a>
### Nested Schema for `sffp_profiles`

Optional:

* `sffp_profiles` (List of String)

<a id="nestedatt--sip_media"></a>
### Nested Schema for `sip_media`

Optional:

* `timeout` (Number) - Sip media timeout value in seconds .Valid values: 30-300.

<a id="nestedatt--sip_ports"></a>
### Nested Schema for `sip_ports`

Optional:

* `ports` (List of Number) - list of TCP/UDP ports. Valid ports 1 - 65535. maximum 10 ports

<a id="nestedatt--sip_session"></a>
### Nested Schema for `sip_session`

Optional:

* `timeout` (Number) - Sip session inactivity timer, value in seconds .Valid values: 30-300.

<a id="nestedatt--sip_tcp_idle_timeout"></a>
### Nested Schema for `sip_tcp_idle_timeout`

Optional:

* `time` (Number) - Sip tcp idle timeout value in seconds .Valid values: 20-600.

<a id="nestedatt--sip_whitelist"></a>
### Nested Schema for `sip_whitelist`

Required:

* `whitelist` (String) - Alias of referenced SIP Whitelist

<a id="nestedatt--ssl_decrypt"></a>
### Nested Schema for `ssl_decrypt`

Required:

* `key_map` (String) - references one of the pre-defined 'SslDecryptionKeyMap' groups

Optional:

* `decrypt_fail_action` (String)
* `enabled` (Boolean)
* `hsm_pkcs11` (Attributes) - GsGroup Ssl Decrypt Hsm Pkcs11 Parameters (see [below for nested schema](#nestedatt--ssl_decrypt--hsm_pkcs11))
* `hsm_timeout` (Number) - in milliseconds
* `key_cache_timeout` (Number) - in seconds
* `non_ssl_traffic` (String)
* `pending_session_timeout` (Number) - in seconds
* `session_timeout` (Number) - in seconds
* `tcp_syn_timeout` (Number) - in seconds
* `ticket_cache_timeout` (Number) - in seconds

<a id="nestedatt--ssl_decrypt--hsm_pkcs11"></a>
### Nested Schema for `ssl_decrypt.hsm_pkcs11`

Optional:

* `debug_level` (Number) - hsm pkcs11 debug level
* `dynamic_object` (Boolean) - hsm pkcs11 dynamic object
* `load_sharing` (Boolean) - hsm pkcs11 load sharing

<a id="nestedatt--xpkt_match"></a>
### Nested Schema for `xpkt_match`

Optional:

* `enabled` (Boolean)

