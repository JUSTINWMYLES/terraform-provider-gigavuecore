action "gigavuecore_redefine_gs_group_tcp_params" "example" {
  config {
    alias        = "example"
    application  = "broadcast"
    load_balance = true
    tcp_control  = "broadcast"
  }
}
