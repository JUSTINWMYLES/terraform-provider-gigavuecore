resource "gigavuecore_snmp_trap_receiver" "example" {
  alias          = "example"
  auth_password  = "example"
  auth_protocol  = "SHA"
  community      = "example"
  ip_address     = "example"
  priv_password  = "example"
  priv_protocol  = "DES"
  security_level = "noAuthNoPriv"
  snmp_port      = 0
  snmp_retries   = 0
  snmp_timeout   = 0
  snmp_version   = "v2c"
  user_name      = "example"
}
