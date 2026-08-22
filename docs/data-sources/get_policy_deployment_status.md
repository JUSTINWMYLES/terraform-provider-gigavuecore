---
page_title: "gigavuecore_get_policy_deployment_status Data Source - gigavuecore"
subcategory: ""
description: |-
  Get policy Deployment in Progress
---

# gigavuecore_get_policy_deployment_status Data Source

Get policy Deployment in Progress

## Example Usage

```terraform
data "gigavuecore_get_policy_deployment_status" "example" {
}
```

## Schema

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `policies` (List(Object({comment, deployed, deployment_error, deployment_percent, dest_port_timestamp, health_state, health_state_reasons, name, policy_id, policy_timestamp, priority, rules, src_port_timestamp, src_ports_info, tags})), computed)

