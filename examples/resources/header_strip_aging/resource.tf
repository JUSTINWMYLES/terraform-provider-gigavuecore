resource "gigavuecore_header_strip_aging" "example" {
  aging_interval = 300
  box_id         = "example"
  dst_port       = 0
  protocol_type  = "none"
}
