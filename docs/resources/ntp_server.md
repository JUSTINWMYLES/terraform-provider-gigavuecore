---
page_title: "gigavuecore_ntp_server Resource - gigavuecore"
subcategory: ""
description: |-
  Add an NTP Server
---

# gigavuecore_ntp_server Resource

Add an NTP Server

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

* `create` (Number) - A create timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `read` (Number) - A read timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `update` (Number) - An update timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `delete` (Number) - A delete timeout in seconds for this operation. Overrides the generator default (1200 seconds).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_ntp_server.example {server}:{cluster_id}
```
