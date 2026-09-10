---
page_title: "gigavuecore_get_all_ntp_server Data Source - gigavuecore"
subcategory: ""
description: |-
  Get all NTP Servers in FM
---

# gigavuecore_get_all_ntp_server Data Source

Get all NTP Servers in FM

## Example Usage

```terraform
data "gigavuecore_get_all_ntp_server" "example" {
  auth_status   = "example"
  fm_ip         = "example"
  page          = "example"
  server_host   = "example"
  server_status = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `auth_status` (String, optional) - Auth status of NTP
* `fm_ip` (String, optional) - FMHighAvailability node IpAddress/domainName
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `server_host` (String, optional) - NTP server/host address
* `server_status` (String, optional) - NTP server status

### Attributes

In addition to all arguments above, the following attributes are exported:

* `enabled` (Boolean, computed)
* `ntp_servers` (Attributes List, computed) (see [below for nested schema](#nestedatt--ntp_servers))

<a id="nestedatt--ntp_servers"></a>
### Nested Schema for `ntp_servers`

Read-Only:

* `auth_required` (Boolean) - authentication enabled status
* `fm_ip` (String) - FM HighAvailability Node address/host
* `is_user_ntp_server` (Boolean) - To differentiate user created ntp server and default ntp server
* `ntp_auth` (Attributes) - NTP Auth Details (see [below for nested schema](#nestedatt--ntp_servers--ntp_auth))
* `server_host` (String) - Ip/Host Address
* `server_status` (Attributes) - Server Status (see [below for nested schema](#nestedatt--ntp_servers--server_status))
* `version` (Number) - version of server

<a id="nestedatt--ntp_servers--ntp_auth"></a>
### Nested Schema for `ntp_servers.ntp_auth`

Read-Only:

* `key` (String) - key for the ntp server
* `type` (String) - Type of algorithm used for ntp server authentication
* `value` (String) - value for the ntp server authentication

<a id="nestedatt--ntp_servers--server_status"></a>
### Nested Schema for `ntp_servers.server_status`

Read-Only:

* `offset` (Number)
* `poll_interval` (String)
* `status` (String) - status of the server
* `stratum` (String) - status of the server

