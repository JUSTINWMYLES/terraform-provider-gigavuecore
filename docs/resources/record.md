---
page_title: "gigavuecore_record Resource - gigavuecore"
subcategory: ""
description: |-
  Find Netflow Template by alias
---

# gigavuecore_record Resource

Find Netflow Template by alias

## Example Usage

```terraform
resource "gigavuecore_record" "example" {
  alias = null
  cluster_id = null
  collect = {}
  collects = []
  description = null
  export_blank_pen = null
  exporters = []
  match = {}
  nf_version = null
  sampling_rate = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required)
* `cluster_id` (String, optional) - id of the defining cluster
* `collect` (Object({collect_fields, intf_neighbor_width, ipv4_addr_dst_mask_len, ipv4_addr_src_mask_len, ipv4_section_header_size, ipv4_section_payload_size, ipv6_addr_dst_mask_len, ipv6_addr_src_mask_len, ipv6_section_header_size, ipv6_section_payload_size, pen_gigamon_url_size, pen_gigamon_user_agent_width, pen_http_host_width, pen_ssl_certificate_issuer_common_name_width, pen_ssl_certificate_issuer_width, pen_ssl_certificate_subject_alt_name_width, pen_ssl_certificate_subject_common_name_width, pen_ssl_certificate_subject_width, pen_ssl_server_name_indication_width, phy_intf_in_width, phy_intf_out_width, tcp_flags}), optional) - Netflow Record Collect spec
  * `collect_fields` (Set(String), required)
  * `intf_neighbor_width` (Number, optional) - Only valid when 'intfNeighbor' is selected. Used for representing input interface name with width
  * `ipv4_addr_dst_mask_len` (Number, optional) - Only valid when 'ipv4AddrDst' is selected. defaults to 32, in which case it matches full address
  * `ipv4_addr_src_mask_len` (Number, optional) - Only valid when 'ipv4AddrSrc' is selected. Defaults to 32, in which case it matches full address
  * `ipv4_section_header_size` (Number, optional) - Number of bytes of raw data starting at the IPv4 header, to use as a key field. Only valid and required when ipv4SectionHeaderSize' is selected
  * `ipv4_section_payload_size` (Number, optional) - Number of bytes of raw data starting at the IPv4 payload, to use as a key field. Only valid and required when 'ipv4SectionPayloadSize' is selected
  * `ipv6_addr_dst_mask_len` (Number, optional) - Only valid when 'ipv6AddrDst' is selected. defaults to 128, in which case it matches full address
  * `ipv6_addr_src_mask_len` (Number, optional) - Only valid when 'ipv6AddrSrc' is selected. Defaults to 128, in which case it matches full address
  * `ipv6_section_header_size` (Number, optional) - Only valid and required when 'ipv6SectionHeaderSize' is selected
  * `ipv6_section_payload_size` (Number, optional) - Only valid and required when 'ipv6SectionPayloadSize' is selected
  * `pen_gigamon_url_size` (Number, optional) - Only valid when 'penGigamonUrl' is selected
  * `pen_gigamon_user_agent_width` (Number, optional) - Only valid when 'penGigamonUserAgent' is selected
  * `pen_http_host_width` (Number, optional) - Only valid when 'penHttpHost' is selected
  * `pen_ssl_certificate_issuer_common_name_width` (Number, optional) - Only valid when 'penSslCertificateIssuerCommonName' is selected
  * `pen_ssl_certificate_issuer_width` (Number, optional) - Only valid when 'penSslCertificateIssuer' is selected
  * `pen_ssl_certificate_subject_alt_name_width` (Number, optional) - Only valid when 'penSslCertificateSubjectAltName' is selected
  * `pen_ssl_certificate_subject_common_name_width` (Number, optional) - Only valid when 'penSslCertificateSubjectCommonName' is selected
  * `pen_ssl_certificate_subject_width` (Number, optional) - Only valid when 'penSslCertificateSubject' is selected
  * `pen_ssl_server_name_indication_width` (Number, optional) - Only valid when 'penSslServerNameIndication' is selected
  * `phy_intf_in_width` (String, optional) - Only valid when 'phyIntfIn' is selected
  * `phy_intf_out_width` (String, optional) - Only valid when 'phyIntfOut' is selected
  * `tcp_flags` (Set(String), optional) - Represents a bit mask. Only valid and required when 'tcpFlags' is selected
* `collects` (List(Dynamic), optional)
* `description` (String, optional)
* `export_blank_pen` (Bool, optional) - if true and there is a mix of private elements and non-private elements and the private elements are blank, export the record
* `exporters` (List(String), optional)
* `match` (Object({description, ipv4_addr_dst_mask_len, ipv4_addr_src_mask_len, ipv4_section_header_size, ipv4_section_payload_size, ipv6_addr_dst_mask_len, ipv6_addr_src_mask_len, ipv6_section_header_size, ipv6_section_payload_size, match_fields, phy_intf_in_width, tcp_flags}), required) - Netflow Record Match spec
  * `description` (String, optional)
  * `ipv4_addr_dst_mask_len` (Number, optional) - Only valid when 'ipv4AddrDst' is selected. defaults to 32, in which case it matches full address
  * `ipv4_addr_src_mask_len` (Number, optional) - Only valid when 'ipv4AddrSrc' is selected. Defaults to 32, in which case it matches full address
  * `ipv4_section_header_size` (Number, optional) - Number of bytes of raw data starting at the IPv4 header, to use as a key field. Only valid and required when ipv4SectionHeaderSize' is selected
  * `ipv4_section_payload_size` (Number, optional) - Number of bytes of raw data starting at the IPv4 payload, to use as a key field. Only valid and required when 'ipv4SectionPayloadSize' is selected
  * `ipv6_addr_dst_mask_len` (Number, optional) - Only valid when 'ipv6AddrDst' is selected. defaults to 128, in which case it matches full address
  * `ipv6_addr_src_mask_len` (Number, optional) - Only valid when 'ipv6AddrSrc' is selected. Defaults to 128, in which case it matches full address
  * `ipv6_section_header_size` (Number, optional) - Only valid and required when 'ipv6SectionHeaderSize' is selected
  * `ipv6_section_payload_size` (Number, optional) - Only valid and required when 'ipv6SectionPayloadSize' is selected
  * `match_fields` (Set(String), required)
  * `phy_intf_in_width` (String, optional) - Only valid when 'phyIntfIn' is selected
  * `tcp_flags` (Set(String), optional) - Represents a bit mask. Only valid and required when 'tcpFlags' is selected
* `nf_version` (String, optional) - v5 is readOnly
* `sampling_rate` (Number, optional) - Packet interval window size. Valid values: 1-16000 (in packets); Associated monitor must 'multi-rate' sampling mode set; 0 is disabled

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `cluster_id` (String, computed) - id of the defining cluster
* `collect` (Object({collect_fields, intf_neighbor_width, ipv4_addr_dst_mask_len, ipv4_addr_src_mask_len, ipv4_section_header_size, ipv4_section_payload_size, ipv6_addr_dst_mask_len, ipv6_addr_src_mask_len, ipv6_section_header_size, ipv6_section_payload_size, pen_gigamon_url_size, pen_gigamon_user_agent_width, pen_http_host_width, pen_ssl_certificate_issuer_common_name_width, pen_ssl_certificate_issuer_width, pen_ssl_certificate_subject_alt_name_width, pen_ssl_certificate_subject_common_name_width, pen_ssl_certificate_subject_width, pen_ssl_server_name_indication_width, phy_intf_in_width, phy_intf_out_width, tcp_flags}), computed) - Netflow Record Collect spec
  * `collect_fields` (Set(String), required)
  * `intf_neighbor_width` (Number, optional) - Only valid when 'intfNeighbor' is selected. Used for representing input interface name with width
  * `ipv4_addr_dst_mask_len` (Number, optional) - Only valid when 'ipv4AddrDst' is selected. defaults to 32, in which case it matches full address
  * `ipv4_addr_src_mask_len` (Number, optional) - Only valid when 'ipv4AddrSrc' is selected. Defaults to 32, in which case it matches full address
  * `ipv4_section_header_size` (Number, optional) - Number of bytes of raw data starting at the IPv4 header, to use as a key field. Only valid and required when ipv4SectionHeaderSize' is selected
  * `ipv4_section_payload_size` (Number, optional) - Number of bytes of raw data starting at the IPv4 payload, to use as a key field. Only valid and required when 'ipv4SectionPayloadSize' is selected
  * `ipv6_addr_dst_mask_len` (Number, optional) - Only valid when 'ipv6AddrDst' is selected. defaults to 128, in which case it matches full address
  * `ipv6_addr_src_mask_len` (Number, optional) - Only valid when 'ipv6AddrSrc' is selected. Defaults to 128, in which case it matches full address
  * `ipv6_section_header_size` (Number, optional) - Only valid and required when 'ipv6SectionHeaderSize' is selected
  * `ipv6_section_payload_size` (Number, optional) - Only valid and required when 'ipv6SectionPayloadSize' is selected
  * `pen_gigamon_url_size` (Number, optional) - Only valid when 'penGigamonUrl' is selected
  * `pen_gigamon_user_agent_width` (Number, optional) - Only valid when 'penGigamonUserAgent' is selected
  * `pen_http_host_width` (Number, optional) - Only valid when 'penHttpHost' is selected
  * `pen_ssl_certificate_issuer_common_name_width` (Number, optional) - Only valid when 'penSslCertificateIssuerCommonName' is selected
  * `pen_ssl_certificate_issuer_width` (Number, optional) - Only valid when 'penSslCertificateIssuer' is selected
  * `pen_ssl_certificate_subject_alt_name_width` (Number, optional) - Only valid when 'penSslCertificateSubjectAltName' is selected
  * `pen_ssl_certificate_subject_common_name_width` (Number, optional) - Only valid when 'penSslCertificateSubjectCommonName' is selected
  * `pen_ssl_certificate_subject_width` (Number, optional) - Only valid when 'penSslCertificateSubject' is selected
  * `pen_ssl_server_name_indication_width` (Number, optional) - Only valid when 'penSslServerNameIndication' is selected
  * `phy_intf_in_width` (String, optional) - Only valid when 'phyIntfIn' is selected
  * `phy_intf_out_width` (String, optional) - Only valid when 'phyIntfOut' is selected
  * `tcp_flags` (Set(String), optional) - Represents a bit mask. Only valid and required when 'tcpFlags' is selected
* `collects` (List(Dynamic), computed)
* `description` (String, computed)
* `export_blank_pen` (Bool, computed) - if true and there is a mix of private elements and non-private elements and the private elements are blank, export the record
* `exporters` (List(String), computed)
* `nf_version` (String, computed) - v5 is readOnly
* `sampling_rate` (Number, computed) - Packet interval window size. Valid values: 1-16000 (in packets); Associated monitor must 'multi-rate' sampling mode set; 0 is disabled

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_record.example {alias}
```
