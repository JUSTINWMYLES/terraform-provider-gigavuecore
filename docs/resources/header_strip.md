---
page_title: "gigavuecore_header_strip Resource - gigavuecore"
subcategory: ""
description: |-
  Load header strip for target box
---

# gigavuecore_header_strip Resource

Load header strip for target box

## Example Usage

```terraform
resource "gigavuecore_header_strip" "example" {
  box_id      = "example"
  cluster_id  = "example"
  mpls_labels = [ "example" ]
}
```

## Schema

### Arguments

The following arguments are supported:

* `box_id` (String, required) - device box id. valid range 1 - 64.
* `cluster_id` (String, required) - Target cluster ID.
* `mpls_labels` (List of String, optional) - mpls ids, valid and required. Range can be specified. Example:1..200

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `mpls_labels` (List of String, computed) - mpls ids, valid and required. Range can be specified. Example:1..200


## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_header_strip.example {box_id}
```
