action "gigavuecore_load_src_port_meta_data" "example" {
  config {
    source_details = [{
      cluster_id = "example"
      components = [{
        ids  = [ "example" ]
        type = "PORT"
      }]
    }]
  }
}
