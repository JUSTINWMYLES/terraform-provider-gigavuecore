resource "gigavuecore_listener" "example" {
  alias               = "example"
  description         = "example"
  gs_group_associated = [ "example" ]
  ip_interface        = [ "example" ]
  l3 = {
    dscp     = 1
    protocol = "example"
    ttl      = 1
  }
  l4 = {
    port_list = [ 1 ]
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
