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
  auth_status = null
  fm_ip = null
  page = null
  server_host = null
  server_status = null
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

* `enabled` (Bool, computed)
* `ntp_servers` (List(Object({auth_required, fm_ip, is_user_ntp_server, ntp_auth, server_host, server_status, version})), computed)

