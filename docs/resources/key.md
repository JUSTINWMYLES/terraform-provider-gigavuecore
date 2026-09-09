---
page_title: "gigavuecore_key Resource - gigavuecore"
subcategory: ""
description: |-
  add a key to the keystore
---

# gigavuecore_key Resource

add a key to the keystore

## Example Usage

```terraform
resource "gigavuecore_key" "example" {
  alias      = "example"
  cluster_id = "example"
  comment    = "example"
  file       = "example"
  file_source = {
    hostname = "example"
    password = "example"
    path     = "example"
    protocol = "scp"
    username = "example"
  }
  key_label  = "example"
  passphrase = "example"
  type       = "pkcs11"
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required)
* `cluster_id` (String, required) - id of the defining cluster
* `comment` (String, optional)
* `file` (String, optional) - The contents of the file. Mutually exclusive with 'fileSource'
* `file_source` (Attributes, optional) - Remote file source or destination (see [below for nested schema](#nestedatt--file_source))
* `key_label` (String, optional) - Luna-hsm key label
* `passphrase` (String, optional) - passphrase applicable for pkcs12 and private only
* `type` (String, required) - hsm pkcs11 key; readonly

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `certificate` (Boolean, computed) - true if a certificate key is configured for this alias; readonly
* `cn` (String, computed) - certificate common name
* `expiry` (String, computed) - certificate expiry date
* `health_state` (String, computed) - Read-only. 'green' indicates certificate participating in flow; 'yellow'  indicates certificate installed but not participating in any flow; 'red'  indicates certificate expired;
* `health_state_reasons` (Attributes List, computed) (see [below for nested schema](#nestedatt--health_state_reasons))
* `installed_on` (String, computed) - certificate installation time
* `o` (String, computed) - organization name
* `ou` (String, computed) - organizational unit - the division of organization handling the certificate
* `private` (Boolean, computed) - true if a private key is configured for this alias; readonly
* `status` (String, computed)

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

<a id="nestedatt--file_source"></a>
### Nested Schema for `file_source`

Required:

* `hostname` (String) - server address
* `path` (String) - file path on server
* `protocol` (String) - Protocol to access server. HTTP and HTTPS only applicable for retrieving file

Optional:

* `password` (String) - password to use for server login
* `username` (String) - user name to use for server login

<a id="nestedatt--health_state_reasons"></a>
### Nested Schema for `health_state_reasons`

Optional:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates certificate participating in flow; 'yellow'  indicates certificate installed but not participating in any flow; 'red'  indicates certificate expired;
<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

* `create` (Number) - A create timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `read` (Number) - A read timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `update` (Number) - An update timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `delete` (Number) - A delete timeout in seconds for this operation. Overrides the generator default (1200 seconds).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_key.example {alias}/{cluster_id}
```
