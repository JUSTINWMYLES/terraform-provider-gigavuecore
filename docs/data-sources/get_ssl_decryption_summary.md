---
page_title: "gigavuecore_get_ssl_decryption_summary Data Source - gigavuecore"
subcategory: ""
description: |-
  Load Ssl Decryption Report Summary
---

# gigavuecore_get_ssl_decryption_summary Data Source

Load Ssl Decryption Report Summary

## Example Usage

```terraform
data "gigavuecore_get_ssl_decryption_summary" "example" {
  alias       = null
  cluster_id  = null
  device_ip   = null
  device_mask = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target GS Group
* `cluster_id` (String, required) - Target Cluster ID
* `device_ip` (String, optional) - deviceIp pattern based active flows
* `device_mask` (String, optional) - deviceIpMask

### Attributes

In addition to all arguments above, the following attributes are exported:

* `gsgroup` (String, computed) - alias of gsgroup
* `session_ids` (Number, computed)
* `ssl30_session` (Number, computed)
* `tickets` (Number, computed)
* `tls10_session` (Number, computed)
* `tls11_session` (Number, computed)
* `tls12_session` (Number, computed)
* `total_session` (Number, computed)


