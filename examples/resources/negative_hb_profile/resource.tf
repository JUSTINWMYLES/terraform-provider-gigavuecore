resource "gigavuecore_negative_hb_profile" "example" {
  alias                   = "example"
  cluster_id              = "example"
  custom_packet           = "example"
  custom_packet_file_name = "example"
  direction               = "aToB"
  period                  = 30
  recovery_time           = 5
}
