---
page_title: "gigavuecore_get_ssl_client_trust_store Data Source - gigavuecore"
subcategory: ""
description: |-
  Get information about the SSL Client trust-store
---

# gigavuecore_get_ssl_client_trust_store Data Source

Get information about the SSL Client trust-store

## Example Usage

```terraform
data "gigavuecore_get_ssl_client_trust_store" "example" {
  alias      = null
  cluster_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Client Trust Store alias
* `cluster_id` (String, required) - id of the defining cluster

### Attributes

In addition to all arguments above, the following attributes are exported:

* `ssl_trust_store_entries` (Attributes List, computed) (see [below for nested schema](#nestedatt--ssl_trust_store_entries))
* `type` (String, computed)

<a id="nestedatt--ssl_trust_store_entries"></a>
### Nested Schema for `ssl_trust_store_entries`

Read-Only:

* `expire_date` (String)
* `fingerprint` (String)
* `issuer_common_name` (String)
* `issuer_name` (String)
* `subject_common_name` (String)
* `subject_name` (String)

