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

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({alias, port_throttles_report})), computed)

