---
page_title: "gigavuecore_load_email_server_configuration Data Source - gigavuecore"
subcategory: ""
description: |-
  Load Email Server Configuration
---

# gigavuecore_load_email_server_configuration Data Source

Load Email Server Configuration

## Example Usage

```terraform
data "gigavuecore_load_email_server_configuration" "example" {
}
```

## Schema

### Attributes

In addition to all arguments above, the following attributes are exported:

* `email_host` (String, computed) - Address of the SMTP Server
* `enable_smtp_auth` (Boolean, computed) - Enable/Disable SMTP Authentication
* `from` (String, computed) - From address for the email
* `password` (String, computed) - Password for the SMTP Server
* `port` (Number, computed) - SMTP Server Port
* `user_name` (String, computed) - Username for the SMTP Server


