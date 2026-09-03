resource "gigavuecore_gta_profile" "example" {
  alias        = "example"
  control_node = "example"
  core_network_nodes = {
    address_list = {
      value = ["example"]
    }
    address_range = {
      ip_ranges = [{
        max_value = "example"
        value     = "example"
      }]
    }
    address_subnet = {
      values = ["example"]
    }
  }
  dst_port  = 0
  src_port  = 0
  user_node = "example"
}
