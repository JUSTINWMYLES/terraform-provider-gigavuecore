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
  alias      = null
  cluster_id = null
  hostname   = null
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
* `ssl_decryption_session_details` (Attributes List, computed) (see [below for nested schema](#nestedatt--ssl_decryption_session_details))

<a id="nestedatt--ssl_decryption_session_details"></a>
### Nested Schema for `ssl_decryption_session_details`

Read-Only:

* `cipher_suite` (String)
* `client_ip` (String)
* `client_port` (Number)
* `decryption_status` (String)
* `duration` (Number)
* `fingerprint` (String)
* `first_error` (Number)
* `first_error_reason` (String)
* `in_pkts` (Number)
* `out_pkts` (Number)
* `server_ip` (String)
* `server_port` (Number)
* `sni` (String)
* `start_time` (String)
* `version` (String)

