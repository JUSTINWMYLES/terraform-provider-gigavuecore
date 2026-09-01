---
page_title: "gigavuecore_ssl_trust_store_cert_file Resource - gigavuecore"
subcategory: ""
description: |-
  Append certificate to the the SSL Client trust-store from local file
---

# gigavuecore_ssl_trust_store_cert_file Resource

Append certificate to the the SSL Client trust-store from local file

~> **Note:** This resource is not yet wired to a remote API endpoint. Invoking it fails with an explicit "not wired" diagnostic instead of calling the API. The OpenAPI operations it was inferred from could not be resolved into a complete mapping; consult the eidos generation warnings for the exact cause.

## Example Usage

```terraform
resource "gigavuecore_ssl_trust_store_cert_file" "example" {
  cluster_id = "example"
  file       = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `file` (String, required) - file to upload to device

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `alias` (String, computed)
* `certificate` (String, computed)
* `id` (String, computed)
* `type` (String, computed)

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

* `create` (String) - A create timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `read` (String) - A read timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).
* `update` (String) - An update timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `delete` (String) - A delete timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).

