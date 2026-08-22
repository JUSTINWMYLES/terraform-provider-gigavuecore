---
page_title: "gigavuecore_get_avisi_policies List Resource - gigavuecore"
subcategory: ""
description: |-
  Get Active Visibility Policies
---

# gigavuecore_get_avisi_policies List Resource

Get Active Visibility Policies

## Example Usage

```terraform
list "gigavuecore_get_avisi_policies" "example" {
  provider = gigavuecore
  limit = 100
}

```

## Schema

### Arguments

The following arguments are supported:


### Identity Attributes

The following identity attributes are exported for each matching result:

* `policy_id` (String, computed)
