---
page_title: "gigavuecore_ssl_profile Resource - gigavuecore"
subcategory: ""
description: |-
  Get Apps Ssl Profile for given alias
---

# gigavuecore_ssl_profile Resource

Get Apps Ssl Profile for given alias

## Example Usage

```terraform
resource "gigavuecore_ssl_profile" "example" {
  alias      = "example"
  cipher     = "example"
  cluster_id = "example"
  mtls       = "example"
  version    = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required)
* `cipher` (String, optional)
* `cluster_id` (String, required) - Target Cluster ID
* `mtls` (String, optional)
* `version` (String, optional)

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `cipher` (String, computed)
* `mtls` (String, computed)
* `version` (String, computed)


## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_ssl_profile.example {alias}
```
