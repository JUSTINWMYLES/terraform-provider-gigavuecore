resource "gigavuecore_notif_meta_config" "example" {
  allow_attachment     = true
  attachment_limit     = 0
  comment              = "example"
  email_subject_prefix = "example"
  enabled              = true
  event_details = [{
    description   = "example"
    display_name  = "example"
    event_type    = "example"
    name          = "example"
    scope         = "example"
    severity      = ["Clear"]
    severity_type = "critical"
    sub_type      = "example"
  }]
  external_trap_receivers = ["example"]
  instant_rate_limit      = 0
  recipients              = ["example"]
  recurring_schedule      = "example"
  send_mail_if_empty      = true
  severity                = ["Clear"]
  tags = [{
    tag_key    = "example"
    tag_values = ["example"]
  }]
  task_id          = "example"
  task_name        = "example"
  template_details = ["example"]
  time_interval    = 1
  time_left        = "example"
  type             = "instant"
}
