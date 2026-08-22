---
page_title: "gigavuecore_get_ssl_client_trust_store_certificate Data Source - gigavuecore"
subcategory: ""
description: |-
  Get the SSL Client Trust Store Certificate
---

# gigavuecore_get_ssl_client_trust_store_certificate Data Source

Get the SSL Client Trust Store Certificate

## Example Usage

```terraform
data "gigavuecore_get_ssl_client_trust_store_certificate" "example" {
  alias = null
  cluster_id = null
  fingerprint = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Client Trust Store alias
* `cluster_id` (String, required) - Target Cluster ID
* `fingerprint` (String, required) - fingerprint for the certificate

### Attributes

In addition to all arguments above, the following attributes are exported:

* `certificate` (String, computed)
* `type` (String, computed)

