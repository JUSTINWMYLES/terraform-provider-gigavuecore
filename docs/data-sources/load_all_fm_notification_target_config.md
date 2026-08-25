---
page_title: "gigavuecore_load_all_fm_notification_target_config Data Source - gigavuecore"
subcategory: ""
description: |-
  Load All Fm Notification Target Config
---

# gigavuecore_load_all_fm_notification_target_config Data Source

Load All Fm Notification Target Config

## Example Usage

```terraform
data "gigavuecore_load_all_fm_notification_target_config" "example" {
  page = null
  sort = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `context` (Attributes, computed) - Gigamon query result context (see [below for nested schema](#nestedatt--context))
* `fm_notification_target_configs` (Attributes List, computed) (see [below for nested schema](#nestedatt--fm_notification_target_configs))

<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `page_no` (Number) - page number of the returned result set
* `page_size` (Number) - page size of the returned result set
* `sort` (List of String) - sorting info of the returned result set. list of fields in the array indicate sorting order
* `total_items` (Number) - total number of items in the queried entity type
<a id="nestedatt--fm_notification_target_configs"></a>
### Nested Schema for `fm_notification_target_configs`

Read-Only:

* `interface_name` (String) - Name of any one of the available network interface names on the FM. If this is chosen, FM registers itself as a notification target on the node with Ipv6 address if both FM and the node have Ipv6 address otherwise Ipv4 address is used. If targetAddress is configured then interfaceName has no effect.
* `interface_type` (String) - Notification target interface type
* `target_address` (String) - Configure FM's DNS name or static IP address to receive the management or data traffic from the node. The configured address is used by FM to register itself as a notification target on the node

