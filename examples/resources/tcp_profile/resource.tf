resource "gigavuecore_tcp_profile" "example" {
  alias            = "example"
  keep_alive_timer = 1
  selective_ack    = "example"
  syn_retries      = 1
}
