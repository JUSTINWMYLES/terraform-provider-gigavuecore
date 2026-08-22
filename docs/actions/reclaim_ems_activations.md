---
page_title: "gigavuecore_reclaim_ems_activations Action - gigavuecore"
subcategory: ""
description: |-
  Reclaim a license activation, return to entitlement the associated quantity(number of licenses)
---

# gigavuecore_reclaim_ems_activations Action

Reclaim a license activation, return to entitlement the associated quantity(number of licenses)

## Example Usage

```terraform
action "gigavuecore_reclaim_ems_activations" "example" {
  config {
    aid = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `aid` (String, required) - Activation ID
