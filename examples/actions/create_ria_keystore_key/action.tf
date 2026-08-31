action "gigavuecore_create_ria_keystore_key" "example" {
  config {
    alias = "example"
    certificate = {
      certificate_file = "example"
      certificate_url = {
        hostname = "example"
        password = "example"
        path     = "example"
        protocol = "scp"
        username = "example"
      }
      passphrase = "example"
      type       = "certificate"
    }
    cluster_names = [ "example" ]
    comment       = "example"
    key = {
      key_file  = "example"
      key_label = "example"
      key_url = {
        hostname = "example"
        password = "example"
        path     = "example"
        protocol = "scp"
        username = "example"
      }
      passphrase = "example"
      type       = "private"
    }
  }
}
