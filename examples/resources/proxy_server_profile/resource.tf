resource "gigavuecore_proxy_server_profile" "example" {
  alias = "example"
  auth_type = "example"
  comment = "example"
  password = "example"
  periodic_ping = "example"
  periodic_ping_failure_retry = 1
  periodic_ping_interval = 1
  periodic_ping_type = "example"
  port = 1
  protocol = "example"
  proxy_address = "example"
  ssl_apps = {
    cluster_name = [ "example" ]
  }
  username = "example"
}
