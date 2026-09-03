---
page_title: "gigavuecore_solution List Resource - gigavuecore"
subcategory: ""
description: |-
  Load all Apps Visibility Solutions
---

# gigavuecore_solution List Resource

Load all Apps Visibility Solutions

-> **Note:** This list resource requires Terraform 1.14 or later and is used through the `terraform query` command, not in configuration files.

## Example Usage

```terraform
list "gigavuecore_solution" "example" {
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

* `cluster_id` (String, required) - clusterId


### Identity Attributes

The following identity attributes are exported for each matching result:

* `solution_alias` (String, computed) - user defined application visibility solution alias


