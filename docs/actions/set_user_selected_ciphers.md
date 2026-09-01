---
page_title: "gigavuecore_set_user_selected_ciphers Action - gigavuecore"
subcategory: ""
description: |-
  Set user-selected ciphers by protocol type
---

# gigavuecore_set_user_selected_ciphers Action

Set user-selected ciphers by protocol type

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_set_user_selected_ciphers" "example" {
  config {
    source_details = [{
      cluster_id = "example"
      components = [{
        ids  = ["example"]
        type = "PORT"
      }]
    }]
    type = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `source_details` (Attributes List, optional) (see [below for nested schema](#nestedatt--source_details))
* `type` (String, required) - Cipher protocol type (for example, TLS, SSH)

<a id="nestedatt--source_details"></a>
### Nested Schema for `source_details`

Required:

* `cluster_id` (String)

Optional:

* `components` (Attributes List) (see [below for nested schema](#nestedatt--source_details--components))

<a id="nestedatt--source_details--components"></a>
### Nested Schema for `source_details.components`

Optional:

* `ids` (List of String)
* `type` (String)

