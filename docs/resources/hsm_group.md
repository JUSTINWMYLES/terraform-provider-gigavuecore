---
page_title: "gigavuecore_hsm_group Resource - gigavuecore"
subcategory: ""
description: |-
  Load HSM Group by alias
---

# gigavuecore_hsm_group Resource

Load HSM Group by alias

## Example Usage

```terraform
resource "gigavuecore_hsm_group" "example" {
  alias              = "example"
  comment            = "example"
  hsms               = [ "example" ]
  operational_status = "unknown"
  type               = "ncipher"
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - hsm group alias
* `comment` (String, optional) - Hsm Group Comment
* `hsms` (List of String, optional) - alias of hsm in hsm group
* `operational_status` (String, optional) - operational status of hsm group
* `type` (String, required)

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

* `create` (String) - A create timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `read` (String) - A read timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).
* `update` (String) - A update timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `delete` (String) - A delete timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_hsm_group.example {alias}
```
