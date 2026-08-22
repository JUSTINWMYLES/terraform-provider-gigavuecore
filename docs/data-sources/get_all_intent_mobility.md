---
page_title: "gigavuecore_get_all_intent_mobility Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all mobility solutions configured
---

# gigavuecore_get_all_intent_mobility Data Source

Load all mobility solutions configured

## Example Usage

```terraform
data "gigavuecore_get_all_intent_mobility" "example" {
}
```

## Schema

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({health_state, health_state_reasons, sites, solution_alias, solution_type, tags, traffic_policies})), computed)

