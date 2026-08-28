---
page_title: "gigavuecore_tcp_profile Resource - gigavuecore"
subcategory: ""
description: |-
  Get Apps TCP Profile for given alias
---

# gigavuecore_tcp_profile Resource

Get Apps TCP Profile for given alias

## Example Usage

```terraform
resource "gigavuecore_tcp_profile" "example" {
  alias            = null
  cluster_id       = null
  keep_alive_timer = null
  selective_ack    = null
  syn_retries      = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required)
* `cluster_id` (String, required) - Target Cluster ID
* `keep_alive_timer` (Number, optional)
* `selective_ack` (String, optional)
* `syn_retries` (Number, optional)

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `id` (String, computed)
* `keep_alive_timer` (Number, computed)
* `selective_ack` (String, computed)
* `syn_retries` (Number, computed)


