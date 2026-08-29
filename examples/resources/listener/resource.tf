resource "gigavuecore_listener" "example" {
  alias               = "example"
  cluster_id          = "example"
  description         = "example"
  gs_group_associated = [ "example" ]
  ip_interface        = [ "example" ]
  l3 = {
    dscp     = 0
    protocol = "example"
    ttl      = 0
  }
  l4 = {
    port_list = [ 0 ]
    protocol  = "example"
  }
  mode        = "example"
  ssl_profile = "example"
  status      = "example"
  tags = [{
    tag_key    = "example"
    tag_values = [ "example" ]
  }]
  tcp_profile = "example"
  type        = "example"
}
