---
page_title: "gigavuecore_endpoint Resource - gigavuecore"
subcategory: ""
description: |-
  Find a configured SSL Decryption Endpoint by alias
---

# gigavuecore_endpoint Resource

Find a configured SSL Decryption Endpoint by alias

## Example Usage

```terraform
resource "gigavuecore_endpoint" "example" {
  address    = "example"
  alias      = "example"
  cluster_id = "example"
  port       = 0
}
```

## Schema

### Arguments

The following arguments are supported:

* `address` (String, required) - IPv4 address of the SSL endpoint (server)
* `alias` (String, required) - Alias of the Endpoint
* `cluster_id` (String, required) - Target Cluster ID
* `port` (Number, required) - Optional port of the SSL endpoint (server). Value of 0 indicates any port


## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_endpoint.example {alias}
```
