---
page_title: "gigavuecore_user Resource - gigavuecore"
subcategory: ""
description: |-
  Create new user
---

# gigavuecore_user Resource

Create new user

~> **Note:** This resource is not yet wired to a remote API endpoint. Invoking it fails with an explicit "not wired" diagnostic instead of calling the API. The OpenAPI operations it was inferred from could not be resolved into a complete mapping; consult the eidos generation warnings for the exact cause.

## Example Usage

```terraform
resource "gigavuecore_user" "example" {
  email_id  = "example"
  enabled   = true
  full_name = "example"
  groups    = ["example"]
  password  = "example-value"
  username  = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `email_id` (String, required) - email ID
* `enabled` (Boolean, optional)
* `full_name` (String, optional) - user's full name
* `groups` (List of String, optional)
* `password` (String, required) - password
* `username` (String, required) - username

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

* `create` (String) - A create timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `read` (String) - A read timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).
* `update` (String) - An update timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `delete` (String) - A delete timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).

