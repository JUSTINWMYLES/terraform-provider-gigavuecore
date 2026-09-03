action "gigavuecore_redefine_gs_group_load_balance_params" "example" {
  config {
    alias      = "example"
    cluster_id = "example"
    failover = {
      enabled               = true
      threshold_lt_bw       = 50
      threshold_lt_pkt_rate = 500
    }
    link_weight_type = "speed"
    replicate_gtpc   = true
  }
}
