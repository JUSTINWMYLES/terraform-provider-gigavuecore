resource "gigavuecore_gpfcp_profile" "example" {
  alias      = "example"
  cluster_id = "example"
  comment    = "example"
  g_profiles = [{
    comment = "example"
    g_interface = {
      ip_addresses = [ "example" ]
    }
    ip_interface = "example"
    node_type    = "example"
    port_list    = [ 0 ]
  }]
}
