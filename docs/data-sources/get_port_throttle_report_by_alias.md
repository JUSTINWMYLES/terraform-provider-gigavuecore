---
page_title: "gigavuecore_get_port_throttle_report_by_alias Data Source - gigavuecore"
subcategory: ""
description: |-
  Load Port Throttle Report
---

# gigavuecore_get_port_throttle_report_by_alias Data Source

Load Port Throttle Report

## Example Usage

```terraform
data "gigavuecore_get_port_throttle_report_by_alias" "example" {
  alias = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - port throttle alias

### Attributes

In addition to all arguments above, the following attributes are exported:

* `port_throttles_report` (List(Object({configured_pps, current_active_sessions, current_pps, last_session_accepted, port_id, session_count_accepted, session_count_rejected})), computed)

