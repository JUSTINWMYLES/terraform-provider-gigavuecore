---
page_title: "gigavuecore_load_all_device_credentials Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all Device Credentials
---

# gigavuecore_load_all_device_credentials Data Source

Load all Device Credentials

## Example Usage

```terraform
data "gigavuecore_load_all_device_credentials" "example" {
  page = null
  sort = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `context` (Attributes, computed) - Gigamon query result context (see [below for nested schema](#nestedatt--context))
* `dev_creds_list` (Attributes List, computed) (see [below for nested schema](#nestedatt--dev_creds_list))

<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `page_no` (Number) - page number of the returned result set
* `page_size` (Number) - page size of the returned result set
* `sort` (List of String) - sorting info of the returned result set. list of fields in the array indicate sorting order
* `total_items` (Number) - total number of items in the queried entity type
<a id="nestedatt--dev_creds_list"></a>
### Nested Schema for `dev_creds_list`

Read-Only:

* `device_address` (String) - value of '0.0.0.0' represents the default (fallback) device credentials
* `hostname` (String) - device configured hostname
* `http_password` (String) - password to use for device login. On reads, '\*\*\*\*\*' is returned
* `http_username` (String) - username to use for device login
* `https_port` (String) - httpsPort to use for device communication. By default 443 is used, If changed the same should be given here
* `snmp_version` (String) - SNMP version to use.

