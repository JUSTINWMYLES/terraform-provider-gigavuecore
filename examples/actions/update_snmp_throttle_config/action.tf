action "gigavuecore_update_snmp_throttle_config" "example" {
  config {
    throttle_config_details = [{
      interval         = 0
      notify_set       = "example"
      report_threshold = 0
      throttle_events  = [ "example" ]
    }]
  }
}
