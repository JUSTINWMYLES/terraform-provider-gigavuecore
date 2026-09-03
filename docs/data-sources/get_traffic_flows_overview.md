---
page_title: "gigavuecore_get_traffic_flows_overview Data Source - gigavuecore"
subcategory: ""
description: |-
  Get overview of all Traffic Flows
---

# gigavuecore_get_traffic_flows_overview Data Source

Get overview of all Traffic Flows

## Example Usage

```terraform
data "gigavuecore_get_traffic_flows_overview" "example" {
}
```

## Schema

### Attributes

In addition to all arguments above, the following attributes are exported:

* `failure` (Number, computed)
* `green` (Number, computed)
* `others` (Number, computed)
* `red` (Number, computed)
* `success` (Number, computed)
* `total` (Number, computed)
* `undeployed` (Number, computed)
* `yellow` (Number, computed)


