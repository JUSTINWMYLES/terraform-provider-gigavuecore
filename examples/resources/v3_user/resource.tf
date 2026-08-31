resource "gigavuecore_v3_user" "example" {
  auth_key      = "example"
  auth_protocol = "md5"
  cluster_id    = "example"
  enabled       = true
  priv_key      = "example"
  priv_protocol = "des"
  read_only     = true
  username      = "example"
}
