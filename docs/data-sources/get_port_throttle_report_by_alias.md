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
  alias = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - port throttle alias

### Attributes

In addition to all arguments above, the following attributes are exported:

* `port_throttles_report` (Attributes List, computed) (see [below for nested schema](#nestedatt--port_throttles_report))

<a id="nestedatt--port_throttles_report"></a>
### Nested Schema for `port_throttles_report`

Read-Only:

* `configured_pps` (Number) - in kpps
* `current_active_sessions` (Number) - in kpps
* `current_pps` (Number) - in kpps
* `last_session_accepted` (Boolean) - last session accepted or rejected
* `port_id` (String)
* `session_count_accepted` (Number) - in kpps
* `session_count_rejected` (Number) - in kpps

