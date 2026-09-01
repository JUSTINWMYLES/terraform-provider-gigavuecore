action "gigavuecore_configure_email_server" "example" {
  config {
    email_host       = "example"
    enable_smtp_auth = true
    from             = "example"
    password         = "example"
    port             = 0
    update_secret    = true
    user_name        = "example"
  }
}
