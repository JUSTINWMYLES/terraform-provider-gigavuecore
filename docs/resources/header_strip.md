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
  box_id = null
  mpls_labels = []
}
```

## Schema

### Arguments

The following arguments are supported:

* `box_id` (String, required) - device box id. valid range 1 - 64.
* `mpls_labels` (List(String), optional) - mpls ids, valid and required. Range can be specified. Example:1..200

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `mpls_labels` (List(String), computed) - mpls ids, valid and required. Range can be specified. Example:1..200

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_header_strip.example {box_id}
```
