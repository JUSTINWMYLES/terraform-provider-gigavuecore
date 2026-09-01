resource "gigavuecore_listener" "example" {
  alias               = "example"
  cluster_id          = "example"
  description         = "example"
  gs_group_associated = ["example"]
  ip_interface        = ["example"]
  l3 = {
    dscp     = 0
    protocol = "ipv4"
    ttl      = 1
  }
  l4 = {
    port_list = [1]
    protocol  = "tcp"
  }
  mode        = "promiscuous"
  ssl_profile = "example"
  status      = "active"
  tags = [{
    tag_key    = "example"
    tag_values = ["example"]
  }]
  tcp_profile = "example"
  type        = "gtp-cups"
}
