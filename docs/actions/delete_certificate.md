---
page_title: "gigavuecore_delete_certificate Action - gigavuecore"
subcategory: ""
description: |-
  Push global delete configuration for the selected certificate to all the devices
---

# gigavuecore_delete_certificate Action

Push global delete configuration for the selected certificate to all the devices

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_delete_certificate" "example" {
  config {
    config = {
      device_ssl_certificate_configs = [{
        issuer              = "example"
        not_after           = "example"
        not_before          = "example"
        operation_type      = "add"
        signature_algorithm = "example"
        subject             = "example"
        trusted_ca = {
          name = "example"
        }
        upload_spec = {
          info = {
            comment    = "example"
            name       = "example"
            passphrase = "example"
            type       = "privateKey"
          }
          pem = "example"
        }
      }]
    }
    config_level       = "GLOBAL"
    config_level_value = ["example"]
    config_type        = "SSL_CERTIFICATE_TEMPLATE"
    modifiable         = true
    ref_count          = 0
    template_name      = "example"
    update_time        = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `config` (Attributes, optional) (see [below for nested schema](#nestedatt--config))
* `config_level` (String, optional) - Scope of the applied FM template
* `config_level_value` (List of String, optional)
* `config_type` (String, optional) - Configuration Type of the FM template
* `modifiable` (Boolean, optional)
* `ref_count` (Number, optional)
* `template_name` (String, optional)
* `update_time` (String, optional)

<a id="nestedatt--config"></a>
### Nested Schema for `config`

Optional:

* `device_ssl_certificate_configs` (Attributes List) - Available when ConfigType is SSL\_CERTIFICATE\_TEMPLATE (see [below for nested schema](#nestedatt--config--device_ssl_certificate_configs))

<a id="nestedatt--config--device_ssl_certificate_configs"></a>
### Nested Schema for `config.device_ssl_certificate_configs`

Required:

* `operation_type` (String)
* `trusted_ca` (Attributes) (see [below for nested schema](#nestedatt--config--device_ssl_certificate_configs--trusted_ca))
* `upload_spec` (Attributes) (see [below for nested schema](#nestedatt--config--device_ssl_certificate_configs--upload_spec))

Optional:

* `issuer` (String) - issuer details of the certificate
* `not_after` (String) - date and time when certificate stops being valid (\[rfc3339\](https://xml2rfc.tools.ietf.org/public/rfc/html/rfc3339.html\#anchor14))
* `not_before` (String) - date and time when certificate starts being valid (\[rfc3339\](https://xml2rfc.tools.ietf.org/public/rfc/html/rfc3339.html\#anchor14))
* `signature_algorithm` (String)
* `subject` (String) - subject name of the certificate

<a id="nestedatt--config--device_ssl_certificate_configs--trusted_ca"></a>
### Nested Schema for `config.device_ssl_certificate_configs.trusted_ca`

Required:

* `name` (String) - name of the certificate

<a id="nestedatt--config--device_ssl_certificate_configs--upload_spec"></a>
### Nested Schema for `config.device_ssl_certificate_configs.upload_spec`

Optional:

* `info` (Attributes) - Certificate info (see [below for nested schema](#nestedatt--config--device_ssl_certificate_configs--upload_spec--info))
* `pem` (String) - contents of the certificate in pem format

<a id="nestedatt--config--device_ssl_certificate_configs--upload_spec--info"></a>
### Nested Schema for `config.device_ssl_certificate_configs.upload_spec.info`

Required:

* `name` (String) - name of the certificate
* `type` (String) - type of the certificate

Optional:

* `comment` (String) - a short description of the certificate
* `passphrase` (String) - used to decrypt pkcs12 and private keys

