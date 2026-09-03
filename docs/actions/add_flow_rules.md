---
page_title: "gigavuecore_add_flow_rules Action - gigavuecore"
subcategory: ""
description: |-
  Add rules within a flow
---

# gigavuecore_add_flow_rules Action

Add rules within a flow

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_add_flow_rules" "example" {
  config {
    alias      = "example"
    flow_alias = "example"
    flow_rules = {
      drop_rules = [{
        gtp = {
          imei      = "*"
          imsi      = "*"
          interface = "Gn"
          msisdn    = "*"
          version   = "any"
        }
        rule_id = 1
      }]
      pass_rules = [{
        gtp = {
          imei      = "*"
          imsi      = "*"
          interface = "Gn"
          msisdn    = "*"
          version   = "any"
        }
        rule_id = 1
      }]
    }
    flow_sample5_g_overlap_rules = {
      pass_rules = [{
        comment = "example"
        flow5_g = {
          dnn      = "example"
          gpsi     = "*"
          nas_5_qi = "0"
          nci      = "*"
          nsiid    = "0"
          pei      = "*"
          plmn_id  = "*"
          supi     = "*"
          tac      = "*"
        }
        percentage = 0
        rule_id    = 1
      }]
    }
    flow_sample5_g_rules = {
      pass_rules = [{
        comment = "example"
        flow5_g = {
          dnn     = "example"
          gpsi    = "*"
          nci     = "a1b2c3d4e"
          nsiid   = "0"
          pei     = "*"
          plmn_id = "123.45"
          supi    = "*"
          tac     = "*"
        }
        percentage = 0
        priority   = 1
        rule_id    = 1
      }]
    }
    flow_sample_diameter_rules = {
      pass_rules = [{
        diameter = {
          user_name = "*"
        }
        interface  = "s6a"
        percentage = 0
        rule_id    = 1
      }]
    }
    flow_sample_overlap_rules = {
      pass_rules = [{
        comment = "example"
        gtp = {
          apn       = "example"
          eci       = "a1b2c3d4"
          imei      = "*"
          imsi      = "*"
          interface = "Gn"
          msisdn    = "*"
          nas_5_qi  = "0"
          nci       = "*"
          plmn_id   = "123.45"
          qci       = 0
          snssai    = "0"
          tac       = "abc1"
          tac_5_g   = "*"
          version   = "any"
        }
        percentage      = 0
        periodic_recalc = true
        priority        = 1
        rule_id         = 1
      }]
    }
    flow_sample_rules = {
      pass_rules = [{
        comment = "example"
        gtp = {
          apn       = "example"
          eci       = "a1b2c3d4"
          imei      = "*"
          imsi      = "*"
          interface = "Gn"
          msisdn    = "*"
          nas_5_qi  = "0"
          nci       = "*"
          plmn_id   = "123.45"
          qci       = 0
          snssai    = "0"
          tac       = "abc1"
          tac_5_g   = "*"
          version   = "any"
        }
        percentage      = 0
        periodic_recalc = true
        priority        = 1
        rule_id         = 1
      }]
    }
    flow_sample_sip_rules = {
      pass_rules = [{
        percentage = 0
        rule_id    = 1
        sip = {
          callee_id = "example"
          callee_id_range = {
            max_value = "example"
            value     = "example"
          }
          caller_id = "example"
          caller_id_range = {
            max_value = "example"
            value     = "example"
          }
          id_range = {
            max_value = "example"
            value     = "example"
          }
        }
      }]
    }
    flow_whitelist5_g_overlap_rules = {
      dnn  = "example"
      type = "example"
    }
    flow_whitelist5_g_rules = {
      dnn                 = "example"
      type                = "example"
      whitelist_databases = ["example"]
    }
    flow_whitelist_overlap_rules = {
      pass_rules = [{
        flow5_g = {
          dnn                 = "example"
          type                = "example"
          whitelist_databases = ["example"]
        }
        gtp = {
          apn                 = "example"
          interface           = "Gn"
          type                = "example"
          version             = "v1"
          whitelist_databases = ["example"]
        }
        rule_id = 1
        sip = {
          type = "all"
        }
      }]
    }
    flow_whitelist_rules = {
      pass_rules = [{
        flow5_g = {
          dnn                 = "example"
          type                = "example"
          whitelist_databases = ["example"]
        }
        gtp = {
          apn                 = "example"
          interface           = "Gn"
          type                = "example"
          version             = "v1"
          whitelist_databases = ["example"]
        }
        rule_id = 1
        sip = {
          type = "all"
        }
      }]
    }
    gs_rules = {
      drop_rules = [{
        comment = "example"
        matches = ["example"]
        rule_id = 1
      }]
      pass_rules = [{
        comment = "example"
        matches = ["example"]
        rule_id = 1
      }]
    }
    rule_type      = "example"
    sub_flow_alias = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Traffic Flows alias or ID
* `flow_alias` (String, required) - Flow alias or ID
* `flow_rules` (Attributes, optional) - Map Flow Rules Container. Private class (see [below for nested schema](#nestedatt--flow_rules))
* `flow_sample5_g_overlap_rules` (Attributes, optional) - Map Flow Sample 5g Overlap Rules Container. Private class (see [below for nested schema](#nestedatt--flow_sample5_g_overlap_rules))
* `flow_sample5_g_rules` (Attributes, optional) - Map Flow Sample 5g Rules Container. Private class (see [below for nested schema](#nestedatt--flow_sample5_g_rules))
* `flow_sample_diameter_rules` (Attributes, optional) - Map Flow Sample Diameter Rules Container. Private class (see [below for nested schema](#nestedatt--flow_sample_diameter_rules))
* `flow_sample_overlap_rules` (Attributes, optional) - Map Flow Sample Overlap Rules Container. Private class (see [below for nested schema](#nestedatt--flow_sample_overlap_rules))
* `flow_sample_rules` (Attributes, optional) - Map Flow Sample Rules Container. Private class (see [below for nested schema](#nestedatt--flow_sample_rules))
* `flow_sample_sip_rules` (Attributes, optional) - Map Flow Sample Sip Rules Container. Private class (see [below for nested schema](#nestedatt--flow_sample_sip_rules))
* `flow_whitelist5_g_overlap_rules` (Attributes, optional) - Map Flow Whitelist 5g Overlap Rule match Definition. Private class (see [below for nested schema](#nestedatt--flow_whitelist5_g_overlap_rules))
* `flow_whitelist5_g_rules` (Attributes, optional) - Map Flow Whitelist 5g Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--flow_whitelist5_g_rules))
* `flow_whitelist_overlap_rules` (Attributes, optional) - Map Flow Whitelist Overlap Rules Container. Private class (see [below for nested schema](#nestedatt--flow_whitelist_overlap_rules))
* `flow_whitelist_rules` (Attributes, optional) - Map Flow Whitelist Rules Container. Private class (see [below for nested schema](#nestedatt--flow_whitelist_rules))
* `gs_rules` (Attributes, optional) - Map GigaSMART Rules Container. Private class (see [below for nested schema](#nestedatt--gs_rules))
* `rule_type` (String, required) - Type of rule
* `sub_flow_alias` (String, required) - Sub-flow alias or ID

<a id="nestedatt--flow_rules"></a>
### Nested Schema for `flow_rules`

Optional:

* `drop_rules` (Attributes Set) (see [below for nested schema](#nestedatt--flow_rules--drop_rules))
* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--flow_rules--pass_rules))

<a id="nestedatt--flow_rules--drop_rules"></a>
### Nested Schema for `flow_rules.drop_rules`

Required:

* `gtp` (Attributes) - Map Flow Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--flow_rules--drop_rules--gtp))
* `rule_id` (Number)

<a id="nestedatt--flow_rules--drop_rules--gtp"></a>
### Nested Schema for `flow_rules.drop_rules.gtp`

Optional:

* `imei` (String) - mutually exclusive with 'imsi' and 'msisdn'. If '\*' is added at the end of the value, it is treated as prefix
* `imsi` (String) - mutually exclusive with 'imei' and 'msisdn'. If '\*' is added at the end of the value, it is treated as prefix
* `interface` (String) - interface type. Mutually exclusive with version
* `msisdn` (String) - mutually exclusive with 'imsi' and 'imei'. If '\*' is added at the end of the value, it is treated as prefix
* `version` (String) - mutually exclusive with interface

<a id="nestedatt--flow_rules--pass_rules"></a>
### Nested Schema for `flow_rules.pass_rules`

Required:

* `gtp` (Attributes) - Map Flow Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--flow_rules--pass_rules--gtp))
* `rule_id` (Number)

<a id="nestedatt--flow_rules--pass_rules--gtp"></a>
### Nested Schema for `flow_rules.pass_rules.gtp`

Optional:

* `imei` (String) - mutually exclusive with 'imsi' and 'msisdn'. If '\*' is added at the end of the value, it is treated as prefix
* `imsi` (String) - mutually exclusive with 'imei' and 'msisdn'. If '\*' is added at the end of the value, it is treated as prefix
* `interface` (String) - interface type. Mutually exclusive with version
* `msisdn` (String) - mutually exclusive with 'imsi' and 'imei'. If '\*' is added at the end of the value, it is treated as prefix
* `version` (String) - mutually exclusive with interface

<a id="nestedatt--flow_sample5_g_overlap_rules"></a>
### Nested Schema for `flow_sample5_g_overlap_rules`

Optional:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--flow_sample5_g_overlap_rules--pass_rules))

<a id="nestedatt--flow_sample5_g_overlap_rules--pass_rules"></a>
### Nested Schema for `flow_sample5_g_overlap_rules.pass_rules`

Required:

* `flow5_g` (Attributes) - Map Flow Sample Overlap Rule 5g match Definition. Private class (see [below for nested schema](#nestedatt--flow_sample5_g_overlap_rules--pass_rules--flow5_g))
* `percentage` (Number)
* `rule_id` (Number)

Optional:

* `comment` (String)

<a id="nestedatt--flow_sample5_g_overlap_rules--pass_rules--flow5_g"></a>
### Nested Schema for `flow_sample5_g_overlap_rules.pass_rules.flow5_g`

Optional:

* `dnn` (String) - Access Point Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `gpsi` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `nas_5_qi` (String) - 5G QoS Indicator. Valid 5qi value <1 - 255>
* `nci` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `nsiid` (String) - If '\*' is added at the end of the SD value, it is treated as prefix
* `pei` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `plmn_id` (String) - If '\*' is added at the end of the MNC value, it is treated as prefix
* `supi` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `tac` (String) - If '\*' is added at the end of the value, it is treated as prefix

<a id="nestedatt--flow_sample5_g_rules"></a>
### Nested Schema for `flow_sample5_g_rules`

Optional:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--flow_sample5_g_rules--pass_rules))

<a id="nestedatt--flow_sample5_g_rules--pass_rules"></a>
### Nested Schema for `flow_sample5_g_rules.pass_rules`

Required:

* `flow5_g` (Attributes) - Map Flow Sample Rule 5g match Definition. Private class (see [below for nested schema](#nestedatt--flow_sample5_g_rules--pass_rules--flow5_g))
* `percentage` (Number)
* `rule_id` (Number)

Optional:

* `comment` (String)
* `priority` (Number)

<a id="nestedatt--flow_sample5_g_rules--pass_rules--flow5_g"></a>
### Nested Schema for `flow_sample5_g_rules.pass_rules.flow5_g`

Optional:

* `dnn` (String) - Access Point Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `gpsi` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `nci` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `nsiid` (String) - If '\*' is added at the end of the SD value, it is treated as prefix
* `pei` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `plmn_id` (String) - If '\*' is added at the end of the MNC value, it is treated as prefix
* `supi` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `tac` (String) - If '\*' is added at the end of the value, it is treated as prefix

<a id="nestedatt--flow_sample_diameter_rules"></a>
### Nested Schema for `flow_sample_diameter_rules`

Optional:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--flow_sample_diameter_rules--pass_rules))

<a id="nestedatt--flow_sample_diameter_rules--pass_rules"></a>
### Nested Schema for `flow_sample_diameter_rules.pass_rules`

Required:

* `diameter` (Attributes) - Map Flow Sample Diameter Rule Definition (see [below for nested schema](#nestedatt--flow_sample_diameter_rules--pass_rules--diameter))
* `interface` (String) - interface type
* `percentage` (Number)
* `rule_id` (Number)

<a id="nestedatt--flow_sample_diameter_rules--pass_rules--diameter"></a>
### Nested Schema for `flow_sample_diameter_rules.pass_rules.diameter`

Optional:

* `user_name` (String) - If '\*' is added at the end of the value, it is treated as prefix

<a id="nestedatt--flow_sample_overlap_rules"></a>
### Nested Schema for `flow_sample_overlap_rules`

Optional:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--flow_sample_overlap_rules--pass_rules))

<a id="nestedatt--flow_sample_overlap_rules--pass_rules"></a>
### Nested Schema for `flow_sample_overlap_rules.pass_rules`

Required:

* `gtp` (Attributes) - Map Flow Sample Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--flow_sample_overlap_rules--pass_rules--gtp))
* `percentage` (Number)
* `rule_id` (Number)

Optional:

* `comment` (String)
* `periodic_recalc` (Boolean) - Enable Periodic Recalc for rotational sampling. Map look up in the data path based on this flag
* `priority` (Number) - maximum value is equal to the number of rules upon completion of the request

<a id="nestedatt--flow_sample_overlap_rules--pass_rules--gtp"></a>
### Nested Schema for `flow_sample_overlap_rules.pass_rules.gtp`

Optional:

* `apn` (String) - Access Point Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `eci` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `imei` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `imsi` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `interface` (String) - interface type. Mutually exclusive with version
* `msisdn` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `nas_5_qi` (String) - 5G QoS Indicator. Valid 5qi value <1 - 255>
* `nci` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `plmn_id` (String) - If '\*' is added at the end of the MNC value, it is treated as prefix
* `qci` (Number) - QoS Class Indicator
* `snssai` (String) - If '\*' is added at the end of the SD value, it is treated as prefix
* `tac` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `tac_5_g` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `version` (String) - mutually exclusive with interface

<a id="nestedatt--flow_sample_rules"></a>
### Nested Schema for `flow_sample_rules`

Optional:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--flow_sample_rules--pass_rules))

<a id="nestedatt--flow_sample_rules--pass_rules"></a>
### Nested Schema for `flow_sample_rules.pass_rules`

Required:

* `gtp` (Attributes) - Map Flow Sample Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--flow_sample_rules--pass_rules--gtp))
* `percentage` (Number)
* `rule_id` (Number)

Optional:

* `comment` (String)
* `periodic_recalc` (Boolean) - Enable Periodic Recalc for rotational sampling. Map look up in the data path based on this flag
* `priority` (Number) - maximum value is equal to the number of rules upon completion of the request

<a id="nestedatt--flow_sample_rules--pass_rules--gtp"></a>
### Nested Schema for `flow_sample_rules.pass_rules.gtp`

Optional:

* `apn` (String) - Access Point Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `eci` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `imei` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `imsi` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `interface` (String) - interface type. Mutually exclusive with version
* `msisdn` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `nas_5_qi` (String) - 5G QoS Indicator. Valid 5qi value <1 - 255>
* `nci` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `plmn_id` (String) - If '\*' is added at the end of the MNC value, it is treated as prefix
* `qci` (Number) - QoS Class Indicator
* `snssai` (String) - If '\*' is added at the end of the SD value, it is treated as prefix
* `tac` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `tac_5_g` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `version` (String) - mutually exclusive with interface

<a id="nestedatt--flow_sample_sip_rules"></a>
### Nested Schema for `flow_sample_sip_rules`

Optional:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--flow_sample_sip_rules--pass_rules))

<a id="nestedatt--flow_sample_sip_rules--pass_rules"></a>
### Nested Schema for `flow_sample_sip_rules.pass_rules`

Required:

* `percentage` (Number)
* `rule_id` (Number)
* `sip` (Attributes) (see [below for nested schema](#nestedatt--flow_sample_sip_rules--pass_rules--sip))

<a id="nestedatt--flow_sample_sip_rules--pass_rules--sip"></a>
### Nested Schema for `flow_sample_sip_rules.pass_rules.sip`

Optional:

* `callee_id` (String) - sip callee id
* `callee_id_range` (Attributes) (see [below for nested schema](#nestedatt--flow_sample_sip_rules--pass_rules--sip--callee_id_range))
* `caller_id` (String) - sip caller id
* `caller_id_range` (Attributes) (see [below for nested schema](#nestedatt--flow_sample_sip_rules--pass_rules--sip--caller_id_range))
* `id_range` (Attributes) (see [below for nested schema](#nestedatt--flow_sample_sip_rules--pass_rules--sip--id_range))

<a id="nestedatt--flow_sample_sip_rules--pass_rules--sip--callee_id_range"></a>
### Nested Schema for `flow_sample_sip_rules.pass_rules.sip.callee_id_range`

Required:

* `max_value` (String)
* `value` (String)

<a id="nestedatt--flow_sample_sip_rules--pass_rules--sip--caller_id_range"></a>
### Nested Schema for `flow_sample_sip_rules.pass_rules.sip.caller_id_range`

Required:

* `max_value` (String)
* `value` (String)

<a id="nestedatt--flow_sample_sip_rules--pass_rules--sip--id_range"></a>
### Nested Schema for `flow_sample_sip_rules.pass_rules.sip.id_range`

Required:

* `max_value` (String)
* `value` (String)

<a id="nestedatt--flow_whitelist5_g_overlap_rules"></a>
### Nested Schema for `flow_whitelist5_g_overlap_rules`

Optional:

* `dnn` (String) - Domain Network Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `type` (String) - Set 5G WL-DB lookup type

<a id="nestedatt--flow_whitelist5_g_rules"></a>
### Nested Schema for `flow_whitelist5_g_rules`

Optional:

* `dnn` (String) - Domain Network Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `type` (String) - Set 5G WL-DB lookup type
* `whitelist_databases` (List of String) - Attach whitelist databases to the map

<a id="nestedatt--flow_whitelist_overlap_rules"></a>
### Nested Schema for `flow_whitelist_overlap_rules`

Optional:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--flow_whitelist_overlap_rules--pass_rules))

<a id="nestedatt--flow_whitelist_overlap_rules--pass_rules"></a>
### Nested Schema for `flow_whitelist_overlap_rules.pass_rules`

Required:

* `rule_id` (Number)

Optional:

* `flow5_g` (Attributes) - Map Flow Whitelist 5g Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--flow_whitelist_overlap_rules--pass_rules--flow5_g))
* `gtp` (Attributes) - Map Flow Whitelist Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--flow_whitelist_overlap_rules--pass_rules--gtp))
* `sip` (Attributes) - Map Flow Whitelist Rule Sip match definition (see [below for nested schema](#nestedatt--flow_whitelist_overlap_rules--pass_rules--sip))

<a id="nestedatt--flow_whitelist_overlap_rules--pass_rules--flow5_g"></a>
### Nested Schema for `flow_whitelist_overlap_rules.pass_rules.flow5_g`

Optional:

* `dnn` (String) - Domain Network Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `type` (String) - Set 5G WL-DB lookup type
* `whitelist_databases` (List of String) - Attach whitelist databases to the map

<a id="nestedatt--flow_whitelist_overlap_rules--pass_rules--gtp"></a>
### Nested Schema for `flow_whitelist_overlap_rules.pass_rules.gtp`

Optional:

* `apn` (String) - Access Point Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `interface` (String) - interface type. Mutually exclusive with version. required till H 5.6
* `type` (String) - Set GTP WL-DB lookup type
* `version` (String) - mutually exclusive with interface
* `whitelist_databases` (List of String) - Attach whitelist databases to the map

<a id="nestedatt--flow_whitelist_overlap_rules--pass_rules--sip"></a>
### Nested Schema for `flow_whitelist_overlap_rules.pass_rules.sip`

Optional:

* `type` (String) - all:Whitelist based on caller/callee/source/destination IP address, bothAddr: Whitelist source/destination IP address, bothId:Whitelist Caller/Callee Id's, calleeId: Whitelist Callee ID, callerId: Whitelist Caller ID, destIp: Whitelist based on destination IP address, srcIp: Whitelist based on Source IP address

<a id="nestedatt--flow_whitelist_rules"></a>
### Nested Schema for `flow_whitelist_rules`

Optional:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--flow_whitelist_rules--pass_rules))

<a id="nestedatt--flow_whitelist_rules--pass_rules"></a>
### Nested Schema for `flow_whitelist_rules.pass_rules`

Required:

* `rule_id` (Number)

Optional:

* `flow5_g` (Attributes) - Map Flow Whitelist 5g Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--flow_whitelist_rules--pass_rules--flow5_g))
* `gtp` (Attributes) - Map Flow Whitelist Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--flow_whitelist_rules--pass_rules--gtp))
* `sip` (Attributes) - Map Flow Whitelist Rule Sip match definition (see [below for nested schema](#nestedatt--flow_whitelist_rules--pass_rules--sip))

<a id="nestedatt--flow_whitelist_rules--pass_rules--flow5_g"></a>
### Nested Schema for `flow_whitelist_rules.pass_rules.flow5_g`

Optional:

* `dnn` (String) - Domain Network Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `type` (String) - Set 5G WL-DB lookup type
* `whitelist_databases` (List of String) - Attach whitelist databases to the map

<a id="nestedatt--flow_whitelist_rules--pass_rules--gtp"></a>
### Nested Schema for `flow_whitelist_rules.pass_rules.gtp`

Optional:

* `apn` (String) - Access Point Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `interface` (String) - interface type. Mutually exclusive with version. required till H 5.6
* `type` (String) - Set GTP WL-DB lookup type
* `version` (String) - mutually exclusive with interface
* `whitelist_databases` (List of String) - Attach whitelist databases to the map

<a id="nestedatt--flow_whitelist_rules--pass_rules--sip"></a>
### Nested Schema for `flow_whitelist_rules.pass_rules.sip`

Optional:

* `type` (String) - all:Whitelist based on caller/callee/source/destination IP address, bothAddr: Whitelist source/destination IP address, bothId:Whitelist Caller/Callee Id's, calleeId: Whitelist Callee ID, callerId: Whitelist Caller ID, destIp: Whitelist based on destination IP address, srcIp: Whitelist based on Source IP address

<a id="nestedatt--gs_rules"></a>
### Nested Schema for `gs_rules`

Optional:

* `drop_rules` (Attributes Set) (see [below for nested schema](#nestedatt--gs_rules--drop_rules))
* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--gs_rules--pass_rules))

<a id="nestedatt--gs_rules--drop_rules"></a>
### Nested Schema for `gs_rules.drop_rules`

Required:

* `matches` (Set of Dynamic) - Set of rule's matching elements. Within a rule, matching elements of the the same type MAY be used more than once, However, their matching positions MUST be unique
* `rule_id` (Number)

Optional:

* `comment` (String)

<a id="nestedatt--gs_rules--pass_rules"></a>
### Nested Schema for `gs_rules.pass_rules`

Required:

* `matches` (Set of Dynamic) - Set of rule's matching elements. Within a rule, matching elements of the the same type MAY be used more than once, However, their matching positions MUST be unique
* `rule_id` (Number)

Optional:

* `comment` (String)

