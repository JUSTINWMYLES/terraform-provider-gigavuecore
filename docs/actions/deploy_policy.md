---
page_title: "gigavuecore_deploy_policy Action - gigavuecore"
subcategory: ""
description: |-
  Deploy a new policy
---

# gigavuecore_deploy_policy Action

Deploy a new policy

## Example Usage

```terraform
action "gigavuecore_deploy_policy" "example" {
  config {
    comment = "example"
    deployed = true
    deployment_error = "example"
    deployment_percent = "example"
    dest_port_timestamp = "example"
    health_state = "example"
    health_state_reasons = null
    name = "example"
    policy_id = "example"
    policy_timestamp = "example"
    priority = true
    rules = null
    src_port_timestamp = "example"
    src_ports_info = null
    tags = null
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `comment` (String, optional)
* `deployed` (Bool, optional) - policy is deployed or not
* `deployment_error` (String, optional) - policy deployment error message
* `deployment_percent` (String, optional) - policy deployment percentage
* `dest_port_timestamp` (String, optional)
* `health_state` (String, optional) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (List(Dynamic), optional)
* `name` (String, required) - policy name
* `policy_id` (String, optional) - generated unique policy ID
* `policy_timestamp` (String, optional)
* `priority` (Bool, optional)
* `rules` (List(Dynamic), optional)
* `src_port_timestamp` (String, optional)
* `src_ports_info` (Dynamic, optional)
* `tags` (List(Dynamic), optional)
