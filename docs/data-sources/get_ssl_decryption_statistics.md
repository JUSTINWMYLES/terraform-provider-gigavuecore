---
page_title: "gigavuecore_get_ssl_decryption_statistics Data Source - gigavuecore"
subcategory: ""
description: |-
  Load Ssl Decryption Report Statistics
---

# gigavuecore_get_ssl_decryption_statistics Data Source

Load Ssl Decryption Report Statistics

## Example Usage

```terraform
data "gigavuecore_get_ssl_decryption_statistics" "example" {
  alias = null
  cluster_id = null
  hostname = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target GS Group
* `cluster_id` (String, required) - Target Cluster ID
* `hostname` (String, optional) - match based on hostname

### Attributes

In addition to all arguments above, the following attributes are exported:

* `gsgroup` (String, computed) - alias of gsgroup
* `ssl_decryption_session_details` (List(Object({cipher_suite, client_ip, client_port, decryption_status, duration, fingerprint, first_error, first_error_reason, in_pkts, out_pkts, server_ip, server_port, sni, start_time, version})), computed)

