---
page_title: "gigavuecore_add_map_flow_whitelist_rule Action - gigavuecore"
subcategory: ""
description: |-
  Add new flowWhitelistRule to a 'secondLevel/flowWhitelist' map
---

# gigavuecore_add_map_flow_whitelist_rule Action

Add new flowWhitelistRule to a 'secondLevel/flowWhitelist' map

## Example Usage

```terraform
action "gigavuecore_add_map_flow_whitelist_rule" "example" {
  config {
    alias = "example"
    flow5_g = {
      dnn                 = "example"
      type                = "example"
      whitelist_databases = [ "example" ]
    }
    gtp = {
      apn                 = "example"
      interface           = "example"
      type                = "example"
      version             = "example"
      whitelist_databases = [ "example" ]
    }
    rule_id = 0
    sip = {
      type = "example"
    }
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target map
* `flow5_g` (Attributes, optional) - Map Flow Whitelist 5g Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--flow5_g))
* `gtp` (Attributes, optional) - Map Flow Whitelist Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--gtp))
* `rule_id` (Number, required)
* `sip` (Attributes, optional) - Map Flow Whitelist Rule Sip match definition (see [below for nested schema](#nestedatt--sip))

<a id="nestedatt--flow5_g"></a>
### Nested Schema for `flow5_g`

Optional:

* `dnn` (String) - Domain Network Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `type` (String) - Set 5G WL-DB lookup type
* `whitelist_databases` (List of String) - Attach whitelist databases to the map

<a id="nestedatt--gtp"></a>
### Nested Schema for `gtp`

Optional:

* `apn` (String) - Access Point Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `interface` (String) - interface type. Mutually exclusive with version. required till H 5.6
* `type` (String) - Set GTP WL-DB lookup type
* `version` (String) - mutually exclusive with interface
* `whitelist_databases` (List of String) - Attach whitelist databases to the map

<a id="nestedatt--sip"></a>
### Nested Schema for `sip`

Optional:

* `type` (String) - all:Whitelist based on caller/callee/source/destination IP address, bothAddr: Whitelist source/destination IP address, bothId:Whitelist Caller/Callee Id's, calleeId: Whitelist Callee ID, callerId: Whitelist Caller ID, destIp: Whitelist based on destination IP address, srcIp: Whitelist based on Source IP address

