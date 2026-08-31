---
page_title: "gigavuecore_ntp_server Resource - gigavuecore"
subcategory: ""
description: |-
  Find NtpServer by address
---

# gigavuecore_ntp_server Resource

Find NtpServer by address

## Example Usage

```terraform
resource "gigavuecore_ntp_server" "example" {
  cluster_id  = "example"
  enabled     = true
  key_enabled = true
  key_number  = 1
  preferred   = true
  server      = "example"
  version     = "v3"
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `enabled` (Boolean, optional)
* `key_enabled` (Boolean, optional)
* `key_number` (Number, optional)
* `preferred` (Boolean, optional)
* `server` (String, required) - ipv4 or ipv6 or hostname
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
terraform import gigavuecore_ntp_server.example {server}:{cluster_id}
```
