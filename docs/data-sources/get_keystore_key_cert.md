---
page_title: "gigavuecore_get_keystore_key_cert Data Source - gigavuecore"
subcategory: ""
description: |-
  Get the certificate
---

# gigavuecore_get_keystore_key_cert Data Source

Get the certificate

## Example Usage

```terraform
data "gigavuecore_get_keystore_key_cert" "example" {
  alias      = null
  cluster_id = null
  raw        = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the key
* `cluster_id` (String, required) - Target Cluster ID
* `raw` (Boolean, optional)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `certificate` (String, computed)


