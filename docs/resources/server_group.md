---
page_title: "gigavuecore_server_group Resource - gigavuecore"
subcategory: ""
description: |-
  Create a new ICAP Server Group
---

# gigavuecore_server_group Resource

Create a new ICAP Server Group

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

* `create` (Number) - A create timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `read` (Number) - A read timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `update` (Number) - An update timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `delete` (Number) - A delete timeout in seconds for this operation. Overrides the generator default (1200 seconds).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_server_group.example {alias}/{cluster_id}
```
