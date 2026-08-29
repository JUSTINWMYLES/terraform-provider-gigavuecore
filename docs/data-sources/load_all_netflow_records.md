---
page_title: "gigavuecore_load_all_netflow_records Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all defined Netflow Templates
---

# gigavuecore_load_all_netflow_records Data Source

Load all defined Netflow Templates

## Example Usage

```terraform
data "gigavuecore_load_all_netflow_records" "example" {
  cluster_id = "example"
  page       = "example"
  sort       = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `alias` (String)
* `cluster_id` (String) - id of the defining cluster
* `collect` (Attributes) - Netflow Record Collect spec (see [below for nested schema](#nestedatt--items--collect))
* `collects` (List of Dynamic)
* `description` (String)
* `export_blank_pen` (Boolean) - if true and there is a mix of private elements and non-private elements and the private elements are blank, export the record
* `exporters` (List of String)
* `match` (Attributes) - Netflow Record Match spec (see [below for nested schema](#nestedatt--items--match))
* `nf_version` (String) - v5 is readOnly
* `sampling_rate` (Number) - Packet interval window size. Valid values: 1-16000 (in packets); Associated monitor must 'multi-rate' sampling mode set; 0 is disabled
<a id="nestedatt--items--collect"></a>
### Nested Schema for `items.collect`

Read-Only:

* `collect_fields` (Set of String)
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
<a id="nestedatt--items--match"></a>
### Nested Schema for `items.match`

Read-Only:

* `description` (String)
* `ipv4_addr_dst_mask_len` (Number) - Only valid when 'ipv4AddrDst' is selected. defaults to 32, in which case it matches full address
* `ipv4_addr_src_mask_len` (Number) - Only valid when 'ipv4AddrSrc' is selected. Defaults to 32, in which case it matches full address
* `ipv4_section_header_size` (Number) - Number of bytes of raw data starting at the IPv4 header, to use as a key field. Only valid and required when ipv4SectionHeaderSize' is selected
* `ipv4_section_payload_size` (Number) - Number of bytes of raw data starting at the IPv4 payload, to use as a key field. Only valid and required when 'ipv4SectionPayloadSize' is selected
* `ipv6_addr_dst_mask_len` (Number) - Only valid when 'ipv6AddrDst' is selected. defaults to 128, in which case it matches full address
* `ipv6_addr_src_mask_len` (Number) - Only valid when 'ipv6AddrSrc' is selected. Defaults to 128, in which case it matches full address
* `ipv6_section_header_size` (Number) - Only valid and required when 'ipv6SectionHeaderSize' is selected
* `ipv6_section_payload_size` (Number) - Only valid and required when 'ipv6SectionPayloadSize' is selected
* `match_fields` (Set of String)
* `phy_intf_in_width` (String) - Only valid when 'phyIntfIn' is selected
* `tcp_flags` (Set of String) - Represents a bit mask. Only valid and required when 'tcpFlags' is selected

