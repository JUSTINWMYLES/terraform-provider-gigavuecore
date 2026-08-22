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
  alias = null
  mappings = []
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Alias of this mappings collection
* `mappings` (List(Object({endpoint, key})), required) - list of SSL Endpoints with corresponding Decryption Keys

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `cluster_id` (String, computed) - id of the defining cluster

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_key_map.example {alias}
```
