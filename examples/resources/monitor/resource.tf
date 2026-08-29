resource "gigavuecore_monitor" "example" {
  alias = "example"
  cache = {
    export_triggers = {
      event            = "example"
      timeout_active   = 0
      timeout_inactive = 0
    }
    type = "example"
  }
  cluster_id  = "example"
  description = "example"
  records     = [ "example" ]
  sampling = {
    mode                 = "example"
    single_sampling_rate = 0
  }
  sampling_space = 0
  ssl_port_restrictions = {
    ports     = [ 0 ]
    ssl_ports = "example"
  }
}
