---
page_title: "gigavuecore_map_group Resource - gigavuecore"
subcategory: ""
description: |-
  Find Map Group by alias
---

# gigavuecore_map_group Resource

Find Map Group by alias

## Example Usage

```terraform
resource "gigavuecore_map_group" "example" {
  alias   = null
  comment = null
  maps    = []
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Map Group alias
* `comment` (String, optional)
* `maps` (List of String, required) - List of maps.SecondLevel maps with flowSample-overlap or flowWhitelist-overlap subtype

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `cluster_id` (String, computed) - id of the defining cluster
* `comment` (String, computed)


## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_map_group.example {alias}
```
