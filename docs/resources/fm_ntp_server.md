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
  auth_required = null
  fm_ip = null
  is_user_ntp_server = null
  ntp_auth = {}
  server_host = null
  server_status = {}
  version = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `auth_required` (Bool, required) - authentication enabled status
* `fm_ip` (String, optional) - FM HighAvailability Node address/host
* `is_user_ntp_server` (Bool, optional) - To differentiate user created ntp server and default ntp server
* `ntp_auth` (Object({key, type, value}), optional) - NTP Auth Details
  * `key` (String, required) - key for the ntp server
  * `type` (String, required) - Type of algorithm used for ntp server authentication
  * `value` (String, required) - value for the ntp server authentication
* `server_host` (String, required) - Ip/Host Address
* `server_status` (Object({offset, poll_interval, status, stratum}), optional) - Server Status
  * `offset` (Number, optional)
  * `poll_interval` (String, optional)
  * `status` (String, required) - status of the server
  * `stratum` (String, optional) - status of the server
* `version` (Number, optional) - version of server

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `fm_ip` (String, computed) - FM HighAvailability Node address/host
* `id` (String, computed)
* `is_user_ntp_server` (Bool, computed) - To differentiate user created ntp server and default ntp server
* `ntp_auth` (Object({key, type, value}), computed) - NTP Auth Details
  * `key` (String, required) - key for the ntp server
  * `type` (String, required) - Type of algorithm used for ntp server authentication
  * `value` (String, required) - value for the ntp server authentication
* `server_status` (Object({offset, poll_interval, status, stratum}), computed) - Server Status
  * `offset` (Number, optional)
  * `poll_interval` (String, optional)
  * `status` (String, required) - status of the server
  * `stratum` (String, optional) - status of the server
* `version` (Number, computed) - version of server

