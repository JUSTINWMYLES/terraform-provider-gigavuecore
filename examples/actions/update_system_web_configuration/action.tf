action "gigavuecore_update_system_web_configuration" "example" {
  config {
    auto_logout_timeout    = 0
    cluster_id             = "example"
    cluster_name           = "example"
    enable_https           = true
    http_port              = 0
    https_port             = 0
    server_ssl_min_version = "tls1"
    session_renewal        = 0
    session_timeout        = 0
  }
}
