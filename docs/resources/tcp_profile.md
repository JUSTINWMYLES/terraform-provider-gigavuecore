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
  alias            = "example"
  cluster_id       = "example"
  keep_alive_timer = 0
  selective_ack    = "example"
  syn_retries      = 0
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

* `keep_alive_timer` (Number, computed)
* `selective_ack` (String, computed)
* `syn_retries` (Number, computed)


## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_tcp_profile.example {alias}
```
