---
page_title: "gigavuecore_get_all_devices Data Source - gigavuecore"
subcategory: ""
description: |-
  Get a list of physical devices on which license can be allocated, the selection criteria for devices determined by the query parameters
---

# gigavuecore_get_all_devices Data Source

Get a list of physical devices on which license can be allocated, the selection criteria for devices determined by the query parameters

## Example Usage

```terraform
data "gigavuecore_get_all_devices" "example" {
  activation_id = null
  device_model = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `activation_id` (String, optional) - activation ID of the license; if provided, devices that are relevant for the SKU / features of the activation are returned
* `device_model` (String, required) - model of device, used only if 'activationId' is null

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({box_id, cards, cluster_name, deleted, healthy, host_id, host_name, model, node_id, product_version, serial_number})), computed)

