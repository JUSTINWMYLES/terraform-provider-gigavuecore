resource "gigavuecore_flex_inline" "example" {
  alias = "example"
  cluster_configs = [{
    cluster_id = "example"
    export_criteria = {
      lsb = 0
    }
    export_type = "example"
    ib_pathway  = "example"
    source = {
      alias = "example"
      type  = "example"
    }
  }]
  keep_on_device = true
  resilient_config = {
    side_a = "example"
    side_b = "example"
  }
  target_traffic_path = "example"
}
