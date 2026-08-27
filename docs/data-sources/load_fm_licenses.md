---
page_title: "gigavuecore_load_fm_licenses Data Source - gigavuecore"
subcategory: ""
description: |-
  Load FM Licenses
---

# gigavuecore_load_fm_licenses Data Source

Load FM Licenses

## Example Usage

```terraform
data "gigavuecore_load_fm_licenses" "example" {
}
```

## Schema

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `active` (Boolean) - license is valid, matched target FM instance and within the valid time period
* `description` (String) - Description of the SKU code this license is issued for
* `end_date` (String) - End date in ISO 8601 format. If omitted, license never expires
* `inactive_reason` (String) - If license is inactive, this field contains the reason
* `license_key` (String)
* `revoked` (Boolean) - indicates whether the license key is revoked
* `start_date` (String) - Start date in ISO 8601 format. If omitted, license is effective from the date of issue
* `target_fm_id` (String) - identifies FM instance this license is issued for
* `valid` (Boolean) - license is well-formed, not revoked and matches the target FM instance
* `well_formed` (Boolean) - indicates that FM is able to parse and interpret the license key string

