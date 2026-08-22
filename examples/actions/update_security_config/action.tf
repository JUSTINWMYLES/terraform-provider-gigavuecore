action "gigavuecore_update_security_config" "example" {
  config {
    allow_blank_password = true
    cluster_id = "example"
    fips_enabled = true
    fips_mode_state = true
    min_password_len = 1
    secure_crypto = true
    secure_crypto_enforced = true
    secure_passwords = true
  }
}
