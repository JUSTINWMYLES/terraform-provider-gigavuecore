data "gigavuecore_get_all_allocation_maps" "example" {
  encode            = true
  exclude_host_name = "example"
  exclude_ip        = "example"
}
