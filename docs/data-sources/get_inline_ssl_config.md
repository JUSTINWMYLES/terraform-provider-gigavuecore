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
  cluster_id = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - id of the defining cluster

### Attributes

In addition to all arguments above, the following attributes are exported:

* `caching` (Attributes, computed) (see [below for nested schema](#nestedatt--caching))
* `dhe_ciphersuit` (String, computed)
* `monitor` (Attributes, computed) - This property is moved to ssl profile for device version >=5.7 (see [below for nested schema](#nestedatt--monitor))
* `resumption` (Attributes, computed) (see [below for nested schema](#nestedatt--resumption))
* `ssl_versions` (Attributes, computed) (see [below for nested schema](#nestedatt--ssl_versions))
* `start_tls` (Attributes, computed) (see [below for nested schema](#nestedatt--start_tls))

<a id="nestedatt--caching"></a>
### Nested Schema for `caching`

Read-Only:

* `persistence` (Attributes) (see [below for nested schema](#nestedatt--caching--persistence))

<a id="nestedatt--caching--persistence"></a>
### Nested Schema for `caching.persistence`

Read-Only:

* `enable` (Boolean)

<a id="nestedatt--monitor"></a>
### Nested Schema for `monitor`

Read-Only:

* `enable` (Boolean)

<a id="nestedatt--resumption"></a>
### Nested Schema for `resumption`

Read-Only:

* `client` (Attributes) - enable client initiated resumption (for debug purposes only) (see [below for nested schema](#nestedatt--resumption--client))

<a id="nestedatt--resumption--client"></a>
### Nested Schema for `resumption.client`

Read-Only:

* `enable` (Boolean)

<a id="nestedatt--ssl_versions"></a>
### Nested Schema for `ssl_versions`

Read-Only:

* `connection_reset_action_for_max_version` (String) - Action to take to reset connection if TLS version is higher than configured
* `connection_reset_action_for_min_version` (String) - Action to take to reset connection if TLS version is lower than configured
* `max_version` (String) - maxVersion must be greater than minVersion
* `min_version` (String)

<a id="nestedatt--start_tls"></a>
### Nested Schema for `start_tls`

Read-Only:

* `enable` (Boolean)

