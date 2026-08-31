resource "gigavuecore_exporter" "example" {
  alias       = "example"
  cluster_id  = "example"
  description = "example"
  destination = {
    l3 = {
      ip = {
        dscp = 0
        ttl  = 1
        ver4 = "example"
        ver6 = "example"
      }
      protocol = "ipv4"
    }
    l4 = {
      port     = 1
      protocol = "tcp"
    }
  }
  gs_group_associated = [ "example" ]
  source = {
    interface = "example"
    l4_port   = 1
  }
  ssl_profile = "example"
  status      = "active"
  tags = [{
    tag_key    = "example"
    tag_values = [ "example" ]
  }]
  tcp_profile = "example"
  type        = "mobility-cups"
}
