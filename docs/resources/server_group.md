---
page_title: "gigavuecore_server_group Resource - gigavuecore"
subcategory: ""
description: |-
  Load ICAP Server Group by Alias
---

# gigavuecore_server_group Resource

Load ICAP Server Group by Alias

## Example Usage

```terraform
resource "gigavuecore_server_group" "example" {
  alias        = null
  comment      = null
  icap_servers = []
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Icap Server Group Alias
* `comment` (String, optional) - Icap Server Group Comment
* `icap_servers` (List of String, optional) - ICAP Servers list separated by comma','

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `cluster_id` (String, computed) - id of the defining cluster
* `comment` (String, computed) - Icap Server Group Comment
* `icap_servers` (List of String, computed) - ICAP Servers list separated by comma','


## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_server_group.example {alias}
```
