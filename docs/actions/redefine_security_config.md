---
page_title: "gigavuecore_redefine_security_config Action - gigavuecore"
subcategory: ""
description: |-
  Redefine Security config
---

# gigavuecore_redefine_security_config Action

Redefine Security config

## Example Usage

```terraform
action "gigavuecore_redefine_security_config" "example" {
  config {
    allow_blank_password   = true
    cluster_id             = "example"
    fips_enabled           = true
    fips_mode_state        = true
    min_password_len       = 0
    secure_crypto          = true
    secure_crypto_enforced = true
    secure_passwords       = true
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `allow_blank_password` (Boolean, optional)
* `cluster_id` (String, optional) - id of the defining cluster
* `fips_enabled` (Boolean, optional) - enable/disable fips mode(pending fips mode), system reload is necessary to activate fips mode
* `fips_mode_state` (Boolean, optional) - current fips mode
* `min_password_len` (Number, optional)
* `secure_crypto` (Boolean, optional) - Secure crypto mode. Value takes effect after reload
* `secure_crypto_enforced` (Boolean, optional) - Secure crypto mode enforced
* `secure_passwords` (Boolean, optional) - Secure passwords mode


