---
page_title: "gigavuecore_load_gtp_whitelist_entry Data Source - gigavuecore"
subcategory: ""
description: |-
  Check if Whitelist Entry exists
---

# gigavuecore_load_gtp_whitelist_entry Data Source

Check if Whitelist Entry exists

## Example Usage

```terraform
data "gigavuecore_load_gtp_whitelist_entry" "example" {
  alias      = null
  cluster_id = null
  imsi       = null
  ran        = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target GTP Whitelist
* `cluster_id` (String, required) - Target Cluster ID
* `imsi` (String, required) - imsi-based whitelist entry being queried
* `ran` (String, optional) - nci value should prefix 0x

### Attributes

In addition to all arguments above, the following attributes are exported:

* `active_sessions` (Number, computed) - Number of active sessions
* `ran` (String, computed) - nci value should prefix 0x


