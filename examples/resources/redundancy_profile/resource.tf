resource "gigavuecore_redundancy_profile" "example" {
  alias           = "example"
  cluster_id      = "example"
  protection_role = "suspended"
  signaling_port  = "example"
}
