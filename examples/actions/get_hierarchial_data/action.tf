action "gigavuecore_get_hierarchial_data" "example" {
  config {
    tags_filter = {
      parent_tags = [{
        tag_key   = "example"
        tag_value = "example"
      }]
      tag_key = "example"
    }
  }
}
