action "gigavuecore_delete_fm_ntp_server" "example" {
  config {
    auth_required      = true
    body_fm_ip         = "example"
    fm_ip              = "example"
    is_user_ntp_server = true
    ntp_auth           = null
    server_host        = "example"
    server_status      = null
    version            = 1
  }
}
