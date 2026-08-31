action "gigavuecore_delete_pcap_file" "example" {
  config {
    cluster_id = "example"
    pcap_delete_list = [{
      box_id    = "example"
      file_name = "example"
    }]
  }
}
