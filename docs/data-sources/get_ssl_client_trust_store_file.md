---
page_title: "gigavuecore_get_ssl_client_trust_store_file Data Source - gigavuecore"
subcategory: ""
description: |-
  Get the SSL Client trust-store file
---

# gigavuecore_get_ssl_client_trust_store_file Data Source

Get the SSL Client trust-store file

## Example Usage

```terraform
data "gigavuecore_get_ssl_client_trust_store_file" "example" {
  alias      = "example"
  cluster_id = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Client Trust Store alias
* `cluster_id` (String, required) - Target Cluster ID

### Attributes

In addition to all arguments above, the following attributes are exported:

* `certificate` (String, computed)
* `type` (String, computed)


