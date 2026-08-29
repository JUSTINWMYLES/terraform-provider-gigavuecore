action "gigavuecore_update_snmp_throttle_config_system_snmp_throttle" "example" {
  config {
    cluster_id = "example"
    throttle_config_details = [{
      interval         = 0
      notify_set       = "example"
      report_threshold = 0
      throttle_events  = [ "example" ]
    }]
  }
}
