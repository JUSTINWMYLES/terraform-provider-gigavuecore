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
  auth_required      = null
  fm_ip              = null
  is_user_ntp_server = null
  ntp_auth           = {}
  server_host        = null
  server_status      = {}
  version            = null
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

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `fm_ip` (String, computed) - FM HighAvailability Node address/host
* `id` (String, computed)
* `is_user_ntp_server` (Boolean, computed) - To differentiate user created ntp server and default ntp server
* `ntp_auth` (Attributes, computed) - NTP Auth Details (see [below for nested schema](#nestedatt--ntp_auth))
* `server_status` (Attributes, computed) - Server Status (see [below for nested schema](#nestedatt--server_status))
* `version` (Number, computed) - version of server

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

