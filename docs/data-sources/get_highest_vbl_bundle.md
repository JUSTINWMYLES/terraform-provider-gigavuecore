---
page_title: "gigavuecore_get_highest_vbl_bundle Data Source - gigavuecore"
subcategory: ""
description: |-
  Returns the currently installed highest VBL bundle family that is not expired or deactivated (CoreVUE < NetVUE < SecureVUEPlus), or 'Unbundled' if no VBL license exists
---

# gigavuecore_get_highest_vbl_bundle Data Source

Returns the currently installed highest VBL bundle family that is not expired or deactivated (CoreVUE < NetVUE < SecureVUEPlus), or 'Unbundled' if no VBL license exists

~> **Note:** This data source is not yet wired to a remote API endpoint. Invoking it fails with an explicit "not wired" diagnostic instead of calling the API. The OpenAPI operations it was inferred from could not be resolved into a complete mapping; consult the eidos generation warnings for the exact cause.

## Example Usage

```terraform
data "gigavuecore_get_highest_vbl_bundle" "example" {
}
```
