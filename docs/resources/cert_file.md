---
page_title: "gigavuecore_cert_file Resource - gigavuecore"
subcategory: ""
description: |-
  Get the Trust Store Certificate
---

# gigavuecore_cert_file Resource

Get the Trust Store Certificate

~> **Note:** The update operation for this resource is not wired to a remote API endpoint: the API spec exposes no usable update mapping. Changing any configuration fails at apply time with an explicit "not wired" diagnostic. To change this resource, replace it (for example `terraform apply -replace=...`); create, read, and delete remain functional.

## Example Usage

```terraform
resource "gigavuecore_cert_file" "example" {
  file = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `file` (String, required) - file to upload to device

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `certificate` (String, computed)
* `fingerprint` (String, computed) - fingerprint for the certificate

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

* `create` (String) - A create timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `read` (String) - A read timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).
* `update` (String) - An update timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `delete` (String) - A delete timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).

