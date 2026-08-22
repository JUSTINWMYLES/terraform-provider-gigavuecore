---
page_title: "gigavuecore_undeploy_policy Action - gigavuecore"
subcategory: ""
description: |-
  Undeploy policy
---

# gigavuecore_undeploy_policy Action

Undeploy policy

## Example Usage

```terraform
action "gigavuecore_undeploy_policy" "example" {
  config {
    name = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `name` (String, required) - policy name
