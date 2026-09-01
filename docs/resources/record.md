---
page_title: "gigavuecore_record Resource - gigavuecore"
subcategory: ""
description: |-
  Create a new Netflow Template
---

# gigavuecore_record Resource

Create a new Netflow Template

## Example Usage

```terraform
resource "gigavuecore_record" "example" {
  alias      = "example"
  cluster_id = "example"
  collect = {
    collect_fields                                = ["counterBytes"]
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
    phy_intf_in_width                             = "2"
    phy_intf_out_width                            = "2"
    tcp_flags                                     = ["cwr"]
  }
  collects         = ["example"]
  description      = "example"
  export_blank_pen = true
  exporters        = ["example"]
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
    match_fields              = ["phyIntfIn"]
    phy_intf_in_width         = "2"
    tcp_flags                 = ["cwr"]
  }
  nf_version    = "v5"
  sampling_rate = 0
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required)
* `cluster_id` (String, required) - id of the defining cluster
* `collect` (Attributes, optional) - Netflow Record Collect spec (see [below for nested schema](#nestedatt--collect))
* `collects` (List of Dynamic, optional)
* `description` (String, optional)
* `export_blank_pen` (Boolean, optional) - if true and there is a mix of private elements and non-private elements and the private elements are blank, export the record
* `exporters` (List of String, optional)
* `match` (Attributes, required) - Netflow Record Match spec (see [below for nested schema](#nestedatt--match))
* `nf_version` (String, optional) - v5 is readOnly
* `sampling_rate` (Number, optional) - Packet interval window size. Valid values: 1-16000 (in packets); Associated monitor must 'multi-rate' sampling mode set; 0 is disabled

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

<a id="nestedatt--collect"></a>
### Nested Schema for `collect`

Required:

* `collect_fields` (Set of String)

Optional:

* `intf_neighbor_width` (Number) - Only valid when 'intfNeighbor' is selected. Used for representing input interface name with width
* `ipv4_addr_dst_mask_len` (Number) - Only valid when 'ipv4AddrDst' is selected. defaults to 32, in which case it matches full address
* `ipv4_addr_src_mask_len` (Number) - Only valid when 'ipv4AddrSrc' is selected. Defaults to 32, in which case it matches full address
* `ipv4_section_header_size` (Number) - Number of bytes of raw data starting at the IPv4 header, to use as a key field. Only valid and required when ipv4SectionHeaderSize' is selected
* `ipv4_section_payload_size` (Number) - Number of bytes of raw data starting at the IPv4 payload, to use as a key field. Only valid and required when 'ipv4SectionPayloadSize' is selected
* `ipv6_addr_dst_mask_len` (Number) - Only valid when 'ipv6AddrDst' is selected. defaults to 128, in which case it matches full address
* `ipv6_addr_src_mask_len` (Number) - Only valid when 'ipv6AddrSrc' is selected. Defaults to 128, in which case it matches full address
* `ipv6_section_header_size` (Number) - Only valid and required when 'ipv6SectionHeaderSize' is selected
* `ipv6_section_payload_size` (Number) - Only valid and required when 'ipv6SectionPayloadSize' is selected
* `pen_gigamon_url_size` (Number) - Only valid when 'penGigamonUrl' is selected
* `pen_gigamon_user_agent_width` (Number) - Only valid when 'penGigamonUserAgent' is selected
* `pen_http_host_width` (Number) - Only valid when 'penHttpHost' is selected
* `pen_ssl_certificate_issuer_common_name_width` (Number) - Only valid when 'penSslCertificateIssuerCommonName' is selected
* `pen_ssl_certificate_issuer_width` (Number) - Only valid when 'penSslCertificateIssuer' is selected
* `pen_ssl_certificate_subject_alt_name_width` (Number) - Only valid when 'penSslCertificateSubjectAltName' is selected
* `pen_ssl_certificate_subject_common_name_width` (Number) - Only valid when 'penSslCertificateSubjectCommonName' is selected
* `pen_ssl_certificate_subject_width` (Number) - Only valid when 'penSslCertificateSubject' is selected
* `pen_ssl_server_name_indication_width` (Number) - Only valid when 'penSslServerNameIndication' is selected
* `phy_intf_in_width` (String) - Only valid when 'phyIntfIn' is selected
* `phy_intf_out_width` (String) - Only valid when 'phyIntfOut' is selected
* `tcp_flags` (Set of String) - Represents a bit mask. Only valid and required when 'tcpFlags' is selected

<a id="nestedatt--match"></a>
### Nested Schema for `match`

Required:

* `match_fields` (Set of String)

Optional:

* `description` (String)
* `ipv4_addr_dst_mask_len` (Number) - Only valid when 'ipv4AddrDst' is selected. defaults to 32, in which case it matches full address
* `ipv4_addr_src_mask_len` (Number) - Only valid when 'ipv4AddrSrc' is selected. Defaults to 32, in which case it matches full address
* `ipv4_section_header_size` (Number) - Number of bytes of raw data starting at the IPv4 header, to use as a key field. Only valid and required when ipv4SectionHeaderSize' is selected
* `ipv4_section_payload_size` (Number) - Number of bytes of raw data starting at the IPv4 payload, to use as a key field. Only valid and required when 'ipv4SectionPayloadSize' is selected
* `ipv6_addr_dst_mask_len` (Number) - Only valid when 'ipv6AddrDst' is selected. defaults to 128, in which case it matches full address
* `ipv6_addr_src_mask_len` (Number) - Only valid when 'ipv6AddrSrc' is selected. Defaults to 128, in which case it matches full address
* `ipv6_section_header_size` (Number) - Only valid and required when 'ipv6SectionHeaderSize' is selected
* `ipv6_section_payload_size` (Number) - Only valid and required when 'ipv6SectionPayloadSize' is selected
* `phy_intf_in_width` (String) - Only valid when 'phyIntfIn' is selected
* `tcp_flags` (Set of String) - Represents a bit mask. Only valid and required when 'tcpFlags' is selected
<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

* `create` (String) - A create timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `read` (String) - A read timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).
* `update` (String) - An update timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `delete` (String) - A delete timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_record.example {alias}/{cluster_id}
```
