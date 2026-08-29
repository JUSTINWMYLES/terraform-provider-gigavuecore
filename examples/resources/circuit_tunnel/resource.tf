resource "gigavuecore_circuit_tunnel" "example" {
  alias       = "example"
  attach      = [ "example" ]
  circuit_ids = [ 0 ]
  cluster_id  = "example"
  comment     = "example"
  dip_address = "example"
  l4_src_port = 0
  mode        = "example"
  type        = "example"
}
