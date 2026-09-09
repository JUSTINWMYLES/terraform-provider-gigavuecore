resource "gigavuecore_monitor" "example" {
  alias = "example"
  cache = {
    export_triggers = {
      event            = "txnEnd"
      timeout_active   = 1
      timeout_inactive = 1
    }
    type = "normal"
  }
  cluster_id  = "example"
  description = "example"
  records     = ["example"]
  sampling = {
    mode                 = "multi-rate"
    single_sampling_rate = 10
  }
  sampling_space = 0
  ssl_port_restrictions = {
    ports     = [0]
    ssl_ports = "all"
  }
}
