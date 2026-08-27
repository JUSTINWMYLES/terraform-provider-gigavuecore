resource "gigavuecore_interface" "example" {
  alias             = "example"
  attach            = [ "example" ]
  cluster_id        = "example"
  comment           = "example"
  gateway           = "example"
  gs_groups         = [ "example" ]
  hw_address        = "example"
  ip_address        = "example"
  ip_mask           = "example"
  mtu               = 1
  netflow_exporters = [ "example" ]
  tags = [{
    tag_key    = "example"
    tag_values = [ "example" ]
  }]
}
