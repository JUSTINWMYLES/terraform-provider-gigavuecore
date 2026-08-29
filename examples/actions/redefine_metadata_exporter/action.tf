action "gigavuecore_redefine_metadata_exporter" "example" {
  config {
    alias                = "example"
    application_profiles = [ "example" ]
    body_alias           = "example"
    cef = {
      active_timeout   = 0
      inactive_timeout = 0
    }
    description = "example"
    destination = {
      dscp         = 0
      ipv4_address = "example"
      l4_port_dst  = 0
      l4_port_src  = 0
      l4_protocol  = "example"
      ttl          = 0
    }
    max_pkt_size = 0
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
      timeout = 0
    }
    netflow = {
      active_timeout   = 0
      inactive_timeout = 0
      template_refresh = 0
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
}
