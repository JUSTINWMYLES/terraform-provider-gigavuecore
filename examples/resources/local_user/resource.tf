resource "gigavuecore_local_user" "example" {
  account_status   = "accountDisabled"
  capability       = "admin"
  cluster_id       = "example"
  current_password = "example"
  enabled          = true
  full_name        = "example"
  roles            = ["example"]
  user_pwd         = "example"
  username         = "example"
}
