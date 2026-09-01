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
  cluster_id = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - id of the defining cluster

### Attributes

In addition to all arguments above, the following attributes are exported:

* `allow_blank_password` (Boolean, computed)
* `fips_enabled` (Boolean, computed) - enable/disable fips mode(pending fips mode), system reload is necessary to activate fips mode
* `fips_mode_state` (Boolean, computed) - current fips mode
* `min_password_len` (Number, computed)
* `secure_crypto` (Boolean, computed) - Secure crypto mode. Value takes effect after reload
* `secure_crypto_enforced` (Boolean, computed) - Secure crypto mode enforced
* `secure_passwords` (Boolean, computed) - Secure passwords mode


