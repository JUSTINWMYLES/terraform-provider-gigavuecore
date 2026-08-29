resource "gigavuecore_netflow_exporter" "example" {
  alias       = "example"
  cluster_id  = "example"
  description = "example"
  destination = {
    address = "example"
    ip_ver  = "example"
  }
  dscp = 0
  filter = {
    rules = [{
      pass_rules = [{
        matches = [ "example" ]
        rule_id = 0
      }]
    }]
  }
  format     = "example"
  nf_version = "example"
  snmp = {
    enabled = true
  }
  template_refresh = 0
  transport = {
    port     = 0
    protocol = "example"
  }
  ttl           = 0
  tunneled_port = "example"
}
