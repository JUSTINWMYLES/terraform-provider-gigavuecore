resource "gigavuecore_sa_apf_profile" "example" {
  alias = "example"
  bidi  = true
  buffering = {
    buffer_count_before_match = 0
    enabled                   = true
    protocol                  = "example"
  }
  cluster_id   = "example"
  packet_count = 0
  session_fields = [{
    pos  = 0
    type = "example"
  }]
  timeout = 0
}
