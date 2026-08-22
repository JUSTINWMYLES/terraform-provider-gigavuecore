---
page_title: "gigavuecore_get_inline_ssl_config Data Source - gigavuecore"
subcategory: ""
description: |-
  Load inline SSL global configuration
---

# gigavuecore_get_inline_ssl_config Data Source

Load inline SSL global configuration

## Example Usage

```terraform
data "gigavuecore_get_inline_ssl_config" "example" {
  cluster_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - id of the defining cluster

### Attributes

In addition to all arguments above, the following attributes are exported:

* `caching` (Object({persistence}), computed)
  * `persistence` (Object({enable}), computed)
    * `enable` (Bool, computed)
* `dhe_ciphersuit` (String, computed)
* `monitor` (Object({enable}), computed)
  * `enable` (Bool, computed)
* `resumption` (Object({client}), computed)
  * `client` (Object({enable}), computed)
    * `enable` (Bool, computed)
* `ssl_versions` (Object({connection_reset_action_for_max_version, connection_reset_action_for_min_version, max_version, min_version}), computed)
  * `connection_reset_action_for_max_version` (String, computed) - Action to take to reset connection if TLS version is higher than configured
  * `connection_reset_action_for_min_version` (String, computed) - Action to take to reset connection if TLS version is lower than configured
  * `max_version` (String, computed) - maxVersion must be greater than minVersion
  * `min_version` (String, computed)
* `start_tls` (Object({enable}), computed)
  * `enable` (Bool, computed)

