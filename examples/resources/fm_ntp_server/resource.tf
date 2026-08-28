resource "gigavuecore_fm_ntp_server" "example" {
  auth_required      = true
  fm_ip              = "example"
  is_user_ntp_server = true
  ntp_auth = {
    key   = "example"
    type  = "example"
    value = "example"
  }
  server_host = "example"
  server_status = {
    offset        = 1
    poll_interval = "example"
    status        = "example"
    stratum       = "example"
  }
  version = 1
}
