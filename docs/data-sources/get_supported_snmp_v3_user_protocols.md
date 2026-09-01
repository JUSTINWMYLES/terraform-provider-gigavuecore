---
page_title: "gigavuecore_get_supported_snmp_v3_user_protocols Data Source - gigavuecore"
subcategory: ""
description: |-
  Find SNMPv3 User Protocols by deviceice swVersion
---

# gigavuecore_get_supported_snmp_v3_user_protocols Data Source

Find SNMPv3 User Protocols by deviceice swVersion

## Example Usage

```terraform
data "gigavuecore_get_supported_snmp_v3_user_protocols" "example" {
  sw_version = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `sw_version` (String, optional) - Device SW Version

### Attributes

In addition to all arguments above, the following attributes are exported:

* `auth_protocol` (String, computed)
* `priv_protocol` (String, computed)


