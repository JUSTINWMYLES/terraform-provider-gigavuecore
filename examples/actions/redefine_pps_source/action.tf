action "gigavuecore_redefine_pps_source" "example" {
  config {
    cluster_id       = "example"
    pps_offset       = 0
    pps_source_admin = "extCoaxial"
    pps_source_oper  = "internal"
  }
}
