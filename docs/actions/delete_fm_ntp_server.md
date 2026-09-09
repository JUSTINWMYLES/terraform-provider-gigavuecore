---
page_title: "gigavuecore_delete_fm_ntp_server Action - gigavuecore"
subcategory: ""
description: |-
  Delete NTP Server by servername
---

# gigavuecore_delete_fm_ntp_server Action

Delete NTP Server by servername

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_delete_fm_ntp_server" "example" {
  config {
    auth_required      = true
    body_fm_ip         = "example"
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
}
```
## Schema

### Arguments

The following arguments are supported:

* `auth_required` (Boolean, required) - authentication enabled status
* `body_fm_ip` (String, optional) - FM HighAvailability Node address/host
* `fm_ip` (String, optional) - FMHighAvailability node IpAddress/domainName
* `is_user_ntp_server` (Boolean, optional) - To differentiate user created ntp server and default ntp server
* `ntp_auth` (Attributes, optional) - NTP Auth Details (see [below for nested schema](#nestedatt--ntp_auth))
* `server_host` (String, required) - Ip/Host Address
* `server_status` (Attributes, optional) - Server Status (see [below for nested schema](#nestedatt--server_status))
* `version` (Number, optional) - version of server

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

