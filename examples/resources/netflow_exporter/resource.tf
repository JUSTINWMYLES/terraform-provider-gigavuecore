resource "gigavuecore_netflow_exporter" "example" {
  alias = "example"
  cluster_id = "example"
  description = "example"
  destination = {
    address = "example"
    ip_ver = "example"
  }
  dscp = 1
  filter = {
    rules = "example"
  }
  format = "example"
  nf_version = "example"
  snmp = {
    enabled = true
  }
  template_refresh = 1
  transport = {
    port = 1
    protocol = "example"
  }
  ttl = 1
  tunneled_port = "example"
}
