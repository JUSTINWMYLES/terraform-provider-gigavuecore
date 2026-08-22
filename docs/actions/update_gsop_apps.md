---
page_title: "gigavuecore_update_gsop_apps Action - gigavuecore"
subcategory: ""
description: |-
  Update GSOP apps configuration
---

# gigavuecore_update_gsop_apps Action

Update GSOP apps configuration

## Example Usage

```terraform
action "gigavuecore_update_gsop_apps" "example" {
  config {
    alias = "example"
    apf = null
    cluster_id = "example"
    dedup = null
    diameter_whitelist = null
    flow_filter = null
    flow_sampling = null
    gseries_header_add = null
    gseries_header_remove = null
    gseries_load_balance = null
    gseries_pattern_match = null
    gtp_whitelist = null
    header_add = null
    header_remove = null
    icap = null
    inline_ssl = null
    load_balance = null
    masking = null
    metadata_export = null
    netflow = null
    sa_apf = null
    sip_whitelist = null
    slicing = null
    ssl_decrypt = null
    trailer_add = null
    trailer_remove = null
    tunnel_decap = null
    tunnel_encap = null
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target GSOP
* `apf` (Dynamic, optional)
* `cluster_id` (String, required) - Target Cluster ID
* `dedup` (Dynamic, optional)
* `diameter_whitelist` (Dynamic, optional)
* `flow_filter` (Dynamic, optional) - GigaSMART 'Flow Filter' Application Configuration
* `flow_sampling` (Dynamic, optional) - GigaSMART 'Flow Sampling' Application Configuration
* `gseries_header_add` (Dynamic, optional) - Only applicable for G-series
* `gseries_header_remove` (Dynamic, optional)
* `gseries_load_balance` (Dynamic, optional) - Only applicable for G-series per-rule GSOP
* `gseries_pattern_match` (Dynamic, optional) - Only applicable for G-series per-rule GSOP
* `gtp_whitelist` (Dynamic, optional)
* `header_add` (Dynamic, optional) - GigaSMART 'Add Header' Application Configuration
* `header_remove` (Dynamic, optional) - GigaSMART 'Remove Header' Application Configuration
* `icap` (Dynamic, optional) - GigaSMART ICAP Configuration
* `inline_ssl` (Dynamic, optional) - GigaSMART Inline SSL Profile Configuration
* `load_balance` (Dynamic, optional) - GigaSMART 'Load Balancing' Application Configuration
* `masking` (Dynamic, optional) - GigaSMART 'Masking' Application Configuration
* `metadata_export` (Dynamic, optional)
* `netflow` (Dynamic, optional)
* `sa_apf` (Dynamic, optional)
* `sip_whitelist` (Dynamic, optional)
* `slicing` (Dynamic, optional) - GigaSMART 'Slicing' Application Configuration
* `ssl_decrypt` (Dynamic, optional) - GigaSMART 'SSL Decrypt' Application Configuration
* `trailer_add` (Dynamic, optional) - GigaSMART 'Add Trailer' Application Configuration
* `trailer_remove` (Dynamic, optional)
* `tunnel_decap` (Dynamic, optional) - GigaSMART 'Decapsulate Tunnel' Application Configuration
* `tunnel_encap` (Dynamic, optional) - GigaSMART 'Encapsulate Tunnel' Application Configuration
