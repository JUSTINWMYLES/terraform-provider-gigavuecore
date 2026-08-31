---
page_title: "gigavuecore_fm_ntp_server Resource - gigavuecore"
subcategory: ""
description: |-
  Get NTP Server for server/host
---

# gigavuecore_fm_ntp_server Resource

Get NTP Server for server/host

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

* `create` (String) - A create timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `read` (String) - A read timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).
* `update` (String) - A update timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `delete` (String) - A delete timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_fm_ntp_server.example {server_host}
```
