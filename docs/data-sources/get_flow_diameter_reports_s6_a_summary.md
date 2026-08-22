---
page_title: "gigavuecore_get_flow_diameter_reports_s6_a_summary Data Source - gigavuecore"
subcategory: ""
description: |-
  Load Flow Diameter S6a Report Summary
---

# gigavuecore_get_flow_diameter_reports_s6_a_summary Data Source

Load Flow Diameter S6a Report Summary

## Example Usage

```terraform
data "gigavuecore_get_flow_diameter_reports_s6_a_summary" "example" {
}
```

## Schema

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({gsgroup, s6_a_messages_stats, s6_a_resource_summary, s6_a_sessions})), computed)

