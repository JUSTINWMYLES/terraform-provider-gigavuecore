---
page_title: "gigavuecore_get_all_flow_sip_summary Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all Flow Sip Report Summary
---

# gigavuecore_get_all_flow_sip_summary Data Source

Load all Flow Sip Report Summary

## Example Usage

```terraform
data "gigavuecore_get_all_flow_sip_summary" "example" {
}
```

## Schema

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({gsgroup, rtp_resource_summary, rtp_sessions, sip_messages_stats, sip_resource_summary, sip_sessions})), computed)

