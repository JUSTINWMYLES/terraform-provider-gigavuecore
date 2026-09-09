---
page_title: "gigavuecore_cert_file Resource - gigavuecore"
subcategory: ""
description: |-
  Append inline SSL trust-store from local file
---

# gigavuecore_cert_file Resource

Append inline SSL trust-store from local file

~> **Note:** The update operation for this resource is not wired to a remote API endpoint: the API spec exposes no usable update mapping. Changing any configuration attribute triggers a resource replacement (all config-settable attributes carry a RequiresReplace plan modifier); create, read, and delete remain functional.

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

* `create` (Number) - A create timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `read` (Number) - A read timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `update` (Number) - An update timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `delete` (Number) - A delete timeout in seconds for this operation. Overrides the generator default (1200 seconds).

