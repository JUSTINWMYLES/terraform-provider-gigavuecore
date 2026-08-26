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
  alias      = null
  cipher     = null
  cluster_id = null
  mtls       = null
  version    = null
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
* `id` (String, computed)
* `mtls` (String, computed)
* `version` (String, computed)


