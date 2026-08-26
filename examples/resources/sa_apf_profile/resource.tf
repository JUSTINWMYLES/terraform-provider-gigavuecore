resource "gigavuecore_sa_apf_profile" "example" {
  alias = "example"
  bidi  = true
  buffering = {
    buffer_count_before_match = 1
    enabled                   = true
    protocol                  = "example"
  }
  packet_count = 1
  session_fields = [{
    pos  = 1
    type = "example"
  }]
  timeout = 1
}
