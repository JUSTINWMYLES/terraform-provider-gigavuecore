action "gigavuecore_redefine_fm_ntp_server" "example" {
  config {
    auth_required      = true
    body_fm_ip         = "example"
    fm_ip              = "example"
    is_user_ntp_server = true
    ntp_auth = {
      key   = "example"
      type  = "example"
      value = "example"
    }
    server_host = "example"
    server_status = {
      offset        = 0
      poll_interval = "example"
      status        = "example"
      stratum       = "example"
    }
    version = 0
  }
}
