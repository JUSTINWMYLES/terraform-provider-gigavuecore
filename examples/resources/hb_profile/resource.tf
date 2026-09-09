resource "gigavuecore_hb_profile" "example" {
  alias                   = "example"
  cluster_id              = "example"
  custom_packet           = "example"
  custom_packet_alias     = "example"
  custom_packet_file_name = "example"
  direction               = "aToB"
  packet_format           = "arp"
  period                  = 30
  recovery_time           = 5
  retries                 = 0
  timeout                 = 20
}
