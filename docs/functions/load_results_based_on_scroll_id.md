---
page_title: "gigavuecore_load_results_based_on_scroll_id Function - gigavuecore"
subcategory: ""
description: |-
  Load results  based on ScrollId
---

# gigavuecore_load_results_based_on_scroll_id Function

Load results  based on ScrollId

## Example Usage

```terraform
# Example: provider::gigavuecore::gigavuecore_load_results_based_on_scroll_id("<scroll_id>", "<response_scroll_timeout>")
output "example" {
  value = provider::gigavuecore::gigavuecore_load_results_based_on_scroll_id("<scroll_id>", "<response_scroll_timeout>")
}
```

## Signature

```text
gigavuecore_load_results_based_on_scroll_id(scroll_id: String, response_scroll_timeout: String) -> String
```

## Arguments

The following arguments are supported:

* `scroll_id` (String) - scrollId which is used in the scroll API in order to retrieve the required batch of results
* `response_scroll_timeout` (String) - Scroll time out to mention how long it should keep the "search context" alive


## Return

The function returns a value of type `String`.
