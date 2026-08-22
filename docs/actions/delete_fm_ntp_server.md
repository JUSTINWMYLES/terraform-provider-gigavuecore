---
page_title: "gigavuecore_delete_fm_ntp_server Action - gigavuecore"
subcategory: ""
description: |-
  Delete NTP Server by servername
---

# gigavuecore_delete_fm_ntp_server Action

Delete NTP Server by servername

## Example Usage

```terraform
action "gigavuecore_delete_fm_ntp_server" "example" {
  config {
    auth_required = true
    body_fm_ip = "example"
    fm_ip = "example"
    is_user_ntp_server = true
    ntp_auth = null
    server_host = "example"
    server_status = null
    version = 1
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `auth_required` (Bool, required) - authentication enabled status
* `body_fm_ip` (String, optional) - FM HighAvailability Node address/host
* `fm_ip` (String, optional) - FMHighAvailability node IpAddress/domainName
* `is_user_ntp_server` (Bool, optional) - To differentiate user created ntp server and default ntp server
* `ntp_auth` (Dynamic, optional) - NTP Auth Details
* `server_host` (String, required) - Ip/Host Address
* `server_status` (Dynamic, optional) - Server Status
* `version` (Number, optional) - version of server
