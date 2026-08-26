resource "gigavuecore_record" "example" {
  alias = "example"
  collect = {
    collect_fields                                = [ "example" ]
    intf_neighbor_width                           = 1
    ipv4_addr_dst_mask_len                        = 1
    ipv4_addr_src_mask_len                        = 1
    ipv4_section_header_size                      = 1
    ipv4_section_payload_size                     = 1
    ipv6_addr_dst_mask_len                        = 1
    ipv6_addr_src_mask_len                        = 1
    ipv6_section_header_size                      = 1
    ipv6_section_payload_size                     = 1
    pen_gigamon_url_size                          = 1
    pen_gigamon_user_agent_width                  = 1
    pen_http_host_width                           = 1
    pen_ssl_certificate_issuer_common_name_width  = 1
    pen_ssl_certificate_issuer_width              = 1
    pen_ssl_certificate_subject_alt_name_width    = 1
    pen_ssl_certificate_subject_common_name_width = 1
    pen_ssl_certificate_subject_width             = 1
    pen_ssl_server_name_indication_width          = 1
    phy_intf_in_width                             = "example"
    phy_intf_out_width                            = "example"
    tcp_flags                                     = [ "example" ]
  }
  collects         = null
  description      = "example"
  export_blank_pen = true
  exporters        = [ "example" ]
  match = {
    description               = "example"
    ipv4_addr_dst_mask_len    = 1
    ipv4_addr_src_mask_len    = 1
    ipv4_section_header_size  = 1
    ipv4_section_payload_size = 1
    ipv6_addr_dst_mask_len    = 1
    ipv6_addr_src_mask_len    = 1
    ipv6_section_header_size  = 1
    ipv6_section_payload_size = 1
    match_fields              = [ "example" ]
    phy_intf_in_width         = "example"
    tcp_flags                 = [ "example" ]
  }
  nf_version    = "example"
  sampling_rate = 1
}
