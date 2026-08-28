---
page_title: "gigavuecore_get_all_hsms Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all HSM
---

# gigavuecore_get_all_hsms Data Source

Load all HSM

## Example Usage

```terraform
data "gigavuecore_get_all_hsms" "example" {
}
```

## Schema

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `alias` (String) - Hsm alias
* `esn` (String) - the ESN of the KNETI key for a given IP address
* `hsm_ip` (String) - hsm server address
* `hsm_port` (Number) - server port number, the port to use when connecting to the given nethsm
* `kneti` (String) - the hash of the KNETI key for a given IP address
* `operational_status` (String) - operational status of hsm
* `partition_label` (String) - hsm partition label
* `partition_password` (String) - hsm partition password
* `server_password` (String) - hsm server password
* `server_username` (String) - hsm server username
* `type` (String)

