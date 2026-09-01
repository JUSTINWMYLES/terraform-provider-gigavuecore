---
page_title: "gigavuecore_get_hsm_key_maps Data Source - gigavuecore"
subcategory: ""
description: |-
  Get HSM key maps
---

# gigavuecore_get_hsm_key_maps Data Source

Get HSM key maps

## Example Usage

```terraform
data "gigavuecore_get_hsm_key_maps" "example" {
  alias      = "example"
  cluster_id = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of target HSM Group
* `cluster_id` (String, required) - Target cluster ID.

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `address` (String) - IPv4 address of the SSL endpoint (server)
* `cluster_id` (String) - id of the defining cluster
* `key_name` (String) - key-name
* `key_token` (String) - key-token
* `port` (Number) - Port of the SSL endpoint (server). Value of 0 indicates any port
* `rfs_match` (String) - Is there a matching RFS key, yes or no
* `rule_id` (String) - Rule Id

