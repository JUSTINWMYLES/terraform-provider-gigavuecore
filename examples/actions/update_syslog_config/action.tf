action "gigavuecore_update_syslog_config" "example" {
  config {
    cluster_id   = "example"
    cluster_name = "example"
    syslog_config_list = [{
      device_ip    = "example"
      log_severity = "example"
      target_hosts = [{
        log_severity      = "example"
        port              = 0
        server            = "example"
        ssh_enabled       = true
        streaming_enabled = true
        username          = "example"
      }]
    }]
  }
}
