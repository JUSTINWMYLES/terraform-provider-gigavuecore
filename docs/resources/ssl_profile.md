---
page_title: "gigavuecore_ssl_profile Resource - gigavuecore"
subcategory: ""
description: |-
  Create Apps Ssl Profile
---

# gigavuecore_ssl_profile Resource

Create Apps Ssl Profile

## Example Usage

```terraform
resource "gigavuecore_ssl_profile" "example" {
  alias      = "example"
  cipher     = "example"
  cluster_id = "example"
  mtls       = "enable"
  version    = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required)
* `cipher` (String, optional)
* `cluster_id` (String, required) - Target Cluster ID
* `mtls` (String, optional)
* `version` (String, optional)

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
terraform import gigavuecore_ssl_profile.example {alias}:{cluster_id}
```
