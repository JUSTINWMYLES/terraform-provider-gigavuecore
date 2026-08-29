action "gigavuecore_issue_acme_certificate" "example" {
  config {
    acme_certificate = [{
      acme_server_alias = "example"
      algorithm         = "example"
      domain            = "example"
      renew_days        = 0
    }]
  }
}
