---
page_title: "gigavuecore_key Resource - gigavuecore"
subcategory: ""
description: |-
  get keystore key by alias
---

# gigavuecore_key Resource

get keystore key by alias

## Example Usage

```terraform
resource "gigavuecore_key" "example" {
  alias = null
  comment = null
  file = null
  file_source = {}
  key_label = null
  passphrase = null
  type = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required)
* `comment` (String, optional)
* `file` (String, optional) - The contents of the file. Mutually exclusive with 'fileSource'
* `file_source` (Object({hostname, password, path, protocol, username}), optional) - Remote file source or destination
  * `hostname` (String, required) - server address
  * `password` (String, optional) - password to use for server login
  * `path` (String, required) - file path on server
  * `protocol` (String, required) - Protocol to access server. HTTP and HTTPS only applicable for retrieving file
  * `username` (String, optional) - user name to use for server login
* `key_label` (String, optional) - Luna-hsm key label
* `passphrase` (String, optional) - passphrase applicable for pkcs12 and private only
* `type` (String, required) - hsm pkcs11 key; readonly

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `certificate` (Bool, computed) - true if a certificate key is configured for this alias; readonly
* `cluster_id` (String, computed) - id of the defining cluster
* `cn` (String, computed) - certificate common name
* `comment` (String, computed)
* `expiry` (String, computed) - certificate expiry date
* `health_state` (String, computed) - Read-only. 'green' indicates certificate participating in flow; 'yellow'  indicates certificate installed but not participating in any flow; 'red'  indicates certificate expired;
* `health_state_reasons` (List(Object({message, severity})), computed)
* `installed_on` (String, computed) - certificate installation time
* `o` (String, computed) - organization name
* `ou` (String, computed) - organizational unit - the division of organization handling the certificate
* `private` (Bool, computed) - true if a private key is configured for this alias; readonly
* `status` (String, computed)

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_key.example {alias}
```
