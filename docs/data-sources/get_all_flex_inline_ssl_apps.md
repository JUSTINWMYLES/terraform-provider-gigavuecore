---
page_title: "gigavuecore_get_all_flex_inline_ssl_apps Data Source - gigavuecore"
subcategory: ""
description: |-
  Get All Inline ssl apps across FM
---

# gigavuecore_get_all_flex_inline_ssl_apps Data Source

Get All Inline ssl apps across FM

## Example Usage

```terraform
data "gigavuecore_get_all_flex_inline_ssl_apps" "example" {
  cluster_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, optional) - If provided, ssl apps only for that cluster are returned

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({alias, app_intent_configs, cluster_name, config_status, config_status_reasons, health_state, health_state_reasons, m_tls, ria_configs, ria_enabled})), computed)

