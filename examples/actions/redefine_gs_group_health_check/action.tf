action "gigavuecore_redefine_gs_group_health_check" "example" {
  config {
    action          = "pass"
    alias           = "example"
    dst_port        = 1
    enabled         = true
    interval        = 5
    protocol        = "icmp"
    rcv_port        = 1
    retries         = 1
    round_trip_time = 1
    src_port        = 1
  }
}
