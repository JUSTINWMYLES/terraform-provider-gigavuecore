---
page_title: "gigavuecore_filter_template Resource - gigavuecore"
subcategory: ""
description: |-
  Create a new Filter Template
---

# gigavuecore_filter_template Resource

Create a new Filter Template

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

* `create` (Number) - A create timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `read` (Number) - A read timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `update` (Number) - An update timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `delete` (Number) - A delete timeout in seconds for this operation. Overrides the generator default (1200 seconds).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_filter_template.example {alias}/{cluster_id}
```
