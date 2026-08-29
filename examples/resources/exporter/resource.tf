resource "gigavuecore_exporter" "example" {
  alias       = "example"
  cluster_id  = "example"
  description = "example"
  destination = {
    l3 = {
      ip = {
        dscp = 0
        ttl  = 0
        ver4 = "example"
        ver6 = "example"
      }
      protocol = "example"
    }
    l4 = {
      port     = 0
      protocol = "example"
    }
  }
  gs_group_associated = [ "example" ]
  source = {
    interface = "example"
    l4_port   = 0
  }
  ssl_profile = "example"
  status      = "example"
  tags = [{
    tag_key    = "example"
    tag_values = [ "example" ]
  }]
  tcp_profile = "example"
  type        = "example"
}
