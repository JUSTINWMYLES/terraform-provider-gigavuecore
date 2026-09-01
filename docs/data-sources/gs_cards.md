---
page_title: "gigavuecore_gs_cards Data Source - gigavuecore"
subcategory: ""
description: |-
  Get Gigasmart card information
---

# gigavuecore_gs_cards Data Source

Get Gigasmart card information

## Example Usage

```terraform
data "gigavuecore_gs_cards" "example" {
  cluster_id = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, optional) - if provided, GS card information of the requested cluster will be returned

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `cards` (Attributes List) (see [below for nested schema](#nestedatt--items--cards))
* `cluster_id` (String)

<a id="nestedatt--items--cards"></a>
### Nested Schema for `items.cards`

Read-Only:

* `box_id` (String)
* `hw_type` (String)
* `product_code` (String)
* `related_gs_groups` (Attributes List) (see [below for nested schema](#nestedatt--items--cards--related_gs_groups))
* `sig_revision` (Attributes List) (see [below for nested schema](#nestedatt--items--cards--sig_revision))
* `slot_id` (String)

<a id="nestedatt--items--cards--related_gs_groups"></a>
### Nested Schema for `items.cards.related_gs_groups`

Read-Only:

* `alias` (String)
* `ports` (List of String)

<a id="nestedatt--items--cards--sig_revision"></a>
### Nested Schema for `items.cards.sig_revision`

Read-Only:

* `engine_id` (String)
* `version` (String)

