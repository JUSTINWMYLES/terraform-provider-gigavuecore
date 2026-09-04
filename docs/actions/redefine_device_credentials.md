---
page_title: "gigavuecore_redefine_device_credentials Action - gigavuecore"
subcategory: ""
description: |-
  Redefine Device Credentials configuration
---

# gigavuecore_redefine_device_credentials Action

Redefine Device Credentials configuration

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

~> **Warning:** This action accepts attributes whose names indicate secrets (httpPassword), but action schemas cannot mark attributes Sensitive, so their values are stored and displayed in plain text. Avoid passing real secrets through these attributes.

## Example Usage

```terraform
action "gigavuecore_redefine_device_credentials" "example" {
  config {
    body_device_address = "example"
    device_address      = "example"
    hostname            = "example"
    http_password       = "example"
    http_username       = "example"
    https_port          = "example"
    snmp_version        = "v2"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `body_device_address` (String, required) - value of '0.0.0.0' represents the default (fallback) device credentials
* `device_address` (String, required) - device address for the target Credentials
* `hostname` (String, optional) - device configured hostname
* `http_password` (String, required) - password to use for device login. On reads, '\*\*\*\*\*' is returned
* `http_username` (String, required) - username to use for device login
* `https_port` (String, optional) - httpsPort to use for device communication. By default 443 is used, If changed the same should be given here
* `snmp_version` (String, optional) - SNMP version to use.


