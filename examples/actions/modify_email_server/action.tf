action "gigavuecore_modify_email_server" "example" {
  config {
    domain_name               = "example"
    enable_auto_support_notif = true
    enable_smtp_auth          = true
    include_hostname          = true
    mail_hub_port             = 1
    password                  = "example"
    return_address            = "example"
    smtp_server               = "example"
    username                  = "example"
  }
}
