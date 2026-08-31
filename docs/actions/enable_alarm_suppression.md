---
page_title: "gigavuecore_enable_alarm_suppression Action - gigavuecore"
subcategory: ""
description: |-
  Enable alarm suppression rules for multiple resources
---

# gigavuecore_enable_alarm_suppression Action

Enable alarm suppression rules for multiple resources

## Example Usage

```terraform
action "gigavuecore_enable_alarm_suppression" "example" {
  config {
    suppressed_entities = [{
      alarm_types      = [ "example" ]
      alias            = "example"
      cluster_id       = "example"
      created_by       = "example"
      created_ts       = "example"
      enable           = true
      expiry_time      = 0
      expiry_time_unit = "minutes"
      expiry_ts        = "example"
      hostname         = "example"
      resource_id      = "example"
      resource_type    = "portPair"
      selected_reason  = "example"
      tags = [{
        tag_key    = "example"
        tag_values = [ "example" ]
      }]
      updated_by = "example"
      updated_ts = "example"
    }]
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `suppressed_entities` (Attributes List, required) - Suppression Rules for resource (see [below for nested schema](#nestedatt--suppressed_entities))

<a id="nestedatt--suppressed_entities"></a>
### Nested Schema for `suppressed_entities`

Required:

* `resource_id` (String) - Id of the resource to be suppressed
* `resource_type` (String)

Optional:

* `alarm_types` (List of String)
* `alias` (String) - Alias
* `cluster_id` (String) - Cluster/Node Id
* `created_by` (String) - Alarm suppression rule created by FM User
* `created_ts` (String) - Alarm suppression rule creation timestamp in ISO 8601 format
* `enable` (Boolean) - FM Alarm Suppression rule state
* `expiry_time` (Number) - Expiry interval for suppression rule
* `expiry_time_unit` (String) - Expiry interval unit
* `expiry_ts` (String) - Expiry Timestamp(UTC) for suppression rule
* `hostname` (String) - FM Hostname
* `selected_reason` (String) - Suppression Reason
* `tags` (Attributes List) (see [below for nested schema](#nestedatt--suppressed_entities--tags))
* `updated_by` (String) - Alarm suppression rule updated by FM User
* `updated_ts` (String) - Alarm suppression rule updated timestamp in ISO 8601 format

<a id="nestedatt--suppressed_entities--tags"></a>
### Nested Schema for `suppressed_entities.tags`

Required:

* `tag_key` (String) - Name of the tag
* `tag_values` (List of String) - All possible values of the tag

