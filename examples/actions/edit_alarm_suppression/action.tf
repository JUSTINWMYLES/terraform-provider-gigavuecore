action "gigavuecore_edit_alarm_suppression" "example" {
  config {
    suppressed_entities = [{
      alarm_types      = [ "example" ]
      alias            = "example"
      cluster_id       = "example"
      created_by       = "example"
      created_ts       = "example"
      enable           = true
      expiry_time      = 0
      expiry_time_unit = "minutes"
      expiry_ts        = "example"
      hostname         = "example"
      resource_id      = "example"
      resource_type    = "portPair"
      selected_reason  = "example"
      tags = [{
        tag_key    = "example"
        tag_values = [ "example" ]
      }]
      updated_by = "example"
      updated_ts = "example"
    }]
  }
}
