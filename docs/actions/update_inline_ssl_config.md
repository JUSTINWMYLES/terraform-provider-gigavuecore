---
page_title: "gigavuecore_update_inline_ssl_config Action - gigavuecore"
subcategory: ""
description: |-
  Update inline SSL global configuration
---

# gigavuecore_update_inline_ssl_config Action

Update inline SSL global configuration

## Example Usage

```terraform
action "gigavuecore_update_inline_ssl_config" "example" {
  config {
    body_cluster_id = "example"
    caching = {
      persistence = {
        enable = true
      }
    }
    cluster_id     = "example"
    dhe_ciphersuit = "example"
    monitor = {
      enable = true
    }
    resumption = {
      client = {
        enable = true
      }
    }
    ssl_versions = {
      connection_reset_action_for_max_version = "example"
      connection_reset_action_for_min_version = "example"
      max_version                             = "example"
      min_version                             = "example"
    }
    start_tls = {
      enable = true
    }
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `body_cluster_id` (String, optional) - id of the defining cluster
* `caching` (Attributes, optional) (see [below for nested schema](#nestedatt--caching))
* `cluster_id` (String, required) - Target Cluster ID
* `dhe_ciphersuit` (String, optional)
* `monitor` (Attributes, optional) - This property is moved to ssl profile for device version >=5.7 (see [below for nested schema](#nestedatt--monitor))
* `resumption` (Attributes, optional) (see [below for nested schema](#nestedatt--resumption))
* `ssl_versions` (Attributes, optional) (see [below for nested schema](#nestedatt--ssl_versions))
* `start_tls` (Attributes, optional) (see [below for nested schema](#nestedatt--start_tls))

<a id="nestedatt--caching"></a>
### Nested Schema for `caching`

Optional:

* `persistence` (Attributes) (see [below for nested schema](#nestedatt--caching--persistence))

<a id="nestedatt--caching--persistence"></a>
### Nested Schema for `caching.persistence`

Optional:

* `enable` (Boolean)

<a id="nestedatt--monitor"></a>
### Nested Schema for `monitor`

Optional:

* `enable` (Boolean)

<a id="nestedatt--resumption"></a>
### Nested Schema for `resumption`

Optional:

* `client` (Attributes) - enable client initiated resumption (for debug purposes only) (see [below for nested schema](#nestedatt--resumption--client))

<a id="nestedatt--resumption--client"></a>
### Nested Schema for `resumption.client`

Optional:

* `enable` (Boolean)

<a id="nestedatt--ssl_versions"></a>
### Nested Schema for `ssl_versions`

Required:

* `max_version` (String) - maxVersion must be greater than minVersion
* `min_version` (String)

Optional:

* `connection_reset_action_for_max_version` (String) - Action to take to reset connection if TLS version is higher than configured
* `connection_reset_action_for_min_version` (String) - Action to take to reset connection if TLS version is lower than configured

<a id="nestedatt--start_tls"></a>
### Nested Schema for `start_tls`

Optional:

* `enable` (Boolean)

