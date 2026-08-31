---
page_title: "gigavuecore_service_state_change Action - gigavuecore"
subcategory: ""
description: |-
  Service State Change Request
---

# gigavuecore_service_state_change Action

Service State Change Request

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_service_state_change" "example" {
  config {
    hostname = "example"
    service  = "application"
    state    = "up"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `hostname` (String, required) - DNS name or IP Address of service state
* `service` (String, required) - special service name which contains all the necessary service for REST API access
* `state` (String, required) - up is for turn on, down is for turn off and in progress is for in-progress


