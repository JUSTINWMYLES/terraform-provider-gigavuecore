action "gigavuecore_add_devices_to_domain" "example" {
  config {
    async = true
    node_add_specs = [{
      https_port   = "example"
      node_address = "example"
      password     = "example"
      snmp_version = "v2"
      username     = "example"
    }]
  }
}
