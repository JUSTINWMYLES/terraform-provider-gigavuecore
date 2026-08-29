---
page_title: "gigavuecore_add_traffic_flows_draft Action - gigavuecore"
subcategory: ""
description: |-
  Add a new traffic flows draft
---

# gigavuecore_add_traffic_flows_draft Action

Add a new traffic flows draft

## Example Usage

```terraform
action "gigavuecore_add_traffic_flows_draft" "example" {
  config {
    alias           = "example"
    comment         = "example"
    deployment_type = "example"
    enable          = true
    flows           = [ "example" ]
    has_draft       = true
    priority_type   = "example"
    sources_and_rules = [{
      alias = "example"
      components = [{
        cluster_id = "example"
        components = [{
          ids  = [ "example" ]
          type = "example"
        }]
      }]
      inline_traffic_path = "example"
      inline_traffic_type = "example"
      ip_rewrite = {
        dst_ip = "example"
        src_ip = "example"
      }
      rewrite = {
        dst_mac = "example"
        src_mac = "example"
      }
      rule_matching = "example"
      rule_type     = "example"
      rules = {
        drop_rules = [{
          bidi    = true
          comment = "example"
          ip_rewrite = {
            dst_ip = "example"
            src_ip = "example"
          }
          matches = [ "example" ]
          rewrite = {
            dst_mac = "example"
            src_mac = "example"
          }
          rule_id = 0
          vlan_tag = {
            tag_protocol_id = "example"
            vlan_action     = "example"
            vlan_id         = 0
          }
        }]
        pass_rules = [{
          bidi    = true
          comment = "example"
          ip_rewrite = {
            dst_ip = "example"
            src_ip = "example"
          }
          matches = [ "example" ]
          rewrite = {
            dst_mac = "example"
            src_mac = "example"
          }
          rule_id = 0
          vlan_tag = {
            tag_protocol_id = "example"
            vlan_action     = "example"
            vlan_id         = 0
          }
        }]
      }
      tags = [{
        tag_key    = "example"
        tag_values = [ "example" ]
      }]
      traffic_type = "example"
      vlan_tag = {
        tag_protocol_id = "example"
        vlan_action     = "example"
        vlan_id         = 0
      }
    }]
    tags = [{
      tag_key    = "example"
      tag_values = [ "example" ]
    }]
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - unique map alias
* `comment` (String, optional)
* `deployment_type` (String, required) - 'regular' maps are from network/hybrid ports to tool/hybrid/gigastream; 'inline' maps are from inline ports to inline/tool/hybrid/gigastream; 'firstLevel' are from network/hybrid ports to vPorts/tool/hybrid/gigastream; 'secondLevel' maps are from vPorts to tool/hybrid/gigastream; 'inlineFirstLevel is from inline network to vport; 'inlineSecondLevel' is from vport to inline tool; 'transitLevel' is from vport to vport
* `enable` (Boolean, optional) - enable/disable map, applicable only to first level maps
* `flows` (List of Dynamic, required)
* `has_draft` (Boolean, optional)
* `priority_type` (String, optional) - Define the map priority to be LOWEST or HIGHEST. Default priority is LOWEST.
* `sources_and_rules` (Attributes List, required) (see [below for nested schema](#nestedatt--sources_and_rules))
* `tags` (Attributes List, optional) (see [below for nested schema](#nestedatt--tags))

<a id="nestedatt--sources_and_rules"></a>
### Nested Schema for `sources_and_rules`

Required:

* `alias` (String) - unique sourcesAndRules alias
* `rule_type` (String) - 'byRule' is applicable to all map types; 'collector' is applicable to 'regular', 'inline' maps; 'passAll' is applicable to 'regular' and 'inline' maps
* `rules` (Attributes) - Map Rules Container. Private class (see [below for nested schema](#nestedatt--sources_and_rules--rules))

Optional:

* `components` (Attributes List) (see [below for nested schema](#nestedatt--sources_and_rules--components))
* `inline_traffic_path` (String) - Only applicable for 'inline' map types, in which case defaults to 'normal'. When set to 'bypass', the 'dstPorts' must be empty
* `inline_traffic_type` (String) - Only applicable for 'inline/passAll' map types, in which case defaults to 'symmetric'. For 'asymmetric' maps, 'srcPort' must be of type 'inline-net'
* `ip_rewrite` (Attributes) - IpRewrite options on the packets (see [below for nested schema](#nestedatt--sources_and_rules--ip_rewrite))
* `rewrite` (Attributes) - Rewrite options on the packets (see [below for nested schema](#nestedatt--sources_and_rules--rewrite))
* `rule_matching` (String) - If set to 'blacklist', packet are passed when no 'drop' rules are matched. This field is only valid for 'regular/byRule' map types. (maps into 'no-rule-match' CLI command)
* `tags` (Attributes List) (see [below for nested schema](#nestedatt--sources_and_rules--tags))
* `traffic_type` (String) - Only applicable for 'firstLevel/byRule' map types, in which case defaults to 'user'
* `vlan_tag` (Attributes) - This field is only valid for 'regular/byRule', 'collector' map types (see [below for nested schema](#nestedatt--sources_and_rules--vlan_tag))

<a id="nestedatt--sources_and_rules--rules"></a>
### Nested Schema for `sources_and_rules.rules`

Optional:

* `drop_rules` (Attributes Set) (see [below for nested schema](#nestedatt--sources_and_rules--rules--drop_rules))
* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--sources_and_rules--rules--pass_rules))

<a id="nestedatt--sources_and_rules--rules--drop_rules"></a>
### Nested Schema for `sources_and_rules.rules.drop_rules`

Required:

* `matches` (Set of Dynamic) - Set of rule's matching elements. Within a rule, matching elements of the the same type MUST NOT be repeated. The 'position' property of each matching element is not relevant for this rule type as only the outer headers are matched
* `rule_id` (Number)

Optional:

* `bidi` (Boolean)
* `comment` (String)
* `ip_rewrite` (Attributes) - IpRewrite options on the packets (see [below for nested schema](#nestedatt--sources_and_rules--rules--drop_rules--ip_rewrite))
* `rewrite` (Attributes) - Rewrite options on the packets (see [below for nested schema](#nestedatt--sources_and_rules--rules--drop_rules--rewrite))
* `vlan_tag` (Attributes) - This field is only  applicable for pass rules (see [below for nested schema](#nestedatt--sources_and_rules--rules--drop_rules--vlan_tag))

<a id="nestedatt--sources_and_rules--rules--drop_rules--ip_rewrite"></a>
### Nested Schema for `sources_and_rules.rules.drop_rules.ip_rewrite`

Optional:

* `dst_ip` (String)
* `src_ip` (String)

<a id="nestedatt--sources_and_rules--rules--drop_rules--rewrite"></a>
### Nested Schema for `sources_and_rules.rules.drop_rules.rewrite`

Optional:

* `dst_mac` (String)
* `src_mac` (String)

<a id="nestedatt--sources_and_rules--rules--drop_rules--vlan_tag"></a>
### Nested Schema for `sources_and_rules.rules.drop_rules.vlan_tag`

Required:

* `vlan_action` (String)

Optional:

* `tag_protocol_id` (String)
* `vlan_id` (Number)

<a id="nestedatt--sources_and_rules--rules--pass_rules"></a>
### Nested Schema for `sources_and_rules.rules.pass_rules`

Required:

* `matches` (Set of Dynamic) - Set of rule's matching elements. Within a rule, matching elements of the the same type MUST NOT be repeated. The 'position' property of each matching element is not relevant for this rule type as only the outer headers are matched
* `rule_id` (Number)

Optional:

* `bidi` (Boolean)
* `comment` (String)
* `ip_rewrite` (Attributes) - IpRewrite options on the packets (see [below for nested schema](#nestedatt--sources_and_rules--rules--pass_rules--ip_rewrite))
* `rewrite` (Attributes) - Rewrite options on the packets (see [below for nested schema](#nestedatt--sources_and_rules--rules--pass_rules--rewrite))
* `vlan_tag` (Attributes) - This field is only  applicable for pass rules (see [below for nested schema](#nestedatt--sources_and_rules--rules--pass_rules--vlan_tag))

<a id="nestedatt--sources_and_rules--rules--pass_rules--ip_rewrite"></a>
### Nested Schema for `sources_and_rules.rules.pass_rules.ip_rewrite`

Optional:

* `dst_ip` (String)
* `src_ip` (String)

<a id="nestedatt--sources_and_rules--rules--pass_rules--rewrite"></a>
### Nested Schema for `sources_and_rules.rules.pass_rules.rewrite`

Optional:

* `dst_mac` (String)
* `src_mac` (String)

<a id="nestedatt--sources_and_rules--rules--pass_rules--vlan_tag"></a>
### Nested Schema for `sources_and_rules.rules.pass_rules.vlan_tag`

Required:

* `vlan_action` (String)

Optional:

* `tag_protocol_id` (String)
* `vlan_id` (Number)

<a id="nestedatt--sources_and_rules--components"></a>
### Nested Schema for `sources_and_rules.components`

Required:

* `cluster_id` (String) - id of the defining cluster

Optional:

* `components` (Attributes List) (see [below for nested schema](#nestedatt--sources_and_rules--components--components))

<a id="nestedatt--sources_and_rules--components--components"></a>
### Nested Schema for `sources_and_rules.components.components`

Required:

* `ids` (List of Dynamic)
* `type` (String)

<a id="nestedatt--sources_and_rules--ip_rewrite"></a>
### Nested Schema for `sources_and_rules.ip_rewrite`

Optional:

* `dst_ip` (String)
* `src_ip` (String)

<a id="nestedatt--sources_and_rules--rewrite"></a>
### Nested Schema for `sources_and_rules.rewrite`

Optional:

* `dst_mac` (String)
* `src_mac` (String)

<a id="nestedatt--sources_and_rules--tags"></a>
### Nested Schema for `sources_and_rules.tags`

Required:

* `tag_key` (String) - Name of the tag
* `tag_values` (List of String) - All possible values of the tag

<a id="nestedatt--sources_and_rules--vlan_tag"></a>
### Nested Schema for `sources_and_rules.vlan_tag`

Required:

* `vlan_action` (String)

Optional:

* `tag_protocol_id` (String)
* `vlan_id` (Number)

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Required:

* `tag_key` (String) - Name of the tag
* `tag_values` (List of String) - All possible values of the tag

