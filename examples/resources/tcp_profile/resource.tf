resource "gigavuecore_tcp_profile" "example" {
  alias            = "example"
  cluster_id       = "example"
  keep_alive_timer = 30
  selective_ack    = "enable"
  syn_retries      = 1
}
