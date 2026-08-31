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
