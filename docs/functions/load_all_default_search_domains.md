---
page_title: "load_all_default_search_domains Function - gigavuecore"
subcategory: ""
description: |-
  Load all default search domains
---

# load_all_default_search_domains Function

Load all default search domains

~> **Note:** This function is not yet wired to a remote API endpoint. Invoking it fails with an explicit "not wired" diagnostic instead of calling the API. The OpenAPI operations it was inferred from could not be resolved into a complete mapping; consult the eidos generation warnings for the exact cause.

## Example Usage

```terraform
# Example: provider::gigavuecore::load_all_default_search_domains()
output "example" {
  value = provider::gigavuecore::load_all_default_search_domains()
}
```

## Signature

```text
load_all_default_search_domains() -> Dynamic
```

## Arguments

The following arguments are supported:



## Return

The function returns a value of type `Dynamic`.
