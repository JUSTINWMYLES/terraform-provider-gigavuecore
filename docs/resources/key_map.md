---
page_title: "gigavuecore_key_map Resource - gigavuecore"
subcategory: ""
description: |-
  Find SSL Decryption KeyMap by alias
---

# gigavuecore_key_map Resource

Find SSL Decryption KeyMap by alias

## Example Usage

```terraform
resource "gigavuecore_key_map" "example" {
  alias      = null
  cluster_id = null
  mappings   = []
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Alias of this mappings collection
* `cluster_id` (String, required) - id of the defining cluster
* `mappings` (Attributes List, required) - list of SSL Endpoints with corresponding Decryption Keys (see [below for nested schema](#nestedatt--mappings))

<a id="nestedatt--mappings"></a>
### Nested Schema for `mappings`

Read-Only:

* `endpoint` (Attributes) - SSL Decryption Endpoint definition (see [below for nested schema](#nestedatt--mappings--endpoint))
* `key` (Attributes) - SSL Decryption Key definition (see [below for nested schema](#nestedatt--mappings--key))
<a id="nestedatt--mappings--endpoint"></a>
### Nested Schema for `mappings.endpoint`

Read-Only:

* `address` (String) - IPv4 address of the SSL endpoint (server)
* `alias` (String) - Alias of the Endpoint
* `port` (Number) - Optional port of the SSL endpoint (server). Value of 0 indicates any port
<a id="nestedatt--mappings--key"></a>
### Nested Schema for `mappings.key`

Read-Only:

* `alias` (String) - unique alias for SSL Decryption Key
* `comment` (String)

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_key_map.example {alias}
```
