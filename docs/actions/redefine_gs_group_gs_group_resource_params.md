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
    alias = "example"
    buffer_asf_size = 1
    cpu = null
    hsm_ssl = null
    inline_ssl = null
    metadata = 1
    packet_buffer = null
    session_overload = null
    tunnel_overload = null
    xpkt_match = null
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target GS Group
* `buffer_asf_size` (Number, optional) - Session aware APF buffer. Valid values: 2-5: in million sessions (0 = disabled). Changes take effect after card or system reboot
* `cpu` (Dynamic, optional)
* `hsm_ssl` (Dynamic, optional) - GsGroup Resource Hsm Ssl Parameters
* `inline_ssl` (Dynamic, optional) - Used to configure other GS apps in addition to Inline SSL on a HC1 box
* `metadata` (Number, optional) - flows in millions, how many flows to support for metadata. 0 to disable
* `packet_buffer` (Dynamic, optional)
* `session_overload` (Dynamic, optional)
* `tunnel_overload` (Dynamic, optional)
* `xpkt_match` (Dynamic, optional) - GsGroup Resource Cross Packet Match Parameters
