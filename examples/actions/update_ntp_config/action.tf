action "gigavuecore_update_ntp_config" "example" {
  config {
    auth_enabled = true
    auth_keys = [{
      key        = "example"
      key_number = 0
      trusted    = true
    }]
    clock_sync = true
    cluster_id = "example"
    enabled    = true
    ntp_servers = [{
      enabled     = true
      key_enabled = true
      key_number  = 0
      preferred   = true
      server      = "example"
      version     = "example"
    }]
    ref_server = "example"
    statuses = [{
      address       = "example"
      last_resp     = 0
      offset        = 1.0
      poll_interval = 0
      ref_clock     = "example"
      status        = "example"
      stratum       = 0
    }]
  }
}
