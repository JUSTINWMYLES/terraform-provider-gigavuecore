resource "gigavuecore_sa_apf_profile" "example" {
  alias = "example"
  bidi  = true
  buffering = {
    buffer_count_before_match = 3
    enabled                   = true
    protocol                  = "tcp"
  }
  cluster_id   = "example"
  packet_count = 0
  session_fields = [{
    pos  = 1
    type = "ipv4Addr"
  }]
  timeout = 10
}
