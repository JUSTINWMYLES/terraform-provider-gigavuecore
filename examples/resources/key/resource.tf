resource "gigavuecore_key" "example" {
  alias      = "example"
  cluster_id = "example"
  comment    = "example"
  file       = "example"
  file_source = {
    hostname = "example"
    password = "example"
    path     = "example"
    protocol = "scp"
    username = "example"
  }
  key_label  = "example"
  passphrase = "example"
  type       = "pkcs11"
}
