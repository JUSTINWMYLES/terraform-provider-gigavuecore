---
page_title: "load_results_based_on_scroll_id Function - gigavuecore"
subcategory: ""
description: |-
  Load results  based on ScrollId
---

# load_results_based_on_scroll_id Function

Load results  based on ScrollId

~> **Note:** This function is not yet wired to a remote API endpoint. Invoking it fails with an explicit "not wired" diagnostic instead of calling the API. The OpenAPI operations it was inferred from could not be resolved into a complete mapping; consult the eidos generation warnings for the exact cause.

## Example Usage

```terraform
# Example: provider::gigavuecore::load_results_based_on_scroll_id("<scroll_id>", "<response_scroll_timeout>")
output "example" {
  value = provider::gigavuecore::load_results_based_on_scroll_id("<scroll_id>", "<response_scroll_timeout>")
}
```

## Signature

```text
load_results_based_on_scroll_id(scroll_id: String, response_scroll_timeout: String) -> String
```

## Arguments

The following arguments are supported:

* `scroll_id` (String) - scrollId which is used in the scroll API in order to retrieve the required batch of results
* `response_scroll_timeout` (String) - Scroll time out to mention how long it should keep the "search context" alive


## Return

The function returns a value of type `String`.
