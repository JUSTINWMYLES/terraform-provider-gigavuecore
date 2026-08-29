---
page_title: "gigavuecore_create_bindings Action - gigavuecore"
subcategory: ""
description: |-
  Create an array of bindings to various license targets (chassis or card type, but not both)
---

# gigavuecore_create_bindings Action

Create an array of bindings to various license targets (chassis or card type, but not both)

## Example Usage

```terraform
action "gigavuecore_create_bindings" "example" {
  config {
    activation_id = "example"
    bindings = [{
      box_id       = 0
      cluster_name = "example"
      license_key  = "example"
      slot_id      = 0
      target_type  = "example"
    }]
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `activation_id` (String, required) - FM License Key bindings to create
* `bindings` (Attributes List, optional) (see [below for nested schema](#nestedatt--bindings))

<a id="nestedatt--bindings"></a>
### Nested Schema for `bindings`

Optional:

* `box_id` (Number)
* `cluster_name` (String)
* `license_key` (String) - encoded string of a license key that looks like LK2-SMT\_HC3-7YF0-86GT-1CEG-LMEU-4KTL-EUPJ-4WHF-L2A4-L3 when decoded
* `slot_id` (Number)
* `target_type` (String)

