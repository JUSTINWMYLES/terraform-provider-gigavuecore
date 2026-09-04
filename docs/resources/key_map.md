---
page_title: "gigavuecore_key_map Resource - gigavuecore"
subcategory: ""
description: |-
  for 4.2 nodes, the 'alias' in the SslDecryptionKeyMapConfigSpec body must reference one of the existing GsGroups
---

# gigavuecore_key_map Resource

for 4.2 nodes, the 'alias' in the SslDecryptionKeyMapConfigSpec body must reference one of the existing GsGroups

## Example Usage

```terraform
resource "gigavuecore_key_map" "example" {
  alias      = "example"
  cluster_id = "example"
  mappings = [{
  }]
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Alias of this mappings collection
* `cluster_id` (String, required) - id of the defining cluster
* `mappings` (Attributes List, required) - list of SSL Endpoints with corresponding Decryption Keys (see [below for nested schema](#nestedatt--mappings))

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

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
terraform import gigavuecore_key_map.example {alias}/{cluster_id}
```
