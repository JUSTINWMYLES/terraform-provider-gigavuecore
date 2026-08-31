---
page_title: "gigavuecore_download_list Data Source - gigavuecore"
subcategory: ""
description: |-
  Download a nodecryptlist or decryptlist
---

# gigavuecore_download_list Data Source

Download a nodecryptlist or decryptlist

## Example Usage

```terraform
data "gigavuecore_download_list" "example" {
  alias      = "example"
  cluster_id = "example"
  list_type  = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Alias of the inline SSL profile
* `cluster_id` (String, required) - Target Cluster ID
* `list_type` (String, required) - specify a nodecryptlist or decryptlist. (deprecated: use nodecryptlist and decryptlist instead of whitelist and blacklist)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `list` (String, computed)


