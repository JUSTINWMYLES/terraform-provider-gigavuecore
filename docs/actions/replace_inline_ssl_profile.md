---
page_title: "gigavuecore_replace_inline_ssl_profile Action - gigavuecore"
subcategory: ""
description: |-
  Replace inline SSL profile
---

# gigavuecore_replace_inline_ssl_profile Action

Replace inline SSL profile

## Example Usage

```terraform
action "gigavuecore_replace_inline_ssl_profile" "example" {
  config {
    alias = "example"
    body_alias = "example"
    body_cluster_id = "example"
    certificate = null
    cluster_id = "example"
    decrypt = null
    default_action = "example"
    high_avail = null
    inbound_tool_early_inspect = null
    key_map = null
    monitor = "example"
    network_group = null
    no_decrypt = null
    non_ssl_tcp = null
    one_arm = "example"
    resilient_inline = null
    rules = null
    split_proxy = null
    start_tls = null
    tcp = null
    tool = null
    tool_l3 = null
    url_cache = null
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Alias of the inline SSL profile
* `body_alias` (String, required) - Alias of the inline SSL profile
* `body_cluster_id` (String, optional) - id of the defining cluster
* `certificate` (Dynamic, optional)
* `cluster_id` (String, required) - Target Cluster ID
* `decrypt` (Dynamic, optional)
* `default_action` (String, optional) - Action to take if none of the profile rules match
* `high_avail` (Dynamic, optional)
* `inbound_tool_early_inspect` (Dynamic, optional)
* `key_map` (List(Dynamic), optional)
* `monitor` (String, optional)
* `network_group` (Dynamic, optional)
* `no_decrypt` (Dynamic, optional)
* `non_ssl_tcp` (Dynamic, optional)
* `one_arm` (String, optional)
* `resilient_inline` (Dynamic, optional)
* `rules` (List(Dynamic), optional) - inline SSL profile rules
* `split_proxy` (Dynamic, optional)
* `start_tls` (Dynamic, optional)
* `tcp` (Dynamic, optional)
* `tool` (Dynamic, optional)
* `tool_l3` (Dynamic, optional)
* `url_cache` (Dynamic, optional)
