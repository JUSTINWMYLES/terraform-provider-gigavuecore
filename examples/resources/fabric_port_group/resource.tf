resource "gigavuecore_fabric_port_group" "example" {
  alias        = "example"
  comment      = "example"
  port_list    = [ "example" ]
  port_weights = [ 1 ]
  smart_lb     = true
  tags = [{
    tag_key    = "example"
    tag_values = [ "example" ]
  }]
}
