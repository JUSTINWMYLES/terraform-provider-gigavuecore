---
page_title: "gigavuecore_redefine_gs_group_session_logging Action - gigavuecore"
subcategory: ""
description: |-
  Redefine GS Group's Session Logging Info
---

# gigavuecore_redefine_gs_group_session_logging Action

Redefine GS Group's Session Logging Info

## Example Usage

```terraform
action "gigavuecore_redefine_gs_group_session_logging" "example" {
  config {
    alias              = "example"
    interface          = "example"
    log_level          = "example"
    remote_syslog_ip   = "example"
    remote_syslog_port = 0
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target GS Group
* `interface` (String, optional) - Associated IP Interface
* `log_level` (String, optional) - Log Level
* `remote_syslog_ip` (String, optional) - Remote Syslog IP
* `remote_syslog_port` (Number, optional) - Remote Syslog Port Number


