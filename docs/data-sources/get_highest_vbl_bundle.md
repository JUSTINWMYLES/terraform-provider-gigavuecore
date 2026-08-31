---
page_title: "gigavuecore_get_highest_vbl_bundle Data Source - gigavuecore"
subcategory: ""
description: |-
  Returns the currently installed highest VBL bundle family that is not expired or deactivated (CoreVUE < NetVUE < SecureVUEPlus), or 'Unbundled' if no VBL license exists
---

# gigavuecore_get_highest_vbl_bundle Data Source

Returns the currently installed highest VBL bundle family that is not expired or deactivated (CoreVUE < NetVUE < SecureVUEPlus), or 'Unbundled' if no VBL license exists

## Example Usage

```terraform
data "gigavuecore_get_highest_vbl_bundle" "example" {
}
```
