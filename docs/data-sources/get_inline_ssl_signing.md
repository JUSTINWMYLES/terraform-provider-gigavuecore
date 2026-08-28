---
page_title: "gigavuecore_get_inline_ssl_signing Data Source - gigavuecore"
subcategory: ""
description: |-
  Get inline SSL certificates and private-key configured state
---

# gigavuecore_get_inline_ssl_signing Data Source

Get inline SSL certificates and private-key configured state

## Example Usage

```terraform
data "gigavuecore_get_inline_ssl_signing" "example" {
  cluster_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - id of the defining cluster

### Attributes

In addition to all arguments above, the following attributes are exported:

* `primary_key_alias` (String, computed) - alias of the key from apps/keystore/keys
* `secondary_key_alias` (String, computed) - alias of the key from apps/keystore/keys


