resource "gigavuecore_vxlan_group" "example" {
  alias     = "example"
  box_id    = "example"
  comment   = "example"
  vxlan_ids = [ 1 ]
}
