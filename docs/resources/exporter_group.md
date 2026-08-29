---
page_title: "gigavuecore_exporter_group Resource - gigavuecore"
subcategory: ""
description: |-
  Get Apps Exporter Group for given alias
---

# gigavuecore_exporter_group Resource

Get Apps Exporter Group for given alias

## Example Usage

```terraform
resource "gigavuecore_exporter_group" "example" {
  alias       = "example"
  cluster_id  = "example"
  description = "example"
  exporters   = [ "example" ]
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required)
* `cluster_id` (String, required) - Target Cluster ID
* `description` (String, optional)
* `exporters` (List of String, optional)

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `description` (String, computed)
* `exporters` (List of String, computed)


## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_exporter_group.example {alias}
```
