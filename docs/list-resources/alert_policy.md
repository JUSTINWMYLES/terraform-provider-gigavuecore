---
page_title: "gigavuecore_alert_policy List Resource - gigavuecore"
subcategory: ""
description: |-
  Alert policy listing
---

# gigavuecore_alert_policy List Resource

Alert policy listing

-> **Note:** This list resource requires Terraform 1.14 or later and is used through the `terraform query` command, not in configuration files.

## Example Usage

```terraform
list "gigavuecore_alert_policy" "example" {
  provider = gigavuecore
  limit    = 100
  config {
    enabled       = true
    page          = "example"
    policy_name   = "example"
    resource_type = "example"
    sort          = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `enabled` (Boolean, optional)
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `policy_name` (String, optional) - Name of the alert policy
* `resource_type` (String, optional)
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)


### Identity Attributes

The following identity attributes are exported for each matching result:

* `policy_name` (String, computed) - Name of the alert policy, like an alias


