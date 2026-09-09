---
page_title: "gigavuecore_get_serial_numbers_of_chassis_requiring_gvos_license Data Source - gigavuecore"
subcategory: ""
description: |-
  Gets the serial numbers of chassis on which GVOS license can be put
---

# gigavuecore_get_serial_numbers_of_chassis_requiring_gvos_license Data Source

Gets the serial numbers of chassis on which GVOS license can be put

## Example Usage

```terraform
data "gigavuecore_get_serial_numbers_of_chassis_requiring_gvos_license" "example" {
  model = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `model` (String, required) - model of chassis

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List of String, computed) - a list of chassis serial numbers


