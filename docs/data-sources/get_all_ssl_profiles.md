---
page_title: "gigavuecore_get_all_ssl_profiles Data Source - gigavuecore"
subcategory: ""
description: |-
  Get all Apps Ssl Profile
---

# gigavuecore_get_all_ssl_profiles Data Source

Get all Apps Ssl Profile

## Example Usage

```terraform
data "gigavuecore_get_all_ssl_profiles" "example" {
  cluster_id = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `alias` (String)
* `cipher` (String)
* `mtls` (String)
* `version` (String)

