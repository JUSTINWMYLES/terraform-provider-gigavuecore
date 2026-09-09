---
page_title: "gigavuecore_change_tunnel_auth_mode Action - gigavuecore"
subcategory: ""
description: |-
  Change FMHA tunnel auth mode
---

# gigavuecore_change_tunnel_auth_mode Action

Change FMHA tunnel auth mode

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_change_tunnel_auth_mode" "example" {
  config {
    tunnel_auth_mode = "PSK"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `tunnel_auth_mode` (String, optional) - Auth mode for setting up FMHA tunnel. PSK is Pre Shared Key. PKI is Public Key Infrastructure(Certificate based)


