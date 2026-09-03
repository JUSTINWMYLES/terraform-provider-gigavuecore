---
page_title: "gigavuecore_get_current_mon_sessions_to_app_tier_map Data Source - gigavuecore"
subcategory: ""
description: |-
  Gives the mapping from current monitoring session id's to corresponding highest app tier as determined from the apps in the monitoring session
---

# gigavuecore_get_current_mon_sessions_to_app_tier_map Data Source

Gives the mapping from current monitoring session id's to corresponding highest app tier as determined from the apps in the monitoring session

~> **Note:** This data source is not yet wired to a remote API endpoint. Invoking it fails with an explicit "not wired" diagnostic instead of calling the API. The OpenAPI operations it was inferred from could not be resolved into a complete mapping; consult the eidos generation warnings for the exact cause.

## Example Usage

```terraform
data "gigavuecore_get_current_mon_sessions_to_app_tier_map" "example" {
}
```
