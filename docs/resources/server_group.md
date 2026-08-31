---
page_title: "gigavuecore_server_group Resource - gigavuecore"
subcategory: ""
description: |-
  Load ICAP Server Group by Alias
---

# gigavuecore_server_group Resource

Load ICAP Server Group by Alias

## Example Usage

```terraform
resource "gigavuecore_server_group" "example" {
  alias        = "example"
  cluster_id   = "example"
  comment      = "example"
  icap_servers = ["example"]
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Icap Server Group Alias
* `cluster_id` (String, required) - id of the defining cluster
* `comment` (String, optional) - Icap Server Group Comment
* `icap_servers` (List of String, optional) - ICAP Servers list separated by comma','

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
terraform import gigavuecore_server_group.example {alias}/{cluster_id}
```
