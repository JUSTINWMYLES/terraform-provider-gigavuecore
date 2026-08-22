---
page_title: "gigavuecore_get_all_flow_filtering_summary Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all Flow Filtering Report Summary
---

# gigavuecore_get_all_flow_filtering_summary Data Source

Load all Flow Filtering Report Summary

## Example Usage

```terraform
data "gigavuecore_get_all_flow_filtering_summary" "example" {
}
```

## Schema

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({control_only_session, control_tunnels, control_user_tunnels, gsgroup, gtp_corelation_statistics, gtp_interface_statistics, gtp_pfcp_statistics, gtp_session_statistics, pending_session})), computed)

