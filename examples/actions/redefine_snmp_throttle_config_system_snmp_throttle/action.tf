action "gigavuecore_redefine_snmp_throttle_config_system_snmp_throttle" "example" {
  config {
    cluster_id = "example"
    throttle_config_details = [{
      interval         = 1
      notify_set       = "all"
      report_threshold = 0
      throttle_events  = ["secondflashboot"]
    }]
  }
}
