---
page_title: "gigavuecore_hsm Resource - gigavuecore"
subcategory: ""
description: |-
  Load HSM by alias
---

# gigavuecore_hsm Resource

Load HSM by alias

## Example Usage

```terraform
resource "gigavuecore_hsm" "example" {
  alias = null
  esn = null
  hsm_ip = null
  hsm_port = null
  kneti = null
  operational_status = null
  partition_label = null
  partition_password = null
  server_password = null
  server_username = null
  type = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Hsm alias
* `esn` (String, optional) - the ESN of the KNETI key for a given IP address
* `hsm_ip` (String, required) - hsm server address
* `hsm_port` (Number, required) - server port number, the port to use when connecting to the given nethsm
* `kneti` (String, optional) - the hash of the KNETI key for a given IP address
* `operational_status` (String, optional) - operational status of hsm
* `partition_label` (String, optional) - hsm partition label
* `partition_password` (String, optional) - hsm partition password
* `server_password` (String, optional) - hsm server password
* `server_username` (String, optional) - hsm server username
* `type` (String, required)

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `esn` (String, computed) - the ESN of the KNETI key for a given IP address
* `kneti` (String, computed) - the hash of the KNETI key for a given IP address
* `operational_status` (String, computed) - operational status of hsm
* `partition_label` (String, computed) - hsm partition label
* `partition_password` (String, computed) - hsm partition password
* `server_password` (String, computed) - hsm server password
* `server_username` (String, computed) - hsm server username

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_hsm.example {alias}
```
