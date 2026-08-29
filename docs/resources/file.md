---
page_title: "gigavuecore_file Resource - gigavuecore"
subcategory: ""
description: |-
  Search inline SSL profile nodecryptlist or decryptlist for a domain
---

# gigavuecore_file Resource

Search inline SSL profile nodecryptlist or decryptlist for a domain

## Example Usage

```terraform
resource "gigavuecore_file" "example" {
  cluster_id = "example"
  file       = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `file` (String, required) - file to upload to device

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `id` (String, computed)
* `result` (String, computed) - (deprecated: use resultAlias instead)
* `result_alias` (String, computed)


