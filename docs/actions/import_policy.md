---
page_title: "gigavuecore_import_policy Action - gigavuecore"
subcategory: ""
description: |-
  Import Policy
---

# gigavuecore_import_policy Action

Import Policy

## Example Usage

```terraform
action "gigavuecore_import_policy" "example" {
  config {
    criteria_bindings              = { "key" = "example" }
    name                           = "example"
    packet_transformation_bindings = { "key" = "example" }
    priority                       = true
    rules = [{
      criteria = [{
        criteria_name = "example"
        filters = [{
          properties = [{
            key    = "example"
            values = [ "example" ]
          }]
          type   = "example"
          values = [ "example" ]
        }]
      }]
      high_priority_drop = true
      no_expansion_tags  = [ "example" ]
      packet_transformation = {
        apps = [{
          parameters = "example"
          type       = "example"
        }]
        engine_ports = [ "example" ]
      }
      rule_name = "example"
      tags = [{
        key    = "example"
        values = [ "example" ]
      }]
      tools = [{
        ports = [ "example" ]
        type  = "example"
      }]
      type = "example"
    }]
    source_bindings = { "key" = "example" }
    sources = [{
      ports = [ "example" ]
      type  = "example"
    }]
    tags = [{
      key    = "example"
      values = [ "example" ]
    }]
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `criteria_bindings` (Map of List of Object, optional)
* `name` (String, optional)
* `packet_transformation_bindings` (Map of List of Object, optional)
* `priority` (Boolean, optional)
* `rules` (Attributes List, optional) (see [below for nested schema](#nestedatt--rules))
* `source_bindings` (Map of List of Object, optional)
* `sources` (Attributes List, optional) (see [below for nested schema](#nestedatt--sources))
* `tags` (Attributes List, optional) (see [below for nested schema](#nestedatt--tags))

<a id="nestedatt--rules"></a>
### Nested Schema for `rules`

Optional:

* `criteria` (Attributes List) (see [below for nested schema](#nestedatt--rules--criteria))
* `high_priority_drop` (Boolean)
* `no_expansion_tags` (List of String)
* `packet_transformation` (Attributes) (see [below for nested schema](#nestedatt--rules--packet_transformation))
* `rule_name` (String)
* `tags` (Attributes List) (see [below for nested schema](#nestedatt--rules--tags))
* `tools` (Attributes List) (see [below for nested schema](#nestedatt--rules--tools))
* `type` (String)

<a id="nestedatt--rules--criteria"></a>
### Nested Schema for `rules.criteria`

Optional:

* `criteria_name` (String)
* `filters` (Attributes List) (see [below for nested schema](#nestedatt--rules--criteria--filters))

<a id="nestedatt--rules--criteria--filters"></a>
### Nested Schema for `rules.criteria.filters`

Optional:

* `properties` (Attributes List) (see [below for nested schema](#nestedatt--rules--criteria--filters--properties))
* `type` (String)
* `values` (List of String)

<a id="nestedatt--rules--criteria--filters--properties"></a>
### Nested Schema for `rules.criteria.filters.properties`

Optional:

* `key` (String)
* `values` (List of Dynamic)

<a id="nestedatt--rules--packet_transformation"></a>
### Nested Schema for `rules.packet_transformation`

Optional:

* `apps` (Attributes List) (see [below for nested schema](#nestedatt--rules--packet_transformation--apps))
* `engine_ports` (List of String)

<a id="nestedatt--rules--packet_transformation--apps"></a>
### Nested Schema for `rules.packet_transformation.apps`

Optional:

* `parameters` (Dynamic)
* `type` (String)

<a id="nestedatt--rules--tags"></a>
### Nested Schema for `rules.tags`

Optional:

* `key` (String)
* `values` (List of String)

<a id="nestedatt--rules--tools"></a>
### Nested Schema for `rules.tools`

Optional:

* `ports` (List of String)
* `type` (String)

<a id="nestedatt--sources"></a>
### Nested Schema for `sources`

Optional:

* `ports` (List of String)
* `type` (String)

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

* `key` (String)
* `values` (List of String)

