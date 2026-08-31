---
page_title: "gigavuecore_ssl_profile Resource - gigavuecore"
subcategory: ""
description: |-
  Get Apps Ssl Profile for given alias
---

# gigavuecore_ssl_profile Resource

Get Apps Ssl Profile for given alias

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

* `create` (String) - A create timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `read` (String) - A read timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).
* `update` (String) - An update timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `delete` (String) - A delete timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_ssl_profile.example {alias}:{cluster_id}
```
