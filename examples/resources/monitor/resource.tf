resource "gigavuecore_monitor" "example" {
  alias = "example"
  cache = {
    export_triggers = {
      event            = "example"
      timeout_active   = 1
      timeout_inactive = 1
    }
    type = "example"
  }
  description = "example"
  records     = [ "example" ]
  sampling = {
    mode                 = "example"
    single_sampling_rate = 1
  }
  sampling_space = 1
  ssl_port_restrictions = {
    ports     = [ 1 ]
    ssl_ports = "example"
  }
}
