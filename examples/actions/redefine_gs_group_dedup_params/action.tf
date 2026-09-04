action "gigavuecore_redefine_gs_group_dedup_params" "example" {
  config {
    action     = "count"
    alias      = "example"
    cluster_id = "example"
    ip_tclass  = "include"
    ip_tos     = "include"
    tcp_seq    = "include"
    timer      = 10
    vlan       = "include"
  }
}
