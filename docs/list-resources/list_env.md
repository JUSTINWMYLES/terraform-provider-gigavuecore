---
page_title: "gigavuecore_list_env List Resource - gigavuecore"
subcategory: ""
description: |-
  List all unified resource environments
---

# gigavuecore_list_env List Resource

List all unified resource environments

## Example Usage

```terraform
list "gigavuecore_list_env" "example" {
  provider = gigavuecore
  limit    = 100
}

```
## Schema

### Identity Attributes

The following identity attributes are exported for each matching result:

* `env_id` (String, computed)


