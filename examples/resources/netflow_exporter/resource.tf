resource "gigavuecore_netflow_exporter" "example" {
  alias       = "example"
  cluster_id  = "example"
  description = "example"
  destination = {
    address = "example"
    ip_ver  = "v4"
  }
  dscp = 0
  filter = {
    rules = [{
      pass_rules = [{
        matches = [ "example" ]
        rule_id = 1
      }]
    }]
  }
  format     = "netflow"
  nf_version = "v5"
  snmp = {
    enabled = true
  }
  template_refresh = 1
  transport = {
    port     = 0
    protocol = "udp"
  }
  ttl           = 1
  tunneled_port = "example"
}
