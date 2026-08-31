---
page_title: "gigavuecore_redefine_gs_group_gs_group_resource_params Action - gigavuecore"
subcategory: ""
description: |-
  Redefine GS Group's resource Params
---

# gigavuecore_redefine_gs_group_gs_group_resource_params Action

Redefine GS Group's resource Params

## Example Usage

```terraform
action "gigavuecore_redefine_gs_group_gs_group_resource_params" "example" {
  config {
    alias           = "example"
    buffer_asf_size = 0
    cpu = {
      overload_threshold = 0
    }
    hsm_ssl = {
      buffer        = 0
      packet_buffer = 20
      session_count = 0
    }
    inline_ssl = {
      standalone = true
    }
    metadata = 0
    packet_buffer = {
      overload_threshold = 0
    }
    session_overload = {
      overload_threshold = 0
    }
    tunnel_overload = {
      overload_threshold = 0
    }
    xpkt_match = {
      flows = 0
    }
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target GS Group
* `buffer_asf_size` (Number, optional) - Session aware APF buffer. Valid values: 2-5: in million sessions (0 = disabled). Changes take effect after card or system reboot
* `cpu` (Attributes, optional) (see [below for nested schema](#nestedatt--cpu))
* `hsm_ssl` (Attributes, optional) - GsGroup Resource Hsm Ssl Parameters (see [below for nested schema](#nestedatt--hsm_ssl))
* `inline_ssl` (Attributes, optional) - Used to configure other GS apps in addition to Inline SSL on a HC1 box (see [below for nested schema](#nestedatt--inline_ssl))
* `metadata` (Number, optional) - flows in millions, how many flows to support for metadata. 0 to disable
* `packet_buffer` (Attributes, optional) (see [below for nested schema](#nestedatt--packet_buffer))
* `session_overload` (Attributes, optional) (see [below for nested schema](#nestedatt--session_overload))
* `tunnel_overload` (Attributes, optional) (see [below for nested schema](#nestedatt--tunnel_overload))
* `xpkt_match` (Attributes, optional) - GsGroup Resource Cross Packet Match Parameters (see [below for nested schema](#nestedatt--xpkt_match))

<a id="nestedatt--cpu"></a>
### Nested Schema for `cpu`

Optional:

* `overload_threshold` (Number) - Valid values 50 - 90%. 0 is disabled. Changes take effect after card or system reboot

<a id="nestedatt--hsm_ssl"></a>
### Nested Schema for `hsm_ssl`

Optional:

* `buffer` (Number) - resource for application hsm-ssl buffer in MB. 0 to disable
* `packet_buffer` (Number) - resource for application hsm-ssl packet-buffer per connection
* `session_count` (Number) - resource for application hsm-ssl buffer session count in million, 0 to disable

<a id="nestedatt--inline_ssl"></a>
### Nested Schema for `inline_ssl`

Optional:

* `standalone` (Boolean) - If enabled , behaves in a normal way. If disabled, can configure other GS apps along side Inline ssl with Inline ssl occupying only 50% of the memory

<a id="nestedatt--packet_buffer"></a>
### Nested Schema for `packet_buffer`

Optional:

* `overload_threshold` (Number) - Valid values 50 - 80%. 0 is disabled. Changes take effect after card or system reboot

<a id="nestedatt--session_overload"></a>
### Nested Schema for `session_overload`

Optional:

* `overload_threshold` (Number) - Session overload threshold value , Default value is 90 and 0 is disabled.

<a id="nestedatt--tunnel_overload"></a>
### Nested Schema for `tunnel_overload`

Optional:

* `overload_threshold` (Number) - Tunnel overload threshold value , Default value is 90 and 0 is disabled.

<a id="nestedatt--xpkt_match"></a>
### Nested Schema for `xpkt_match`

Optional:

* `flows` (Number) - num in 100K flows. 0 is disable

