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
    app_tcp = null
    cluster_id = "example"
    dedup = null
    diameter_packet = null
    diameter_s6_a_session = null
    diameter_whitelist = null
    eflow = null
    engine_watchdog_timer = null
    erspan3 = null
    flow_mask = null
    flow_sampling = null
    generic_session_timeout = null
    gpfcp_profiles = null
    gs_group_system = null
    gta_profiles = null
    gtp_control_sampling = null
    gtp_flow = null
    gtp_gpfcp_delay = null
    gtp_persistence = null
    gtp_random_sampling = null
    gtp_whitelist = null
    health_check = null
    hsm_group = null
    ip_frag = null
    load_balance = null
    netflow = null
    node_role = null
    port_throttle_sip = null
    resource = null
    rtp_ports = null
    sa_apf = null
    session_logging = null
    sffp_profiles = null
    sip_media = null
    sip_ports = null
    sip_session = null
    sip_tcp_idle_timeout = null
    sip_whitelist = null
    ssl_decrypt = null
    xpkt_match = null
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target GS Group
* `app_tcp` (Dynamic, optional) - GsGroup TCP Parameters
* `cluster_id` (String, required) - Target Cluster ID
* `dedup` (Dynamic, optional) - GsGroup Dedup Parameters
* `diameter_packet` (Dynamic, optional) - GsGroup Diameter Packet Timeout Parameters
* `diameter_s6_a_session` (Dynamic, optional) - GsGroup Diameter s6a Session Timeout Parameters
* `diameter_whitelist` (Dynamic, optional) - GsGroup Diameter Whitelist Parameters
* `eflow` (Dynamic, optional) - GsGroup Eflow Parameters
* `engine_watchdog_timer` (Dynamic, optional) - GsGroup EngineWatchdogTimer Parameters
* `erspan3` (Dynamic, optional) - GsGroup ERSPAN III Parameters
* `flow_mask` (Dynamic, optional) - GsGroup Flow Mask Parameters
* `flow_sampling` (Dynamic, optional) - GsGroup Flow Sampling Parameters
* `generic_session_timeout` (Dynamic, optional) - GsGroup Generic Session Timeout Parameters
* `gpfcp_profiles` (Dynamic, optional) - Enriched CUPS Gpfcp profile aliases
* `gs_group_system` (Dynamic, optional) - GsGroup System Monitoring Parameters
* `gta_profiles` (Dynamic, optional) - 3GPP CUPS gta profile aliases
* `gtp_control_sampling` (Dynamic, optional) - GsGroup Gtp Control Sampling Parameters
* `gtp_flow` (Dynamic, optional) - GsGroup Gtp Flow Parameters
* `gtp_gpfcp_delay` (Dynamic, optional) - GsGroup Gtp GPFCP Delay time
* `gtp_persistence` (Dynamic, optional) - GsGroup Gtp Persistence Parameters
* `gtp_random_sampling` (Dynamic, optional) - GsGroup Gtp Random Sampling Parameters
* `gtp_whitelist` (Dynamic, optional) - GsGroup GTP Whitelist Parameters
* `health_check` (Dynamic, optional) - health check configuration
* `hsm_group` (Dynamic, optional) - GsGroup Hsm Group Parameters
* `ip_frag` (Dynamic, optional) - GsGroup IP Fragmentation Parameters
* `load_balance` (Dynamic, optional) - GsGroup Load Balancing Parameters
* `netflow` (Dynamic, optional) - GsGroup Netflow Parameters
* `node_role` (Dynamic, optional)
* `port_throttle_sip` (Dynamic, optional) - GsGroup SIP Port Throttle Parameters
* `resource` (Dynamic, optional) - GsGroup Resource Parameters
* `rtp_ports` (Dynamic, optional) - GsGroup Rtp Ports Parameters
* `sa_apf` (Dynamic, optional) - GsGroup Session Aware APF Parameters
* `session_logging` (Dynamic, optional) - GsGroup Session Logging Configuration
* `sffp_profiles` (Dynamic, optional) - 3GPP CUPS sffp profile aliases
* `sip_media` (Dynamic, optional) - GsGroup Sip Media Parameters
* `sip_ports` (Dynamic, optional) - GsGroup Sip Ports Parameters
* `sip_session` (Dynamic, optional) - GsGroup Sip Session Parameters
* `sip_tcp_idle_timeout` (Dynamic, optional) - GsGroup Sip Tcp Idle Parameters
* `sip_whitelist` (Dynamic, optional) - GsGroup SIP Whitelist Parameters
* `ssl_decrypt` (Dynamic, optional) - GsGroup SSL Decrypt Parameters
* `xpkt_match` (Dynamic, optional) - cross packet match configuration
