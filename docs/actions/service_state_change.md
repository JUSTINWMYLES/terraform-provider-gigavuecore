---
page_title: "gigavuecore_service_state_change Action - gigavuecore"
subcategory: ""
description: |-
  Service State Change Request
---

# gigavuecore_service_state_change Action

Service State Change Request

## Example Usage

```terraform
action "gigavuecore_service_state_change" "example" {
  config {
    hostname = "example"
    service = "example"
    state = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `hostname` (String, required) - DNS name or IP Address of service state
* `service` (String, required) - special service name which contains all the necessary service for REST API access
* `state` (String, required) - up is for turn on, down is for turn off and in progress is for in-progress
