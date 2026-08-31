---
page_title: "gigavuecore_get_all_tcp_profiles Data Source - gigavuecore"
subcategory: ""
description: |-
  Get all Apps TCP Profile
---

# gigavuecore_get_all_tcp_profiles Data Source

Get all Apps TCP Profile

## Example Usage

```terraform
data "gigavuecore_get_all_tcp_profiles" "example" {
  cluster_id = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `alias` (String)
* `keep_alive_timer` (Number)
* `selective_ack` (String)
* `syn_retries` (Number)

