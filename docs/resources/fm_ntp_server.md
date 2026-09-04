---
page_title: "gigavuecore_fm_ntp_server Resource - gigavuecore"
subcategory: ""
description: |-
  Add NTP Servers in FM
---

# gigavuecore_fm_ntp_server Resource

Add NTP Servers in FM

~> **Note:** The update operation for this resource is not wired to a remote API endpoint: the API spec exposes no usable update mapping. Changing any configuration attribute triggers a resource replacement (all config-settable attributes carry a RequiresReplace plan modifier); create, read, and delete remain functional.

## Example Usage

```terraform
resource "gigavuecore_fm_ntp_server" "example" {
  auth_required      = true
  fm_ip              = "example"
  is_user_ntp_server = true
  ntp_auth = {
    key   = "example"
    type  = "example"
    value = "example"
  }
  server_host = "example"
  server_status = {
    offset        = 0
    poll_interval = "example"
    status        = "example"
    stratum       = "example"
  }
  version = 0
}
```

## Schema

### Arguments

The following arguments are supported:

* `auth_required` (Boolean, required) - authentication enabled status
* `fm_ip` (String, optional) - FM HighAvailability Node address/host
* `is_user_ntp_server` (Boolean, optional) - To differentiate user created ntp server and default ntp server
* `ntp_auth` (Attributes, optional) - NTP Auth Details (see [below for nested schema](#nestedatt--ntp_auth))
* `server_host` (String, required) - Ip/Host Address
* `server_status` (Attributes, optional) - Server Status (see [below for nested schema](#nestedatt--server_status))
* `version` (Number, optional) - version of server

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

<a id="nestedatt--ntp_auth"></a>
### Nested Schema for `ntp_auth`

Required:

* `key` (String) - key for the ntp server
* `type` (String) - Type of algorithm used for ntp server authentication
* `value` (String) - value for the ntp server authentication

<a id="nestedatt--server_status"></a>
### Nested Schema for `server_status`

Required:

* `status` (String) - status of the server

Optional:

* `offset` (Number)
* `poll_interval` (String)
* `stratum` (String) - status of the server
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
terraform import gigavuecore_fm_ntp_server.example {server_host}
```
