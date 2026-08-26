---
page_title: "gigavuecore_load_all_netflow_monitors Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all Netwflow Monitors
---

# gigavuecore_load_all_netflow_monitors Data Source

Load all Netwflow Monitors

## Example Usage

```terraform
data "gigavuecore_load_all_netflow_monitors" "example" {
  cluster_id = null
  page       = null
  sort       = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `alias` (String)
* `cache` (Attributes) - Netflow Monitor Cache (see [below for nested schema](#nestedatt--items--cache))
* `cluster_id` (String) - id of the defining cluster
* `description` (String)
* `records` (Set of String) - Aliases of referenced Netflow Records defined on the GsGroup. Up to 5 records.
* `sampling` (Attributes) - monitor sampling (see [below for nested schema](#nestedatt--items--sampling))
* `sampling_space` (Number) - DEPRECATED: use 'sampling'
* `ssl_port_restrictions` (Attributes) - Port restrictions for Netflow/SSL sessions (see [below for nested schema](#nestedatt--items--ssl_port_restrictions))
<a id="nestedatt--items--cache"></a>
### Nested Schema for `items.cache`

Read-Only:

* `export_triggers` (Attributes) - Netflow Monitor Cache Export Triggers (see [below for nested schema](#nestedatt--items--cache--export_triggers))
* `type` (String)
<a id="nestedatt--items--cache--export_triggers"></a>
### Nested Schema for `items.cache.export_triggers`

Read-Only:

* `event` (String)
* `timeout_active` (Number) - in seconds. max value is 7 days. default is 30 min
* `timeout_inactive` (Number) - in seconds. max value is 7 days. default is 15 sec
<a id="nestedatt--items--sampling"></a>
### Nested Schema for `items.sampling`

Read-Only:

* `mode` (String)
* `single_sampling_rate` (Number) - Packet interval window size. Valid values: 10-16000 (in packets)
<a id="nestedatt--items--ssl_port_restrictions"></a>
### Nested Schema for `items.ssl_port_restrictions`

Read-Only:

* `ports` (List of Number) - The list of TCP ports whose packets will be sent to the SSL module. Not valid if sslPort is 'all'. if 'wellKnownPorts' is selected for 'sslPorts' ports will be \[993, 995, 465, 636, 563, 443\]
* `ssl_ports` (String) - Ports whose packets will be sent to the SSL module. 'all' ports; 'wellKnownPorts' is \[993, 995, 465, 636, 563, 443\]; or 'ports' - specify upto 10 ports

