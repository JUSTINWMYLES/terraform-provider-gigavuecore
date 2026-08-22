---
page_title: "gigavuecore_latest_import_tag_result Data Source - gigavuecore"
subcategory: ""
description: |-
  get the latest import result all tags
---

# gigavuecore_latest_import_tag_result Data Source

get the latest import result all tags

## Example Usage

```terraform
data "gigavuecore_latest_import_tag_result" "example" {
}
```

## Schema

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `completed` (Number, computed) - number of operations completed
* `end_time` (Number, computed) - Operation end timestamp in UTC milliseconds
* `errors_count` (Number, computed) - number of operations failed
* `file_type` (String, computed) - type of the file imported
* `operation_type` (String, computed)
* `skipped_count` (Number, computed) - number of operations skipped
* `start_time` (Number, computed) - Operation start timestamp in UTC milliseconds
* `status` (String, computed) - operation status
* `total` (Number, computed) - total number of operations

