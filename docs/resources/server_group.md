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
  alias        = "example"
  cluster_id   = "example"
  comment      = "example"
  icap_servers = [ "example" ]
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Icap Server Group Alias
* `cluster_id` (String, required) - id of the defining cluster
* `comment` (String, optional) - Icap Server Group Comment
* `icap_servers` (List of String, optional) - ICAP Servers list separated by comma','

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `comment` (String, computed) - Icap Server Group Comment
* `icap_servers` (List of String, computed) - ICAP Servers list separated by comma','


## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_server_group.example {alias}
```
