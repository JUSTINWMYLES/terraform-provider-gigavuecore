---
page_title: "gigavuecore_ssl_profile List Resource - gigavuecore"
subcategory: ""
description: |-
  Get all Apps Ssl Profile
---

# gigavuecore_ssl_profile List Resource

Get all Apps Ssl Profile

-> **Note:** This list resource requires Terraform 1.14 or later and is used through the `terraform query` command, not in configuration files.

## Example Usage

```terraform
list "gigavuecore_ssl_profile" "example" {
  provider = gigavuecore
  limit    = 100
  config {
    cluster_id = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID


### Identity Attributes

The following identity attributes are exported for each matching result:

* `alias` (String, computed)


