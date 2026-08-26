---
page_title: "gigavuecore_get_policy_deployment_log Data Source - gigavuecore"
subcategory: ""
description: |-
  Get policy Deployment log
---

# gigavuecore_get_policy_deployment_log Data Source

Get policy Deployment log

## Example Usage

```terraform
data "gigavuecore_get_policy_deployment_log" "example" {
  name = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `name` (String, required) - policy name

### Attributes

In addition to all arguments above, the following attributes are exported:

* `message` (String, computed)
* `policy_id` (String, computed)


