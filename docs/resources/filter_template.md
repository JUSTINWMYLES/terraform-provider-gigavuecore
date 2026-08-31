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
  alias         = "example"
  cluster_id    = "example"
  qualifier_set = ["mplsLabelTtl"]
  qualifiers    = ["ipsrc"]
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Filter Template alias
* `cluster_id` (String, required) - Target Cluster ID
* `qualifier_set` (List of String, optional)
* `qualifiers` (List of String, optional) - 'ip6fl', 'ipfrag', 'tcpctl', 'tos', 'ttl' are available only for default filter templates. 'circuit-id' is readonly and will be available on all filter templates by default

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

* `create` (String) - A create timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `read` (String) - A read timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).
* `update` (String) - An update timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `delete` (String) - A delete timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_filter_template.example {alias}/{cluster_id}
```
