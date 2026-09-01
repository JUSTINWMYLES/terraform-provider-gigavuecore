---
page_title: "gigavuecore_hsm Resource - gigavuecore"
subcategory: ""
description: |-
  Create a new HSM
---

# gigavuecore_hsm Resource

Create a new HSM

## Example Usage

```terraform
resource "gigavuecore_hsm" "example" {
  alias              = "example"
  esn                = "example"
  hsm_ip             = "example"
  hsm_port           = 0
  kneti              = "example"
  operational_status = "unknown"
  partition_label    = "example"
  partition_password = "example"
  server_password    = "example"
  server_username    = "example"
  type               = "ncipher"
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Hsm alias
* `esn` (String, optional) - the ESN of the KNETI key for a given IP address
* `hsm_ip` (String, required) - hsm server address
* `hsm_port` (Number, required) - server port number, the port to use when connecting to the given nethsm
* `kneti` (String, optional) - the hash of the KNETI key for a given IP address
* `operational_status` (String, optional) - operational status of hsm
* `partition_label` (String, optional) - hsm partition label
* `partition_password` (String, optional) - hsm partition password
* `server_password` (String, optional) - hsm server password
* `server_username` (String, optional) - hsm server username
* `type` (String, required)

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
terraform import gigavuecore_hsm.example {alias}
```
