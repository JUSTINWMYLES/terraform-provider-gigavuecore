---
page_title: "gigavuecore_update_email_receivers Action - gigavuecore"
subcategory: ""
description: |-
  Update the optional email receivers, separate lists for VBL reports and license expiry (for upcoming renewals)
---

# gigavuecore_update_email_receivers Action

Update the optional email receivers, separate lists for VBL reports and license expiry (for upcoming renewals)

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_update_email_receivers" "example" {
  config {
    upcoming_renewals = ["example"]
    vbl_report        = ["example"]
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `upcoming_renewals` (List of String, optional)
* `vbl_report` (List of String, optional)


