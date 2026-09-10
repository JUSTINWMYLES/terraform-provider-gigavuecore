action "gigavuecore_configure_acme_certificate_details_of_device" "example" {
  config {
    acme_certificate = [{
      acme_server_url = "example"
      algorithm       = "rsa-2048"
      domain          = "example"
      renew_days      = 0
    }]
    operation_type = "example"
  }
}
