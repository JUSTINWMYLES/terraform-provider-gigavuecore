---
page_title: "gigavuecore_update_email_receivers Action - gigavuecore"
subcategory: ""
description: |-
  Update the optional email receivers, separate lists for VBL reports and license expiry (for upcoming renewals)
---

# gigavuecore_update_email_receivers Action

Update the optional email receivers, separate lists for VBL reports and license expiry (for upcoming renewals)

## Example Usage

```terraform
action "gigavuecore_update_email_receivers" "example" {
  config {
    upcoming_renewals = [ "example" ]
    vbl_report = [ "example" ]
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `upcoming_renewals` (List(String), optional)
* `vbl_report` (List(String), optional)
