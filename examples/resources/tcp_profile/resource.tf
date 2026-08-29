resource "gigavuecore_tcp_profile" "example" {
  alias            = "example"
  cluster_id       = "example"
  keep_alive_timer = 0
  selective_ack    = "example"
  syn_retries      = 0
}
