resource "gigavuecore_profile" "example" {
  alias               = "example"
  cluster_id          = "example"
  exceed_action       = "drop"
  http_req_buf        = 10
  inactivity_timeout  = 2
  preview             = 0
  resp_mod            = "enable"
  resp_timeout        = 5
  resp_timeout_action = "drop"
  server_group        = "example"
  src_max_l4_port     = 10000
  src_min_l4_port     = 10000
}
