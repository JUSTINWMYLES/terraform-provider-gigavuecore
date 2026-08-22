---
page_title: "gigavuecore_redefine_snmp_throttle_config Action - gigavuecore"
subcategory: ""
description: |-
  Redefine Snmp Throttle config
---

# gigavuecore_redefine_snmp_throttle_config Action

Redefine Snmp Throttle config

## Example Usage

```terraform
action "gigavuecore_redefine_snmp_throttle_config" "example" {
  config {
    throttle_config_details = null
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `throttle_config_details` (List(Dynamic), optional) - list of SNMP Throttle Config Details on the node
