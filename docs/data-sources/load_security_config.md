---
page_title: "gigavuecore_load_security_config Data Source - gigavuecore"
subcategory: ""
description: |-
  Load Security config
---

# gigavuecore_load_security_config Data Source

Load Security config

## Example Usage

```terraform
data "gigavuecore_load_security_config" "example" {
  cluster_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - id of the defining cluster

### Attributes

In addition to all arguments above, the following attributes are exported:

* `allow_blank_password` (Bool, computed)
* `fips_enabled` (Bool, computed) - enable/disable fips mode(pending fips mode), system reload is necessary to activate fips mode
* `fips_mode_state` (Bool, computed) - current fips mode
* `min_password_len` (Number, computed)
* `secure_crypto` (Bool, computed) - Secure crypto mode. Value takes effect after reload
* `secure_crypto_enforced` (Bool, computed) - Secure crypto mode enforced
* `secure_passwords` (Bool, computed) - Secure passwords mode

