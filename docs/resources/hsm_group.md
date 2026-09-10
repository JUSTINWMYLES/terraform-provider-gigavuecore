---
page_title: "gigavuecore_hsm_group Resource - gigavuecore"
subcategory: ""
description: |-
  Create a new HSM Group
---

# gigavuecore_hsm_group Resource

Create a new HSM Group

## Example Usage

```terraform
resource "gigavuecore_hsm_group" "example" {
  alias              = "example"
  comment            = "example"
  hsms               = ["example"]
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

* `create` (Number) - A create timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `read` (Number) - A read timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `update` (Number) - An update timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `delete` (Number) - A delete timeout in seconds for this operation. Overrides the generator default (1200 seconds).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_hsm_group.example {alias}
```
