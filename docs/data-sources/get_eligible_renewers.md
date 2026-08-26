---
page_title: "gigavuecore_get_eligible_renewers Data Source - gigavuecore"
subcategory: ""
description: |-
  Get eligible renewable licenses
---

# gigavuecore_get_eligible_renewers Data Source

Get eligible renewable licenses

## Example Usage

```terraform
data "gigavuecore_get_eligible_renewers" "example" {
  renewee_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `renewee_id` (String, required) - Renewee ID

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `aid` (String) - activation id as encoded in the license
* `binding_errors` (Boolean) - true if there are heartbeat misses (from FM to device) or any operational error on license bindings to device
* `bindings` (Attributes List) - provides information on the chassis/cards where licenses from this floating pool are assigned to (see [below for nested schema](#nestedatt--items--bindings))
* `bundle` (String) - this field is unused in floating licenses
* `description` (String) - a description of this subset of features
* `eid` (String) - entitlement id as encoded in the license
* `end_date` (String) - end date of the license, when it transitions out of ACTIVE state
* `grace_days` (Number) - number of days of grace period beyond the license end date that allows the functionality provided by the license to be used as if the license was still active
* `license_status` (String) - status of the license
* `license_type` (String) - type of the license
* `num_licenses` (Number) - licenses provided by the SKU is multiplied by this factor
* `registration_date` (String) - date when the license was imported into FM
* `revocation_code` (String) - explains why the license was revoked
* `revocation_date` (String) - date when the license was revoked
* `revoked` (Boolean) - was the license revoked
* `sku` (String) - the license SKU (Stock Keeping Unit) or product code that uniquely identifies the license type in Gigamon catalog
* `start_date` (String) - start date of the license, when it transitions into ACTIVE state
* `total_volume` (String) - unused for floating licenses
<a id="nestedatt--items--bindings"></a>
### Nested Schema for `items.bindings`

Read-Only:

* `box_id` (Number)
* `cluster_name` (String)
* `license_key` (String) - encoded string of a license key that looks like LK2-SMT\_HC3-7YF0-86GT-1CEG-LMEU-4KTL-EUPJ-4WHF-L2A4-L3 when decoded
* `slot_id` (Number)
* `target_type` (String)

