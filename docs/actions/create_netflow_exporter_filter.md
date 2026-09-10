---
page_title: "gigavuecore_create_netflow_exporter_filter Action - gigavuecore"
subcategory: ""
description: |-
  Create Netflow Exporter Filter
---

# gigavuecore_create_netflow_exporter_filter Action

Create Netflow Exporter Filter

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_create_netflow_exporter_filter" "example" {
  config {
    alias      = "example"
    cluster_id = "example"
    rules = [{
      pass_rules = [{
        matches = ["example"]
        rule_id = 1
      }]
    }]
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target Netflow Exporter
* `cluster_id` (String, required) - Target Cluster ID
* `rules` (Attributes List, required) (see [below for nested schema](#nestedatt--rules))

<a id="nestedatt--rules"></a>
### Nested Schema for `rules`

Optional:

* `pass_rules` (Attributes List) (see [below for nested schema](#nestedatt--rules--pass_rules))

<a id="nestedatt--rules--pass_rules"></a>
### Nested Schema for `rules.pass_rules`

Required:

* `rule_id` (Number)

Optional:

* `matches` (List of Dynamic) - Set of rule's matching elements. Within a rule, matching elements of the the same type MAY be used more than once, However, their matching positions MUST be unique

