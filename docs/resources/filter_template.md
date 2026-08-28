---
page_title: "gigavuecore_filter_template Resource - gigavuecore"
subcategory: ""
description: |-
  Find Filter Template by alias
---

# gigavuecore_filter_template Resource

Find Filter Template by alias

## Example Usage

```terraform
resource "gigavuecore_filter_template" "example" {
  alias         = null
  cluster_id    = null
  qualifier_set = []
  qualifiers    = []
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Filter Template alias
* `cluster_id` (String, required) - Target Cluster ID
* `qualifier_set` (List of String, optional)
* `qualifiers` (List of String, optional) - 'ip6fl', 'ipfrag', 'tcpctl', 'tos', 'ttl' are available only for default filter templates. 'circuit-id' is readonly and will be available on all filter templates by default

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `qualifier_set` (List of String, computed)
* `qualifiers` (List of String, computed) - 'ip6fl', 'ipfrag', 'tcpctl', 'tos', 'ttl' are available only for default filter templates. 'circuit-id' is readonly and will be available on all filter templates by default


## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_filter_template.example {alias}
```
