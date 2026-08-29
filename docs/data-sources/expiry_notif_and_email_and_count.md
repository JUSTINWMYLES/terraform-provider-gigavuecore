---
page_title: "gigavuecore_expiry_notif_and_email_and_count Data Source - gigavuecore"
subcategory: ""
description: |-
  Returns expiry count for floating (90 day, expiring soon, expired) then same for node-locked licenses, after generating events and sending email for those
---

# gigavuecore_expiry_notif_and_email_and_count Data Source

Returns expiry count for floating (90 day, expiring soon, expired) then same for node-locked licenses, after generating events and sending email for those

## Example Usage

```terraform
data "gigavuecore_expiry_notif_and_email_and_count" "example" {
}
```

## Schema

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List of Number, computed) - first three numbers are for floating licenses, second three are for node-locked licenses; the numbers signify count expiring in 90 days, expiring soon (within notification period), and expired (beyond end date)


