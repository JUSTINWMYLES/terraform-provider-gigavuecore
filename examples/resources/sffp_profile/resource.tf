resource "gigavuecore_sffp_profile" "example" {
  alias = "example"
  profiles = [{
    ip_interface = "example"
    node_type    = "example"
    port_list    = [ 0 ]
    sx_interface = {
      ip_addresses = [ "example" ]
    }
  }]
}
