action "gigavuecore_redefine_snmp_throttle_config" "example" {
  config {
    throttle_config_details = [{
      interval         = 1
      notify_set       = "all"
      report_threshold = 0
      throttle_events  = [ "secondflashboot" ]
    }]
  }
}
