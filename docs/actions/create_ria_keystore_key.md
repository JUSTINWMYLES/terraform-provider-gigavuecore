---
page_title: "gigavuecore_create_ria_keystore_key Action - gigavuecore"
subcategory: ""
description: |-
  add key and/or certificate to the keystore for two devices needed for Resilient Inline SSL
---

# gigavuecore_create_ria_keystore_key Action

add key and/or certificate to the keystore for two devices needed for Resilient Inline SSL

## Example Usage

```terraform
action "gigavuecore_create_ria_keystore_key" "example" {
  config {
    alias = "example"
    certificate = {
      certificate_file = "example"
      certificate_url = {
        hostname = "example"
        password = "example"
        path     = "example"
        protocol = "scp"
        username = "example"
      }
      passphrase = "example"
      type       = "certificate"
    }
    cluster_names = [ "example" ]
    comment       = "example"
    key = {
      key_file  = "example"
      key_label = "example"
      key_url = {
        hostname = "example"
        password = "example"
        path     = "example"
        protocol = "scp"
        username = "example"
      }
      passphrase = "example"
      type       = "private"
    }
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required)
* `certificate` (Attributes, optional) - Remote location to fetch file from. Mutually exclusive with 'file' (see [below for nested schema](#nestedatt--certificate))
* `cluster_names` (List of String, optional)
* `comment` (String, optional)
* `key` (Attributes, optional) - Remote location to fetch file from. Mutually exclusive with 'file' (see [below for nested schema](#nestedatt--key))

<a id="nestedatt--certificate"></a>
### Nested Schema for `certificate`

Optional:

* `certificate_file` (String) - The contents of the certificate file. Mutually exclusive with 'fileSource'
* `certificate_url` (Attributes) - Remote file source or destination (see [below for nested schema](#nestedatt--certificate--certificate_url))
* `passphrase` (String) - passphrase applicable for pkcs12 and private only
* `type` (String)

<a id="nestedatt--certificate--certificate_url"></a>
### Nested Schema for `certificate.certificate_url`

Required:

* `hostname` (String) - server address
* `path` (String) - file path on server
* `protocol` (String) - Protocol to access server. HTTP and HTTPS only applicable for retrieving file

Optional:

* `password` (String) - password to use for server login
* `username` (String) - user name to use for server login

<a id="nestedatt--key"></a>
### Nested Schema for `key`

Optional:

* `key_file` (String) - The contents of the key file. Mutually exclusive with 'fileSource'
* `key_label` (String) - Key Label
* `key_url` (Attributes) - Remote file source or destination (see [below for nested schema](#nestedatt--key--key_url))
* `passphrase` (String) - passphrase applicable for pkcs12 and private only
* `type` (String)

<a id="nestedatt--key--key_url"></a>
### Nested Schema for `key.key_url`

Required:

* `hostname` (String) - server address
* `path` (String) - file path on server
* `protocol` (String) - Protocol to access server. HTTP and HTTPS only applicable for retrieving file

Optional:

* `password` (String) - password to use for server login
* `username` (String) - user name to use for server login

