---
page_title: "gigavuecore_get_serial_number_of_card Data Source - gigavuecore"
subcategory: ""
description: |-
  Gets the serial numbers of card specified by input parameters
---

# gigavuecore_get_serial_number_of_card Data Source

Gets the serial numbers of card specified by input parameters

~> **Note:** This data source is not yet wired to a remote API endpoint. Invoking it fails with an explicit "not wired" diagnostic instead of calling the API. The OpenAPI operations it was inferred from could not be resolved into a complete mapping; consult the eidos generation warnings for the exact cause.

## Example Usage

```terraform
data "gigavuecore_get_serial_number_of_card" "example" {
  box_slash_slot_id = "example"
  cluster_name      = "example"
  node_id           = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `box_slash_slot_id` (String, required) - <box-id>/<slot-id>
* `cluster_name` (String, required) - name of chassis cluster
* `node_id` (String, required) - node ID, the IP address of the chassis


