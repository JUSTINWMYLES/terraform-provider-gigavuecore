---
page_title: "gigavuecore_load_all_gs_groups Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all GS Groups
---

# gigavuecore_load_all_gs_groups Data Source

Load all GS Groups

## Example Usage

```terraform
data "gigavuecore_load_all_gs_groups" "example" {
  cluster_id = "example"
  page       = "example"
  sort       = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `alias` (String) - GS Group alias
* `cluster_id` (String) - id of the defining cluster
* `hash` (String) - gsgroup hashing parameters
* `health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--items--health_state_reasons))
* `params` (Attributes) - GsGroup Parameters (see [below for nested schema](#nestedatt--items--params))
* `ports` (List of String) - list of the member GigaSMART e-ports
<a id="nestedatt--items--health_state_reasons"></a>
### Nested Schema for `items.health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type
<a id="nestedatt--items--params"></a>
### Nested Schema for `items.params`

Read-Only:

* `app_tcp` (Attributes) - GsGroup TCP Parameters (see [below for nested schema](#nestedatt--items--params--app_tcp))
* `dedup` (Attributes) - GsGroup Dedup Parameters (see [below for nested schema](#nestedatt--items--params--dedup))
* `diameter_packet` (Attributes) - GsGroup Diameter Packet Timeout Parameters (see [below for nested schema](#nestedatt--items--params--diameter_packet))
* `diameter_s6_a_session` (Attributes) - GsGroup Diameter s6a Session Timeout Parameters (see [below for nested schema](#nestedatt--items--params--diameter_s6_a_session))
* `diameter_whitelist` (Attributes) - GsGroup Diameter Whitelist Parameters (see [below for nested schema](#nestedatt--items--params--diameter_whitelist))
* `eflow` (Attributes) - GsGroup Eflow Parameters (see [below for nested schema](#nestedatt--items--params--eflow))
* `engine_watchdog_timer` (Attributes) - GsGroup EngineWatchdogTimer Parameters (see [below for nested schema](#nestedatt--items--params--engine_watchdog_timer))
* `erspan3` (Attributes) - GsGroup ERSPAN III Parameters (see [below for nested schema](#nestedatt--items--params--erspan3))
* `flow_mask` (Attributes) - GsGroup Flow Mask Parameters (see [below for nested schema](#nestedatt--items--params--flow_mask))
* `flow_sampling` (Attributes) - GsGroup Flow Sampling Parameters (see [below for nested schema](#nestedatt--items--params--flow_sampling))
* `generic_session_timeout` (Attributes) - GsGroup Generic Session Timeout Parameters (see [below for nested schema](#nestedatt--items--params--generic_session_timeout))
* `gpfcp_profiles` (Attributes) - Enriched CUPS Gpfcp profile aliases (see [below for nested schema](#nestedatt--items--params--gpfcp_profiles))
* `gs_group_system` (Attributes) - GsGroup System Monitoring Parameters (see [below for nested schema](#nestedatt--items--params--gs_group_system))
* `gta_profiles` (Attributes) - 3GPP CUPS gta profile aliases (see [below for nested schema](#nestedatt--items--params--gta_profiles))
* `gtp_control_sampling` (Attributes) - GsGroup Gtp Control Sampling Parameters (see [below for nested schema](#nestedatt--items--params--gtp_control_sampling))
* `gtp_flow` (Attributes) - GsGroup Gtp Flow Parameters (see [below for nested schema](#nestedatt--items--params--gtp_flow))
* `gtp_gpfcp_delay` (Attributes) - GsGroup Gtp GPFCP Delay time (see [below for nested schema](#nestedatt--items--params--gtp_gpfcp_delay))
* `gtp_persistence` (Attributes) - GsGroup Gtp Persistence Parameters (see [below for nested schema](#nestedatt--items--params--gtp_persistence))
* `gtp_random_sampling` (Attributes) - GsGroup Gtp Random Sampling Parameters (see [below for nested schema](#nestedatt--items--params--gtp_random_sampling))
* `gtp_whitelist` (Attributes) - GsGroup GTP Whitelist Parameters (see [below for nested schema](#nestedatt--items--params--gtp_whitelist))
* `health_check` (Attributes) - health check configuration (see [below for nested schema](#nestedatt--items--params--health_check))
* `hsm_group` (Attributes) - GsGroup Hsm Group Parameters (see [below for nested schema](#nestedatt--items--params--hsm_group))
* `ip_frag` (Attributes) - GsGroup IP Fragmentation Parameters (see [below for nested schema](#nestedatt--items--params--ip_frag))
* `load_balance` (Attributes) - GsGroup Load Balancing Parameters (see [below for nested schema](#nestedatt--items--params--load_balance))
* `netflow` (Attributes) - GsGroup Netflow Parameters (see [below for nested schema](#nestedatt--items--params--netflow))
* `node_role` (Attributes) (see [below for nested schema](#nestedatt--items--params--node_role))
* `port_throttle_sip` (Attributes) - GsGroup SIP Port Throttle Parameters (see [below for nested schema](#nestedatt--items--params--port_throttle_sip))
* `resource` (Attributes) - GsGroup Resource Parameters (see [below for nested schema](#nestedatt--items--params--resource))
* `rtp_ports` (Attributes) - GsGroup Rtp Ports Parameters (see [below for nested schema](#nestedatt--items--params--rtp_ports))
* `sa_apf` (Attributes) - GsGroup Session Aware APF Parameters (see [below for nested schema](#nestedatt--items--params--sa_apf))
* `session_logging` (Attributes) - GsGroup Session Logging Configuration (see [below for nested schema](#nestedatt--items--params--session_logging))
* `sffp_profiles` (Attributes) - 3GPP CUPS sffp profile aliases (see [below for nested schema](#nestedatt--items--params--sffp_profiles))
* `sip_media` (Attributes) - GsGroup Sip Media Parameters (see [below for nested schema](#nestedatt--items--params--sip_media))
* `sip_ports` (Attributes) - GsGroup Sip Ports Parameters (see [below for nested schema](#nestedatt--items--params--sip_ports))
* `sip_session` (Attributes) - GsGroup Sip Session Parameters (see [below for nested schema](#nestedatt--items--params--sip_session))
* `sip_tcp_idle_timeout` (Attributes) - GsGroup Sip Tcp Idle Parameters (see [below for nested schema](#nestedatt--items--params--sip_tcp_idle_timeout))
* `sip_whitelist` (Attributes) - GsGroup SIP Whitelist Parameters (see [below for nested schema](#nestedatt--items--params--sip_whitelist))
* `ssl_decrypt` (Attributes) - GsGroup SSL Decrypt Parameters (see [below for nested schema](#nestedatt--items--params--ssl_decrypt))
* `xpkt_match` (Attributes) - cross packet match configuration (see [below for nested schema](#nestedatt--items--params--xpkt_match))
<a id="nestedatt--items--params--app_tcp"></a>
### Nested Schema for `items.params.app_tcp`

Read-Only:

* `application` (String) - To choose the action on Unknown Application Data
* `load_balance` (Boolean) - Enables TCP loadbalancing for the Tool Ports. Supports SIP & N11 only
* `tcp_control` (String) - To choose the action on TCP Control messages
<a id="nestedatt--items--params--dedup"></a>
### Nested Schema for `items.params.dedup`

Read-Only:

* `action` (String)
* `ip_tclass` (String)
* `ip_tos` (String)
* `tcp_seq` (String)
* `timer` (Number) - in microseconds
* `vlan` (String)
<a id="nestedatt--items--params--diameter_packet"></a>
### Nested Schema for `items.params.diameter_packet`

Read-Only:

* `timeout` (Number) - timeout in seconds to remove OOO or fragmented packet
<a id="nestedatt--items--params--diameter_s6_a_session"></a>
### Nested Schema for `items.params.diameter_s6_a_session`

Read-Only:

* `limit` (Number) - Number of sessions to allocate for Diameter S6A
* `timeout` (Number) - timeout in seconds used to clean inactive sessions
<a id="nestedatt--items--params--diameter_whitelist"></a>
### Nested Schema for `items.params.diameter_whitelist`

Read-Only:

* `whitelist` (String) - Alias of referenced diameter Whitelist
<a id="nestedatt--items--params--eflow"></a>
### Nested Schema for `items.params.eflow`

Read-Only:

* `enabled` (Boolean) - Enable/Disable elephant flow detection and handling
* `interval` (Number) - time interval in seconds
* `log_enabled` (Boolean) - Enable/Disable logging of elephant flow parameters into gs logs
* `packet_count` (Number) - Number of packets to be received by a flow
* `packet_ratio` (Number) - Percentage of packets in a flow vs overall packet count
<a id="nestedatt--items--params--engine_watchdog_timer"></a>
### Nested Schema for `items.params.engine_watchdog_timer`

Read-Only:

* `time` (Number) - Engine Watchdog Timer value, time to restart in seconds .Valid values: 60-600. 0 is disable
<a id="nestedatt--items--params--erspan3"></a>
### Nested Schema for `items.params.erspan3`

Read-Only:

* `timestamp_format` (String) - erspan III tunnelDecap timestamp format
<a id="nestedatt--items--params--flow_mask"></a>
### Nested Schema for `items.params.flow_mask`

Read-Only:

* `enabled` (Boolean)
* `length` (Number)
* `offset` (Number)
<a id="nestedatt--items--params--flow_sampling"></a>
### Nested Schema for `items.params.flow_sampling`

Read-Only:

* `ip_ranges` (List of Dynamic)
* `rate` (Number) - in percent
* `timeout` (Number) - in minutes
* `type` (String)
<a id="nestedatt--items--params--generic_session_timeout"></a>
### Nested Schema for `items.params.generic_session_timeout`

Read-Only:

* `time` (Number) - Maximum timeout for session entry
<a id="nestedatt--items--params--gpfcp_profiles"></a>
### Nested Schema for `items.params.gpfcp_profiles`

Read-Only:

* `g_pfcp_profiles` (List of String)
<a id="nestedatt--items--params--gs_group_system"></a>
### Nested Schema for `items.params.gs_group_system`

Read-Only:

* `cpu_load_alarm_threshold` (Number) - CPU rising threshold percentage 20-99%. Once this threshold is crossed for 5 seconds, an alarm is generated
<a id="nestedatt--items--params--gta_profiles"></a>
### Nested Schema for `items.params.gta_profiles`

Read-Only:

* `gta_profiles` (List of String)
<a id="nestedatt--items--params--gtp_control_sampling"></a>
### Nested Schema for `items.params.gtp_control_sampling`

Read-Only:

* `enabled` (Boolean) - When enabled, GTP Control plane traffic for the subscribers contained within the subscriber sample defined in GTP flowsample maps will be sent to tools. When disabled, all GTP Control plane traffic for all subscribers in GTP flowsample maps will be sent to tools.
<a id="nestedatt--items--params--gtp_flow"></a>
### Nested Schema for `items.params.gtp_flow`

Read-Only:

* `timeout` (Number) - Session Timeout. in units of 10 minutes. Default of 48 is 8 hours
<a id="nestedatt--items--params--gtp_gpfcp_delay"></a>
### Nested Schema for `items.params.gtp_gpfcp_delay`

Read-Only:

* `timeout` (Number)
<a id="nestedatt--items--params--gtp_persistence"></a>
### Nested Schema for `items.params.gtp_persistence`

Read-Only:

* `enabled` (Boolean) - GTP Persistence Status
* `file_age_timeout` (Number) - GTP Persistence File Age Timeout(mins)
* `interval` (Number) - GTP Persistence Interval(mins)
* `restart_age_time` (Number) - GTP Persistence Restart Age Time(mins)
<a id="nestedatt--items--params--gtp_random_sampling"></a>
### Nested Schema for `items.params.gtp_random_sampling`

Read-Only:

* `enabled` (Boolean) - When enabled, sampling of subscriber's sessions happens in random fashion
* `interval` (Number) - Rotation Interval in multiples of 12 (hrs)
<a id="nestedatt--items--params--gtp_whitelist"></a>
### Nested Schema for `items.params.gtp_whitelist`

Read-Only:

* `multi_whitelists` (List of String) - Alias/Aliases of referenced GTP Whitelists.
* `whitelist` (String) - Alias of referenced GTP Whitelist.Deprecated since H 5.12
<a id="nestedatt--items--params--health_check"></a>
### Nested Schema for `items.params.health_check`

Read-Only:

* `action` (String)
* `dst_port` (Number)
* `enabled` (Boolean)
* `interval` (Number)
* `protocol` (String)
* `rcv_port` (Number)
* `retries` (Number)
* `round_trip_time` (Number)
* `src_port` (Number)
<a id="nestedatt--items--params--hsm_group"></a>
### Nested Schema for `items.params.hsm_group`

Read-Only:

* `hsm_group` (String) - alias of Hsm Group
<a id="nestedatt--items--params--ip_frag"></a>
### Nested Schema for `items.params.ip_frag`

Read-Only:

* `forward` (Boolean)
* `head_session_timeout` (Number) - in seconds
* `timeout` (Number) - in seconds
<a id="nestedatt--items--params--load_balance"></a>
### Nested Schema for `items.params.load_balance`

Read-Only:

* `failover` (Attributes) - Private class. Failover part of the GsGroup Load Balancing Parameters (see [below for nested schema](#nestedatt--items--params--load_balance--failover))
* `link_weight_type` (String) - Controls LB port weight resolution for 'weighted' load balancing (wtRoundRobin, wtLeastBw, wtLeastPktRate, wtLeastConn, wtLeastTotalTraffic,wtImsi, wtSupi). For 'assigned', explicit user-configured port weights on the port group are used. For 'speed', link speed is used and user-assigned port weights are ignored. Ex: if a port group consists of four tool ports, and one of them is 100G and the others are 10G, the 100G link will be selected about 10 times more than the 10G links
* `replicate_gtpc` (Boolean) - Enables replication of GTP control packets (GTP-c)
<a id="nestedatt--items--params--load_balance--failover"></a>
### Nested Schema for `items.params.load_balance.failover`

Read-Only:

* `enabled` (Boolean) - Enables or disables failover when tool ports are down or thresholds to other tool ports in the load balancing port group are exceeded. A GigaSMART application failover will occur no more than once in 30 seconds
* `threshold_lt_bw` (Number) - Mutually exclusive with 'thresholdLtPktRate'. Failover threshold for Least Bandwidth. In percent's of the maximum bandwidth of a tool port. Ex: for a 1G port, a failover threshold of 90% means that failover to another tool port occurs when the bandwidth reaches 900Mbps. The default is 80%
* `threshold_lt_pkt_rate` (Number) - Mutually exclusive with 'thresholdLtBw'. Failover threshold for Least Packet Rate. In In kilo-packets-per-second: \[500k..5M\] pps. A tool port will failover to another tool port when the packet rate is over the specified threshold, in packets per second. The default is 1M
<a id="nestedatt--items--params--netflow"></a>
### Nested Schema for `items.params.netflow`

Read-Only:

* `monitor` (String) - Alias of referenced Netflow Monitor
<a id="nestedatt--items--params--node_role"></a>
### Nested Schema for `items.params.node_role`

Read-Only:

* `mob5_g_limit` (Number) - Number of sessions to allocate for Control 5G Node
* `mob_lte_limit` (Number) - Number of sessions to allocate for LTE CPN / UPN Node
* `stand_alone_mode` (Boolean) - Enables UPN Stand-alone mode
* `type` (String)
<a id="nestedatt--items--params--port_throttle_sip"></a>
### Nested Schema for `items.params.port_throttle_sip`

Read-Only:

* `port_throttle` (String) - Alias of referenced Port Throttle
<a id="nestedatt--items--params--resource"></a>
### Nested Schema for `items.params.resource`

Read-Only:

* `buffer_asf_size` (Number) - Session aware APF buffer. Valid values: 2-5: in million sessions (0 = disabled). Changes take effect after card or system reboot
* `cpu` (Attributes) (see [below for nested schema](#nestedatt--items--params--resource--cpu))
* `hsm_ssl` (Attributes) - GsGroup Resource Hsm Ssl Parameters (see [below for nested schema](#nestedatt--items--params--resource--hsm_ssl))
* `inline_ssl` (Attributes) - Used to configure other GS apps in addition to Inline SSL on a HC1 box (see [below for nested schema](#nestedatt--items--params--resource--inline_ssl))
* `metadata` (Number) - flows in millions, how many flows to support for metadata. 0 to disable
* `packet_buffer` (Attributes) (see [below for nested schema](#nestedatt--items--params--resource--packet_buffer))
* `session_overload` (Attributes) (see [below for nested schema](#nestedatt--items--params--resource--session_overload))
* `tunnel_overload` (Attributes) (see [below for nested schema](#nestedatt--items--params--resource--tunnel_overload))
* `xpkt_match` (Attributes) - GsGroup Resource Cross Packet Match Parameters (see [below for nested schema](#nestedatt--items--params--resource--xpkt_match))
<a id="nestedatt--items--params--resource--cpu"></a>
### Nested Schema for `items.params.resource.cpu`

Read-Only:

* `overload_threshold` (Number) - Valid values 50 - 90%. 0 is disabled. Changes take effect after card or system reboot
<a id="nestedatt--items--params--resource--hsm_ssl"></a>
### Nested Schema for `items.params.resource.hsm_ssl`

Read-Only:

* `buffer` (Number) - resource for application hsm-ssl buffer in MB. 0 to disable
* `packet_buffer` (Number) - resource for application hsm-ssl packet-buffer per connection
* `session_count` (Number) - resource for application hsm-ssl buffer session count in million, 0 to disable
<a id="nestedatt--items--params--resource--inline_ssl"></a>
### Nested Schema for `items.params.resource.inline_ssl`

Read-Only:

* `standalone` (Boolean) - If enabled , behaves in a normal way. If disabled, can configure other GS apps along side Inline ssl with Inline ssl occupying only 50% of the memory
<a id="nestedatt--items--params--resource--packet_buffer"></a>
### Nested Schema for `items.params.resource.packet_buffer`

Read-Only:

* `overload_threshold` (Number) - Valid values 50 - 80%. 0 is disabled. Changes take effect after card or system reboot
<a id="nestedatt--items--params--resource--session_overload"></a>
### Nested Schema for `items.params.resource.session_overload`

Read-Only:

* `overload_threshold` (Number) - Session overload threshold value , Default value is 90 and 0 is disabled.
<a id="nestedatt--items--params--resource--tunnel_overload"></a>
### Nested Schema for `items.params.resource.tunnel_overload`

Read-Only:

* `overload_threshold` (Number) - Tunnel overload threshold value , Default value is 90 and 0 is disabled.
<a id="nestedatt--items--params--resource--xpkt_match"></a>
### Nested Schema for `items.params.resource.xpkt_match`

Read-Only:

* `flows` (Number) - num in 100K flows. 0 is disable
<a id="nestedatt--items--params--rtp_ports"></a>
### Nested Schema for `items.params.rtp_ports`

Read-Only:

* `range` (Attributes) - GsGroup Rtp Port Range Parameters (see [below for nested schema](#nestedatt--items--params--rtp_ports--range))
<a id="nestedatt--items--params--rtp_ports--range"></a>
### Nested Schema for `items.params.rtp_ports.range`

Read-Only:

* `port` (Number)
* `port_max` (Number) - If specified should be greater than 'port'
<a id="nestedatt--items--params--sa_apf"></a>
### Nested Schema for `items.params.sa_apf`

Read-Only:

* `buffer_size` (Number) - Session aware APF buffer. Valid values: 2-5: in million sessions (0 = disabled). Changes take effect after card or system reboot
<a id="nestedatt--items--params--session_logging"></a>
### Nested Schema for `items.params.session_logging`

Read-Only:

* `interface` (String) - Associated IP Interface
* `log_level` (String) - Log Level
* `remote_syslog_ip` (String) - Remote Syslog IP
* `remote_syslog_port` (Number) - Remote Syslog Port Number
<a id="nestedatt--items--params--sffp_profiles"></a>
### Nested Schema for `items.params.sffp_profiles`

Read-Only:

* `sffp_profiles` (List of String)
<a id="nestedatt--items--params--sip_media"></a>
### Nested Schema for `items.params.sip_media`

Read-Only:

* `timeout` (Number) - Sip media timeout value in seconds .Valid values: 30-300.
<a id="nestedatt--items--params--sip_ports"></a>
### Nested Schema for `items.params.sip_ports`

Read-Only:

* `ports` (List of Number) - list of TCP/UDP ports. Valid ports 1 - 65535. maximum 10 ports
<a id="nestedatt--items--params--sip_session"></a>
### Nested Schema for `items.params.sip_session`

Read-Only:

* `timeout` (Number) - Sip session inactivity timer, value in seconds .Valid values: 30-300.
<a id="nestedatt--items--params--sip_tcp_idle_timeout"></a>
### Nested Schema for `items.params.sip_tcp_idle_timeout`

Read-Only:

* `time` (Number) - Sip tcp idle timeout value in seconds .Valid values: 20-600.
<a id="nestedatt--items--params--sip_whitelist"></a>
### Nested Schema for `items.params.sip_whitelist`

Read-Only:

* `whitelist` (String) - Alias of referenced SIP Whitelist
<a id="nestedatt--items--params--ssl_decrypt"></a>
### Nested Schema for `items.params.ssl_decrypt`

Read-Only:

* `decrypt_fail_action` (String)
* `enabled` (Boolean)
* `hsm_pkcs11` (Attributes) - GsGroup Ssl Decrypt Hsm Pkcs11 Parameters (see [below for nested schema](#nestedatt--items--params--ssl_decrypt--hsm_pkcs11))
* `hsm_timeout` (Number) - in milliseconds
* `key_cache_timeout` (Number) - in seconds
* `key_map` (String) - references one of the pre-defined 'SslDecryptionKeyMap' groups
* `non_ssl_traffic` (String)
* `pending_session_timeout` (Number) - in seconds
* `session_timeout` (Number) - in seconds
* `tcp_syn_timeout` (Number) - in seconds
* `ticket_cache_timeout` (Number) - in seconds
<a id="nestedatt--items--params--ssl_decrypt--hsm_pkcs11"></a>
### Nested Schema for `items.params.ssl_decrypt.hsm_pkcs11`

Read-Only:

* `debug_level` (Number) - hsm pkcs11 debug level
* `dynamic_object` (Boolean) - hsm pkcs11 dynamic object
* `load_sharing` (Boolean) - hsm pkcs11 load sharing
<a id="nestedatt--items--params--xpkt_match"></a>
### Nested Schema for `items.params.xpkt_match`

Read-Only:

* `enabled` (Boolean)

