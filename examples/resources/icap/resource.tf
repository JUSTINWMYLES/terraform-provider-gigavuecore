resource "gigavuecore_icap" "example" {
  alias                 = "example"
  cluster_id            = "example"
  config_status         = "example"
  config_status_reasons = [ "example" ]
  gs_engines            = [ "example" ]
  gs_grp_alias          = "example"
  gsop_alias            = "example"
  health_state          = "example"
  health_state_reasons = [{
    message                               = "example"
    severity                              = "example"
    traffic_health_state_computation_type = "example"
  }]
  icap_map_alias = "example"
  icap_profile_config = {
    alias               = "example"
    cluster_id          = "example"
    exceed_action       = "example"
    http_req_buf        = 0
    inactivity_timeout  = 0
    preview             = 0
    resp_mod            = "example"
    resp_timeout        = 0
    resp_timeout_action = "example"
    server_group        = "example"
    src_max_l4_port     = 0
    src_min_l4_port     = 0
  }
  icap_server_grp_alias = "example"
  icap_servers = [{
    alias               = "example"
    cluster_id          = "example"
    comment             = "example"
    l3_address          = "example"
    l4_port             = 0
    options_service_url = "example"
    reqmod_service_url  = "example"
    respmod_service_url = "example"
  }]
  ing_alias          = "example"
  inline_networks    = [ "example" ]
  ip_interface_alias = "example"
}
