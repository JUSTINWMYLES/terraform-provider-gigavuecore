---
page_title: "gigavuecore_transfer_bindings Action - gigavuecore"
subcategory: ""
description: |-
  Transfer an array of bindings to various license targets (Gigasmart cards) based on a list of pairs consisting of source bindings to be deleted and the destination to be transferred to (s1, d1, s2, d2, ... where "s" is source and "d" is destination)
---

# gigavuecore_transfer_bindings Action

Transfer an array of bindings to various license targets (Gigasmart cards) based on a list of pairs consisting of source bindings to be deleted and the destination to be transferred to (s1, d1, s2, d2, ... where "s" is source and "d" is destination)

## Example Usage

```terraform
action "gigavuecore_transfer_bindings" "example" {
  config {
    activation_id = "example"
    bindings      = null
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `activation_id` (String, required) - ID of the feature activation
* `bindings` (List of Dynamic, optional)


