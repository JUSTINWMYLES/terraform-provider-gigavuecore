resource "gigavuecore_record" "example" {
  alias      = "example"
  cluster_id = "example"
  collect = {
    collect_fields                                = [ "example" ]
    intf_neighbor_width                           = 0
    ipv4_addr_dst_mask_len                        = 0
    ipv4_addr_src_mask_len                        = 0
    ipv4_section_header_size                      = 0
    ipv4_section_payload_size                     = 0
    ipv6_addr_dst_mask_len                        = 0
    ipv6_addr_src_mask_len                        = 0
    ipv6_section_header_size                      = 0
    ipv6_section_payload_size                     = 0
    pen_gigamon_url_size                          = 0
    pen_gigamon_user_agent_width                  = 0
    pen_http_host_width                           = 0
    pen_ssl_certificate_issuer_common_name_width  = 0
    pen_ssl_certificate_issuer_width              = 0
    pen_ssl_certificate_subject_alt_name_width    = 0
    pen_ssl_certificate_subject_common_name_width = 0
    pen_ssl_certificate_subject_width             = 0
    pen_ssl_server_name_indication_width          = 0
    phy_intf_in_width                             = "example"
    phy_intf_out_width                            = "example"
    tcp_flags                                     = [ "example" ]
  }
  collects         = [ "example" ]
  description      = "example"
  export_blank_pen = true
  exporters        = [ "example" ]
  match = {
    description               = "example"
    ipv4_addr_dst_mask_len    = 0
    ipv4_addr_src_mask_len    = 0
    ipv4_section_header_size  = 0
    ipv4_section_payload_size = 0
    ipv6_addr_dst_mask_len    = 0
    ipv6_addr_src_mask_len    = 0
    ipv6_section_header_size  = 0
    ipv6_section_payload_size = 0
    match_fields              = [ "example" ]
    phy_intf_in_width         = "example"
    tcp_flags                 = [ "example" ]
  }
  nf_version    = "example"
  sampling_rate = 0
}
