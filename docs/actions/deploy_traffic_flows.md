---
page_title: "gigavuecore_deploy_traffic_flows Action - gigavuecore"
subcategory: ""
description: |-
  Deploy a Traffic Flows configuration
---

# gigavuecore_deploy_traffic_flows Action

Deploy a Traffic Flows configuration

## Example Usage

```terraform
action "gigavuecore_deploy_traffic_flows" "example" {
  config {
    alias = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Traffic Flows alias or ID


