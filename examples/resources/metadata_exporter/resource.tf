resource "gigavuecore_metadata_exporter" "example" {
  alias                = "example"
  application_profiles = [ "example" ]
  cef = {
    active_timeout   = 1
    inactive_timeout = 1
  }
  description = "example"
  destination = {
    dscp         = 1
    ipv4_address = "example"
    l4_port_dst  = 1
    l4_port_src  = 1
    l4_protocol  = "example"
    ttl          = 1
  }
  max_pkt_size = 1
  mobility_sam = {
    encoding        = "example"
    encoding_format = "example"
    event_enable = {
      modify = true
      update = true
    }
    trigger = "example"
  }
  monitor = {
    timeout = 1
  }
  netflow = {
    active_timeout   = 1
    inactive_timeout = 1
    template_refresh = 1
    template_type    = "example"
    version          = "example"
  }
  snmp = {
    enabled = true
  }
  source = {
    ip_interface = "example"
  }
  type = "example"
}
