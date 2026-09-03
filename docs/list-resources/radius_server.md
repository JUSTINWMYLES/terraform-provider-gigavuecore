---
page_title: "gigavuecore_radius_server List Resource - gigavuecore"
subcategory: ""
description: |-
  Load all RADIUS Servers
---

# gigavuecore_radius_server List Resource

Load all RADIUS Servers

-> **Note:** This list resource requires Terraform 1.14 or later and is used through the `terraform query` command, not in configuration files.

~> **Warning:** This list resource accepts attributes whose names indicate secrets (secret_key), but list resource schemas cannot mark attributes Sensitive, so their values are stored and displayed in plain text. Avoid passing real secrets through these attributes.

## Example Usage

```terraform
list "gigavuecore_radius_server" "example" {
  provider = gigavuecore
  limit    = 100
  config {
    cluster_id = "example"
    page       = "example"
    sort       = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)


### Identity Attributes

The following identity attributes are exported for each matching result:

* `server_address` (String, computed) - ipv4 or ipv6 or hostname


