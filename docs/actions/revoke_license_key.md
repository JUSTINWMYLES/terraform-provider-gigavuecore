---
page_title: "gigavuecore_revoke_license_key Action - gigavuecore"
subcategory: ""
description: |-
  Revoke Node-Locked Device License
---

# gigavuecore_revoke_license_key Action

Revoke Node-Locked Device License

## Example Usage

```terraform
action "gigavuecore_revoke_license_key" "example" {
  config {
    activation_id = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `activation_id` (String, required) - ID of the feature activation
