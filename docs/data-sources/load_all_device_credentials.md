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
  page = "example"
  sort = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `device_address` (String) - value of '0.0.0.0' represents the default (fallback) device credentials
* `hostname` (String) - device configured hostname
* `http_password` (String) - password to use for device login. On reads, '\*\*\*\*\*' is returned
* `http_username` (String) - username to use for device login
* `https_port` (String) - httpsPort to use for device communication. By default 443 is used, If changed the same should be given here
* `snmp_version` (String) - SNMP version to use.

