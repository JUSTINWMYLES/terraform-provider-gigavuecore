---
page_title: "gigavuecore_get_port_throttle_report Data Source - gigavuecore"
subcategory: ""
description: |-
  Load All Port Throttle Reports
---

# gigavuecore_get_port_throttle_report Data Source

Load All Port Throttle Reports

## Example Usage

```terraform
data "gigavuecore_get_port_throttle_report" "example" {
}
```

## Schema

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `alias` (String) - port throttle alias
* `port_throttles_report` (Attributes List) (see [below for nested schema](#nestedatt--items--port_throttles_report))

<a id="nestedatt--items--port_throttles_report"></a>
### Nested Schema for `items.port_throttles_report`

Read-Only:

* `configured_pps` (Number) - in kpps
* `current_active_sessions` (Number) - in kpps
* `current_pps` (Number) - in kpps
* `last_session_accepted` (Boolean) - last session accepted or rejected
* `port_id` (String)
* `session_count_accepted` (Number) - in kpps
* `session_count_rejected` (Number) - in kpps

