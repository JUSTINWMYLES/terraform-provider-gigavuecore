resource "gigavuecore_proxy_server_profile" "example" {
  alias                       = "example"
  auth_type                   = "none"
  comment                     = "example"
  password                    = "example"
  periodic_ping               = "enable"
  periodic_ping_failure_retry = 1
  periodic_ping_interval      = 1
  periodic_ping_type          = "http-connect"
  port                        = 1
  protocol                    = "http"
  proxy_address               = "example"
  ssl_apps = {
    cluster_name = [ "example" ]
  }
  username = "example"
}
