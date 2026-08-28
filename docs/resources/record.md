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
  alias            = null
  cluster_id       = null
  collect          = {}
  collects         = []
  description      = null
  export_blank_pen = null
  exporters        = []
  match            = {}
  nf_version       = null
  sampling_rate    = null
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

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `collect` (Attributes, computed) - Netflow Record Collect spec (see [below for nested schema](#nestedatt--collect))
* `collects` (List of Dynamic, computed)
* `description` (String, computed)
* `export_blank_pen` (Boolean, computed) - if true and there is a mix of private elements and non-private elements and the private elements are blank, export the record
* `exporters` (List of String, computed)
* `nf_version` (String, computed) - v5 is readOnly
* `sampling_rate` (Number, computed) - Packet interval window size. Valid values: 1-16000 (in packets); Associated monitor must 'multi-rate' sampling mode set; 0 is disabled

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

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_record.example {alias}
```
