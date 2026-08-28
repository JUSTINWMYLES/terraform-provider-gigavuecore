resource "gigavuecore_exporter" "example" {
  alias       = "example"
  cluster_id  = "example"
  description = "example"
  destination = {
    l3 = {
      ip = {
        dscp = 1
        ttl  = 1
        ver4 = "example"
        ver6 = "example"
      }
      protocol = "example"
    }
    l4 = {
      port     = 1
      protocol = "example"
    }
  }
  gs_group_associated = [ "example" ]
  source = {
    interface = "example"
    l4_port   = 1
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
