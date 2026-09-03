---
page_title: "gigavuecore_policy List Resource - gigavuecore"
subcategory: ""
description: |-
  Get Active Visibility Policies
---

# gigavuecore_policy List Resource

Get Active Visibility Policies

-> **Note:** This list resource requires Terraform 1.14 or later and is used through the `terraform query` command, not in configuration files.

## Example Usage

```terraform
list "gigavuecore_policy" "example" {
  provider = gigavuecore
  limit    = 100
}
```
## Schema

### Identity Attributes

The following identity attributes are exported for each matching result:

* `policy_id` (String, computed)


