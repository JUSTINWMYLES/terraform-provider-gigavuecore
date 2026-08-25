---
page_title: "gigavuecore_rollback_traffic_flows Action - gigavuecore"
subcategory: ""
description: |-
  Roll back Traffic Flow migration
---

# gigavuecore_rollback_traffic_flows Action

Roll back Traffic Flow migration

## Example Usage

```terraform
action "gigavuecore_rollback_traffic_flows" "example" {
  config {
    traffic_flows = [ "example" ]
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `traffic_flows` (List of String, optional)


